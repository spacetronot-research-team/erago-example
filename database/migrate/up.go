package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spacetronot-research-team/erago-example/database"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	db, err := database.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := migrationUp(db); err != nil {
		log.Fatal(err)
	}

	log.Println("migrate up success")
}

func migrationUp(db *gorm.DB) error {
	migrate.SetTable("migrations")

	sql, err := db.DB()
	if err != nil {
		return fmt.Errorf("'*gorm.DB' fail return '*sql.DB': %v", err)
	}

	_, err = migrate.Exec(sql, "postgres", getFileMigrationSource(), migrate.Up)
	if err != nil {
		return fmt.Errorf("fail execute migrations: %v", err)
	}

	return nil
}

func getFileMigrationSource() *migrate.FileMigrationSource {
	migrations := &migrate.FileMigrationSource{
		Dir: filepath.Join("database", "schema_migration"),
	}
	return migrations
}
