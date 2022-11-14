package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id     int
	Name   string
	Salary int
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Welcome1"
	dbName := "hrd"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM employees ORDER BY id")
	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	res := []Employee{}
	for selDB.Next() {
		err = selDB.Scan(&employee.Id, &employee.Name, &employee.Salary)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, employee)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM employees WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	employee := Employee{}
	for selDB.Next() {
		err = selDB.Scan(&employee.Id, &employee.Name, &employee.Salary)
		if err != nil {
			panic(err.Error())
		}
	}
	tmpl.ExecuteTemplate(w, "Show", employee)
	defer db.Close()
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.ListenAndServe(":8080", nil)
}
