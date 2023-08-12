package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AccountTestSuite struct {
	suite.Suite
	customer *Customer
}

func (suite *AccountTestSuite) SetupTest() {
	customer, _ := NewCustomer("Josimar Zimermann", "josimarz@yahoo.com.br")
	suite.customer = customer
}

func (suite *AccountTestSuite) TestNewAccount() {
	account := NewAccount(suite.customer)
	assert.NotNil(suite.T(), account)
	assert.Equal(suite.T(), suite.customer, account.Customer)
}

func (suite *AccountTestSuite) TestAccount_Deposit() {
	account := NewAccount(suite.customer)
	assert.NotNil(suite.T(), account)

	amount := 1000.99
	err := account.Deposit(amount)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), amount, account.Balance)
}

func (suite *AccountTestSuite) TestAccount_Deposit_WithNegativeAmount() {
	account := NewAccount(suite.customer)
	assert.NotNil(suite.T(), account)

	err := account.Deposit(-999.8)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to deposito: negative amount")
}

func (suite *AccountTestSuite) TestWithdraw() {
	account := NewAccount(suite.customer)
	assert.NotNil(suite.T(), account)

	err := account.Deposit(1999.99)
	assert.Nil(suite.T(), err)

	err = account.Withdraw(1599.0)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1999.99-1599.0, account.Balance)
}

func (suite *AccountTestSuite) TestWithdraw_WithInsufficientFunds() {
	account := NewAccount(suite.customer)
	assert.NotNil(suite.T(), account)

	err := account.Withdraw(10.9)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to withdraw: insufficient funds")
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
