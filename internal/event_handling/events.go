package eventhandling

import "time"

type TransactionCreatedEvent struct {
	name     string
	dateTime time.Time
	payload  interface{}
}

func NewTransactionCreatedEvent() *TransactionCreatedEvent {
	return &TransactionCreatedEvent{
		name:     "transaction.created",
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

type BalancesUpdatedEvent struct {
	name     string
	dateTime time.Time
	payload  interface{}
}

func NewBalancesUpdatedEvent() *BalancesUpdatedEvent {
	return &BalancesUpdatedEvent{
		name:     "balances.updated",
		dateTime: time.Now(),
	}
}

func (e *BalancesUpdatedEvent) GetName() string {
	return e.name
}

func (e *BalancesUpdatedEvent) GetPayload() interface{} {
	return e.payload
}

func (e *BalancesUpdatedEvent) SetPayload(payload interface{}) {
	e.payload = payload
}

func (e *BalancesUpdatedEvent) GetDateTime() time.Time {
	return e.dateTime
}
