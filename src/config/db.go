package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB
var Url_users string

func InitDB() {

	url := Env.URLServiceUsers
	dbUser := Env.DBUser
	password := Env.DBAppUserPassword
	host := Env.DBHostAuth
	dbname := Env.DBNameAuth
	port := Env.DBPort

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
