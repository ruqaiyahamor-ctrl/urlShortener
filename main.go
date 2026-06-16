/* Starting point for my applciation:

1. Loads the .env file
2. Reads database details
3. Connects to PostgreSQL
4. Creates/migrates database tables
5. Sets up routes
6. Starts the server
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"urlShortener/models"
	"urlShortener/routes"
)

func main() {
	//STEP 1: Load the .env file
	godotenv.Load() 

	//STEP 2: Get database info from .env (these files connect to PostgresSQL)
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	
	//STEP 3: Build the postgres connection string (this puts the database details into a string for GORM)
	connStr := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s", host, user, password, port, name, sslmode)

	//STEP 4: Open/connect to the db using GORM
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	//STEP 5: Create/update db tables and run migration (AutoMigrate creates the users and urls tables)
	err = db.AutoMigrate(&models.User{}, &models.URL{})
	if err != nil {
		log.Fatal("Failed to auto-migrate database:, err")
	}
	//STEP 6: Set up the routes (connects endpoints)
	r := routes.Router(db)
	
	//STEP 7: Start the server (this will run on http://localhost:8080)
	fmt.Println("Server is running")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}	