package controller

import (
	"fmt"
	"net/http"
	"web-go/dao"
)

func ListDorms(w http.ResponseWriter, r *http.Request) {
	flag, _ := dao.IsLogin(r)
	if flag{
		fmt.Fprintf(w, "宿舍名\t楼号\t床位数\t空余床位数\n")
		dorms, _ := dao.GetDorms()
		for _,dorm:= range dorms{
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n",dorm.DormName,dorm.BuildingName,dorm.BedsCount,dorm.AvailableBedsCount)
		}
	}else{
		fmt.Fprintf(w,"请先登录")
	}
}
