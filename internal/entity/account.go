package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Entity
	Customer *Customer
	Balance  float64
}

func NewAccount(customer *Customer) *Account {
	return &Account{
		Entity: Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Balance:  0,
		Customer: customer,
	}
}

func (e *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("unable to deposito: negative amount")
	}
	e.Balance += amount
	e.UpdatedAt = time.Now()
	return nil
}

func (e *Account) Withdraw(amount float64) error {
	if amount > e.Balance {
		return errors.New("unable to withdraw: insufficient funds")
	}
	e.Balance -= amount
	e.UpdatedAt = time.Now()
	return nil
}
