package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/josimarz/fc-eda-challenge/configs"
	eventhandling "github.com/josimarz/fc-eda-challenge/internal/event_handling"
	"github.com/josimarz/fc-eda-challenge/internal/gateway"
	"github.com/josimarz/fc-eda-challenge/internal/infra/database/mysql"
	"github.com/josimarz/fc-eda-challenge/internal/infra/kafka"
	"github.com/josimarz/fc-eda-challenge/internal/infra/webserver"
	"github.com/josimarz/fc-eda-challenge/internal/usecase"
	"github.com/josimarz/fc-eda-challenge/pkg/events"
)

var (
	config                      *configs.Config
	server                      *webserver.Server
	walletCoreDB                *sql.DB
	transactionsDB              *sql.DB
	customerGateway             gateway.CustomerGateway
	accountGateway              gateway.AccountGateway
	transactionGateway          gateway.TransactionGateway
	createCustomerUseCase       *usecase.CreateCustomerUseCase
	findCustomerUseCase         *usecase.FindCustomerUseCase
	listCustomersUseCase        *usecase.ListCustomersUseCase
	updateCustomerUseCase       *usecase.UpdateCustomerUseCase
	deleteCustomersUseCase      *usecase.DeleteCustomerUseCase
	createAccountUseCase        *usecase.CreateAccountUseCase
	listCustomerAccountsUseCase *usecase.ListCustomerAccountsUseCase
	depositUseCase              *usecase.DepositUseCase
	withdrawUseCase             *usecase.WithdrawUseCase
	showAccountBalanceUseCase   *usecase.ShowAccountBalanceUseCase
	createTransactionUseCase    *usecase.CreateTransactionUseCase
	createCustomerHandler       *webserver.CreateCustomerHandler
	findCustomerHandler         *webserver.FindCustomerHandler
	listCustomersHandler        *webserver.ListCustomersHandler
	updateCustomerHandler       *webserver.UpdateCustomerHandler
	deleteCustomerHandler       *webserver.DeleteCustomerHandler
	createAccountHandler        *webserver.CreateAccountHandler
	listCustomerAccountsHandler *webserver.ListCustomerAccountsHandler
	depositHandler              *webserver.DepositHandler
	withdrawHandler             *webserver.WithdrawHandler
	showAccountBalanceHandler   *webserver.ShowAccountBalanceHandler
	createTransactionHandler    *webserver.CreateTransactionHandler
	producer                    *kafka.Producer
	consumer                    *kafka.Consumer
	eventDispatcher             *events.EventDispatcher
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

	startEventProducer()
	go startEventConsumer()
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

func startEventProducer() {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": config.KafkaDSN,
		"group.id":          "wallet",
	}
	producer = kafka.NewProducer(&configMap)
	eventDispatcher = events.NewEventDispatcher()
	eventDispatcher.Register("balance.updated", eventhandling.NewBalanceUpdatedHandler(producer))
}

func startEventConsumer() {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": config.KafkaDSN,
		"group.id":          "wallet",
	}
	consumer = kafka.NewConsumer(&configMap, []string{"transactions"})
	ch := make(chan *ckafka.Message)
	go consumer.Consume(ch)
	for {
		message := <-ch
		output := usecase.CreateTransactionOutput{}
		if err := json.Unmarshal(message.Value, &output); err == nil {
			depositInput := &usecase.DepositInput{
				Id:     output.To.Id,
				Amount: output.Amount,
			}
			if _, err := depositUseCase.Execute(depositInput); err != nil {
				continue
			}
			withdrawInput := &usecase.WithdrawInput{
				Id:     output.From.Id,
				Amount: output.Amount,
			}
			if _, err := withdrawUseCase.Execute(withdrawInput); err != nil {
				continue
			}
		}
	}
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
	createAccountUseCase = usecase.NewCreateAccountUseCase(accountGateway, customerGateway)
	listCustomerAccountsUseCase = usecase.NewListCustomerAccountsUseCase(accountGateway, customerGateway)
	depositUseCase = usecase.NewDepositUseCase(accountGateway)
	withdrawUseCase = usecase.NewWithdrawUseCase(accountGateway)
	showAccountBalanceUseCase = usecase.NewShowAccountBalanceUseCase(accountGateway)
	createTransactionUseCase = usecase.NewCreateTransactionUseCase(transactionGateway, accountGateway, eventDispatcher)
}

func createHandlers() {
	createCustomerHandler = webserver.NewCreateCustomerHandler(createCustomerUseCase)
	findCustomerHandler = webserver.NewFindCustomerHandler(findCustomerUseCase)
	listCustomersHandler = webserver.NewListCustomersHandler(listCustomersUseCase)
	updateCustomerHandler = webserver.NewUpdateCustomerHandler(updateCustomerUseCase)
	deleteCustomerHandler = webserver.NewDeleteCustomerHandler(deleteCustomersUseCase)
	createAccountHandler = webserver.NewCreateAccountHandler(createAccountUseCase)
	listCustomerAccountsHandler = webserver.NewListCustomerAccountsHandler(listCustomerAccountsUseCase)
	depositHandler = webserver.NewDepositHandler(depositUseCase)
	withdrawHandler = webserver.NewWithdrawHandler(withdrawUseCase)
	showAccountBalanceHandler = webserver.NewShowAccountBalanceHandler(showAccountBalanceUseCase)
	createTransactionHandler = webserver.NewCreateTransactionHandler(createTransactionUseCase)
}

func startServer() error {
	server = webserver.NewServer(config.Port)
	server.AddHandler(createCustomerHandler)
	server.AddHandler(findCustomerHandler)
	server.AddHandler(listCustomersHandler)
	server.AddHandler(updateCustomerHandler)
	server.AddHandler(deleteCustomerHandler)
	server.AddHandler(createAccountHandler)
	server.AddHandler(listCustomerAccountsHandler)
	server.AddHandler(depositHandler)
	server.AddHandler(withdrawHandler)
	server.AddHandler(showAccountBalanceHandler)
	server.AddHandler(createTransactionHandler)

	ch := make(chan error)
	go func() {
		fmt.Printf("[Web Server] Starting on port %s\n", config.Port)
		if err := server.Start(); err != nil {
			ch <- err
		}
	}()
	return <-ch
}
