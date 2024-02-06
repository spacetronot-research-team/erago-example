package http

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/erago-example/internal/service"
)

var (
	Err7647940019703865992 = errors.New("err jklasjd")
)

type HelloWorldController struct {
	helloWorldService service.HelloWorld
}

func NewHelloWorldController(helloWorldService service.HelloWorld) *HelloWorldController {
	return &HelloWorldController{
		helloWorldService: helloWorldService,
	}
}

// Qux babibu.
func (hwc *HelloWorldController) Qux(ctx *gin.Context) {
	if err := hwc.helloWorldService.Bar(ctx); err != nil {
		err = errors.Join(err, Err7647940019703865992)
		log.Println(err)
		return
	}
	log.Println("^.^")
}
