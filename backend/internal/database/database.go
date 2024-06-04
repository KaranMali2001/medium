package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Println(err)
		return
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Println("db url is not set in .env")
		return
	}

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	sqlDB, err := Db.DB()
	if err != nil {
		log.Println("failed to connect database/sql", err)
		return
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	err = sqlDB.Ping()
	if err != nil {
		log.Println("failed to ping th db", err)
		return
	}
	Start()

}
func Start() {
	Db.AutoMigrate(&User{}, &Blog{})
	fmt.Println("database connected successfully")
}
