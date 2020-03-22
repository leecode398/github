package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:li000000@tcp(127.0.0.1:3306)/WebTest?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("select * from user_info")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var password string
		err = rows.Scan(&id, &name, &password)
		fmt.Printf("rows id = %d, value = %s, password:%s\n", id, name, password)
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
}
