package service

import (
	"context"
	"errors"

	"github.com/spacetronot-research-team/erago-example/internal/repository"
)

//go:generate mockgen -source=hello_world.go -destination=mockservice/hello_world.go -package=mockservice

var (
	Err5974383538316647993 = errors.New("err jasdfsefs")
	Err2409322033459614448  = errors.New("err jasdf")
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
	if err := hws.helloWorldRepository.Foo(ctx); err != nil {
		return errors.Join(err, Err5974383538316647993)
	}

	if err := hws.helloWorldRepository.Baz(ctx); err != nil {
		return errors.Join(err, Err2409322033459614448)
	}

	return nil
}
