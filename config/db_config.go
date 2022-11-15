package config

//dependency
import (
	/*	"fmt"
		"os"
	*/
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connect mysql database
func Connect() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	/*
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")*/

	dsn := "vqrhxopx_mahamudul:mahamudul0308412@tcp(surokkha-gov-bd-foreigner.info:2083)/vqrhxopx_mahamudul_hasan?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create conection with mysql database")
	}
	//db.AutoMigrate()
	return db
}

// for close database conection
func DisConnect(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection")
	}
	dbSQL.Close()
}
