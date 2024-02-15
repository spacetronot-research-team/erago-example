package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ginEngine *gin.Engine, db *gorm.DB) {
	helloWorldController := getHelloWorldController(db)

	ginEngine.GET("", helloWorldController.Qux)
}
