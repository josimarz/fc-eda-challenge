package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/josimarz/fc-eda-challenge/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CustomerTestSuite struct {
	suite.Suite
	mockCustomerGateway   *MockCustomerGateway
	createCustomerUseCase *CreateCustomerUseCase
	findCustomerUseCase   *FindCustomerUseCase
	listCustomersUseCase  *ListCustomersUseCase
	updateCustomerUseCase *UpdateCustomerUseCase
	deleteCustomerUseCase *DeleteCustomerUseCase
}

func (suite *CustomerTestSuite) SetupTest() {
	suite.mockCustomerGateway = &MockCustomerGateway{}
	suite.createCustomerUseCase = NewCreateCustomerUseCase(suite.mockCustomerGateway)
	suite.findCustomerUseCase = NewFindCustomerUseCase(suite.mockCustomerGateway)
	suite.listCustomersUseCase = NewListCustomersUseCase(suite.mockCustomerGateway)
	suite.updateCustomerUseCase = NewUpdateCustomerUseCase(suite.mockCustomerGateway)
	suite.deleteCustomerUseCase = NewDeleteCustomerUseCase(suite.mockCustomerGateway)
}

func (suite *CustomerTestSuite) TestCreateCustomerUseCase_Execute() {
	name := "Josimar Zimermann"
	email := "josimarz@yahoo.com.br"

	suite.mockCustomerGateway.On("Create", mock.Anything).Return(nil)
	input := &CreateCustomerInput{
		Name:  name,
		Email: email,
	}
	output, err := suite.createCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)
	assert.Equal(suite.T(), name, output.Name)
	assert.Equal(suite.T(), email, output.Email)
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "Create", 1)
}

func (suite *CustomerTestSuite) TestCreateCustomerUseCase_Execute_WithInvalidCustomerName() {
	input := &CreateCustomerInput{
		Name:  "",
		Email: "josimarz@yahoo.com.br",
	}
	output, err := suite.createCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "name is required")
}

func (suite *CustomerTestSuite) TestCreateCustomerUseCase_Execute_WithGatewayError() {
	suite.mockCustomerGateway.On("Create", mock.Anything).Return(errors.New("unable to create customer"))

	input := &CreateCustomerInput{
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	output, err := suite.createCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to create customer")
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "Create", 1)
}

func (suite *CustomerTestSuite) TestFindCustomerUseCase_Execute() {
	customer := &entity.Customer{
		Entity: entity.Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	suite.mockCustomerGateway.On("FindById", mock.Anything).Return(customer, nil)

	input := &FindCustomerInput{Id: ""}
	output, err := suite.findCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)
	assert.Equal(suite.T(), customer.Id, output.Id)
	assert.Equal(suite.T(), customer.Name, output.Name)
	assert.Equal(suite.T(), customer.Email, output.Email)
	assert.Equal(suite.T(), customer.CreatedAt, output.CreatedAt)
	assert.Equal(suite.T(), customer.UpdatedAt, output.UpdatedAt)
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindById", 1)
}

func (suite *CustomerTestSuite) TestFindCustomerUseCase_Execute_WithGatewayError() {
	suite.mockCustomerGateway.On("FindById", mock.Anything).Return(&entity.Customer{}, errors.New("unable to find customer"))
	input := &FindCustomerInput{Id: ""}
	output, err := suite.findCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to find customer")
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindById", 1)
}

func (suite *CustomerTestSuite) TestListCustomersUseCase_Execute() {
	customers := []*entity.Customer{
		{
			Entity: entity.Entity{
				Id:        uuid.NewString(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:  "Gustavo Kuerten",
			Email: "guga@itf.com",
		},
		{
			Entity: entity.Entity{
				Id:        uuid.NewString(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:  "Roger Federer",
			Email: "federer@itf.com",
		},
		{
			Entity: entity.Entity{
				Id:        uuid.NewString(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:  "Rafael Nadal",
			Email: "nadal@itf.com",
		},
	}
	suite.mockCustomerGateway.On("FindAll", mock.Anything).Return(customers, nil)
	output, err := suite.listCustomersUseCase.Execute(&ListCustomersInput{})

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)
	assert.Len(suite.T(), output.Customers, 3)
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindAll", 1)
}

func (suite *CustomerTestSuite) TestListCustomersUseCase_Execute_WithGatewayError() {
	suite.mockCustomerGateway.On("FindAll", mock.Anything).Return([]*entity.Customer{}, errors.New("unable to find customers"))
	output, err := suite.listCustomersUseCase.Execute(&ListCustomersInput{})

	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to find customers")
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindAll", 1)
}

func (suite *CustomerTestSuite) TestUpdateCustomerUseCase_Execute() {
	customer := &entity.Customer{
		Entity: entity.Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	suite.mockCustomerGateway.On("FindById", mock.Anything).Return(customer, nil)
	suite.mockCustomerGateway.On("Update", mock.Anything).Return(nil)
	input := &UpdateCustomerInput{
		Id:    customer.Id,
		Name:  "Ana Ivanovic",
		Email: "ivanovic@wta.com",
	}
	output, err := suite.updateCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)
	assert.Equal(suite.T(), customer.Id, output.Id)
	assert.Equal(suite.T(), customer.Name, output.Name)
	assert.Equal(suite.T(), customer.Email, output.Email)
	assert.Equal(suite.T(), customer.CreatedAt, output.CreatedAt)
	assert.Equal(suite.T(), customer.UpdatedAt, output.UpdatedAt)
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindById", 1)
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "Update", 1)
}

func (suite *CustomerTestSuite) TestUpdateCustomerUseCase_Execute_WithFindError() {
	suite.mockCustomerGateway.On("FindById", mock.Anything).Return(&entity.Customer{}, errors.New("unable to find customer"))
	input := &UpdateCustomerInput{
		Id:    "",
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	output, err := suite.updateCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to find customer")
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindById", 1)
}

func (suite *CustomerTestSuite) TestUpdateCustomerUseCase_Execute_WithEmptyName() {
	customer := &entity.Customer{
		Entity: entity.Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	suite.mockCustomerGateway.On("FindById", mock.Anything).Return(customer, nil)
	input := &UpdateCustomerInput{
		Id:    customer.Id,
		Name:  "",
		Email: "josimarz@yahoo.com.br",
	}
	output, err := suite.updateCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "name is required")
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindById", 1)
}

func (suite *CustomerTestSuite) TestUpdateCustomerUseCase_Execute_WithUpdateError() {
	customer := &entity.Customer{
		Entity: entity.Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	suite.mockCustomerGateway.On("FindById", mock.Anything).Return(customer, nil)
	suite.mockCustomerGateway.On("Update", mock.Anything).Return(errors.New("unable to update customer"))
	input := &UpdateCustomerInput{
		Id:    customer.Id,
		Name:  "Ana Ivanovic",
		Email: "ivanovic@wta.com",
	}
	output, err := suite.updateCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to update customer")
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindById", 1)
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "Update", 1)
}

func (suite *CustomerTestSuite) TestDeleteCustomerUseCase_Execute() {
	customer := &entity.Customer{
		Entity: entity.Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	suite.mockCustomerGateway.On("FindById", mock.Anything).Return(customer, nil)
	suite.mockCustomerGateway.On("Delete", mock.Anything).Return(nil)
	input := &DeleteCustomerInput{Id: customer.Id}
	output, err := suite.deleteCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), output)
	assert.Equal(suite.T(), customer.Id, output.Id)
	assert.Equal(suite.T(), customer.Name, output.Name)
	assert.Equal(suite.T(), customer.Email, output.Email)
	assert.Equal(suite.T(), customer.CreatedAt, output.CreatedAt)
	assert.Equal(suite.T(), customer.UpdatedAt, output.UpdatedAt)
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindById", 1)
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "Delete", 1)
}

func (suite *CustomerTestSuite) TestDeleteCustomerUseCase_Execute_WithFindError() {
	suite.mockCustomerGateway.On("FindById", mock.Anything).Return(&entity.Customer{}, errors.New("unable to find customer"))
	input := &DeleteCustomerInput{Id: ""}
	output, err := suite.deleteCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to find customer")
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindById", 1)
}

func (suite *CustomerTestSuite) TestDeleteCustomerUseCase_Execute_WithDeleteError() {
	customer := &entity.Customer{
		Entity: entity.Entity{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Josimar Zimermann",
		Email: "josimarz@yahoo.com.br",
	}
	suite.mockCustomerGateway.On("FindById", mock.Anything).Return(customer, nil)
	suite.mockCustomerGateway.On("Delete", mock.Anything).Return(errors.New("unable to delete customer"))
	input := &DeleteCustomerInput{Id: customer.Id}
	output, err := suite.deleteCustomerUseCase.Execute(input)

	assert.Nil(suite.T(), output)
	assert.NotNil(suite.T(), err)
	assert.EqualError(suite.T(), err, "unable to delete customer")
	suite.mockCustomerGateway.AssertExpectations(suite.T())
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "FindById", 1)
	suite.mockCustomerGateway.AssertNumberOfCalls(suite.T(), "Delete", 1)
}

func TestCustomerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerTestSuite))
}
