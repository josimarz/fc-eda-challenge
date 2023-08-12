package gateway

import "github.com/josimarz/fc-eda-challenge/internal/entity"

type AccountGateway interface {
	Create(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
	FindByCustomer(customer *entity.Customer) ([]*entity.Account, error)
	Update(account *entity.Account) error
}
