package controller

import (
	"net/http"
	"strconv"
	"time"
	"web-go/src/model"
	"web-go/src/model/dormmodel"
	"web-go/src/model/ordermodel"
	"web-go/src/model/usermodel"
	"web-go/src/utils"
)

func ChooseDorm(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session := model.IsLogin(r)
	if !flag {
		// TODO: login
		return
	}

	//判断是否填了信息
	auth := usermodel.GetAuthByUserID(session.UserID)
	if auth == nil {
		// TODO: studentinfo
		return
	}

	// 解析请求
	buildingName := r.PostFormValue("building")
	building := dormmodel.GetBuildingByName(buildingName)
	if building == nil {
		model.Response(w, false, 500, "格式错误", nil)
		return
	}
	var students []*usermodel.Student
	stu := usermodel.GetStudentByID(auth.StudentID)
	students = append(students, stu)
	for i := 0; i < 3; i++ {
		key := "code" + strconv.Itoa(i)
		codeStr := r.PostFormValue(key)
		if codeStr != "" {
			code, err := strconv.Atoi(codeStr)
			if err != nil {
				model.Response(w, false, 500, "格式错误", nil)
				return
			}
			auth := usermodel.GetAuthByCode(code)
			if auth == nil {
				model.Response(w, false, 500, "同住人不存在", map[string]string{"code": codeStr})
				return
			}
			stu := usermodel.GetStudentByID(auth.StudentID)
			students = append(students, stu)
		}
	}

	//创建生成订单的时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	//生成订单
	order := &ordermodel.Order{
		ID:         utils.CreateUUID(),
		UserID:     session.UserID,
		Count:      len(students),
		BuildingID: building.ID,
		Gender:     stu.Gender,
		CreateTime: timeStr,
		State:      0,
	}
	ordermodel.AddOrder(order)

	for _, stu := range students {
		orderItem := &ordermodel.OrderItem{
			OrderID:   order.ID,
			StudentID: stu.ID,
		}
		ordermodel.AddOrderItem(orderItem)
	}
}

func GetResult(w http.ResponseWriter, r *http.Request) {
	// 判断是否登录
	flag, session := model.IsLogin(r)
	if !flag {
		// TODO: login
		return
	}
	// 判断是否填了信息
	auth := usermodel.GetAuthByUserID(session.UserID)
	if auth == nil {
		// TODO: studentinfo
		return
	}
	// 判断是否选宿舍成功
	stuDorm := model.GetStuDormByStudentID(auth.StudentID)
	if stuDorm == nil {
		// TODO: choose dorm
		return
	}

	dorm := dormmodel.GetDormByID(stuDorm.DormID)
	unit := dormmodel.GetUnitByID(dorm.UnitID)
	building := dormmodel.GetBuildingByID(unit.BuildingID)

	data := map[string]string{"dorm": dorm.Name, "building": building.Name}
	model.Response(w, true, 200, "成功", data)
	return
}

func ProcessOrder(timeStr string) {
	orders := ordermodel.GetUnprocessedOrdersBefore(timeStr)
	for _, order := range orders {
		var dorms []*dormmodel.Dorm
		units := dormmodel.GetUnitsByBuilding(order.BuildingID)
		for _, unit := range units {
			dorms = append(dorms, dormmodel.GetAvailableDorms(unit.ID, order.Count, order.Gender)...)
		}
		if len(dorms) > 0 {
			dorm := dorms[0]
			availableBeds := dorm.AvailableBeds - order.Count
			dormmodel.UpdateDormAvailableBeds(dorm.ID, availableBeds)
			ordermodel.UpdateOrderState(order.ID, 1)
		} else {
			ordermodel.UpdateOrderState(order.ID, 2)
		}
	}

}
