package databases

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Init(connectionString string) (*sql.DB, error) {
	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(3)

	log.Println("Database connected successfully")
	return db, nil
}
