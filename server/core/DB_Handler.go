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
	Logger.Info("Initializing DB Connection...")

	db.config = &mysql.Config{}
	param := map[string]string{
		"allowNativePasswords": "true",
	}

	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	net := os.Getenv("DB_NET")
	addr := os.Getenv("DB_ADDR")

	// Set initial config from environment variables
	db.config.User = user
	db.config.Passwd = passwd
	db.config.DBName = dbName
	db.config.Net = net
	db.config.Addr = addr
	db.config.Params = param

	var err error

	// Retry connecting to MySQL (up to 10 attempts)
	for i := 0; i < 10; i++ {
		// Open the connection
		db.sql, err = sql.Open("mysql", db.config.FormatDSN())
		if err != nil {
			Logger.Error("Failed to open database connection (attempt %d): %v", i+1, err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Verify the connection
		err = db.sql.Ping()
		if err == nil {
			Logger.Info("Successfully connected to the database!")
			return
		}

		Logger.Error("Failed to ping database (attempt %d): %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	Logger.Error("Could not establish a connection to the database after multiple attempts")
	Logger.Info("Switching to fallback configuration")

	// Fallback configuration
	db.config.User = "root"
	db.config.Passwd = ""
	db.config.DBName = "ZooDaba"
	db.config.Net = "tcp"
	db.config.Addr = "localhost"
	db.config.Params = param

	// Retry connecting using the fallback configuration
	for i := 0; i < 10; i++ {
		db.sql, err = sql.Open("mysql", db.config.FormatDSN())
		if err != nil {
			Logger.Error("Failed to open database connection (attempt %d): %v", i+1, err)
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.sql.Ping()
		if err == nil {
			Logger.Info("Successfully connected to the fallback database!")
			return
		}

		Logger.Error("Failed to ping fallback database (attempt %d): %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	log.Fatal("Could not establish a connection to the fallback database. Exiting...")
}

func (db *DB_Handler) Close() {
	if db.sql != nil {
		Logger.Info("Closing DB Connection...")
		err := db.sql.Close()
		if err != nil {
			Logger.Error("Error closing DB connection: %v", err)
		} else {
			Logger.Info("Successfully closed DB Connection!")
		}
	} else {
		Logger.Warn("No active DB connection to close.")
	}
}

func (db *DB_Handler) Query(query string, args ...interface{}) *sql.Rows {
	if db.sql == nil {
		Logger.Error("DB connection is nil, cannot perform query: %v", query)
		return nil
	}
	rows, err := db.sql.Query(query, args...)
	if err != nil {
		Logger.Error("Error executing query: %v, Args: %v, SQL Error: %v", query, args, err)
		return nil
	}
	return rows
}

func (db *DB_Handler) QueryRow(query string, args ...interface{}) *sql.Row {
	if db.sql == nil {
		Logger.Error("DB connection is nil, cannot execute QueryRow: %v", query)
		return nil
	}
	return db.sql.QueryRow(query, args...)
}

func (db *DB_Handler) Exec(query string, args ...interface{}) {
	if db.sql == nil {
		Logger.Error("DB connection is nil, cannot execute query: %v", query)
		return
	}
	_, err := db.sql.Exec(query, args...)
	if err != nil {
		Logger.Error("Error executing query: %v, Args: %v, SQL Error: %v", query, args, err)
	}
}
