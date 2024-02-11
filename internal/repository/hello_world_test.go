package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewHelloWorldRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm stub database connection", err)
	}

	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want HelloWorld
	}{
		{
			name: "Initialize a new HelloWorld repository instance",
			args: args{
				db: gormDB,
			},
			want: &helloWorldRepository{
				db: gormDB,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHelloWorldRepository(tt.args.db)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_helloWorldRepository_Bar(t *testing.T) {
	expectedQuery := "SELECT 1"
	type fields struct {
		db sqlmock.Sqlmock
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		mock func(f fields)
		want int32
	}{
		{
			name: "Success to retrieve query result",
			args: args{
				ctx: context.TODO(),
			},
			mock: func(f fields) {
				f.db.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"column"}).
						AddRow(1))
			},
			want: 1,
		},
		{
			name: "Error when retrieve query result",
			args: args{
				ctx: context.TODO(),
			},
			mock: func(f fields) {
				f.db.ExpectQuery(expectedQuery).WillReturnError(assert.AnError)
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a gorm stub database connection", err)
			}

			f := fields{
				db: mock,
			}
			tt.mock(f)

			hwr := &helloWorldRepository{
				db: gormDB,
			}
			got := hwr.Bar(tt.args.ctx)
			assert.Equal(t, tt.want, got)
		})
	}
}
