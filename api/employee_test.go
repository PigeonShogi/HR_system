package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/PigeonShogi/HR_system/db/mock"
	db "github.com/PigeonShogi/HR_system/db/sqlc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	employee := createEmployee()

	testCases := []struct {
		name          string
		employeeID    int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "OK",
			employeeID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployee(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(employee, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchEmployee(t, recorder.Body, employee)
			},
		},
		{
			name:       "NotFound",
			employeeID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployee(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(db.Employee{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:       "InternalError",
			employeeID: employee.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployee(gomock.Any(), gomock.Eq(employee.ID)).
					Times(1).
					Return(db.Employee{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:       "InvalidID",
			employeeID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetEmployee(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		// TODO add more cases

	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// 啟動測試伺服器並發送請求
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/employees/%d", tc.employeeID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			// check response
			tc.checkResponse(t, recorder)
		})
	}
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

func requireBodyMatchEmployee(t *testing.T, body *bytes.Buffer, employee db.Employee) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotEmployee db.Employee
	err = json.Unmarshal(data, &gotEmployee)
	require.NoError(t, err)
	require.Equal(t, employee, gotEmployee)
}
