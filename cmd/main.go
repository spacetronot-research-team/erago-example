package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago-example/database"
	"github.com/spacetronot-research-team/erago-example/internal/router"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err)
	}

	db, err := database.InitializeDB()
	if err != nil {
		logrus.Fatal("err initialize db")
	}

	ginEngine := gin.Default()

	router.Register(ginEngine, db)

	if err := ginEngine.Run(); err != nil {
		logrus.Fatal(err)
	}
}
