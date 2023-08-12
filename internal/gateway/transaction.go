package gateway

import "github.com/josimarz/fc-eda-challenge/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
