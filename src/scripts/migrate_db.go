package main

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

func main() {
    // Open a connection to the database
    db, err := sql.Open("sqlite3", "./data/main.db")
    if err != nil {
        log.Fatalf("failed to connect to the database: %v", err)
    }
    defer db.Close()

    // Migrate the database schema
    err = migrateDatabase(db)
    if err != nil {
        log.Fatalf("failed to migrate the database: %v", err)
    }

    log.Println("Database migration completed successfully.")
}

// migrateDatabase applies the necessary schema changes to the database
func migrateDatabase(db *sql.DB) error {
    // Example migration: Create sections table
    createSectionsTable := `
    CREATE TABLE IF NOT EXISTS sections (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT
    );`

    _, err := db.Exec(createSectionsTable)
    if err != nil {
        return err
    }

    // Example migration: Create topics table
    createTopicsTable := `
    CREATE TABLE IF NOT EXISTS topics (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        section_id INTEGER,
        title TEXT NOT NULL,
        FOREIGN KEY (section_id) REFERENCES sections (id)
    );`

    _, err = db.Exec(createTopicsTable)
    if err != nil {
        return err
    }

    // Example migration: Create test_tasks table
    createTestTasksTable := `
    CREATE TABLE IF NOT EXISTS test_tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        topic_id INTEGER,
        question TEXT NOT NULL,
        answer TEXT NOT NULL,
        FOREIGN KEY (topic_id) REFERENCES topics (id)
    );`

    _, err = db.Exec(createTestTasksTable)
    if err != nil {
        return err
    }

    // Example migration: Create user_stats table
    createUserStatsTable := `
    CREATE TABLE IF NOT EXISTS user_stats (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        section_id INTEGER,
        score INTEGER,
        FOREIGN KEY (section_id) REFERENCES sections (id)
    );`

    _, err = db.Exec(createUserStatsTable)
    if err != nil {
        return err
    }

    return nil
}