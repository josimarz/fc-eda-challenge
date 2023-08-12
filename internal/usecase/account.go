package usecase

import (
	"time"

	"github.com/josimarz/fc-eda-challenge/internal/entity"
	"github.com/josimarz/fc-eda-challenge/internal/gateway"
)

type CreateAccountInput struct {
	CustomerId string
}

type CreateAccountOutput struct {
	Id        string         `json:"id"`
	Balance   float64        `json:"balance"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Customer  CustomerOutput `json:"customer"`
}

type CreateAccountUseCase struct {
	accountGateway  gateway.AccountGateway
	customerGateway gateway.CustomerGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, customerGateway gateway.CustomerGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{accountGateway, customerGateway}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInput) (*CreateAccountOutput, error) {
	customer, err := uc.customerGateway.FindById(input.CustomerId)
	if err != nil {
		return nil, err
	}
	account := entity.NewAccount(customer)
	if err := uc.accountGateway.Create(account); err != nil {
		return nil, err
	}
	return &CreateAccountOutput{
		Id:        account.Id,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		Customer: CustomerOutput{
			Id:        customer.Id,
			Name:      customer.Name,
			Email:     customer.Email,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
		},
	}, nil
}

type ListCustomerAccountsInput struct {
	CustomerId string
}

type AccountOutput struct {
	Id        string    `json:"id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ListCustomerAccountsOutput struct {
	Accounts []*AccountOutput `json:"accounts"`
}

type ListCustomerAccountsUseCase struct {
	accountGateway  gateway.AccountGateway
	customerGateway gateway.CustomerGateway
}

func NewListCustomerAccountsUseCase(accountGateway gateway.AccountGateway, customerGateway gateway.CustomerGateway) *ListCustomerAccountsUseCase {
	return &ListCustomerAccountsUseCase{accountGateway, customerGateway}
}

func (uc *ListCustomerAccountsUseCase) Execute(input ListCustomerAccountsInput) (*ListCustomerAccountsOutput, error) {
	customer, err := uc.customerGateway.FindById(input.CustomerId)
	if err != nil {
		return nil, err
	}
	accounts, err := uc.accountGateway.FindByCustomer(customer)
	if err != nil {
		return nil, err
	}
	output := &ListCustomerAccountsOutput{}
	for _, account := range accounts {
		output.Accounts = append(output.Accounts, &AccountOutput{
			Id:        account.Id,
			Balance:   account.Balance,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		})
	}
	return output, nil
}

type DepositInput struct {
	Id     string
	Amount float64 `json:"amount"`
}

type DepositOutput struct {
	Id        string         `json:"id"`
	Balance   float64        `json:"balance"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"udpatedAt"`
	Customer  CustomerOutput `json:"customer"`
}

type DepositUseCase struct {
	accountGateway gateway.AccountGateway
}

func NewDepositUseCase(accountGateway gateway.AccountGateway) *DepositUseCase {
	return &DepositUseCase{accountGateway}
}

func (uc *DepositUseCase) Execute(input *DepositInput) (*DepositOutput, error) {
	account, err := uc.accountGateway.FindById(input.Id)
	if err != nil {
		return nil, err
	}
	if err := account.Deposit(input.Amount); err != nil {
		return nil, err
	}
	if err := uc.accountGateway.Update(account); err != nil {
		return nil, err
	}
	return &DepositOutput{
		Id:        account.Id,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		Customer: CustomerOutput{
			Id:        account.Customer.Id,
			Name:      account.Customer.Name,
			Email:     account.Customer.Email,
			CreatedAt: account.Customer.CreatedAt,
			UpdatedAt: account.Customer.UpdatedAt,
		},
	}, nil
}

type WithdrawInput struct {
	Id     string
	Amount float64 `json:"amount"`
}

type WithdrawOutput struct {
	Id        string         `json:"id"`
	Balance   float64        `json:"balance"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"udpatedAt"`
	Customer  CustomerOutput `json:"customer"`
}

type WithdrawUseCase struct {
	accountGateway gateway.AccountGateway
}

func NewWithdrawUseCase(accountGateway gateway.AccountGateway) *WithdrawUseCase {
	return &WithdrawUseCase{accountGateway}
}

func (uc *WithdrawUseCase) Execute(input *WithdrawInput) (*WithdrawOutput, error) {
	account, err := uc.accountGateway.FindById(input.Id)
	if err != nil {
		return nil, err
	}
	if err := account.Withdraw(input.Amount); err != nil {
		return nil, err
	}
	if err := uc.accountGateway.Update(account); err != nil {
		return nil, err
	}
	return &WithdrawOutput{
		Id:        account.Id,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		Customer: CustomerOutput{
			Id:        account.Customer.Id,
			Name:      account.Customer.Name,
			Email:     account.Customer.Email,
			CreatedAt: account.Customer.CreatedAt,
			UpdatedAt: account.Customer.UpdatedAt,
		},
	}, nil
}
