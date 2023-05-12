package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEmployeeFromIdentities(t *testing.T) {
	employee, err := testQueries.GetEmployeeFromIdentities(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, employee)
	require.Equal(t, employee.Name, "employee")
}

func TestGetHrAdminFromIdentities(t *testing.T) {
	employee, err := testQueries.GetHrAdminFromIdentities(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, employee)
	require.Equal(t, employee.Name, "HR-Admin")
}
