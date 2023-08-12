package usecase

import (
	"github.com/josimarz/fc-eda-challenge/internal/entity"
	"github.com/stretchr/testify/mock"
)

type MockCustomerGateway struct {
	mock.Mock
}

func (m *MockCustomerGateway) Create(customer *entity.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerGateway) FindById(id string) (*entity.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Customer), args.Error(1)
}

func (m *MockCustomerGateway) FindAll() ([]*entity.Customer, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Customer), args.Error(1)
}

func (m *MockCustomerGateway) Update(customer *entity.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerGateway) Delete(customer *entity.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}
