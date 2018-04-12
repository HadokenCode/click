// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kamilsk/click/server (interfaces: Service)

// Package server_test is a generated GoMock package.
package server_test

import (
	gomock "github.com/golang/mock/gomock"
	domain "github.com/kamilsk/click/domain"
	transfer "github.com/kamilsk/click/transfer"
	v1 "github.com/kamilsk/click/transfer/api/v1"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// HandleGetV1 mocks base method
func (m *MockService) HandleGetV1(arg0 v1.GetRequest) v1.GetResponse {
	ret := m.ctrl.Call(m, "HandleGetV1", arg0)
	ret0, _ := ret[0].(v1.GetResponse)
	return ret0
}

// HandleGetV1 indicates an expected call of HandleGetV1
func (mr *MockServiceMockRecorder) HandleGetV1(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleGetV1", reflect.TypeOf((*MockService)(nil).HandleGetV1), arg0)
}

// HandleRedirect mocks base method
func (m *MockService) HandleRedirect(arg0 transfer.RedirectRequest) transfer.RedirectResponse {
	ret := m.ctrl.Call(m, "HandleRedirect", arg0)
	ret0, _ := ret[0].(transfer.RedirectResponse)
	return ret0
}

// HandleRedirect indicates an expected call of HandleRedirect
func (mr *MockServiceMockRecorder) HandleRedirect(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleRedirect", reflect.TypeOf((*MockService)(nil).HandleRedirect), arg0)
}

// LogRedirectEvent mocks base method
func (m *MockService) LogRedirectEvent(arg0 domain.Log) {
	m.ctrl.Call(m, "LogRedirectEvent", arg0)
}

// LogRedirectEvent indicates an expected call of LogRedirectEvent
func (mr *MockServiceMockRecorder) LogRedirectEvent(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogRedirectEvent", reflect.TypeOf((*MockService)(nil).LogRedirectEvent), arg0)
}