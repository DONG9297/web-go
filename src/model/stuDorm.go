package model

import (
	"web-go/src/utils"
)

type StuDorm struct {
	ID        int
	StudentID int
	DormID    int
}

func AddStudentIntoDorm(studentID, dormID int) (err error) {
	sqlStr := "insert into stu_dorm(stu_id, dorm_id) values(?,?)"
	_, err = utils.Db.Exec(sqlStr, studentID, dormID)
	return err
}

func DeleteStudentFromDorm(studentID int) (err error) {
	sqlStr := "delete from stu_dorm where stu_id = ?"
	_, err = utils.Db.Exec(sqlStr, studentID)
	return err
}

func GetStuDormsByDormID(dormID int) (stuDorms []*StuDorm) {
	sqlStr := "select id, stu_id, dorm_id from stu_dorm where dorm_id = ?"
	rows, err := utils.Db.Query(sqlStr, dormID)
	if err != nil {
		return nil
	}
	for rows.Next() {
		stuDorm := &StuDorm{}
		err = rows.Scan(&stuDorm.ID, &stuDorm.StudentID, &stuDorm.DormID)
		if err != nil {
			return stuDorms
		}
		stuDorms = append(stuDorms, stuDorm)
	}
	return stuDorms
}

func GetStuDormByStudentID(studentID int) (stuDorm *StuDorm) {
	sqlStr := "select id, stu_id, dorm_id from stu_dorm where stu_id = ?"
	stuDorm = &StuDorm{}
	row := utils.Db.QueryRow(sqlStr, studentID)
	err := row.Scan(&stuDorm.ID, &stuDorm.StudentID, &stuDorm.DormID)
	if err != nil {
		return nil
	}
	return stuDorm
}

func HasStudentChosenDorm(studentID int) bool {
	stuDorm := GetStuDormsByDormID(studentID)
	if stuDorm == nil {
		return false
	}
	return true
}
