package core

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DB_Handler struct {
	config *mysql.Config
	sql    *sql.DB
}

func (db *DB_Handler) Init() {

	db.config.User = os.Getenv("DB_USER")
	db.config.Passwd = os.Getenv("DB_PASSWORD")
	db.config.DBName = os.Getenv("DB_NAME")
	db.config.Net = os.Getenv("DB_NET")
	db.config.Addr = os.Getenv("DB_ADDR")
	db.config.Params = map[string]string{
		"allowNativePasswords": "true",
	}

	var err error
	// Retry connecting to MySQL (up to 10 attempts)
	for i := 0; i < 10; i++ {
		db.sql, err = sql.Open("mysql", db.config.FormatDSN())
		if err != nil {
			log.Printf("Failed to open database connection (attempt %d): %v", i+1, err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Verify the connection with a Ping
		err = db.sql.Ping()
		if err == nil {
			log.Println("Successfully connected to the database!")
			return
		}
		log.Printf("Failed to ping database (attempt %d): %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	log.Fatal("Could not establish a connection to the database after multiple attempts")
}

func (db *DB_Handler) Close() {
	db.sql.Close()
}

func (db *DB_Handler) Query(query string, args ...interface{}) *sql.Rows {
	rows, err := db.sql.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func (db *DB_Handler) QueryRow(query string, args ...interface{}) *sql.Row {
	row := db.sql.QueryRow(query, args...)
	return row
}

func (db *DB_Handler) Exec(query string, args ...interface{}) {
	_, err := db.sql.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB_Handler) Prepare(query string) *sql.Stmt {
	stmt, err := db.sql.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	return stmt
}
