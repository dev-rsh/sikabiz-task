package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

var (
	dbConn *gorm.DB
	dbOnce sync.Once
)

func openDBConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234567890 dbname=sikabiz port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not establish connection to DB. Err is %s", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not get db instance. Err is %s", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(15)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	return db
}

func GetDBConn() *gorm.DB {
	dbOnce.Do(func() {
		dbConn = openDBConnection()
	})

	return dbConn
}
