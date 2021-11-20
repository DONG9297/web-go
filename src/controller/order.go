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
	gender := stu.Gender
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
				model.Response(w, false, 500, "同住人不存在", map[string]string{"auth_code": codeStr})
				return
			}
			stu := usermodel.GetStudentByID(auth.StudentID)
			if stu == nil || stu.Gender != gender {
				model.Response(w, false, 500, "同住人性别错误", map[string]string{"auth_code": codeStr})
				return
			}
			if model.HasStudentChosenDorm(stu.ID) {
				model.Response(w, false, 500, "同住人已选宿舍", map[string]string{"auth_code": codeStr})
				return
			}
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
	model.Response(w, true, 200, "成功", nil)
	return
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

// ProcessOrder 处理时间早于timeStr发生的订单
func ProcessOrder(timeStr string) {
	// 获取待处理订单
	orders := ordermodel.GetUnprocessedOrdersBefore(timeStr)
	for _, order := range orders {
		// 获取满足订单条件的宿舍列表
		var dorms []*dormmodel.Dorm
		units := dormmodel.GetUnitsByBuilding(order.BuildingID)
		for _, unit := range units {
			dorms = append(dorms, dormmodel.GetAvailableDorms(unit.ID, order.Count, order.Gender)...)
		}
		// 如果宿舍列表不为空
		if len(dorms) > 0 {
			dorm := dorms[0]
			availableBeds := dorm.AvailableBeds - order.Count
			// 更新宿舍空床数
			dormmodel.UpdateDormAvailableBeds(dorm.ID, availableBeds)
			// 将选宿舍信息加入学生宿舍表
			items := ordermodel.GetItemsByOrderID(order.ID)
			for _, item := range items {
				model.AddStudentIntoDorm(item.StudentID, dorm.ID)
			}
			// 更新订单状态
			ordermodel.UpdateOrderState(order.ID, 1)
		} else {
			ordermodel.UpdateOrderState(order.ID, 2)
		}
	}
}
