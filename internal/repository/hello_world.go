package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockgen -source=hello_world.go -destination=mockrepository/hello_world.go -package=mockrepository

var (
	Err7647940019703865992  = errors.New("err blabla")
	Err6907370662868939236 = errors.New("err babibu")
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
		return errors.Join(err, Err7647940019703865992)
	}
	return nil
}

// Baz blablablablabla.
func (hwr *helloWorldRepository) Baz(ctx context.Context) error {
	err := gorm.ErrRecordNotFound // error from query
	if err != nil {
		return errors.Join(err, Err6907370662868939236)
	}
	return nil
}
