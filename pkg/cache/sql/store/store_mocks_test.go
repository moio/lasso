// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/lasso/pkg/cache/sql/store (interfaces: DBClient)
//
// Generated by this command:
//
//	mockgen --build_flags=--mod=mod -package store -destination ./store_mocks_test.go github.com/rancher/lasso/pkg/cache/sql/store DBClient
//

// Package store is a generated GoMock package.
package store

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	db "github.com/rancher/lasso/pkg/cache/sql/db"
	transaction "github.com/rancher/lasso/pkg/cache/sql/db/transaction"
	gomock "go.uber.org/mock/gomock"
)

// MockDBClient is a mock of DBClient interface.
type MockDBClient struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientMockRecorder
}

// MockDBClientMockRecorder is the mock recorder for MockDBClient.
type MockDBClientMockRecorder struct {
	mock *MockDBClient
}

// NewMockDBClient creates a new mock instance.
func NewMockDBClient(ctrl *gomock.Controller) *MockDBClient {
	mock := &MockDBClient{ctrl: ctrl}
	mock.recorder = &MockDBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBClient) EXPECT() *MockDBClientMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockDBClient) BeginTx(arg0 context.Context, arg1 bool) (db.TXClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", arg0, arg1)
	ret0, _ := ret[0].(db.TXClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockDBClientMockRecorder) BeginTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockDBClient)(nil).BeginTx), arg0, arg1)
}

// CloseStmt mocks base method.
func (m *MockDBClient) CloseStmt(arg0 db.Closable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseStmt", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseStmt indicates an expected call of CloseStmt.
func (mr *MockDBClientMockRecorder) CloseStmt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseStmt", reflect.TypeOf((*MockDBClient)(nil).CloseStmt), arg0)
}

// Prepare mocks base method.
func (m *MockDBClient) Prepare(arg0 string) *sql.Stmt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", arg0)
	ret0, _ := ret[0].(*sql.Stmt)
	return ret0
}

// Prepare indicates an expected call of Prepare.
func (mr *MockDBClientMockRecorder) Prepare(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockDBClient)(nil).Prepare), arg0)
}

// QueryForRows mocks base method.
func (m *MockDBClient) QueryForRows(arg0 context.Context, arg1 transaction.Stmt, arg2 ...any) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryForRows", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryForRows indicates an expected call of QueryForRows.
func (mr *MockDBClientMockRecorder) QueryForRows(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryForRows", reflect.TypeOf((*MockDBClient)(nil).QueryForRows), varargs...)
}

// ReadInt mocks base method.
func (m *MockDBClient) ReadInt(arg0 db.Rows) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadInt", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadInt indicates an expected call of ReadInt.
func (mr *MockDBClientMockRecorder) ReadInt(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInt", reflect.TypeOf((*MockDBClient)(nil).ReadInt), arg0)
}

// ReadObjects mocks base method.
func (m *MockDBClient) ReadObjects(arg0 db.Rows, arg1 reflect.Type, arg2 bool) ([]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadObjects", arg0, arg1, arg2)
	ret0, _ := ret[0].([]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadObjects indicates an expected call of ReadObjects.
func (mr *MockDBClientMockRecorder) ReadObjects(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadObjects", reflect.TypeOf((*MockDBClient)(nil).ReadObjects), arg0, arg1, arg2)
}

// ReadStrings mocks base method.
func (m *MockDBClient) ReadStrings(arg0 db.Rows) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadStrings", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadStrings indicates an expected call of ReadStrings.
func (mr *MockDBClientMockRecorder) ReadStrings(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadStrings", reflect.TypeOf((*MockDBClient)(nil).ReadStrings), arg0)
}

// Upsert mocks base method.
func (m *MockDBClient) Upsert(arg0 db.TXClient, arg1 *sql.Stmt, arg2 string, arg3 any, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockDBClientMockRecorder) Upsert(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockDBClient)(nil).Upsert), arg0, arg1, arg2, arg3, arg4)
}
