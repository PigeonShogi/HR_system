package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePunch(t *testing.T) {
	argEmployee := ListEmployeesParams{
		Limit:  1,
		Offset: 0,
	}
	employees, err := testQueries.ListEmployees(context.Background(), argEmployee)
	require.NoError(t, err)
	require.NotEmpty(t, employees)

	status, err := testQueries.GetStatusByName(context.Background(), "工時未達標準")
	require.NoError(t, err)
	require.NotEmpty(t, status)
	require.Equal(t, "工時未達標準", status.Name)

	argPunch := CreatePunchParams{
		EmployeeID:   employees[0].ID,
		WorkingHours: 0,
		StatusID:     int16(status.ID),
	}
	punch, err := testQueries.CreatePunch(context.Background(), argPunch)
	fmt.Println(punch)
	require.NoError(t, err)
	require.NotEmpty(t, punch)
	require.Equal(t, punch.EmployeeID, employees[0].ID)
	require.Equal(t, punch.StatusID, int16(status.ID))
	require.Zero(t, punch.WorkingHours)
}
