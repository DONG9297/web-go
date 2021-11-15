package main

import (
	"net/http"
	"time"
	"web-go/src/controller"
)

func main() {
	// 首页
	//http.HandleFunc("/", )
	// 注册
	http.HandleFunc("/register", controller.Register)
	// 登陆
	http.HandleFunc("/login", controller.Login)
	// 退出
	http.HandleFunc("/logout", controller.Logout)
	// 填写学生信息，获取认证码
	http.HandleFunc("/studentInfo", controller.GetAuth)
	// 选宿舍
	http.HandleFunc("/chooseDorm", controller.ChooseDorm)
	// 查询结果
	http.HandleFunc("/result", controller.GetResult)

	http.ListenAndServe(":10700", nil)

	// 每10秒处理一次订单
	ticker := time.NewTicker(10 * time.Second)
	for _ = range ticker.C {
		m, _ := time.ParseDuration("-5s") //当前时间减5s，防止读写冲突
		timeStr := time.Now().Add(m).Format("2006-01-02 15:04:05")
		controller.ProcessOrder(timeStr)
	}
}
