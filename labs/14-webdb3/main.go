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
	selDB, err := db.Query("SELECT * FROM Employees ORDER BY id")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id, salary int
		var name string
		err = selDB.Scan(&id, &name, &salary)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Salary = salary
		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employees WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id, salary int
		var name string
		err = selDB.Scan(&id, &name, &salary)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Salary = salary
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
	selDB, err := db.Query("SELECT * FROM Employees WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id, salary int
		var name string
		err = selDB.Scan(&id, &name, &salary)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Salary = salary
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		salary := r.FormValue("salary")
		insForm, err := db.Prepare("INSERT INTO Employees(name, salary) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, salary)
		log.Println("INSERT: Name: " + name + " | Salary: " + salary)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		salary := r.FormValue("salary")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employees SET name=?, salary=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, salary, id)
		log.Println("UPDATE: Name: " + name + " | Salary: " + salary)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	_, err := db.Query("DELETE FROM Employees WHERE id=?", nId)
	log.Println("DELETE")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/remove", Delete)
	http.ListenAndServe(":8080", nil)
}
