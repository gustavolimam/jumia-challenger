package model

import "regexp"

// PhoneNumber struct used to treat the phone data
type PhoneNumber struct {
	Country     string `json:"country"`
	CountryCode int    `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
	State       bool   `json:"state"`
}

//CountryDictionary structure binding the country name and regex
type CountryDictionary struct {
	Name  string
	Regex *regexp.Regexp
}

type CountriesDictionary map[int]CountryDictionary

var (
	PhonesList = CountriesDictionary{
		237: CountryDictionary{"Cameroon", regexp.MustCompile(`\((237)\)\ ?([2368]\d{7,8}$)`)},
		251: CountryDictionary{"Ethiopia", regexp.MustCompile(`\((251)\)\ ?([1-59]\d{8}$)`)},
		212: CountryDictionary{"Moroco", regexp.MustCompile(`\((212)\)\ ?([5-9]\d{8}$)`)},
		258: CountryDictionary{"Mozambique", regexp.MustCompile(`\((258)\)\ ?([28]\d{7,8}$)`)},
		256: CountryDictionary{"Uganda", regexp.MustCompile(`\((256)\)\ ?(\d{9}$)`)},
	}
)
