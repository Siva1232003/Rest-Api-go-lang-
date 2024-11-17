package main

import (
	"fmt"
	"log"
	"net/http"
	"task-api/db"
	"task-api/routes"
)

func main() {
    // Initialize database connection
    db.InitDB()
	
	if db.DB == nil {
		log.Println("Database connection is nil")
	}
	
	defer db.DB.Close() //  database connection is close or not
	fmt.Print("hello")
	router := routes.RegisterRoutes()

    // Server start here
    log.Fatal(http.ListenAndServe(":8080", router))
}


