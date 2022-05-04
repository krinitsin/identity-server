// Code generated by MockGen. DO NOT EDIT.
// Source: sparket/pkg/utils (interfaces: ITimeService)

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockITimeService is a mock of ITimeService interface.
type MockITimeService struct {
	ctrl     *gomock.Controller
	recorder *MockITimeServiceMockRecorder
}

// MockITimeServiceMockRecorder is the mock recorder for MockITimeService.
type MockITimeServiceMockRecorder struct {
	mock *MockITimeService
}

// NewMockITimeService creates a new mock instance.
func NewMockITimeService(ctrl *gomock.Controller) *MockITimeService {
	mock := &MockITimeService{ctrl: ctrl}
	mock.recorder = &MockITimeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITimeService) EXPECT() *MockITimeServiceMockRecorder {
	return m.recorder
}

// GetTimeFromUnixMilliEpoch mocks base method.
func (m *MockITimeService) GetTimeFromUnixMilliEpoch(arg0 int64) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeFromUnixMilliEpoch", arg0)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimeFromUnixMilliEpoch indicates an expected call of GetTimeFromUnixMilliEpoch.
func (mr *MockITimeServiceMockRecorder) GetTimeFromUnixMilliEpoch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeFromUnixMilliEpoch", reflect.TypeOf((*MockITimeService)(nil).GetTimeFromUnixMilliEpoch), arg0)
}

// GetTimeISONow mocks base method.
func (m *MockITimeService) GetTimeISONow() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeISONow")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTimeISONow indicates an expected call of GetTimeISONow.
func (mr *MockITimeServiceMockRecorder) GetTimeISONow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeISONow", reflect.TypeOf((*MockITimeService)(nil).GetTimeISONow))
}

// GetTimeNow mocks base method.
func (m *MockITimeService) GetTimeNow() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeNow")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimeNow indicates an expected call of GetTimeNow.
func (mr *MockITimeServiceMockRecorder) GetTimeNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeNow", reflect.TypeOf((*MockITimeService)(nil).GetTimeNow))
}

// GetUTCTimeISONow mocks base method.
func (m *MockITimeService) GetUTCTimeISONow() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUTCTimeISONow")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUTCTimeISONow indicates an expected call of GetUTCTimeISONow.
func (mr *MockITimeServiceMockRecorder) GetUTCTimeISONow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUTCTimeISONow", reflect.TypeOf((*MockITimeService)(nil).GetUTCTimeISONow))
}

// GetUTCTimeNow mocks base method.
func (m *MockITimeService) GetUTCTimeNow() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUTCTimeNow")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetUTCTimeNow indicates an expected call of GetUTCTimeNow.
func (mr *MockITimeServiceMockRecorder) GetUTCTimeNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUTCTimeNow", reflect.TypeOf((*MockITimeService)(nil).GetUTCTimeNow))
}

// GetUnixEpochMilliFromTime mocks base method.
func (m *MockITimeService) GetUnixEpochMilliFromTime(arg0 time.Time) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnixEpochMilliFromTime", arg0)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetUnixEpochMilliFromTime indicates an expected call of GetUnixEpochMilliFromTime.
func (mr *MockITimeServiceMockRecorder) GetUnixEpochMilliFromTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnixEpochMilliFromTime", reflect.TypeOf((*MockITimeService)(nil).GetUnixEpochMilliFromTime), arg0)
}