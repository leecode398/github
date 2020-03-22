package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

//UserInfo user information
type UserInfo struct {
	ID     int
	Title  string
	Auther string
	Date   time.Time
}

func mysql() {
	db, err := sql.Open("mysql", "root:li000000@/WebTest?charset=utf8")
	checkErr(err)
	orm := beedb.New(db)
	var testUser UserInfo
	testUser.Title = "test beedb"
	testUser.Auther = "lee"
	testUser.Date = time.Now()
	orm.Save(&testUser)

	//查询
	var user UserInfo
	orm.Where("Auther=?", "lee").Find(&user)
	fmt.Println(user)
	// query
	rows, err := db.Query("SELECT * FROM web_table")
	checkErr(err)
	fmt.Printf("\nQuery:\n")
	fmt.Printf("%-5s%-15s%-10s%-20s\n","id", "title", "auther", "date")
	for rows.Next() {
		var id int
		var title string
		var auther string
		var date string

		err = rows.Scan(&id, &title, &auther, &date)
		checkErr(err)
		fmt.Printf("%-5d%-15s%-10s%-20s\n",id, title, auther, date)
	}

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	mysql()
}
