// Code generated by MockGen. DO NOT EDIT.
// Source: core/metric.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMetric is a mock of Metric interface
type MockMetric struct {
	ctrl     *gomock.Controller
	recorder *MockMetricMockRecorder
}

// MockMetricMockRecorder is the mock recorder for MockMetric
type MockMetricMockRecorder struct {
	mock *MockMetric
}

// NewMockMetric creates a new mock instance
func NewMockMetric(ctrl *gomock.Controller) *MockMetric {
	mock := &MockMetric{ctrl: ctrl}
	mock.recorder = &MockMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetric) EXPECT() *MockMetricMockRecorder {
	return m.recorder
}

// InjectCounter mocks base method
func (m *MockMetric) InjectCounter(metricName string, labels []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InjectCounter", metricName, labels)
}

// InjectCounter indicates an expected call of InjectCounter
func (mr *MockMetricMockRecorder) InjectCounter(metricName, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectCounter", reflect.TypeOf((*MockMetric)(nil).InjectCounter), metricName, labels)
}

// InjectHistogram mocks base method
func (m *MockMetric) InjectHistogram(metricName string, labels []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InjectHistogram", metricName, labels)
}

// InjectHistogram indicates an expected call of InjectHistogram
func (mr *MockMetricMockRecorder) InjectHistogram(metricName, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InjectHistogram", reflect.TypeOf((*MockMetric)(nil).InjectHistogram), metricName, labels)
}

// Inc mocks base method
func (m *MockMetric) Inc(metricName string, labels map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inc", metricName, labels)
	ret0, _ := ret[0].(error)
	return ret0
}

// Inc indicates an expected call of Inc
func (mr *MockMetricMockRecorder) Inc(metricName, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inc", reflect.TypeOf((*MockMetric)(nil).Inc), metricName, labels)
}

// Observe mocks base method
func (m *MockMetric) Observe(metricName string, value float64, labels map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Observe", metricName, value, labels)
	ret0, _ := ret[0].(error)
	return ret0
}

// Observe indicates an expected call of Observe
func (mr *MockMetricMockRecorder) Observe(metricName, value, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Observe", reflect.TypeOf((*MockMetric)(nil).Observe), metricName, value, labels)
}