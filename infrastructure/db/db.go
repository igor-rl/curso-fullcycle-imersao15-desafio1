package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/igorlage/curso/fullcycle/desafio1/domain/model"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}
}

func ConnectDB(env string) *gorm.DB {
	var dsn string
	var db *gorm.DB
	var err error

	if env != "test" {
		dsn = os.Getenv("dsn")
	} else {
		dsn = os.Getenv("dsnTest")
	}

	db, err = gorm.Open("sqlite3", dsn)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	if os.Getenv("debug") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AutoMigrateDb") == "true" {
		db.AutoMigrate(&model.Product{})
	}

	return db
}
