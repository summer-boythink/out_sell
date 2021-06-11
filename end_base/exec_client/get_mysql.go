package exec_client

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //init()
)

func GetMysql() (db *sql.DB) {
	dsn := "root:root@tcp(121.5.118.127:3311)/out_sell"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	return db
}