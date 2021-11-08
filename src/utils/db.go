package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "test:123456*@tcp(39.107.35.228:3306)/test")
	if err != nil {
		panic(err.Error())
	}
}
