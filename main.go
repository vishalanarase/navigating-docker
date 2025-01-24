package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnection() (*sql.DB, error) {
	// Connect to MySQL database
	dsn := "root:my-secret-pw@tcp(mysql-container:3306)/my_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db, err := dbConnection()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error connecting to the database: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query the database for a list of users
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching data from the database: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Display the results
	var users []string
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, fmt.Sprintf("Error reading rows: %v", err), http.StatusInternalServerError)
			return
		}
		users = append(users, fmt.Sprintf("%d: %s", id, name))
	}

	// Respond with the list of users
	fmt.Fprintf(w, "Users: \n%s", users)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
