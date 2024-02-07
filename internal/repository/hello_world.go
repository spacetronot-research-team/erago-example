package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockgen -source=hello_world.go -destination=mockrepository/hello_world.go -package=mockrepository

var (
	Err5974383538316647993  = errors.New("err blabla")
	Err2409322033459614448 = errors.New("err babibu")
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
		return errors.Join(err, Err5974383538316647993)
	}
	return nil
}

// Baz blablablablabla.
func (hwr *helloWorldRepository) Baz(ctx context.Context) error {
	err := gorm.ErrRecordNotFound // error from query
	if err != nil {
		return errors.Join(err, Err2409322033459614448)
	}
	return nil
}
