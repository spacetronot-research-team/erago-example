package http

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/erago-example/internal/service"
)

var (
	Err5974383538316647993 = errors.New("err jklasjd")
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
		err = errors.Join(err, Err5974383538316647993)
		log.Println(err)
		return
	}
	log.Println("^.^")
}
