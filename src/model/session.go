package model

import (
	"net/http"
	"web-go/src/utils"
)

// Session 结构
type Session struct {
	SessionID string
	UserID    int
}

//AddSession 向数据库中添加Session
func AddSession(sessionID string, userID int) error {
	sqlStr := "insert into sessions (session_id, user_id) values(?,?)"
	_, err := utils.Db.Exec(sqlStr, sessionID, userID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 删除数据库中的Session
func DeleteSession(sessionID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessionID)
	if err != nil {
		return err
	}
	return nil
}

//GetSession 根据session的Id值从数据库中查询Session
func GetSession(sessionID string) (sess *Session) {
	sqlStr := "select session_id, user_id from sessions where session_id = ?"
	row := utils.Db.QueryRow(sqlStr, sessionID)
	//扫描数据库中的字段值为Session的字段赋值
	sess = &Session{}
	err := row.Scan(&sess.SessionID, &sess.UserID)
	if err != nil {
		return nil
	}
	return sess
}

//IsLogin 判断用户是否已经登录 false 没有登录 true 已经登录
func IsLogin(r *http.Request) (bool, *Session) {
	//根据Cookie的name获取Cookie
	cookie, _ := r.Cookie("dorm_user")
	if cookie != nil {
		//获取Cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库中查询与之对应的Session
		session := GetSession(cookieValue)
		if session == nil {
			return false, nil
		}
		return true, session
	}
	//没有登录
	return false, nil
}
