// Code generated by MockGen. DO NOT EDIT.
// Source: node_visitor.go

// Package mock_agent is a generated GoMock package.
package mock_agent

import (
	types "chasqi-go/types"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNodeVisitor is a mock of NodeVisitor interface.
type MockNodeVisitor struct {
	ctrl     *gomock.Controller
	recorder *MockNodeVisitorMockRecorder
}

// MockNodeVisitorMockRecorder is the mock recorder for MockNodeVisitor.
type MockNodeVisitorMockRecorder struct {
	mock *MockNodeVisitor
}

// NewMockNodeVisitor creates a new mock instance.
func NewMockNodeVisitor(ctrl *gomock.Controller) *MockNodeVisitor {
	mock := &MockNodeVisitor{ctrl: ctrl}
	mock.recorder = &MockNodeVisitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeVisitor) EXPECT() *MockNodeVisitorMockRecorder {
	return m.recorder
}

// Visit mocks base method.
func (m *MockNodeVisitor) Visit(method, url string, body io.Reader, headers map[string][]string) (*types.ResponseResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Visit", method, url, body, headers)
	ret0, _ := ret[0].(*types.ResponseResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Visit indicates an expected call of Visit.
func (mr *MockNodeVisitorMockRecorder) Visit(method, url, body, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Visit", reflect.TypeOf((*MockNodeVisitor)(nil).Visit), method, url, body, headers)
}
