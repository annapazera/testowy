package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fmt.Println("Hello Docker Tutorial")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello from Docker")
	// })

	// fmt.Println("Listening on :8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

	connStr := fmt.Sprintf("user=postgresUser password=123456 dbname=test sslmode=disable")

	db, err := sql.Open("postgres", connStr)

	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Succesfully connected")
	defer db.Close()
	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")

	// Create table.
	_, err = db.Exec("CREATE TABLE products (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table")

	fmt.Println("Inserting values")

	// Insert some data into table.
	sqlStatement := "INSERT INTO products (name, quantity) VALUES ($1, $2);"
	_, err = db.Exec(sqlStatement, "banana", 150)
	checkError(err)
	_, err = db.Exec(sqlStatement, "orange", 154)
	checkError(err)
	_, err = db.Exec(sqlStatement, "apple", 100)
	checkError(err)
	fmt.Println("Inserted 3 rows of data")

	// Read rows from table.
	var id int
	var name string
	var quantity int

	sqlStatement = "SELECT * from products;"
	rows, err := db.Query(sqlStatement)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &quantity); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
		default:
			checkError(err)
		}
	}

}
