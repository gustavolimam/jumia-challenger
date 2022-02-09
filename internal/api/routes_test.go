package api

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jumia-challenger/internal/model"
	"github.com/jumia-challenger/internal/service"
	mockService "github.com/jumia-challenger/internal/service/mocks"
	"github.com/labstack/echo/v4"
)

func newEcho(path string, body io.Reader, method string) echo.Context {
	e := echo.New()
	req := httptest.NewRequest(method, "/", body)
	res := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, res)

	c.SetPath(path)

	return c
}

func TestRouter_ReadPhoneNumbers(t *testing.T) {
	path := "/phones"
	type fields struct {
		base    *echo.Echo
		service service.Service
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{echo.New(), &mockService.Service{}},
			args:    args{newEcho(path, nil, http.MethodGet)},
			wantErr: false,
		},
		{
			name:    "error",
			fields:  fields{echo.New(), &mockService.Service{}},
			args:    args{newEcho(path, nil, http.MethodGet)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Router{
				base:    tt.fields.base,
				service: tt.fields.service,
			}

			switch tt.name {
			case "success":
				tt.fields.service.(*mockService.Service).On("ReadPhoneNumbers").Return([]model.PhoneNumber{}, nil)
			case "error":
				tt.fields.service.(*mockService.Service).On("ReadPhoneNumbers").Return(nil, errors.New("errorOnTest"))
			}

			if err := r.ReadPhoneNumbers(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Router.ReadPhoneNumbers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
