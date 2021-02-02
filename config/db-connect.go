package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


var db *sql.DB

func CreateConnection() *sql.DB {
	var err error
	err = godotenv.Load(os.ExpandEnv(".env"))
	
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sql.Open(os.Getenv("DB_DRIVER"), psqlInfo)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sukses koneksi ke db...")
	return db
}