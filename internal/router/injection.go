package router

import (
	"gorm.io/gorm"
	"github.com/spacetronot-research-team/erago-example/internal/controller/http"
	"github.com/spacetronot-research-team/erago-example/internal/repository"
	"github.com/spacetronot-research-team/erago-example/internal/service"
)

func getHelloWorldController(db *gorm.DB) *http.HelloWorldController {
	helloWorldRepository := repository.NewHelloWorldRepository(db)
	helloWorldService := service.NewHelloWorldService(helloWorldRepository)
	helloWorldController := http.NewHelloWorldController(helloWorldService)
	return helloWorldController
}
