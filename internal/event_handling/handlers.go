package eventhandling

import (
	"fmt"
	"sync"

	"github.com/josimarz/fc-eda-challenge/internal/infra/kafka"
	"github.com/josimarz/fc-eda-challenge/pkg/events"
)

type TransactionCreatedHandler struct {
	producer *kafka.Producer
}

func NewTransactionCreatedHandler(producer *kafka.Producer) *TransactionCreatedHandler {
	return &TransactionCreatedHandler{producer}
}

func (h *TransactionCreatedHandler) Handle(message events.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	h.producer.Publish(message.GetPayload(), nil, "transactions")
	fmt.Println("TransactionCreatedHandler called")
}

type BalancesUpdatedHandler struct {
	producer *kafka.Producer
}

func NewBalancesUpdatedHandler(producer *kafka.Producer) *BalancesUpdatedHandler {
	return &BalancesUpdatedHandler{producer}
}

func (h *BalancesUpdatedHandler) Handle(message events.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	h.producer.Publish(message.GetPayload(), nil, "balances")
	fmt.Println("BalancesUpdatedHandler called")
}
