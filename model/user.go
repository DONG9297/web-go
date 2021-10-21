package model

// User 用户格式
type User struct {
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
