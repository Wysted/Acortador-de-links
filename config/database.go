package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Wysted/shortLink/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = 123123
	dbname = "slink"
)

func InitDatabase() *gorm.DB {
	Db_user := os.Getenv("DB_USER")
	Db_passowrd := os.Getenv("DB_PASSWORD")
	Db_name := os.Getenv("DB_NAME")
	
	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=localhost sslmode=disable",Db_user,Db_passowrd,Db_name)

	// Initialize the GORM database connection
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión a la base de datos exitosa.")

	if err := db.Migrations(database); err != nil {
		log.Fatal(err)
	}
	return database
}
