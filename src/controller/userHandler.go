package controller

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"web-go/src/dao"
	"web-go/src/model"
	"web-go/src/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	// 解析请求
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&user)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	// 判断格式
	reg1, _ := regexp.MatchString(`^1\d{10}`, user.Phone)          // 手机号
	reg2, _ := regexp.MatchString(`^[^0-9][\w_]{3,11}`, user.Name) // 用户名必须是4-12位字母、数字或下划线，不能以数字开头
	reg3, _ := regexp.MatchString(`^[\w_]{6,20}`, user.Password)   // 密码必须是6-20位的字母、数字或下划线
	if !(reg1 && reg2 && reg3) {
		// 格式不正确
		rst := &model.Result{
			Code: 500,
			Msg:  "输入格式不正确",
			Data: []string{},
		}
		response, _ := json.Marshal(rst)  // json化结果集
		fmt.Fprintln(w, string(response)) // 返回结果
		return
	} else {
		//加密
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
		//添加用户
		err := dao.AddUser(&user)
		if err != nil {
			//手机号存在
			rst := &model.Result{
				Code: 500,
				Msg:  "手机号已注册",
				Data: []string{},
			}
			response, _ := json.Marshal(rst)  // json化结果集
			fmt.Fprintln(w, string(response)) // 返回结果
			return
		}
		rst := &model.Result{
			Code: 200,
			Msg:  "注册成功",
			Data: []string{strconv.Itoa(user.ID)},
		}
		response, _ := json.Marshal(rst)
		fmt.Fprintln(w, string(response))

	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	// 解析请求
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&user)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	// 判断格式
	reg1, _ := regexp.MatchString(`^1\d{10}`, user.Phone)        // 手机号
	reg3, _ := regexp.MatchString(`^[\w_]{6,20}`, user.Password) // 密码必须是6-20位的字母、数字或下划线
	if !(reg1 && reg3) {
		//格式有误
		rst := model.Result{
			Code: 500,
			Msg:  "登陆失败，手机号或密码不正确",
			Data: []string{},
		}
		response, _ := json.Marshal(rst)  // json化结果集
		fmt.Fprintln(w, string(response)) // 返回结果
		return
	} else {
		//格式正确，查询手机号和密码是否匹配
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
		user, err := dao.CheckUserPhoneAndPassword(user.Phone, user.Password)
		if err != nil {
			//手机号或密码不正确
			rst := model.Result{
				Code: 500,
				Msg:  "登陆失败，手机号或密码不正确",
				Data: []string{},
			}
			response, _ := json.Marshal(rst)  // json化结果集
			fmt.Fprintln(w, string(response)) // 返回结果
			return
		} else {

			//用户名和密码正确
			token, err := utils.GenerateToken(user.Phone, user.Name)
			if err != nil {
				rst := model.Result{
					Code: 500,
					Msg:  "未成功生成Token",
					Data: []string{},
				}
				response, _ := json.Marshal(rst)  // json化结果集
				fmt.Fprintln(w, string(response)) // 返回结果
			}

			cookie := http.Cookie{
				Name:  "dorm_user",
				Value: token,
				//HttpOnly: true,
			}

			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			//返回成功消息
			rst := model.Result{
				Code: 200,
				Msg:  "登陆成功",
				Data: []string{token},
			}
			response, _ := json.Marshal(rst)  // json化结果集
			fmt.Fprintln(w, string(response)) // 返回结果
			return
			////用户名和密码正确
			////生成UUID作为Session的id
			//uuid := utils.CreateUUID()
			////创建一个Session
			//sess := &model.Session{
			//	SessionID: uuid,
			//	UserID:    user.ID,
			//}
			////将Session保存到数据库中
			//dao.AddSession(sess)
			////创建一个Cookie，让它与Session相关联
			//cookie := http.Cookie{
			//	Name:  "dorm_user",
			//	Value: uuid,
			//	//HttpOnly: true,
			//}
			////将cookie发送给浏览器
			//http.SetCookie(w, &cookie)
			////返回成功消息
			//rst := model.Result{
			//	Code: 200,
			//	Msg:  "登陆成功",
			//	Data: []string{sess.SessionID},
			//}
			//response, _ := json.Marshal(rst)  // json化结果集
			//fmt.Fprintln(w, string(response)) // 返回结果
			//return
		}
	}
}
