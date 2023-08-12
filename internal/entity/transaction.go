package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Entity
	To     *Account
	From   *Account
	Amount float64
}

func NewTransaction(to, from *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		Entity: Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		To:     to,
		From:   from,
		Amount: amount,
	}
	if err := transaction.IsValid(); err != nil {
		return nil, err
	}
	transaction.Commit()
	return transaction, nil
}

func (e *Transaction) IsValid() error {
	if e.Amount < 0 {
		return errors.New("unable to execute transaction: negative amount is not allowed")
	}
	if e.Amount > e.From.Balance {
		return errors.New("unable to execute transaction: insufficient funds")
	}
	return nil
}

func (e *Transaction) Commit() {
	e.To.Deposit(e.Amount)
	e.From.Withdraw(e.Amount)
}
