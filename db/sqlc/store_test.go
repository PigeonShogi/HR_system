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

	fmt.Println(">> before:", employee0.Stock, employee1.Stock)

	// 並行 n 個交易，預設五次（次數不宜太低，否則可能偵測不到 deadlock 等漏洞）
	n := 5
	amount := int64(1000)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {

		go func() {
			ctx := context.Background()
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromEmployeeID: int32(employee0.ID),
				ToEmployeeID:   int32(employee1.ID),
				Amount:         amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

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

		// check employees
		fromEmployee := result.FromEmployee
		require.NotEmpty(t, fromEmployee)
		require.Equal(t, employee0.ID, fromEmployee.ID)

		toEmployee := result.ToEmployee
		require.NotEmpty(t, toEmployee)
		require.Equal(t, employee1.ID, toEmployee.ID)

		// check employee's stocks
		fmt.Println(">> tx:", fromEmployee.Stock, toEmployee.Stock)
		// 取得股票者交易前的股數 - 讓出股票者交易後的股數 ＝ 單筆成交量的倍數
		diff1 := employee0.Stock - fromEmployee.Stock
		// 取得股票者交易後的股數 - 讓出股票者交易前的股數 ＝ 單筆成交量的倍數
		diff2 := toEmployee.Stock - employee1.Stock
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		// 每筆測試交易的成交量固定，因此：取得股票者交易前的股數 - 讓出股票者交易後的股數 ＝ 單筆成交量的倍數
		require.True(t, diff1%amount == 0) // 1 * amount, 2 * amount, 3 * amount,... n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// 檢查最終更新的 stock
	updatedEmployee0, err := testQueries.GetEmployee(context.Background(), employee0.ID)
	require.NoError(t, err)

	updatedEmployee1, err := testQueries.GetEmployee(context.Background(), employee1.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedEmployee0.Stock, updatedEmployee1.Stock)
	require.Equal(t, employee0.Stock-int64(n)*amount, updatedEmployee0.Stock)
	require.Equal(t, employee1.Stock+int64(n)*amount, updatedEmployee1.Stock)
}
