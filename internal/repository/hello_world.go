package repository

import (
	"context"
	"errors"

	otel "github.com/erajayatech/go-opentelemetry"
	"github.com/spacetronot-research-team/erago-example/pkg/funcs"
	"gorm.io/gorm"
)

//go:generate mockgen -source=hello_world.go -destination=mock/hello_world.go -package=mock

var (
	ErrKPbpe = errors.New("err blabla")
	ErrUyqru = errors.New("err babibu")
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
	ctx, span := otel.NewSpan(ctx, funcs.GetMyName(), "")
	defer span.End()

	var quuz int32
	err := hwr.db.WithContext(ctx).Raw("SELECT 1").Scan(&quuz).Error
	if err != nil {
		return errors.Join(err, ErrKPbpe)
	}
	return nil
}

// Baz blablablablabla.
func (hwr *helloWorldRepository) Baz(ctx context.Context) error {
	_, span := otel.NewSpan(ctx, funcs.GetMyName(), "")
	defer span.End()

	err := gorm.ErrRecordNotFound // error from query
	if err != nil {
		return errors.Join(err, ErrUyqru)
	}
	return nil
}
