// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: employee.sql

package db

import (
	"context"
)

const addEmployeeStock = `-- name: AddEmployeeStock :one
UPDATE employees
SET stock = stock + $1
WHERE id = $2
RETURNING id, identity_id, code, full_name, password, stock, created_at, updated_at
`

type AddEmployeeStockParams struct {
	Amount int64 `json:"amount"`
	ID     int32 `json:"id"`
}

func (q *Queries) AddEmployeeStock(ctx context.Context, arg AddEmployeeStockParams) (Employee, error) {
	row := q.db.QueryRowContext(ctx, addEmployeeStock, arg.Amount, arg.ID)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.IdentityID,
		&i.Code,
		&i.FullName,
		&i.Password,
		&i.Stock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO employees (
  identity_id, code, full_name
) VALUES (
  $1, $2, $3
)
RETURNING id, identity_id, code, full_name, password, stock, created_at, updated_at
`

type CreateEmployeeParams struct {
	IdentityID int32  `json:"identity_id"`
	Code       string `json:"code"`
	FullName   string `json:"full_name"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error) {
	row := q.db.QueryRowContext(ctx, createEmployee, arg.IdentityID, arg.Code, arg.FullName)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.IdentityID,
		&i.Code,
		&i.FullName,
		&i.Password,
		&i.Stock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteEmployeeById = `-- name: DeleteEmployeeById :exec
DELETE FROM employees
WHERE id = $1
`

func (q *Queries) DeleteEmployeeById(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEmployeeById, id)
	return err
}

const getEmployee = `-- name: GetEmployee :one
SELECT id, identity_id, code, full_name, password, stock, created_at, updated_at FROM employees
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEmployee(ctx context.Context, id int32) (Employee, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.IdentityID,
		&i.Code,
		&i.FullName,
		&i.Password,
		&i.Stock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getEmployeeForUpdate = `-- name: GetEmployeeForUpdate :one
SELECT id, identity_id, code, full_name, password, stock, created_at, updated_at FROM employees
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetEmployeeForUpdate(ctx context.Context, id int32) (Employee, error) {
	row := q.db.QueryRowContext(ctx, getEmployeeForUpdate, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.IdentityID,
		&i.Code,
		&i.FullName,
		&i.Password,
		&i.Stock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listEmployees = `-- name: ListEmployees :many
SELECT id, identity_id, code, full_name, password, stock, created_at, updated_at FROM employees
ORDER BY code ASC
LIMIT $1
OFFSET $2
`

type ListEmployeesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEmployees(ctx context.Context, arg ListEmployeesParams) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, listEmployees, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Employee{}
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.IdentityID,
			&i.Code,
			&i.FullName,
			&i.Password,
			&i.Stock,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEmployeeWithStock = `-- name: UpdateEmployeeWithStock :one
UPDATE employees
SET stock = $2
WHERE id = $1
RETURNING id, identity_id, code, full_name, password, stock, created_at, updated_at
`

type UpdateEmployeeWithStockParams struct {
	ID    int32 `json:"id"`
	Stock int64 `json:"stock"`
}

func (q *Queries) UpdateEmployeeWithStock(ctx context.Context, arg UpdateEmployeeWithStockParams) (Employee, error) {
	row := q.db.QueryRowContext(ctx, updateEmployeeWithStock, arg.ID, arg.Stock)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.IdentityID,
		&i.Code,
		&i.FullName,
		&i.Password,
		&i.Stock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
