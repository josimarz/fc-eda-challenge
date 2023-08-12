package entity

import (
	"errors"
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Entity
	Name  string
	Email string
}

func NewCustomer(name, email string) (*Customer, error) {
	customer := &Customer{
		Entity: Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  name,
		Email: email,
	}
	if err := customer.IsValid(); err != nil {
		return nil, err
	}
	return customer, nil
}

func (e *Customer) IsValid() error {
	var errs []error
	if strings.Trim(e.Name, "") == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if _, err := mail.ParseAddress(e.Email); err != nil {
		errs = append(errs, errors.New("invalid email address"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func (e *Customer) Update(name, email string) error {
	e.Name = name
	e.Email = email
	e.UpdatedAt = time.Now()
	if err := e.IsValid(); err != nil {
		return err
	}
	return nil
}
