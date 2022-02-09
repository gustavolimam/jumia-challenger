package model

type Customer struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Phone string `json:"phone" db:"phone"`
}
