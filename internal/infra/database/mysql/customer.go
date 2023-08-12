package mysql

import (
	"database/sql"

	"github.com/josimarz/fc-eda-challenge/internal/entity"
)

type CustomerGateway struct {
	db *sql.DB
}

func NewCustomerGateway(db *sql.DB) *CustomerGateway {
	return &CustomerGateway{db}
}

func (g *CustomerGateway) Create(customer *entity.Customer) error {
	stmt, err := g.db.Prepare("insert into `customer` (id, name, email, created_at, updated_at) values (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	args := []any{
		customer.Id,
		customer.Name,
		customer.Email,
		customer.CreatedAt,
		customer.UpdatedAt,
	}
	if _, err := stmt.Exec(args...); err != nil {
		return err
	}
	return nil
}

func (g *CustomerGateway) FindById(id string) (*entity.Customer, error) {
	stmt, err := g.db.Prepare("select id, name, email, created_at, updated_at from `customer` where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var customer *entity.Customer
	dest := []any{
		&customer.Id,
		&customer.Name,
		&customer.Email,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	}
	if err := stmt.QueryRow(id).Scan(dest...); err != nil {
		return nil, err
	}
	return customer, nil
}

func (g *CustomerGateway) FindAll() ([]*entity.Customer, error) {
	stmt, err := g.db.Prepare("select id, name, email, created_at, updated_at from `customer`")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []*entity.Customer
	for rows.Next() {
		var customer *entity.Customer
		dest := []any{
			&customer.Id,
			&customer.Name,
			&customer.Email,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		}
		if err := rows.Scan(dest...); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (g *CustomerGateway) Update(customer *entity.Customer) error {
	stmt, err := g.db.Prepare("update `customer` set name = ?, email = ?, created_at = ?, updated_at = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	args := []any{
		customer.Name,
		customer.Email,
		customer.CreatedAt,
		customer.UpdatedAt,
		customer.Id,
	}
	if _, err := stmt.Exec(args...); err != nil {
		return err
	}
	return nil
}

func (g *CustomerGateway) Delete(customer *entity.Customer) error {
	stmt, err := g.db.Prepare("delete from `customer` where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(customer.Id); err != nil {
		return err
	}
	return nil
}
