package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josimarz/fc-eda-challenge/configs"
	"github.com/josimarz/fc-eda-challenge/internal/gateway"
	"github.com/josimarz/fc-eda-challenge/internal/infra/database/mysql"
	"github.com/josimarz/fc-eda-challenge/internal/infra/webserver"
	"github.com/josimarz/fc-eda-challenge/internal/usecase"
)

var (
	config                 *configs.Config
	server                 *webserver.Server
	walletCoreDB           *sql.DB
	transactionsDB         *sql.DB
	customerGateway        gateway.CustomerGateway
	accountGateway         gateway.AccountGateway
	transactionGateway     gateway.TransactionGateway
	createCustomerUseCase  *usecase.CreateCustomerUseCase
	findCustomerUseCase    *usecase.FindCustomerUseCase
	listCustomersUseCase   *usecase.ListCustomersUseCase
	updateCustomerUseCase  *usecase.UpdateCustomerUseCase
	deleteCustomersUseCase *usecase.DeleteCustomerUseCase
	createCustomerHandler  *webserver.CreateCustomerHandler
	findCustomerHandler    *webserver.FindCustomerHandler
	listCustomersHandler   *webserver.ListCustomersHandler
	updateCustomerHandler  *webserver.UpdateCustomerHandler
	deleteCustomerHandler  *webserver.DeleteCustomerHandler
)

func main() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = openWalletCoreDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = openTransactionsDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	createGateways()
	createUseCases()
	createHandlers()

	if err := startServer(); err != nil {
		log.Fatal(err.Error())
	}
}

func loadConfig() (err error) {
	config, err = configs.LoadConfig(".")
	return err
}

func openWalletCoreDB() (err error) {
	walletCoreDB, err = sql.Open("mysql", config.WalletCoreDSN)
	return err
}

func openTransactionsDB() (err error) {
	transactionsDB, err = sql.Open("mysql", config.TransactionsDSN)
	return err
}

func createGateways() {
	customerGateway = mysql.NewCustomerGateway(walletCoreDB)
	accountGateway = mysql.NewAccountGateway(walletCoreDB)
	transactionGateway = mysql.NewTransactionGateway(transactionsDB)
}

func createUseCases() {
	createCustomerUseCase = usecase.NewCreateCustomerUseCase(customerGateway)
	findCustomerUseCase = usecase.NewFindCustomerUseCase(customerGateway)
	listCustomersUseCase = usecase.NewListCustomersUseCase(customerGateway)
	updateCustomerUseCase = usecase.NewUpdateCustomerUseCase(customerGateway)
	deleteCustomersUseCase = usecase.NewDeleteCustomerUseCase(customerGateway)
}

func createHandlers() {
	createCustomerHandler = webserver.NewCreateCustomerHandler(createCustomerUseCase)
	findCustomerHandler = webserver.NewFindCustomerHandler(findCustomerUseCase)
	listCustomersHandler = webserver.NewListCustomersHandler(listCustomersUseCase)
	updateCustomerHandler = webserver.NewUpdateCustomerHandler(updateCustomerUseCase)
	deleteCustomerHandler = webserver.NewDeleteCustomerHandler(deleteCustomersUseCase)
}

func startServer() error {
	server = webserver.NewServer(config.Port)
	server.AddHandler(createCustomerHandler)
	server.AddHandler(findCustomerHandler)
	server.AddHandler(listCustomersHandler)
	server.AddHandler(updateCustomerHandler)
	server.AddHandler(deleteCustomerHandler)

	ch := make(chan error)
	go func() {
		fmt.Printf("[Web Server] Starting on port %s\n", config.Port)
		if err := server.Start(); err != nil {
			ch <- err
		}
	}()
	return <-ch
}
