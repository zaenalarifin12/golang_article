package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zaenalarifin12/golang_article/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env")
	}

	//load env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	//	migrate
	db.AutoMigrate(&entity.User{}, &entity.Article{}, &entity.Comment{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB)  {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close database")
	}
	dbSql.Close()
}