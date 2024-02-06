package service

import (
	"context"
	"fmt"

	"github.com/spacetronot-research-team/erago-example/internal/repository"
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
		return fmt.Errorf("err babibu: %v", err)
	}

	if err := hws.helloWorldRepository.Baz(ctx); err != nil {
		return fmt.Errorf("err zzzzzz: %v", err)
	}

	return nil
}
