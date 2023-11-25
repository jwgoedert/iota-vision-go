// database connection
package utils

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Substitute your DB details
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "iota_vision_db"
)

func GetConnection() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	log.Println("DB Connection established...")
	return db
}

func CreateTablesIfNotExist() {
	db := GetConnection()
	defer db.Close()

	// Create Post table for initial testing
	// Create table
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS public.posts (
		id SERIAL PRIMARY KEY,
		title TEXT COLLATE pg_catalog."default",
		body TEXT COLLATE pg_catalog."default"
	) WITH (
		OIDS = FALSE
	) TABLESPACE pg_default;
`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		// Handle error
	}

	// Alter table ownership
	alterTableQuery := `ALTER TABLE IF EXISTS public.posts OWNER TO postgres;`
	_, err = db.Exec(alterTableQuery)
	if err != nil {
		// Handle error
	}

	// Add table comment
	commentQuery := `COMMENT ON TABLE public.posts IS 'Create Posts Table for Tests';`
	_, err = db.Exec(commentQuery)
	if err != nil {
		// Handle error
	}

}
func CreatePostTable() {

}
