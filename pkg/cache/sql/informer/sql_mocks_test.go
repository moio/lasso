// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/lasso/pkg/cache/sql/informer (interfaces: Store)

// Package informer is a generated GoMock package.
package informer

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/rancher/lasso/pkg/cache/sql/db"
	transaction "github.com/rancher/lasso/pkg/cache/sql/db/transaction"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStore) Add(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockStoreMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStore)(nil).Add), arg0)
}

// Begin mocks base method.
func (m *MockStore) Begin() (db.TXClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(db.TXClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockStoreMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockStore)(nil).Begin))
}

// CloseStmt mocks base method.
func (m *MockStore) CloseStmt(arg0 db.Closable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseStmt", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseStmt indicates an expected call of CloseStmt.
func (mr *MockStoreMockRecorder) CloseStmt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseStmt", reflect.TypeOf((*MockStore)(nil).CloseStmt), arg0)
}

// Delete mocks base method.
func (m *MockStore) Delete(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStoreMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockStore) Get(arg0 interface{}) (interface{}, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockStoreMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), arg0)
}

// GetByKey mocks base method.
func (m *MockStore) GetByKey(arg0 string) (interface{}, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKey", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByKey indicates an expected call of GetByKey.
func (mr *MockStoreMockRecorder) GetByKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKey", reflect.TypeOf((*MockStore)(nil).GetByKey), arg0)
}

// GetName mocks base method.
func (m *MockStore) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockStoreMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockStore)(nil).GetName))
}

// GetShouldEncrypt mocks base method.
func (m *MockStore) GetShouldEncrypt() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShouldEncrypt")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetShouldEncrypt indicates an expected call of GetShouldEncrypt.
func (mr *MockStoreMockRecorder) GetShouldEncrypt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShouldEncrypt", reflect.TypeOf((*MockStore)(nil).GetShouldEncrypt))
}

// GetType mocks base method.
func (m *MockStore) GetType() reflect.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(reflect.Type)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockStoreMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockStore)(nil).GetType))
}

// List mocks base method.
func (m *MockStore) List() []interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]interface{})
	return ret0
}

// List indicates an expected call of List.
func (mr *MockStoreMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStore)(nil).List))
}

// ListKeys mocks base method.
func (m *MockStore) ListKeys() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockStoreMockRecorder) ListKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockStore)(nil).ListKeys))
}

// NewConnection mocks base method.
func (m *MockStore) NewConnection() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConnection")
	ret0, _ := ret[0].(error)
	return ret0
}

// NewConnection indicates an expected call of NewConnection.
func (mr *MockStoreMockRecorder) NewConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConnection", reflect.TypeOf((*MockStore)(nil).NewConnection))
}

// Prepare mocks base method.
func (m *MockStore) Prepare(arg0 string) *sql.Stmt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", arg0)
	ret0, _ := ret[0].(*sql.Stmt)
	return ret0
}

// Prepare indicates an expected call of Prepare.
func (mr *MockStoreMockRecorder) Prepare(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockStore)(nil).Prepare), arg0)
}

// QueryForRows mocks base method.
func (m *MockStore) QueryForRows(arg0 context.Context, arg1 transaction.Stmt, arg2 ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryForRows", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryForRows indicates an expected call of QueryForRows.
func (mr *MockStoreMockRecorder) QueryForRows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryForRows", reflect.TypeOf((*MockStore)(nil).QueryForRows), varargs...)
}

// ReadObjects mocks base method.
func (m *MockStore) ReadObjects(arg0 db.Rows, arg1 reflect.Type, arg2 bool) ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadObjects", arg0, arg1, arg2)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadObjects indicates an expected call of ReadObjects.
func (mr *MockStoreMockRecorder) ReadObjects(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadObjects", reflect.TypeOf((*MockStore)(nil).ReadObjects), arg0, arg1, arg2)
}

// ReadStrings mocks base method.
func (m *MockStore) ReadStrings(arg0 db.Rows) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadStrings", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadStrings indicates an expected call of ReadStrings.
func (mr *MockStoreMockRecorder) ReadStrings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadStrings", reflect.TypeOf((*MockStore)(nil).ReadStrings), arg0)
}

// RegisterAfterDelete mocks base method.
func (m *MockStore) RegisterAfterDelete(arg0 func(string, db.TXClient) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterAfterDelete", arg0)
}

// RegisterAfterDelete indicates an expected call of RegisterAfterDelete.
func (mr *MockStoreMockRecorder) RegisterAfterDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAfterDelete", reflect.TypeOf((*MockStore)(nil).RegisterAfterDelete), arg0)
}

// RegisterAfterUpsert mocks base method.
func (m *MockStore) RegisterAfterUpsert(arg0 func(string, interface{}, db.TXClient) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterAfterUpsert", arg0)
}

// RegisterAfterUpsert indicates an expected call of RegisterAfterUpsert.
func (mr *MockStoreMockRecorder) RegisterAfterUpsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAfterUpsert", reflect.TypeOf((*MockStore)(nil).RegisterAfterUpsert), arg0)
}

// Replace mocks base method.
func (m *MockStore) Replace(arg0 []interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Replace indicates an expected call of Replace.
func (mr *MockStoreMockRecorder) Replace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockStore)(nil).Replace), arg0, arg1)
}

// Resync mocks base method.
func (m *MockStore) Resync() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resync")
	ret0, _ := ret[0].(error)
	return ret0
}

// Resync indicates an expected call of Resync.
func (mr *MockStoreMockRecorder) Resync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resync", reflect.TypeOf((*MockStore)(nil).Resync))
}

// Update mocks base method.
func (m *MockStore) Update(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStoreMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), arg0)
}

// Upsert mocks base method.
func (m *MockStore) Upsert(arg0 db.TXClient, arg1 *sql.Stmt, arg2 string, arg3 interface{}, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockStoreMockRecorder) Upsert(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockStore)(nil).Upsert), arg0, arg1, arg2, arg3, arg4)
}
