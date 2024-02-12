package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/erago-example/internal/service/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHelloWorldController_Qux(t *testing.T) {
	type fields struct {
		helloWorldService *mock.MockHelloWorld
	}
	type args struct {
		// Add structs, dtos, or anything
	}
	tests := []struct {
		name     string
		mock     func(f fields)
		args     args
		wantCode int
	}{
		{
			name: "OK",
			mock: func(f fields) {
				f.helloWorldService.EXPECT().
					Bar(gomock.Any()).
					Return(nil)
			},
			wantCode: http.StatusOK,
		},
		{
			name: "Internal Server Error",
			mock: func(f fields) {
				f.helloWorldService.EXPECT().
					Bar(gomock.Any()).
					Return(assert.AnError)
			},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockHelloWorldService := mock.NewMockHelloWorld(ctrl)

			c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

			f := fields{
				helloWorldService: mockHelloWorldService,
			}
			tt.mock(f)

			hwc := &HelloWorldController{
				helloWorldService: f.helloWorldService,
			}
			hwc.Qux(c)

			assert.Equal(t, tt.wantCode, rr.Code)
		})
	}
}
