package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Mysql() {
	db, err := sql.Open("mysql", "root:lx000000@/test?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT user_info SET id=?,name=?")
	checkErr(err)

	res, err := stmt.Exec(15, "qq")
	checkErr(err)

	// update
	stmt, err = db.Prepare("update user_info set name=? where id=?")
	checkErr(err)

	res, err = stmt.Exec("wangshubo_qeuqpdate", 4)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM user_info")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string

		err = rows.Scan(&uid, &username)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
	}

	// delete
	stmt, err = db.Prepare("delete from user_info where id=?")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	// query
	rows, err = db.Query("SELECT * FROM user_info")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string

		err = rows.Scan(&uid, &username)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
	}

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
