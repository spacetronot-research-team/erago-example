package http

import (
	"errors"
	"net/http"

	otel "github.com/erajayatech/go-opentelemetry"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spacetronot-research-team/erago-example/internal/service"
)

var (
	ErrKPbpe = errors.New("[erago-example@CmeTc] err jklasjd")
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
func (hwc *HelloWorldController) Qux(c *gin.Context) {
	ctx, span := otel.Start(c)
	defer span.End()

	if err := hwc.helloWorldService.Bar(ctx); err != nil {
		err = errors.Join(err, ErrKPbpe)

		otel.AddSpanError(span, err)
		otel.FailSpan(span, err.Error())

		logrus.Info(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  "success qux",
		"error": nil,
	})
}
