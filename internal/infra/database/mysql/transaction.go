package mysql

import (
	"database/sql"

	"github.com/josimarz/fc-eda-challenge/internal/entity"
)

type TransactionGateway struct {
	db *sql.DB
}

func NewTransactionGateway(db *sql.DB) *TransactionGateway {
	return &TransactionGateway{db}
}

func (g *TransactionGateway) Create(transaction *entity.Transaction) error {
	stmt, err := g.db.Prepare("insert into transaction (id, from_id, to_id, amount, created_at, updated_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	args := []any{
		transaction.Id,
		transaction.From.Id,
		transaction.To.Id,
		transaction.Amount,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	}
	if _, err := stmt.Exec(args...); err != nil {
		return err
	}
	return nil
}
