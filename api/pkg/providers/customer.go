//go:generate mockgen -source=customer.go -destination=..\mocks\customer_mock.go -package=mocks
package providers

import (
	"customer-demo/pkg/models"
	"gorm.io/gorm"
)

type CustomerProvider interface {
	GetAll() ([]*models.Customer, error)
	GetById(id int) (*models.Customer, error)
}

type customerProvider struct {
	db *gorm.DB
}

func (c *customerProvider) GetAll() ([]*models.Customer, error) {
	var customers []*models.Customer
	err := c.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *customerProvider) GetById(id int) (*models.Customer, error) {
	var customer models.Customer
	err := c.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func NewCustomerProvider(db *gorm.DB) CustomerProvider {
	return &customerProvider{
		db: db,
	}
}
