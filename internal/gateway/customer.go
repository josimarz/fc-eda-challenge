package gateway

import "github.com/josimarz/fc-eda-challenge/internal/entity"

type CustomerGateway interface {
	Create(customer *entity.Customer) error
	FindById(id string) (*entity.Customer, error)
	FindAll() ([]*entity.Customer, error)
	Update(customer *entity.Customer) error
	Delete(customer *entity.Customer) error
}
