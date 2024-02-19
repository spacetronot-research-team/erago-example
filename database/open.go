package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

func InitializeDB() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")
	tz := os.Getenv("TZ")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, sslMode, tz,
	)

	logger := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger})
	if err != nil {
		return nil, fmt.Errorf("fail initialize db session: %v", err)
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		return nil, fmt.Errorf("err gorm db use opentelemetry tracing plugin: %v", err)
	}

	return db, err
}
