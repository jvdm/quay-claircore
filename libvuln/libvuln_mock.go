// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quay/claircore/libvuln (interfaces: Libvuln)

// Package libvuln is a generated GoMock package.
package libvuln

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	claircore "github.com/quay/claircore"
	reflect "reflect"
)

// MockLibvuln is a mock of Libvuln interface
type MockLibvuln struct {
	ctrl     *gomock.Controller
	recorder *MockLibvulnMockRecorder
}

// MockLibvulnMockRecorder is the mock recorder for MockLibvuln
type MockLibvulnMockRecorder struct {
	mock *MockLibvuln
}

// NewMockLibvuln creates a new mock instance
func NewMockLibvuln(ctrl *gomock.Controller) *MockLibvuln {
	mock := &MockLibvuln{ctrl: ctrl}
	mock.recorder = &MockLibvulnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLibvuln) EXPECT() *MockLibvulnMockRecorder {
	return m.recorder
}

// Scan mocks base method
func (m *MockLibvuln) Scan(arg0 context.Context, arg1 *claircore.ScanReport) (*claircore.VulnerabilityReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", arg0, arg1)
	ret0, _ := ret[0].(*claircore.VulnerabilityReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scan indicates an expected call of Scan
func (mr *MockLibvulnMockRecorder) Scan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockLibvuln)(nil).Scan), arg0, arg1)
}
