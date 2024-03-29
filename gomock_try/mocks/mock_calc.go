// Code generated by MockGen. DO NOT EDIT.
// Source: nagase/doer (interfaces: Calculator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCalculator is a mock of Calculator interface.
type MockCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockCalculatorMockRecorder
}

// MockCalculatorMockRecorder is the mock recorder for MockCalculator.
type MockCalculatorMockRecorder struct {
	mock *MockCalculator
}

// NewMockCalculator creates a new mock instance.
func NewMockCalculator(ctrl *gomock.Controller) *MockCalculator {
	mock := &MockCalculator{ctrl: ctrl}
	mock.recorder = &MockCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculator) EXPECT() *MockCalculatorMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCalculator) Add(arg0, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockCalculatorMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCalculator)(nil).Add), arg0, arg1)
}

// DoSomething mocks base method.
func (m *MockCalculator) DoSomething(arg0 int, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoSomething", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSomething indicates an expected call of DoSomething.
func (mr *MockCalculatorMockRecorder) DoSomething(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoSomething", reflect.TypeOf((*MockCalculator)(nil).DoSomething), arg0, arg1)
}
