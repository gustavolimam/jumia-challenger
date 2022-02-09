package service

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jumia-challenger/internal/model"
)

var (
	regexCode = regexp.MustCompile(`\((\d+)\).+`)
)

func validatePhoneNumber(customer []model.Customer) []model.PhoneNumber {
	phones := []model.PhoneNumber{}

	// Goes through the entire contents of the customers array
	for _, customer := range customer {
		phone := model.PhoneNumber{}

		// validate if the string has some country code
		countryInfo := regexCode.FindStringSubmatch(customer.Phone)
		if countryInfo == nil {
			continue
		}
		phoneNumberStr := strings.Split(countryInfo[0], " ")
		phone.PhoneNumber = phoneNumberStr[1]

		// convert country code to type int
		countryCodeInt, err := strconv.Atoi(countryInfo[1])
		if err != nil {
			continue
		}
		phone.CountryCode = countryCodeInt

		// Verify if the country code is equal to some expected
		countryDetected, founded := model.PhonesList[countryCodeInt]
		if !founded {
			continue
		}
		phone.Country = countryDetected.Name

		// Run a list of countries expected and validate if the number is correctly or not
		if matches := countryDetected.Regex.FindStringSubmatch(customer.Phone); matches != nil {

			phone.PhoneNumber = matches[2]
			phone.State = true
		}

		phones = append(phones, phone)
	}

	return phones
}
