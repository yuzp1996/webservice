// Code generated by MockGen. DO NOT EDIT.
// Source: webservice/Http/server/interface (interfaces: Animal)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	interface0 "webservice/Http/server/interface"
)

// MockAnimal is a mock of Animal interface
type MockAnimal struct {
	ctrl     *gomock.Controller
	recorder *MockAnimalMockRecorder
}

// MockAnimalMockRecorder is the mock recorder for MockAnimal
type MockAnimalMockRecorder struct {
	mock *MockAnimal
}

// NewMockAnimal creates a new mock instance
func NewMockAnimal(ctrl *gomock.Controller) *MockAnimal {
	mock := &MockAnimal{ctrl: ctrl}
	mock.recorder = &MockAnimalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnimal) EXPECT() *MockAnimalMockRecorder {
	return m.recorder
}

// Myname mocks base method
func (m *MockAnimal) Myname() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Myname")
	ret0, _ := ret[0].(string)
	return ret0
}

// Myname indicates an expected call of Myname
func (mr *MockAnimalMockRecorder) Myname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Myname", reflect.TypeOf((*MockAnimal)(nil).Myname))
}

// Say mocks base method
func (m *MockAnimal) Say(arg0 interface0.Animal) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Say", arg0)
}

// Say indicates an expected call of Say
func (mr *MockAnimalMockRecorder) Say(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Say", reflect.TypeOf((*MockAnimal)(nil).Say), arg0)
}
