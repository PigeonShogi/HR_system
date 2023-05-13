// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  from_employee_id,
  to_employee_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING id, from_employee_id, to_employee_id, amount, created_at
`

type CreateTransferParams struct {
	FromEmployeeID int32 `json:"from_employee_id"`
	ToEmployeeID   int32 `json:"to_employee_id"`
	Amount         int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromEmployeeID, arg.ToEmployeeID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromEmployeeID,
		&i.ToEmployeeID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_employee_id, to_employee_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromEmployeeID,
		&i.ToEmployeeID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_employee_id, to_employee_id, amount, created_at FROM transfers
WHERE 
    from_employee_id = $1 OR
    to_employee_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransfersParams struct {
	FromEmployeeID int32 `json:"from_employee_id"`
	ToEmployeeID   int32 `json:"to_employee_id"`
	Limit          int32 `json:"limit"`
	Offset         int32 `json:"offset"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers,
		arg.FromEmployeeID,
		arg.ToEmployeeID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromEmployeeID,
			&i.ToEmployeeID,
			&i.Amount,
			&i.CreatedAt,
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
