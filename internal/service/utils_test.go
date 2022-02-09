package service

import (
	"reflect"
	"testing"

	"github.com/jumia-challenger/internal/model"
)

func Test_validatePhoneNumber(t *testing.T) {
	customers := []model.Customer{
		{
			ID:    1,
			Name:  "Test1",
			Phone: "(212) 698054317",
		},
		{
			ID:    2,
			Name:  "Test2",
			Phone: "(258) 847651504",
		},
	}

	phones := []model.PhoneNumber{
		{
			Country:     "Moroco",
			CountryCode: 212,
			PhoneNumber: "698054317",
			State:       true,
		},
		{
			Country:     "Mozambique",
			CountryCode: 258,
			PhoneNumber: "847651504",
			State:       true,
		},
	}

	type args struct {
		customer []model.Customer
	}
	tests := []struct {
		name string
		args args
		want []model.PhoneNumber
	}{
		{
			name: "success",
			args: args{customers},
			want: phones,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePhoneNumber(tt.args.customer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validatePhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
