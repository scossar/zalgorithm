package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func PrepareDB() {
	db, err := sql.Open("sqlite3", "db/notes.db")
	checkErr(err)
	defer db.Close()

	ensureTablesExist(db)
}

func ensureTablesExist(db *sql.DB) {
	createTableQuery := `
  CREATE TABLE IF NOT EXISTS notes (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      title TEXT NOT NULL,
      slug TEXT NOT NULL,
      markdown TEXT NOT NULL
  );
  `
	_, err := db.Exec(createTableQuery)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("database error: %v", err)
	}
}
