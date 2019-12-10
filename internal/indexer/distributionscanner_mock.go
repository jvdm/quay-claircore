// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quay/claircore/internal/indexer (interfaces: DistributionScanner)

// Package indexer is a generated GoMock package.
package indexer

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	claircore "github.com/quay/claircore"
	reflect "reflect"
)

// MockDistributionScanner is a mock of DistributionScanner interface
type MockDistributionScanner struct {
	ctrl     *gomock.Controller
	recorder *MockDistributionScannerMockRecorder
}

// MockDistributionScannerMockRecorder is the mock recorder for MockDistributionScanner
type MockDistributionScannerMockRecorder struct {
	mock *MockDistributionScanner
}

// NewMockDistributionScanner creates a new mock instance
func NewMockDistributionScanner(ctrl *gomock.Controller) *MockDistributionScanner {
	mock := &MockDistributionScanner{ctrl: ctrl}
	mock.recorder = &MockDistributionScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDistributionScanner) EXPECT() *MockDistributionScannerMockRecorder {
	return m.recorder
}

// Kind mocks base method
func (m *MockDistributionScanner) Kind() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(string)
	return ret0
}

// Kind indicates an expected call of Kind
func (mr *MockDistributionScannerMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockDistributionScanner)(nil).Kind))
}

// Name mocks base method
func (m *MockDistributionScanner) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockDistributionScannerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockDistributionScanner)(nil).Name))
}

// Scan mocks base method
func (m *MockDistributionScanner) Scan(arg0 context.Context, arg1 *claircore.Layer) ([]*claircore.Distribution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", arg0, arg1)
	ret0, _ := ret[0].([]*claircore.Distribution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scan indicates an expected call of Scan
func (mr *MockDistributionScannerMockRecorder) Scan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockDistributionScanner)(nil).Scan), arg0, arg1)
}

// Version mocks base method
func (m *MockDistributionScanner) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockDistributionScannerMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockDistributionScanner)(nil).Version))
}