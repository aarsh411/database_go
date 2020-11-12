package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id         int    `json:id`
	Name       string `json:name`
	Department string `json:department`
	Address     int    `json:address`
	Salary      int   `json:salary`
}
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "aarsh"
	dbPass := "1234"
	dbName := "sql_go"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	router := gin.Default()

	router.POST("/add", func(c *gin.Context) {
		var e Employee

		//name := c.Query("name")
		//department := c.Query("department")
		//address := c.Query("address")
		//salary := c.Query("salary")
		if c.BindJSON(&e)== nil{
		

		c.JSON(200, gin.H{
			"name":       e.Name,
			"department": e.Department,
			"address":     e.Address,
			"salary":   e.Salary,
		})
		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO developer_team (name, department, address, salary) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(e.Name, e.Department, e.Address, e.Salary)
		fmt.Printf("name: %s; department: %s; address: %s; salary: %d", e.Name, e.Department, e.Address, e.Salary)
		}
	})
	router.DELETE("/delete", func(c *gin.Context) {
		var e Employee
		//id := c.Query("id")
		if c.BindJSON(&e) == nil {

			db := dbConn()
	
			delForm, err := db.Prepare("DELETE FROM developer_team WHERE id=?")
			if err != nil {
				panic(err.Error())
			}
			delForm.Exec(e.Id)
			log.Println("DELETE")
			defer db.Close()
		}
		
	})

	
	router.PUT("/update", func(c *gin.Context) {
		//id := c.Query("id")
		//name := c.Query("name")
		//department := c.Query("department")
		//address := c.Query("address")
		//salary := c.Query("salary")
		var e Employee
	if c.BindJSON(&e) == nil {
		db := dbConn()

		c.JSON(200, gin.H{
			"name":       e.Name,
			"department": e.Department,
			"address":     e.Address,
			"salary":   e.Salary,
		})
	
		insForm, err := db.Prepare("UPDATE developer_team SET name=?, department=?, address=?, salary=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(e.Name, e.Department, e.Address, e.Salary, e.Id)
		//fmt.Printf("name: %s; department: %s; address: %s; salary: %d", e.Name, e.Department, e.Address, e.Salary)
		}
	})

	router.GET("/get", func(c *gin.Context) {
		//id := c.Query("id")
		db := dbConn()
		var e Employee
		if c.BindJSON(&e) == nil {
		selDB, err := db.Query("SELECT * FROM developer_team WHERE id=?", e.Id)
		if err != nil {
			panic(err.Error())
		}
		var name, address, department string
		var salary int
		for selDB.Next(){

			err = selDB.Scan(&e.Id, &e.Name, &e.Department, &e.Address, &e.Salary)
			if err != nil {
				panic(err.Error())
			}
		}
		fmt.Printf("name: %s; department: %s; address: %s; salary: %d", name, department, address, salary)

		c.JSON(200, gin.H{
			"id":         e.Id,
			"name":       e.Name,
			"department": e.Department,
			"address":     e.Address,
			"salary": e.Salary,
		})
	
	}

	})

	router.Run(":8080")
}