package controller

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"regexp"
	"web-go/src/model"
	"web-go/src/model/usermodel"
)

func Register(w http.ResponseWriter, r *http.Request) {

	// 解析请求
	phone := r.PostFormValue("phone")
	name := r.PostFormValue("name")
	password := r.PostFormValue("password")
	if phone == "" || name == "" || password == "" {
		model.Response(w, false, 500, "解析失败", nil)
		return
	}

	// 检查格式
	matched1, _ := regexp.MatchString(`^1\d{10}`, phone)          // 手机号
	matched2, _ := regexp.MatchString(`^[^0-9][\w_]{3,11}`, name) // 用户名必须是4-12位字母、数字或下划线，不能以数字开头
	matched3, _ := regexp.MatchString(`^[\w_]{6,20}`, password)   // 密码必须是6-20位的字母、数字或下划线
	if !(matched1 && matched2 && matched3) {
		// 格式不正确
		model.Response(w, false, 500, "输入格式不正确", nil)
		return
	}

	//加密
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))

	//添加用户
	user := usermodel.GetUserByPhone(phone)
	if user != nil && user.Phone == phone {
		// 手机号存在
		model.Response(w, false, 500, "手机号已注册", nil)
		return
	}
	err := usermodel.AddUser(&usermodel.User{
		Phone:    phone,
		Name:     name,
		Password: password,
	})
	if err != nil {
		// 添加用户失败
		model.Response(w, false, 500, "添加用户失败", nil)
		return
	}
	// 添加用户成功
	model.Response(w, true, 200, "成功", nil)
}
