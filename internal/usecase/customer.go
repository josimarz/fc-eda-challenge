package usecase

import (
	"time"

	"github.com/josimarz/fc-eda-challenge/internal/entity"
	"github.com/josimarz/fc-eda-challenge/internal/gateway"
)

type CreateCustomerInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateCustomerOutput struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateCustomerUseCase struct {
	customerGateway gateway.CustomerGateway
}

func NewCreateCustomerUseCase(customerGateway gateway.CustomerGateway) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{customerGateway}
}

func (uc *CreateCustomerUseCase) Execute(input *CreateCustomerInput) (*CreateCustomerOutput, error) {
	customer, err := entity.NewCustomer(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	if err := uc.customerGateway.Create(customer); err != nil {
		return nil, err
	}
	return &CreateCustomerOutput{
		Id:        customer.Id,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}, nil
}

type FindCustomerInput struct {
	Id string
}

type FindCustomerOutput struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FindCustomerUseCase struct {
	customerGateway gateway.CustomerGateway
}

func NewFindCustomerUseCase(customerGateway gateway.CustomerGateway) *FindCustomerUseCase {
	return &FindCustomerUseCase{customerGateway}
}

func (uc *FindCustomerUseCase) Execute(input *FindCustomerInput) (*FindCustomerOutput, error) {
	customer, err := uc.customerGateway.FindById(input.Id)
	if err != nil {
		return nil, err
	}
	return &FindCustomerOutput{
		Id:        customer.Id,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}, nil
}

type ListCustomersInput struct{}

type CustomerOutput struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ListCustomersOutput struct {
	Customers []*CustomerOutput `json:"customers"`
}

type ListCustomersUseCase struct {
	customerGateway gateway.CustomerGateway
}

func NewListCustomersUseCase(customerGateway gateway.CustomerGateway) *ListCustomersUseCase {
	return &ListCustomersUseCase{customerGateway}
}

func (uc *ListCustomersUseCase) Execute(input *ListCustomersInput) (*ListCustomersOutput, error) {
	customers, err := uc.customerGateway.FindAll()
	if err != nil {
		return nil, err
	}
	output := &ListCustomersOutput{}
	for _, customer := range customers {
		item := &CustomerOutput{
			Id:        customer.Id,
			Name:      customer.Name,
			Email:     customer.Email,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		}
		output.Customers = append(output.Customers, item)
	}
	return output, nil
}

type UpdateCustomerInput struct {
	Id    string
	Name  string
	Email string
}

type UpdateCustomerOutput struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateCustomerUseCase struct {
	customerGateway gateway.CustomerGateway
}

func NewUpdateCustomerUseCase(customerGateway gateway.CustomerGateway) *UpdateCustomerUseCase {
	return &UpdateCustomerUseCase{customerGateway}
}

func (uc *UpdateCustomerUseCase) Execute(input *UpdateCustomerInput) (*UpdateCustomerOutput, error) {
	customer, err := uc.customerGateway.FindById(input.Id)
	if err != nil {
		return nil, err
	}
	if err := customer.Update(input.Name, input.Email); err != nil {
		return nil, err
	}
	if err := uc.customerGateway.Update(customer); err != nil {
		return nil, err
	}
	return &UpdateCustomerOutput{
		Id:        customer.Id,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}, nil
}

type DeleteCustomerInput struct {
	Id string
}

type DeleteCustomerOutput struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DeleteCustomerUseCase struct {
	customerGateway gateway.CustomerGateway
}

func NewDeleteCustomerUseCase(customerGateway gateway.CustomerGateway) *DeleteCustomerUseCase {
	return &DeleteCustomerUseCase{customerGateway}
}

func (uc *DeleteCustomerUseCase) Execute(input *DeleteCustomerInput) (*DeleteCustomerOutput, error) {
	customer, err := uc.customerGateway.FindById(input.Id)
	if err != nil {
		return nil, err
	}
	if err := uc.customerGateway.Delete(customer); err != nil {
		return nil, err
	}
	return &DeleteCustomerOutput{
		Id:        customer.Id,
		Name:      customer.Name,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}, nil
}
