package controller

import (
	"net/http"
	"strconv"
	"web-go/src/model"
	"web-go/src/model/usermodel"
)

func GetAuth(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session := model.IsLogin(r)
	if !flag {
		// TODO: login
		return
	}
	// 解析请求
	no := r.PostFormValue("stu_no")
	name := r.PostFormValue("stu_name")
	gender := r.PostFormValue("stu_gender")
	email := r.PostFormValue("stu_email")
	if no == "" || name == "" || gender == "" || email == "" {
		model.Response(w, false, 500, "解析失败", nil)
		return
	}
	stu := usermodel.GetStudentByNo(no)
	if stu == nil || stu.Name != name {
		model.Response(w, false, 500, "填入信息有误", nil)
		return
	}
	// 将UserID填入学生表
	stu.Gender = gender
	stu.Email = email
	err := usermodel.UpdateStudent(stu)
	if err != nil {
		model.Response(w, false, 500, "修改信息失败", nil)
		return
	}
	stu = usermodel.GetStudentByNo(no)
	code, err := usermodel.AddAuth(session.UserID, stu.ID)
	if err != nil {
		model.Response(w, false, 500, "添加失败", nil)
		return
	}

	// 添加成功
	data := map[string]string{"auth_code": strconv.Itoa(code)}
	model.Response(w, true, 200, "成功", data)
}
