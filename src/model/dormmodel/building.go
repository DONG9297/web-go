package dormmodel

import "web-go/src/utils"

type Building struct {
	ID   int
	Name string
}

func GetBuildingByName(name string) (building *Building) {
	sqlStr := "select building_id, name from buildings where name = ?"
	row := utils.Db.QueryRow(sqlStr, name)
	building = &Building{}
	err := row.Scan(&building.ID, &building.Name)
	if err != nil {
		return nil
	}
	return building
}

func GetBuildingByID(ID int) (building *Building) {
	sqlStr := "select building_id, name from buildings where building_id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	building = &Building{}
	err := row.Scan(&building.ID, &building.Name)
	if err != nil {
		return nil
	}
	return building
}
