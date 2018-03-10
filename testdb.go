package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "base:imkloKuLiqNMc6Cn@tcp(172.31.215.37:3306)/innotree_data_online?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT id,date,downNum,ftpStartDate FROM pdf_down_parse_record limit 3")
	checkErr(err)

	for rows.Next() {
		var id int
		var date string
		var downNum int
		var ftpStartDate string
		err = rows.Scan(&id, &date, &downNum, &ftpStartDate)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(date)
		fmt.Println(downNum)
		fmt.Println(ftpStartDate)
	}
}
