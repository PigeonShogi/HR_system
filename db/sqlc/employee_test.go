package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// 測試能否在 employees 資料表建立新資料
func TestCreateEmployee(t *testing.T) {
	identityEmployee, _ := testQueries.GetEmployeeFromIdentities(context.Background())
	identityHrAdmin, _ := testQueries.GetHrAdminFromIdentities(context.Background())

	employeeArg := CreateEmployeeParams{
		IdentityID: identityEmployee.ID,
		// code 始於 S、姓名後加註 (seed) 以表示該記錄為種子資料
		Code:     "S2023050003",
		FullName: "林小明(seed)",
	}

	hrAdminArg := CreateEmployeeParams{
		IdentityID: identityHrAdmin.ID,
		Code:       "S2023050004",
		FullName:   "林怡君(seed)",
	}

	employee, err := testQueries.CreateEmployee(context.Background(), employeeArg)
	require.NoError(t, err)
	require.NotEmpty(t, employee)
	require.Zero(t, employee.Stock)
	require.Equal(t, employee.Code, employeeArg.Code)
	require.Equal(t, employee.IdentityID, employeeArg.IdentityID)
	require.Equal(t, employee.FullName, employeeArg.FullName)

	hrAdmin, err := testQueries.CreateEmployee(context.Background(), hrAdminArg)
	require.NoError(t, err)
	require.NotEmpty(t, hrAdmin)
	require.Zero(t, hrAdmin.Stock)
	require.Equal(t, hrAdmin.Code, hrAdminArg.Code)
	require.Equal(t, hrAdmin.IdentityID, hrAdminArg.IdentityID)
	require.Equal(t, hrAdmin.FullName, hrAdminArg.FullName)

	// 以上測試若沒有出錯，就刪除本次測試產出的 employees 記錄
	testQueries.DeleteEmployeeById(context.Background(), employee.ID)
	testQueries.DeleteEmployeeById(context.Background(), hrAdmin.ID)
}

// 測試能否找出 employees 資料表的資料
func TestListEmployees(t *testing.T) {
	args := ListEmployeesParams{
		Limit:  10,
		Offset: 0,
	}
	employees, err := testQueries.ListEmployees(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, employees)

}
