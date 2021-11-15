package usermodel

import (
	"math/rand"
	"time"
	"web-go/src/utils"
)

type Auth struct {
	UserID    int
	StudentID int
	Code      int
}

func AddAuth(userID, studentID int) (code int, err error) {
	for {
		rand.Seed(time.Now().Unix())
		code = rand.Intn(1e8)
		row := utils.Db.QueryRow("Select code from auth_codes where code=?", code)
		var temp int
		err := row.Scan(&temp)
		if err != nil {
			break
		}
	}
	sqlStr := "insert into auth_codes(user_id, stu_id, code) values(?,?,?)"
	_, err = utils.Db.Exec(sqlStr, userID, studentID, code)
	return code, err
}

func GetAuthByStudentID(studentID int) (auth *Auth) {
	sqlStr := "select user_id, stu_id, code from auth_codes where stu_id = ?"
	row := utils.Db.QueryRow(sqlStr, studentID)
	auth = &Auth{}
	err := row.Scan(&auth.UserID, &auth.StudentID, &auth.Code)
	if err != nil {
		return nil
	}
	return auth
}
func GetAuthByUserID(userID int) (auth *Auth) {
	sqlStr := "select user_id, stu_id, code from auth_codes where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, userID)
	auth = &Auth{}
	err := row.Scan(&auth.UserID, &auth.StudentID, &auth.Code)
	if err != nil {
		return nil
	}
	return auth
}
func GetAuthByCode(code int) (auth *Auth) {
	sqlStr := "select user_id, stu_id, code from auth_codes where code = ?"
	row := utils.Db.QueryRow(sqlStr, code)
	auth = &Auth{}
	err := row.Scan(&auth.UserID, &auth.StudentID, &auth.Code)
	if err != nil {
		return nil
	}
	return auth
}
