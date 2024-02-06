package service

import (
	"context"
	"errors"

	"github.com/spacetronot-research-team/erago-example/internal/repository"
)

//go:generate mockgen -source=hello_world.go -destination=mockservice/hello_world.go -package=mockservice

var (
	Err7647940019703865992 = errors.New("err jasdfsefs")
	Err6907370662868939236  = errors.New("err jasdf")
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
		return errors.Join(err, Err7647940019703865992)
	}

	if err := hws.helloWorldRepository.Baz(ctx); err != nil {
		return errors.Join(err, Err6907370662868939236)
	}

	return nil
}
