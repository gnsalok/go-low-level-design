package main

import (
	"database/sql"
	"ddd-go/internal/domain/repositories"
	"ddd-go/internal/infrastructure/database"
	"fmt"
	"log"

	"github.com/gorilla/mux"
)

func main() {

	var userRepo repositories.UserRepository

	// Example: Use in-memory implementation
	useInMemory := true
	if useInMemory {
		userRepo = database.NewUserRepositoryMemory()
	} else {
		// Example: Use SQL implementation
		db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
		if err != nil {
			fmt.Println("Failed to connect to database:", err)
			return
		}
		userRepo = database.NewUserRepositorySQL(db)
	}

	// Use `userRepo` in your application
	fmt.Println("User repository initialized:", userRepo)

	router := mux.NewRouter()

	// Start the server
	log.Println("Starting server on :8080")
	if err := router.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
