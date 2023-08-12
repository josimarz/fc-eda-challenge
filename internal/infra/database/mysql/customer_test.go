package mysql

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/josimarz/fc-eda-challenge/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerGatewayTestSuite struct {
	suite.Suite
	mock            sqlmock.Sqlmock
	db              *sql.DB
	customerGateway *CustomerGateway
}

func (suite *CustomerGatewayTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatal(err.Error())
	}
	suite.db = db
	suite.mock = mock
	suite.customerGateway = NewCustomerGateway(suite.db)
}

func (suite *CustomerGatewayTestSuite) TearDownTest() {
	suite.db.Close()
}

func (suite *CustomerGatewayTestSuite) Create() {
	customer := &entity.Customer{
		Entity: entity.Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	suite.mock.
		ExpectPrepare("insert into `customer` (id, name, email, created_at, updated_at) values (?, ?, ?, ?, ?)").
		ExpectExec().
		WithArgs(5).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := suite.customerGateway.Create(customer)
	assert.Nil(suite.T(), err)
}

func (suite *CustomerGatewayTestSuite) Create_WithPrepareError() {
	suite.mock.
		ExpectPrepare("insert into `customer` (id, name, email, created_at, updated_at) values (?, ?, ?, ?, ?)").
		WillReturnError(errors.New("error"))
	err := suite.customerGateway.Create(nil)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "error")
}

func (suite *CustomerGatewayTestSuite) Create_WithExecError() {
	customer := &entity.Customer{
		Entity: entity.Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	suite.mock.
		ExpectPrepare("insert into `customer` (id, name, email, created_at, updated_at) values (?, ?, ?, ?, ?)").
		ExpectExec().
		WillReturnError(errors.New("error"))
	err := suite.customerGateway.Create(customer)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "error")
}

func TestCustomerGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerGatewayTestSuite))
}
