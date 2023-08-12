package eventhandling

import "time"

type TransactionCreatedEvent struct {
	name     string
	dateTime time.Time
	payload  interface{}
}

func NewTransactionCreatedEvent() *TransactionCreatedEvent {
	return &TransactionCreatedEvent{
		name:     "balance.updated",
		dateTime: time.Now(),
	}
}

func (e *TransactionCreatedEvent) GetName() string {
	return e.name
}

func (e *TransactionCreatedEvent) GetPayload() interface{} {
	return e.payload
}

func (e *TransactionCreatedEvent) SetPayload(payload interface{}) {
	e.payload = payload
}

func (e *TransactionCreatedEvent) GetDateTime() time.Time {
	return e.dateTime
}
