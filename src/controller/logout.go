package controller

import (
	"net/http"
	"web-go/src/model"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("dorm_user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		err := model.DeleteSession(cookieValue)
		if err != nil {
			model.Response(w, false, 500, "未登录", nil)
			return
		}
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改之后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//返回成功消息
	model.Response(w, true, 200, "退出成功", nil)
	return
}
