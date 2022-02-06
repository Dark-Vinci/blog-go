package config

import (
	"fmt"
	entity "new-proj/entities"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBConnection () *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		panic("failed to loa app env")
	}

	dbServer := os.Getenv("DB_SERVER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbServer, dbPass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Book{})

	return db
}

func CloseDBConnection (db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic ("failed to close connecttion")
	}

	dbSQL.Close()
}