package ordermodel

import "web-go/src/utils"

type Order struct {
	ID         string
	UserID     int
	Count      int
	BuildingID int
	Gender     string
	CreateTime string
	State      int // 0 未处理， 1 成功， 2 失败
}

func AddOrder(order *Order) error {
	sql := "insert into orders(order_id, user_id, count, building_id, gender, create_time, state) values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, order.ID, order.UserID, order.Count, order.BuildingID, order.Gender, order.CreateTime, order.State)
	if err != nil {
		return err
	}
	return nil
}

func GetAllOrders() (orders []*Order) {
	sql := "select order_id, user_id, count, building_id, gender, create_time, state from orders"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil
	}
	for rows.Next() {
		order := &Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Count, &order.BuildingID, &order.Gender, &order.CreateTime, &order.State)
		if err != nil {
			return orders
		}
		orders = append(orders, order)
	}
	return orders
}
func GetUnprocessedOrdersBefore(timeStr string) (orders []*Order) {
	sql := "select order_id, user_id, count, building_id, gender, create_time, state from orders where create_time<? and state = 0 ORDER BY  create_time ASC"
	rows, err := utils.Db.Query(sql, timeStr)
	if err != nil {
		return nil
	}
	for rows.Next() {
		order := &Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Count, &order.BuildingID, &order.Gender, &order.CreateTime, &order.State)
		if err != nil {
			return orders
		}
		orders = append(orders, order)
	}
	return orders
}

func GetOrdersByUserID(userID int) (orders []*Order) {
	sql := "select order_id, user_id, count, building_id, gender, create_time, state from orders where user_id = ?"
	rows, err := utils.Db.Query(sql, userID)
	if err != nil {
		return nil
	}
	for rows.Next() {
		order := &Order{}
		err := rows.Scan(&order.ID, &order.UserID, &order.Count, &order.BuildingID, &order.Gender, &order.CreateTime, &order.State)
		if err != nil {
			return orders
		}
		orders = append(orders, order)
	}
	return orders
}

func UpdateOrderState(orderID string, state int) error {
	//写sql语句
	sql := "update orders set state = ? where order_id = ?"
	//执行
	_, err := utils.Db.Exec(sql, state, orderID)
	return err
}
