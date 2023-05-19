package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/PigeonShogi/HR_system/db/sqlc"
	"github.com/golang/mock/gomock"
	mockdb "github.com/pigeonshogi/HR_system/db/mock"
	"github.com/stretchr/testify/require"
)

func TestAccountAPI(t *testing.T) {
	employee := createEmployee()

	ctrl := gomock.NewController(t)
	// New in go1.14+, if you are passing a *testing.T into this function you no
	// longer need to call ctrl.Finish() in your test methods.
	// defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	// build stubs
	store.EXPECT().
		GetEmployee(gomock.Any(), gomock.Eq(employee.ID)).
		Times(1).
		Return(employee, nil)

	// 啟動測試伺服器並發送請求
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", employee.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

func createEmployee() db.Employee {
	return db.Employee{
		ID:         1,
		IdentityID: 1,
		Code:       "M001",
		FullName:   "涂小明",
		Stock:      10000,
	}
}
