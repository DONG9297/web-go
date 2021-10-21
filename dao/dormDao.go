package dao

import (
	"web-go/model"
	"web-go/utils"
)

func GetDorms()([]*model.Dorm,error){
	//写sql语句
	sql := "select dorm_name, building_name, beds_amount, availiable_beds_count from dorms"
	//执行
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var dorms []*model.Dorm
	for rows.Next() {
		dorm := &model.Dorm{}
		rows.Scan(&dorm.DormName, &dorm.BuildingName, &dorm.BedsCount, &dorm.AvailableBedsCount)
		dorms = append(dorms, dorm)
	}
	return dorms, nil
}