package main

import (
	"database/sql"
	"fmt"
	"log"
	"zalgorithm/db"
)

func main() {
	db.PrepareDB()
	notesDB, err := sql.Open("sqlite3", "db/notes.db")
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}
	defer notesDB.Close()

	rows, err := notesDB.Query("SELECT id, title, slug , markdown FROM notes")
	if err != nil {
		log.Fatalf("error querying db: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title, slug, markdown string
		err := rows.Scan(&id, &title, &slug, &markdown)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		fmt.Printf("id: %d\n", id)
		fmt.Printf("title: %s\n", title)
		fmt.Printf("slug: %s\n", slug)
		fmt.Printf("markdown: %s\n", markdown)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error scanning rows: %v", err)
	}
}
