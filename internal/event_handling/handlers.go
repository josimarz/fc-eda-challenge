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

func NewBalanceUpdatedHandler(producer *kafka.Producer) *TransactionCreatedHandler {
	return &TransactionCreatedHandler{producer}
}

func (h *TransactionCreatedHandler) Handle(message events.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	h.producer.Publish(message.GetPayload(), nil, "transactions")
	fmt.Println("TransactionCreatedHandler called")
}
