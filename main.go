package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

var database *sql.DB
var err error

//Product type
type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func setDatabase(dba *sql.DB) {
	database = dba
}

func main() {
	fmt.Println("Hello Docker Tutorial")

	connStr := fmt.Sprintf("host=db port=5432 user=postgresUser password=123456 dbname=test sslmode=disable")

	db, err := sql.Open("postgres", connStr)
	setDatabase(db)
	checkError(err)
	defer db.Close()
	err = db.Ping()
	fmt.Println("Succesfully connected")

	// Drop previous table of same name if one exists.
	_, err = database.Exec("DROP TABLE IF EXISTS products;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")

	// Create table.
	_, err = db.Exec("CREATE TABLE products (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table")
	fmt.Println("Inserting values")
	sqlStatement := "INSERT INTO products (name, quantity) VALUES ($1,$2);"
	_, err = db.Exec(sqlStatement, "banana", 150)
	checkError(err)
	_, err = db.Exec(sqlStatement, "orange", 154)
	checkError(err)
	_, err = db.Exec(sqlStatement, "apple", 100)
	checkError(err)

	router := mux.NewRouter()
	router.HandleFunc("/products", createProduct).Methods("POST")
	router.HandleFunc("/products", getProducts).Methods("GET")

	http.ListenAndServe(":8080", router)

}
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var products []Product

	result, err := database.Query("SELECT * from products")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var product Product
		err := result.Scan(&product.ID, &product.Name, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}

	json.NewEncoder(w).Encode(products)
}
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Product
	json.NewDecoder(r.Body).Decode(&product)

	query, err := database.Prepare("INSERT INTO products(name,quantity) VALUES($1,$2)")
	if err != nil {
		panic(err.Error())
	}
	_, err = query.Exec(product.Name, product.Quantity)

	fmt.Fprintf(w, "New product was created")
}

func test1(w http.ResponseWriter, r *http.Request) {
	// Handles about page.
	// ... Get the path from the URL of the request.
	path := html.EscapeString(r.URL.Path)
	fmt.Fprintf(w, "Now you are on: %q", path)
}
