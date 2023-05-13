package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var listEmployeesParams = ListEmployeesParams{
	Limit: 2, Offset: 0,
}

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	employees, err := testQueries.ListEmployees(context.Background(), listEmployeesParams)
	if err != nil {
		fmt.Println(err)
	}
	employee0 := employees[0]
	employee1 := employees[1]

	// 並行 n 個交易
	n := 5
	amount := int64(1000)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromEmployeeID: int32(employee0.ID),
				ToEmployeeID:   int32(employee1.ID),
				Amount:         amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, employee0.ID, transfer.FromEmployeeID)
		require.Equal(t, employee1.ID, transfer.ToEmployeeID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, employee0.ID, fromEntry.EmployeeID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, employee1.ID, toEntry.EmployeeID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
	}
}
