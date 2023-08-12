package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	name := "Josimar Zimermann"
	email := "josimarz@yahoo.com.br"
	customer, err := NewCustomer(name, email)
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, name, customer.Name)
	assert.Equal(t, email, customer.Email)
}

func TestNewCustomer_WithEmptyName(t *testing.T) {
	customer, err := NewCustomer("", "josimarz@yahoo.com.br")
	assert.Nil(t, customer)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "name is required")
}

func TestNewCustomer_WithInvalidEmail(t *testing.T) {
	customer, err := NewCustomer("Josimar Zimermann", "josimarz.yahoo.com.br")
	assert.Nil(t, customer)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid email address")
}

func TestNewCustomer_WithEmptyNameAndInvalidEmail(t *testing.T) {
	customer, err := NewCustomer("", "josimarz.yahoo.com.br")
	assert.Nil(t, customer)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "name is required\ninvalid email address")
}

func TestCustomer_Update(t *testing.T) {
	customer, err := NewCustomer("Josimar Zimermann", "josimarz@yahoo.com.br")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	name := "Ana Ivanovic"
	email := "ivanovi@wta.com"
	err = customer.Update(name, email)
	assert.Nil(t, err)
	assert.Equal(t, name, customer.Name)
	assert.Equal(t, email, customer.Email)
}

func TestCustomer_Update_WithEmptyName(t *testing.T) {
	customer, err := NewCustomer("Josimar Zimermann", "josimarz@yahoo.com.br")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	err = customer.Update("", "ivanovic@wta.com")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "name is required")
}

func TestCustomer_Update_WithInvalidEmail(t *testing.T) {
	customer, err := NewCustomer("Josimar Zimermann", "josimarz@yahoo.com.br")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	err = customer.Update("Ana Ivanovic", "ivanovic.wta.com")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid email address")
}

func TestCustomer_Update_WithEmptyNameAndInvalidEmail(t *testing.T) {
	customer, err := NewCustomer("Josimar Zimermann", "josimarz@yahoo.com.br")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	err = customer.Update("", "ivanovic.wta.com")
	assert.NotNil(t, err)
	assert.EqualError(t, err, "name is required\ninvalid email address")
}
