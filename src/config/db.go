package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var Url_users string

func InitDB() {

	_ = godotenv.Load()

	url := os.Getenv("URL_SERVICE_USERS")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_APP_USER_PASSWORD")
	host := os.Getenv("DB_HOST_AUTH")
	dbname := os.Getenv("DB_NAME_AUTH")
	port := os.Getenv("DB_PORT")

	if dbUser == "" || password == "" || host == "" || dbname == "" || port == "" || url == "" {
		log.Fatal("❌ Some variables not found")
	}

	Url_users = url

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("❌ Could not connect to DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("❌ Could not ping to DB: %v", err)
	}

	fmt.Println("Successfully connected!")
	DB = db
}
