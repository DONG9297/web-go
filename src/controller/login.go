package controller

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"regexp"
	"web-go/src/model"
	"web-go/src/model/usermodel"
	"web-go/src/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 如果已经登陆
	flag, session := model.IsLogin(r)
	if flag {
		//返回成功消息
		model.Response(w, true, 200, "登陆成功", map[string]string{"session id": session.SessionID})
		return
	}

	// 解析请求
	phone := r.PostFormValue("phone")
	password := r.PostFormValue("password")
	if phone == "" || password == "" {
		model.Response(w, false, 500, "解析失败", nil)
		return
	}

	// 检查格式
	matched1, _ := regexp.MatchString(`^1\d{10}`, phone)        // 手机号
	matched2, _ := regexp.MatchString(`^[\w_]{6,20}`, password) // 密码必须是6-20位的字母、数字或下划线
	if !(matched1 && matched2) {
		// 格式不正确
		model.Response(w, false, 500, "登陆失败，手机号或密码不正确", nil)
		return
	}

	//格式正确，查询手机号和密码是否匹配
	//加密
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	user := usermodel.GetUserByPhone(phone)
	if user == nil || user.Password != password {
		//手机号或密码不正确
		model.Response(w, false, 500, "登陆失败，手机号或密码不正确", nil)
		return
	}

	//用户名和密码正确
	//生成UUID作为Session的id
	uuid := utils.CreateUUID()
	//创建一个Cookie，让它与Session相关联
	cookie := http.Cookie{
		Name:  "dorm_user",
		Value: uuid,
	}
	//将cookie发送给浏览器
	http.SetCookie(w, &cookie)
	//将Session保存到数据库中
	model.AddSession(uuid, user.ID)

	//返回成功消息
	model.Response(w, true, 200, "登陆成功", map[string]string{"session id": uuid})
	return
}
