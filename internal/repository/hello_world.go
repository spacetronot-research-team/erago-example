package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockgen -source=hello_world.go -destination=mockrepository/hello_world.go -package=mockrepository

var (
	ErrEvhqw  = errors.New("err blabla")
	ErrQgogf = errors.New("err babibu")
)

type HelloWorld interface {
	// Foo blablabla.
	Foo(ctx context.Context) error
	// Baz blablablabla.
	Baz(ctx context.Context) error
}

type helloWorldRepository struct {
	db *gorm.DB
}

func NewHelloWorldRepository(db *gorm.DB) HelloWorld {
	return &helloWorldRepository{
		db: db,
	}
}

// Foo blablablablabla.
func (hwr *helloWorldRepository) Foo(ctx context.Context) error {
	err := gorm.ErrRecordNotFound // error from query
	if err != nil {
		return errors.Join(err, ErrEvhqw)
	}
	return nil
}

// Baz blablablablabla.
func (hwr *helloWorldRepository) Baz(ctx context.Context) error {
	err := gorm.ErrRecordNotFound // error from query
	if err != nil {
		return errors.Join(err, ErrQgogf)
	}
	return nil
}
