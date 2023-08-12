package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionTestSuite struct {
	suite.Suite
	to   *Account
	from *Account
}

func (suite *TransactionTestSuite) SetupTest() {
	customer, _ := NewCustomer("Ana Ivanovic", "ivanovic@wta.com")
	suite.to = NewAccount(customer)

	customer, _ = NewCustomer("Maria Sharapova", "sharapova@wta.com")
	suite.from = NewAccount(customer)
	suite.from.Deposit(1999.9)
}

func (suite *TransactionTestSuite) TestNewTransaction() {
	transaction, err := NewTransaction(suite.to, suite.from, 200.8)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), transaction)
	assert.Equal(suite.T(), suite.to, transaction.To)
	assert.Equal(suite.T(), suite.from, transaction.From)
	assert.Equal(suite.T(), 200.8, transaction.Amount)
}

func (suite *TransactionTestSuite) TestNewTransaction_WithNegativeAmount() {
	transaction, err := NewTransaction(suite.to, suite.from, -100.9)
	assert.Nil(suite.T(), transaction)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to execute transaction: negative amount is not allowed")
}

func (suite *TransactionTestSuite) TestNewTransaction_WithInsufficientFunds() {
	transaction, err := NewTransaction(suite.to, suite.from, 2000.0)
	assert.Nil(suite.T(), transaction)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to execute transaction: insufficient funds")
}

func (suite *TransactionTestSuite) TestTransaction_Commit() {
	transaction, err := NewTransaction(suite.to, suite.from, 99.9)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), transaction)
	transaction.Commit()
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionTestSuite))
}
