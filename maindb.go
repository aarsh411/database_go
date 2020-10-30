package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	db, err := sql.Open("mysql", "aarsh:1234@(127.0.0.1:3306)/kloudonedata")
	fmt.Println(db)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	//insert, err := db.Query("INSERT INTO users VALUES ( 'aarsh' )")
	insert, err := db.Query("INSERT INTO golang VALUES ( 2,'Dhiraj','2020-09-28'),(3,'divyang','2020-09-28')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
