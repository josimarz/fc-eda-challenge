package mysql

import (
	"database/sql"

	"github.com/josimarz/fc-eda-challenge/internal/entity"
)

type AccountGateway struct {
	db *sql.DB
}

func NewAccountGateway(db *sql.DB) *AccountGateway {
	return &AccountGateway{db}
}

func (g *AccountGateway) Create(account *entity.Account) error {
	stmt, err := g.db.Prepare("insert into `account` (id, customer_id, balance, created_at, updated_at) values (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	args := []any{
		account.Id,
		account.Customer.Id,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	}
	if _, err := stmt.Exec(args...); err != nil {
		return err
	}
	return nil
}

func (g *AccountGateway) FindById(id string) (*entity.Account, error) {
	stmt, err := g.db.Prepare(`
		select
			a.id,
			a.balance,
			a.created_at,
			a.updated_at,
			c.id,
			c.name,
			c.email,
			c.created_at,
			c.updated_at
		from
			account a
				join customer c on (a.customer_id = c.id)
		where
			id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var customer *entity.Customer
	var account *entity.Account
	dest := []any{
		&account.Id,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
		&customer.Id,
		&customer.Name,
		&customer.Email,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	}
	if err := stmt.QueryRow(id).Scan(dest...); err != nil {
		return nil, err
	}
	account.Customer = customer
	return account, nil
}

func (g *AccountGateway) FindByCustomer(customer *entity.Customer) ([]*entity.Account, error) {
	stmt, err := g.db.Prepare("select id, balance, created_at, updated_at from `account` where customer_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(customer.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []*entity.Account
	for rows.Next() {
		var account *entity.Account
		dest := []any{
			&account.Id,
			&account.Balance,
			&account.CreatedAt,
			&account.UpdatedAt,
		}
		if err := rows.Scan(dest...); err != nil {
			return nil, err
		}
		account.Customer = customer
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (g *AccountGateway) Update(account *entity.Account) error {
	stmt, err := g.db.Prepare("update `account` set balance = ?, created_at = ?, updated_at = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	args := []any{
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
		account.Id,
	}
	if _, err := stmt.Exec(args...); err != nil {
		return err
	}
	return nil
}
