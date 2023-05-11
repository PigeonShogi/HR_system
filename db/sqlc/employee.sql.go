// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: employee.sql

package db

import (
	"context"
)

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO employees (
  identity_id, code, full_name, password
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, identity_id, code, full_name, password, created_at, updated_at
`

type CreateEmployeeParams struct {
	IdentityID int32  `json:"identity_id"`
	Code       string `json:"code"`
	FullName   string `json:"full_name"`
	Password   string `json:"password"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error) {
	row := q.db.QueryRowContext(ctx, createEmployee,
		arg.IdentityID,
		arg.Code,
		arg.FullName,
		arg.Password,
	)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.IdentityID,
		&i.Code,
		&i.FullName,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listEmployees = `-- name: ListEmployees :many
SELECT id, identity_id, code, full_name, password, created_at, updated_at FROM employees
ORDER BY code
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
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.IdentityID,
			&i.Code,
			&i.FullName,
			&i.Password,
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
