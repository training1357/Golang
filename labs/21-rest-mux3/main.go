package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func main() {
	// Initialize Database
	Connect("root:Welcome1@tcp(127.0.0.1:3306)/sales")

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	log.Println("Starting Server on port 8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8080), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", GetProducts).Methods("GET")
}

func Connect(connectionString string) {
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product
	DB.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
