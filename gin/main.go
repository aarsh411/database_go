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

		name := c.Query("name")
		department := c.Query("department")
		address := c.Query("address")

		c.JSON(200, gin.H{
			"name":       name,
			"department": department,
			"address":     address,
		})
		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO developer_team (name, department, address) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, department, address)
		fmt.Printf("name: %s; department: %s; address: %s", name, department, address)
	})
	router.GET("/delete", func(c *gin.Context) {
		id := c.Query("id")
		db := dbConn()
	
		delForm, err := db.Prepare("DELETE FROM developer_team WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(id)
		log.Println("DELETE")
		defer db.Close()
		
	})

	
	router.PUT("/update", func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")
		department := c.Query("department")
		address := c.Query("address")
		db := dbConn()

		c.JSON(200, gin.H{
			"name":       name,
			"department": department,
			"address":     address,
		})
	
		insForm, err := db.Prepare("UPDATE developer_team SET name=?, department=?, address=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, department, address, id)
		fmt.Printf("name: %s; department: %s; address: %s", name, department, address)
	})

	router.GET("/get", func(c *gin.Context) {
		id := c.Query("id")
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM developer_team WHERE id=?", id)
		if err != nil {
			panic(err.Error())
		}
		var name, address, department string
		for selDB.Next() {

			err = selDB.Scan(&id, &name, &department, &address)
			if err != nil {
				panic(err.Error())
			}
		}
		fmt.Printf("name: %s; department: %s; address: %s; salary: %d", name, department, address)

		c.JSON(200, gin.H{
			"id":         id,
			"name":       name,
			"department": department,
			"address":     address,
		})

	})

	router.Run(":8080")
}