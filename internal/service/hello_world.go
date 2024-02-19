package service

import (
	"context"
	"errors"

	otel "github.com/erajayatech/go-opentelemetry"
	"github.com/spacetronot-research-team/erago-example/internal/repository"
	"github.com/spacetronot-research-team/erago-example/pkg/funcs"
)

//go:generate mockgen -source=hello_world.go -destination=mock/hello_world.go -package=mock

var (
	ErrKPbpe = errors.New("err jasdfsefs")
	ErrUyqru = errors.New("err jasdf")
)

type HelloWorld interface {
	// Bar blablabla
	Bar(ctx context.Context) error
}

type helloWorldService struct {
	helloWorldRepository repository.HelloWorld
}

func NewHelloWorldService(helloWorldRepository repository.HelloWorld) HelloWorld {
	return &helloWorldService{
		helloWorldRepository: helloWorldRepository,
	}
}

// Bar blablabla.
func (hws *helloWorldService) Bar(ctx context.Context) error {
	ctx, span := otel.NewSpan(ctx, funcs.GetMyName(), "")
	defer span.End()

	if err := hws.helloWorldRepository.Foo(ctx); err != nil {
		err = errors.Join(err, ErrKPbpe)
		otel.AddSpanError(span, err)
		otel.FailSpan(span, err.Error())
		return err
	}

	if err := hws.helloWorldRepository.Baz(ctx); err != nil {
		err = errors.Join(err, ErrUyqru)
		otel.AddSpanError(span, err)
		otel.FailSpan(span, err.Error())
		return err
	}

	return nil
}
