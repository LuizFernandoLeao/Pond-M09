package utils

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

var db *sql.DB

// InitializeDB initializes the database connection
func InitializeDB(host, port, user, password, dbname string) {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var err error
    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    log.Println("Database connection established")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
    return db
}

// CloseDB closes the database connection
func CloseDB() {
    if db != nil {
        err := db.Close()
        if err != nil {
            log.Fatalf("Error closing database: %v", err)
        }
        log.Println("Database connection closed")
    }
}