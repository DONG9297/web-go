package usermodel

import (
	"web-go/src/utils"
)

// User 用户格式
type User struct {
	ID       int
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
	//IsDelete bool   `json:"is_delete"`
}

func GetUserByPhone(phone string) (user *User) {
	sqlStr := "select user_id, phone, name, password from users where phone = ?"
	row := utils.Db.QueryRow(sqlStr, phone)
	user = &User{}
	err := row.Scan(&user.ID, &user.Phone, &user.Name, &user.Password)
	if err != nil {
		return nil
	}
	return user
}

func GetUserByID(id int) (user *User) {
	sqlStr := "select user_id,  phone, name, password from users where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, id)
	user = &User{}
	err := row.Scan(&user.ID, &user.Phone, &user.Name, &user.Password)
	if err != nil {
		return nil
	}
	return user
}

func AddUser(user *User) (err error) {
	sqlStr := "insert into users(phone, name, password) values(?,?,?,?)"
	_, err = utils.Db.Exec(sqlStr, user.Phone, user.Name, user.Password)
	return err
}
