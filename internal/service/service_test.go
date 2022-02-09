package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/jumia-challenger/internal/model"
	"github.com/jumia-challenger/internal/repository"
	mockRepo "github.com/jumia-challenger/internal/repository/mocks"
)

func Test_service_ReadPhoneNumbers(t *testing.T) {
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

	type fields struct {
		repository repository.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.PhoneNumber
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{&mockRepo.Repository{}},
			want:    phones,
			wantErr: false,
		},
		{
			name:    "error",
			fields:  fields{&mockRepo.Repository{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repository: tt.fields.repository,
			}

			switch tt.name {
			case "success":
				tt.fields.repository.(*mockRepo.Repository).On("ListCustomers").Return(customers, nil)
			case "error":
				tt.fields.repository.(*mockRepo.Repository).On("ListCustomers").Return(nil, errors.New("errorOnTest"))
			}

			got, err := s.ReadPhoneNumbers()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ReadPhoneNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.ReadPhoneNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
