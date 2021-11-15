package ordermodel

import "web-go/src/utils"

type OrderItem struct {
	ID        int
	OrderID   string
	StudentID int
}

func AddOrderItem(orderItem *OrderItem) error {
	sql := "insert into order_items(order_id, stu_id) values(?,?)"
	_, err := utils.Db.Exec(sql, orderItem.OrderID, orderItem.StudentID)
	return err
}

func GetItemsByOrderID(OrderID int) (orderItems []*OrderItem) {
	sqlStr := "select item_id, order_id, stu_id from order_items where order_id = ?"
	rows, err := utils.Db.Query(sqlStr, OrderID)
	if err != nil {
		return nil
	}
	for rows.Next() {
		item := &OrderItem{}
		err = rows.Scan(&item.ID, &item.OrderID, &item.StudentID)
		if err != nil {
			return orderItems
		}
		orderItems = append(orderItems, item)
	}
	return orderItems
}
