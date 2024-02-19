package main

import (
	"context"

	otel "github.com/erajayatech/go-opentelemetry"
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

	ctx := context.Background()
	if err := initOpentelemetry(ctx); err != nil {
		logrus.Fatalf("err init opentelemetry: %v", err)
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

func initOpentelemetry(ctx context.Context) error {
	otelTracerService := otel.ConstructOtelTracer()
	otelTracerServiceErr := otelTracerService.SetTraceProviderNewRelic(ctx)
	if otelTracerServiceErr != nil {
		return otelTracerServiceErr
	}
	return nil
}
