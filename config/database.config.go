package config

import (
	"fmt"
	"os"

	"github.com/ydhnwb/go_checkpoint/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDBConnection() *gorm.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic("Env not detected")
	// }

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.LogEntity{})
	println("Database connected!")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
