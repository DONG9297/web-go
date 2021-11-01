package dao

import (
	"web-go/src/model"
	"web-go/src/utils"
)

func CheckUserPhoneAndPassword(phone string, password string) (*model.User, error) {

	sqlStr := "select id, phone, name, password from users where phone = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, phone, password)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Phone, &user.Name, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func AddUser(user *model.User) error {
	sqlStr := "insert into users(phone, name, password) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, user.Phone, user.Name, user.Password)
	if err != nil {
		return err
	}
	utils.Db.QueryRow("SELECT id FROM users WHERE phone=?", user.Phone).Scan(&user.ID)
	return nil

}

func GetUserByID(id int) (*model.User, error) {
	sqlStr := "select id, phone, name, password from users where id = ?"
	row := utils.Db.QueryRow(sqlStr, id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Phone, &user.Name, &user.Password)
	if err != nil {
		return user, err
	}
	return user, err
}
