package usermodel

import (
	"web-go/src/utils"
)

type Student struct {
	ID     int
	No     string `json:"stu_no"`
	Name   string `json:"stu_name"`
	Gender string `json:"stu_gender"`
	Email  string `json:"stu_email"`
}

func GetStudentByNo(stuNo string) (stu *Student) {
	sqlStr := "select stu_id, stu_no, stu_name, stu_gender, stu_email from students where stu_no = ?"
	row := utils.Db.QueryRow(sqlStr, stuNo)
	stu = &Student{}
	err := row.Scan(&stu.ID, &stu.No, &stu.Name, &stu.Gender, &stu.Email)
	if err != nil {
		return nil
	}
	return stu
}
func GetStudentByID(stuID int) (stu *Student) {
	sqlStr := "select stu_id, stu_no, stu_name, stu_gender, stu_email from students where stu_id = ?"
	row := utils.Db.QueryRow(sqlStr, stuID)
	stu = &Student{}
	err := row.Scan(&stu.ID, &stu.No, &stu.Name, &stu.Gender, &stu.Email)
	if err != nil {
		return nil
	}
	return stu
}

func AddStudent(stu *Student) (err error) {
	sqlStr := "insert into students(stu_no, stu_name, stu_gender, stu_email) values(?,?,?,?,?)"
	_, err = utils.Db.Exec(sqlStr, stu.No, stu.Name, stu.Gender, stu.Email)
	if err != nil {
		return err
	}
	return nil
}

func UpdateStudent(stu *Student) (err error) {
	sqlStr := "update students set stu_gender = ?, stu_email = ?  where stu_no = ? and stu_name = ?"
	_, err = utils.Db.Exec(sqlStr, stu.Gender, stu.Email, stu.No, stu.Name)
	if err != nil {
		return err
	}
	return nil
}
