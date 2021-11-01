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
	Db, err = sql.Open("mysql", "root:123456@(userdb:3307)/my_db")
	if err != nil {
		panic(err.Error())
	}
}
