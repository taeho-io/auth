// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/taeho-io/idl/gen/go/auth (interfaces: AuthClient)

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	auth "github.com/taeho-io/idl/gen/go/auth"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAuthClient is a mock of AuthClient interface
type MockAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthClientMockRecorder
}

// MockAuthClientMockRecorder is the mock recorder for MockAuthClient
type MockAuthClientMockRecorder struct {
	mock *MockAuthClient
}

// NewMockAuthClient creates a new mock instance
func NewMockAuthClient(ctrl *gomock.Controller) *MockAuthClient {
	mock := &MockAuthClient{ctrl: ctrl}
	mock.recorder = &MockAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthClient) EXPECT() *MockAuthClientMockRecorder {
	return m.recorder
}

// Auth mocks base method
func (m *MockAuthClient) Auth(arg0 context.Context, arg1 *auth.AuthRequest, arg2 ...grpc.CallOption) (*auth.AuthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Auth", varargs...)
	ret0, _ := ret[0].(*auth.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth
func (mr *MockAuthClientMockRecorder) Auth(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockAuthClient)(nil).Auth), varargs...)
}

// Jwks mocks base method
func (m *MockAuthClient) Jwks(arg0 context.Context, arg1 *auth.JwksRequest, arg2 ...grpc.CallOption) (*auth.JwksResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Jwks", varargs...)
	ret0, _ := ret[0].(*auth.JwksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Jwks indicates an expected call of Jwks
func (mr *MockAuthClientMockRecorder) Jwks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jwks", reflect.TypeOf((*MockAuthClient)(nil).Jwks), varargs...)
}

// Parse mocks base method
func (m *MockAuthClient) Parse(arg0 context.Context, arg1 *auth.ParseRequest, arg2 ...grpc.CallOption) (*auth.ParseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Parse", varargs...)
	ret0, _ := ret[0].(*auth.ParseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse
func (mr *MockAuthClientMockRecorder) Parse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockAuthClient)(nil).Parse), varargs...)
}

// Refresh mocks base method
func (m *MockAuthClient) Refresh(arg0 context.Context, arg1 *auth.RefreshRequest, arg2 ...grpc.CallOption) (*auth.RefreshResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Refresh", varargs...)
	ret0, _ := ret[0].(*auth.RefreshResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh
func (mr *MockAuthClientMockRecorder) Refresh(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockAuthClient)(nil).Refresh), varargs...)
}

// Verify mocks base method
func (m *MockAuthClient) Verify(arg0 context.Context, arg1 *auth.VerifyRequest, arg2 ...grpc.CallOption) (*auth.VerifyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Verify", varargs...)
	ret0, _ := ret[0].(*auth.VerifyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockAuthClientMockRecorder) Verify(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockAuthClient)(nil).Verify), varargs...)
}
