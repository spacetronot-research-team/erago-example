package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spacetronot-research-team/erago-example/database"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	db, err := database.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
}
