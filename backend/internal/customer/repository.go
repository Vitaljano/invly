package customer

import "github.com/Vitaljano/invly/backend/pkg/db"

type CustomerRepository struct {
	Database *db.Db
}

func NewCustomerRepository(db *db.Db) *CustomerRepository {
	return &CustomerRepository{
		Database: db,
	}
}

func (repo *CustomerRepository) Create(customer *Customer) (*Customer, error) {
	result := repo.Database.Create(customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}
