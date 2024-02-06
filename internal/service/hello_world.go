package service

import (
	"context"
	"errors"

	"github.com/spacetronot-research-team/erago-example/internal/repository"
)

var (
	Err7624329967302134303 = errors.New("err jasdfsefs")
	Err7969868174642337230  = errors.New("err jasdf")
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
		return errors.Join(err, Err7624329967302134303)
	}

	if err := hws.helloWorldRepository.Baz(ctx); err != nil {
		return errors.Join(err, Err7969868174642337230)
	}

	return nil
}
