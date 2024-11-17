package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

var err error

func InitDB() {
	godotenv.Load()

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
	)

	fmt.Print("Connection :", connStr)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	tableCheck := CheckIfTableExists()
	if !tableCheck {
		createTable()
	}
}

func CheckIfTableExists() bool {
	var exists bool
	log.Println("db checker came")
	// Query to check if the table exists
	err := DB.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'tasks');").Scan(&exists)
	if err != nil {

		log.Println("Error checking if table exists: ", err)
		return false
	}

	return exists

}

func createTable() {
	// SQL query to create the task table
	createTableSQL := `CREATE TABLE tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FALSE
	);`

	fmt.Print(createTableSQL)

	_, err := DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating task table: %v", err)
	}
	log.Println("Task table created successfully!")
}
