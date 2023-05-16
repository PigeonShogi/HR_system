package db

import (
	"context"
	"database/sql"
	"fmt"
)

// 透過 Store 執行資料庫查詢及交易的函式
type Store struct {
	db *sql.DB
	*Queries
}

// 透過 NewStore 建立新的 Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx 在資料庫交易中執行函式
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// 先開啟交易，若交易發生錯誤，提前結束函式
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 若交易沒發生錯誤，建立一個 sql query
	q := New(tx)
	// 將 sql query 傳入查詢用的函式
	err = fn(q)
	// 若調用查詢函式後發生錯誤，調用 Rollback 函式。
	if err != nil {
		// 若 Rollback 函式執行後發生錯誤，總共需回傳兩個錯誤，否則只回傳查詢函式的錯誤
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams 包含股票轉移交易所需輸入的參數
type TransferTxParams struct {
	FromEmployeeID int32 `json:"from_employee_id"`
	ToEmployeeID   int32 `json:"to_employee_id"`
	Amount         int64 `json:"amount"`
}

// TransferTxResult 為股票轉移交易的結果
type TransferTxResult struct {
	Transfer     Transfer `json:"transfer"`
	FromEmployee Employee `json:"from_employee"`
	ToEmployee   Employee `json:"to_employee"`
	FromEntry    Entry    `json:"from_entry"`
	ToEntry      Entry    `json:"to_entry"`
}

// TransferTx 用於將持股從某員工轉移至另一員工。
// 此函式會使用單筆資料庫交易建立股票轉移記錄、員工持股增減記錄、員工持股餘額。
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromEmployeeID: arg.FromEmployeeID,
			ToEmployeeID:   arg.ToEmployeeID,
			Amount:         arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			EmployeeID: arg.FromEmployeeID,
			Amount:     -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			EmployeeID: arg.ToEmployeeID,
			Amount:     arg.Amount,
		})
		if err != nil {
			return err
		}

		// get account -> update its stocks
		employee1, err := q.GetEmployeeForUpdate(ctx, arg.FromEmployeeID)
		if err != nil {
			return err
		}

		result.FromEmployee, err = q.UpdateEmployeeWithStock(ctx, UpdateEmployeeWithStockParams{
			ID:    arg.FromEmployeeID,
			Stock: employee1.Stock - arg.Amount,
		})
		if err != nil {
			return err
		}

		employee2, err := q.GetEmployeeForUpdate(ctx, arg.ToEmployeeID)
		if err != nil {
			return err
		}

		result.ToEmployee, err = q.UpdateEmployeeWithStock(ctx, UpdateEmployeeWithStockParams{
			ID:    arg.ToEmployeeID,
			Stock: employee2.Stock + arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
