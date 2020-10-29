package main

import (
    "database/sql"
    "log"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type developer_team struct {
    Id    int
    Name  string
    Department string
	Address	string
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

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM developer_team ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }
    emp := developer_team{}
    res := []developer_team{}
    for selDB.Next() {
        var id int
        var name, department, address string
        err = selDB.Scan(&id, &name, &department, &address)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Department = department
		emp.Address = address
        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM developer_team WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := developer_team{}
    for selDB.Next() {
        var id int
        var name, department, address string
        err = selDB.Scan(&id, &name, &department, &address)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Department = department
		emp.Address = address
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM developer_team WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := developer_team{}
    for selDB.Next() {
        var id int
        var name, department, address string
        err = selDB.Scan(&id, &name, &department, &address)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Department = department
		emp.Address = address
    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        department := r.FormValue("department")
		address := r.FormValue("address")

        insForm, err := db.Prepare("INSERT INTO developer_team(name, department, address) VALUES(?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, department, address)
        log.Println("INSERT: Name: " + name + " | department: " + department +  " | Address" + address)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        department := r.FormValue("department")
		address := r.FormValue("address")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE developer_team SET name=?, department=?, address=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, department, address, id)
        log.Println("INSERT: Name: " + name + " | department: " + department + " | Address" + address)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM developer_team WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {

    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}