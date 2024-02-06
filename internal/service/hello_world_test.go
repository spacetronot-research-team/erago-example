package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	repository "github.com/spacetronot-research-team/erago-example/internal/repository/mock"
)

func Test_helloWorldService_Bar(t *testing.T) {
	type fields struct {
		helloWorldRepository *repository.MockHelloWorld
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func(f fields)
		args    args
		wantErr bool
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				helloWorldRepository: repository.NewMockHelloWorld(ctrl),
			}
			tt.mock(f)

			hws := &helloWorldService{
				helloWorldRepository: f.helloWorldRepository,
			}

			err := hws.Bar(tt.args.ctx)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
