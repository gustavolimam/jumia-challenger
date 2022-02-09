package repository

import (
	"github.com/jumia-challenger/internal/model"
)

type CustomerQueries interface {
	ListCustomers() ([]model.Customer, error)
}

// ListCustomers function to execute a query to return a list of customer saved on database
func (r *repository) ListCustomers() ([]model.Customer, error) {
	var customers []model.Customer

	query := `SELECT * FROM customer;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer model.Customer

		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Phone); err != nil {

			return customers, err
		}

		customers = append(customers, customer)
	}

	if err = rows.Err(); err != nil {
		return customers, err
	}

	return customers, nil
}
