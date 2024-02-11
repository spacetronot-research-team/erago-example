package service

import (
	"context"
	"testing"

	"github.com/spacetronot-research-team/erago-example/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_helloWorldService_Bar(t *testing.T) {
	type fields struct {
		helloWorldRepository *mock.MockHelloWorld
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		wantErr error
	}{
		{
			name: "bar err foo",
			mock: func(f fields) {
				f.helloWorldRepository.EXPECT().
					Foo(nil).Return(assert.AnError)
			},
			args: args{
				ctx: nil,
			},
			wantErr: ErrEvhqw,
		},
		{
			name: "bar err baz",
			mock: func(f fields) {
				f.helloWorldRepository.EXPECT().
					Foo(nil).Return(nil)

				f.helloWorldRepository.EXPECT().
					Baz(nil).Return(assert.AnError)
			},
			args: args{
				ctx: nil,
			},
			wantErr: ErrQgogf,
		},
		{
			name: "bar success",
			mock: func(f fields) {
				f.helloWorldRepository.EXPECT().
					Foo(nil).Return(nil)

				f.helloWorldRepository.EXPECT().
					Baz(nil).Return(nil)
			},
			args: args{
				ctx: nil,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				helloWorldRepository: mock.NewMockHelloWorld(ctrl),
			}
			tt.mock(f)

			hws := &helloWorldService{
				helloWorldRepository: f.helloWorldRepository,
			}

			err := hws.Bar(tt.args.ctx)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
