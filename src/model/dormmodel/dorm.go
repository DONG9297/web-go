package dormmodel

import "web-go/src/utils"

type Dorm struct {
	ID            int
	Name          string
	Gender        string
	TotalBeds     int
	AvailableBeds int
	UnitID        int
}

func GetAvailableDorms(unitID, available_beds int, gender string) (dorms []*Dorm) {
	sqlStr := "select dorm_id, name, gender, total_beds, available_beds, unit_id from dorms where uint_id = ? and gender =? and available_beds >= ?"
	rows, err := utils.Db.Query(sqlStr, unitID, gender, available_beds)
	if err != nil {
		return nil
	}
	for rows.Next() {
		dorm := &Dorm{}
		err = rows.Scan(&dorm.ID, &dorm.Name, &dorm.Gender, &dorm.TotalBeds, &dorm.AvailableBeds, &dorm.UnitID)
		if err != nil {
			return dorms
		}
		dorms = append(dorms, dorm)
	}
	return dorms
}

func GetDormByID(ID int) (dorm *Dorm) {
	sqlStr := "select dorm_id, name, gender, total_beds, available_beds, unit_id from dorms where dorm_id = ?"
	row := utils.Db.QueryRow(sqlStr, ID)
	dorm = &Dorm{}
	err := row.Scan(&dorm.ID, &dorm.Name, &dorm.Gender, &dorm.TotalBeds, &dorm.AvailableBeds, &dorm.UnitID)
	if err != nil {
		return nil
	}
	return dorm
}
func UpdateDormAvailableBeds(dormID, availableBeds int) (err error) {
	sqlStr := "update dorms set available_beds = ? where dorm_id = ?"
	_, err = utils.Db.Exec(sqlStr, availableBeds, dormID)
	if err != nil {
		return err
	}
	return nil

}
