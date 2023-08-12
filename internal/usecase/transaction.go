package usecase

import (
	"github.com/josimarz/fc-eda-challenge/internal/entity"
	eventhandling "github.com/josimarz/fc-eda-challenge/internal/event_handling"
	"github.com/josimarz/fc-eda-challenge/internal/gateway"
	"github.com/josimarz/fc-eda-challenge/pkg/events"
)

type CreateTransactionInput struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type CreateTransactionOutput struct {
	From   AccountOutput `json:"from"`
	To     AccountOutput `json:"to"`
	Amount float64       `json:"amount"`
}

type CreateTransactionUseCase struct {
	transactionGateway gateway.TransactionGateway
	accountGateway     gateway.AccountGateway
	eventDispatcher    *events.EventDispatcher
}

func NewCreateTransactionUseCase(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	eventDispatcher *events.EventDispatcher,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{transactionGateway, accountGateway, eventDispatcher}
}

func (uc *CreateTransactionUseCase) Execute(input *CreateTransactionInput) (*CreateTransactionOutput, error) {
	from, err := uc.accountGateway.FindById(input.From)
	if err != nil {
		return nil, err
	}
	to, err := uc.accountGateway.FindById(input.To)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(to, from, input.Amount)
	if err != nil {
		return nil, err
	}
	if err := uc.transactionGateway.Create(transaction); err != nil {
		return nil, err
	}
	output := &CreateTransactionOutput{
		From: AccountOutput{
			Id:        from.Id,
			Balance:   from.Balance,
			CreatedAt: from.CreatedAt,
			UpdatedAt: from.UpdatedAt,
		},
		To: AccountOutput{
			Id:        to.Id,
			Balance:   to.Balance,
			CreatedAt: to.CreatedAt,
			UpdatedAt: to.UpdatedAt,
		},
		Amount: transaction.Amount,
	}
	event := eventhandling.NewTransactionCreatedEvent()
	event.SetPayload(output)
	uc.eventDispatcher.Dispatch(event)
	return output, nil
}
