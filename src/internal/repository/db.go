package repository

import (
    "database/sql"
    "log"
    "sync"

    _ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

var (
    dbInstance *sql.DB
    once       sync.Once
)

// GetDBInstance returns a singleton instance of the database connection
func GetDBInstance(dataSourceName string) *sql.DB {
    once.Do(func() {
        var err error
        dbInstance, err = sql.Open("sqlite3", dataSourceName)
        if err != nil {
            log.Fatalf("Failed to open database: %v", err)
        }

        if err = dbInstance.Ping(); err != nil {
            log.Fatalf("Failed to connect to database: %v", err)
        }
    })
    return dbInstance
}

// CloseDB closes the database connection
func CloseDB() {
    if dbInstance != nil {
        err := dbInstance.Close()
        if err != nil {
            log.Printf("Error closing database: %v", err)
        }
    }
}