package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/josimarz/fc-eda-challenge/configs"
	eventhandling "github.com/josimarz/fc-eda-challenge/internal/event_handling"
	"github.com/josimarz/fc-eda-challenge/internal/gateway"
	"github.com/josimarz/fc-eda-challenge/internal/infra/database/mysql"
	"github.com/josimarz/fc-eda-challenge/internal/infra/kafka"
	"github.com/josimarz/fc-eda-challenge/internal/usecase"
	"github.com/josimarz/fc-eda-challenge/pkg/events"
)

var (
	config                   *configs.Config
	walletCoreDB             *sql.DB
	transactionsDB           *sql.DB
	accountGateway           gateway.AccountGateway
	transactionGateway       gateway.TransactionGateway
	createTransactionUseCase *usecase.CreateTransactionUseCase
	producer                 *kafka.Producer
	consumer                 *kafka.Consumer
	eventDispatcher          *events.EventDispatcher
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
	createGateways()
	createUseCases()
	startEventConsumer()
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
	eventDispatcher.Register("balances.updated", eventhandling.NewBalancesUpdatedHandler(producer))
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
		var input usecase.CreateTransactionInput
		if err := json.Unmarshal(message.Value, &input); err == nil {
			createTransactionUseCase.Execute(&input)
		}
	}
}

func createGateways() {
	accountGateway = mysql.NewAccountGateway(walletCoreDB)
	transactionGateway = mysql.NewTransactionGateway(transactionsDB)
}

func createUseCases() {
	createTransactionUseCase = usecase.NewCreateTransactionUseCase(transactionGateway, accountGateway, eventDispatcher)
}
