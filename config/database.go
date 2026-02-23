package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS students (
	id TEXT PRIMARY KEY,
	name TEXT,
	major TEXT,
	gpa REAL
)
`)
if err != nil {
	log.Fatal("Create table error:", err)
}

	return db
}