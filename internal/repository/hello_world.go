package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

//go:generate mockgen -source=hello_world.go -destination=mock/hello_world.go -package=mock

var (
	ErrEvhqw = errors.New("err blabla")
	ErrQgogf = errors.New("err babibu")
)

type HelloWorld interface {
	Foo(ctx context.Context) error
	Baz(ctx context.Context) error
	Bar(ctx context.Context) int32
}

type helloWorldRepository struct {
	db *gorm.DB
}

func NewHelloWorldRepository(db *gorm.DB) HelloWorld {
	return &helloWorldRepository{
		db: db,
	}
}

func (hwr *helloWorldRepository) Foo(ctx context.Context) error {
	err := gorm.ErrRecordNotFound // error from query
	if err != nil {
		return errors.Join(err, ErrEvhqw)
	}
	return nil
}

func (hwr *helloWorldRepository) Baz(ctx context.Context) error {
	err := gorm.ErrRecordNotFound // error from query
	if err != nil {
		return errors.Join(err, ErrQgogf)
	}
	return nil
}

func (hwr *helloWorldRepository) Bar(ctx context.Context) int32 {
	var result int32
	if err := hwr.db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return 0
	}
	return result
}
