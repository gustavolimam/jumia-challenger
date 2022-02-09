package service

import (
	"github.com/jumia-challenger/internal/model"
	"github.com/jumia-challenger/internal/repository"
	"github.com/rs/zerolog/log"
)

type Service interface {
	ReadPhoneNumbers() ([]model.PhoneNumber, error)
}

type service struct {
	repository repository.Repository
}

func New() Service {
	return &service{
		repository.New(),
	}
}

// ReadPhoneNumbers function to get a list of customer and separate the data about phone number
func (s *service) ReadPhoneNumbers() ([]model.PhoneNumber, error) {
	// Get a list of customer data
	customers, err := s.repository.ListCustomers()
	if err != nil {
		log.Error().Err(err).Msg("Error trying to get a list of customers")

		return nil, err
	}

	// Call the function to validate and separate the information of phone number
	phones := validatePhoneNumber(customers)

	return phones, nil
}
