package dbstore

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	once sync.Once
	db   *gorm.DB
)

func connectDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	timeZone := os.Getenv("TIMEZONE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s", host, port, user, pass, dbName, timeZone)

	log.Println("Connecting with DSN: ", dsn)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	return dbConn, err
}

func connect() {
	dbConn, err := connectDB()

	if err != nil {
		log.Println("Failed connecting to db", err)
		log.Fatalln("Failed connecting to db", err)
	}

	log.Println("Success connecting to db")

	db = dbConn
}

func Get() *gorm.DB {
	once.Do(func() {
		connect()
	})

	return db
}
