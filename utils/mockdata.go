// package main
package utils

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

const (
	connectionString = "postgresql://postgres:postgres@localhost:5432/iota_vision_db?sslmode=disable" // Update with your connection string
	tableName        = "posts"
	dataCount        = 10 // Number of rows to generate
)

func mock() {
	// func main() {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= dataCount; i++ {
		title := fmt.Sprintf("Title %d", i)
		body := fmt.Sprintf("Body %d", i)

		_, err := db.Exec("INSERT INTO "+tableName+" (title, body) VALUES ($1, $2)", title, body)
		if err != nil {
			log.Fatal("Error inserting data:", err)
		}
	}

	fmt.Println("Mock data insertion completed.")
}
