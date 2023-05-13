package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// 測試能否從 identities 資料表取得 name 欄位為 employee 的資料
func TestGetEmployeeFromIdentities(t *testing.T) {
	employee, err := testQueries.GetEmployeeFromIdentities(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, employee)
	require.Equal(t, employee.Name, "employee")
}

// 測試能否從 identities 資料表取得 name 欄位為 HR-Admin 的資料
func TestGetHrAdminFromIdentities(t *testing.T) {
	employee, err := testQueries.GetHrAdminFromIdentities(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, employee)
	require.Equal(t, employee.Name, "HR-Admin")
}
