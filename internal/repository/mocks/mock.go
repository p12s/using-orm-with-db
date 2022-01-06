// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/p12s/using-orm-with-db/internal/repository (interfaces: Auther)

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/p12s/using-orm-with-db/internal/domain"
)

// MockAuther is a mock of Auther interface.
type MockAuther struct {
	ctrl     *gomock.Controller
	recorder *MockAutherMockRecorder
}

// MockAutherMockRecorder is the mock recorder for MockAuther.
type MockAutherMockRecorder struct {
	mock *MockAuther
}

// NewMockAuther creates a new mock instance.
func NewMockAuther(ctrl *gomock.Controller) *MockAuther {
	mock := &MockAuther{ctrl: ctrl}
	mock.recorder = &MockAutherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuther) EXPECT() *MockAutherMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockAuther) CreateAccount(arg0 domain.SignUpInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockAutherMockRecorder) CreateAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAuther)(nil).CreateAccount), arg0)
}

// GetAccountById mocks base method.
func (m *MockAuther) GetAccountById(arg0 int) (domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountById", arg0)
	ret0, _ := ret[0].(domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountById indicates an expected call of GetAccountById.
func (mr *MockAutherMockRecorder) GetAccountById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountById", reflect.TypeOf((*MockAuther)(nil).GetAccountById), arg0)
}

// GetByCredentials mocks base method.
func (m *MockAuther) GetByCredentials(arg0, arg1 string) (domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCredentials", arg0, arg1)
	ret0, _ := ret[0].(domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCredentials indicates an expected call of GetByCredentials.
func (mr *MockAutherMockRecorder) GetByCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCredentials", reflect.TypeOf((*MockAuther)(nil).GetByCredentials), arg0, arg1)
}
