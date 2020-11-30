// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/environment/environmentiface/iface.go

// Package mock_environmentiface is a generated GoMock package.
package mock_environmentiface

import (
	gomock "github.com/golang/mock/gomock"
	environmentiface "github.com/jecolasurdo/marsrover/pkg/environment/environmentiface"
	environmenttypes "github.com/jecolasurdo/marsrover/pkg/environment/environmenttypes"
	objectiface "github.com/jecolasurdo/marsrover/pkg/objects/objectiface"
	spatial "github.com/jecolasurdo/marsrover/pkg/spatial"
	reflect "reflect"
)

// MockEnvironmentBuilder is a mock of EnvironmentBuilder interface
type MockEnvironmentBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentBuilderMockRecorder
}

// MockEnvironmentBuilderMockRecorder is the mock recorder for MockEnvironmentBuilder
type MockEnvironmentBuilderMockRecorder struct {
	mock *MockEnvironmentBuilder
}

// NewMockEnvironmentBuilder creates a new mock instance
func NewMockEnvironmentBuilder(ctrl *gomock.Controller) *MockEnvironmentBuilder {
	mock := &MockEnvironmentBuilder{ctrl: ctrl}
	mock.recorder = &MockEnvironmentBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentBuilder) EXPECT() *MockEnvironmentBuilderMockRecorder {
	return m.recorder
}

// NewEnvironment mocks base method
func (m *MockEnvironmentBuilder) NewEnvironment(arg0 spatial.Point) (environmentiface.Environmenter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEnvironment", arg0)
	ret0, _ := ret[0].(environmentiface.Environmenter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewEnvironment indicates an expected call of NewEnvironment
func (mr *MockEnvironmentBuilderMockRecorder) NewEnvironment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEnvironment", reflect.TypeOf((*MockEnvironmentBuilder)(nil).NewEnvironment), arg0)
}

// MockEnvironmenter is a mock of Environmenter interface
type MockEnvironmenter struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmenterMockRecorder
}

// MockEnvironmenterMockRecorder is the mock recorder for MockEnvironmenter
type MockEnvironmenterMockRecorder struct {
	mock *MockEnvironmenter
}

// NewMockEnvironmenter creates a new mock instance
func NewMockEnvironmenter(ctrl *gomock.Controller) *MockEnvironmenter {
	mock := &MockEnvironmenter{ctrl: ctrl}
	mock.recorder = &MockEnvironmenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmenter) EXPECT() *MockEnvironmenterMockRecorder {
	return m.recorder
}

// GetDimensions mocks base method
func (m *MockEnvironmenter) GetDimensions() spatial.Point {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensions")
	ret0, _ := ret[0].(spatial.Point)
	return ret0
}

// GetDimensions indicates an expected call of GetDimensions
func (mr *MockEnvironmenterMockRecorder) GetDimensions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensions", reflect.TypeOf((*MockEnvironmenter)(nil).GetDimensions))
}

// PlaceObject mocks base method
func (m *MockEnvironmenter) PlaceObject(arg0 objectiface.Objecter, arg1 spatial.Point) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlaceObject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PlaceObject indicates an expected call of PlaceObject
func (mr *MockEnvironmenterMockRecorder) PlaceObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlaceObject", reflect.TypeOf((*MockEnvironmenter)(nil).PlaceObject), arg0, arg1)
}

// RecordMovement mocks base method
func (m *MockEnvironmenter) RecordMovement(arg0 objectiface.Objecter, arg1 spatial.Point) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordMovement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordMovement indicates an expected call of RecordMovement
func (mr *MockEnvironmenterMockRecorder) RecordMovement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordMovement", reflect.TypeOf((*MockEnvironmenter)(nil).RecordMovement), arg0, arg1)
}

// ShowObjects mocks base method
func (m *MockEnvironmenter) ShowObjects() map[spatial.Point][]objectiface.Objecter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowObjects")
	ret0, _ := ret[0].(map[spatial.Point][]objectiface.Objecter)
	return ret0
}

// ShowObjects indicates an expected call of ShowObjects
func (mr *MockEnvironmenterMockRecorder) ShowObjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowObjects", reflect.TypeOf((*MockEnvironmenter)(nil).ShowObjects))
}

// FindObject mocks base method
func (m *MockEnvironmenter) FindObject(arg0 objectiface.Objecter) (bool, *environmenttypes.ObjectPosition) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindObject", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*environmenttypes.ObjectPosition)
	return ret0, ret1
}

// FindObject indicates an expected call of FindObject
func (mr *MockEnvironmenterMockRecorder) FindObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindObject", reflect.TypeOf((*MockEnvironmenter)(nil).FindObject), arg0)
}

// InspectPosition mocks base method
func (m *MockEnvironmenter) InspectPosition(arg0 spatial.Point) (bool, []objectiface.Objecter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectPosition", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]objectiface.Objecter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// InspectPosition indicates an expected call of InspectPosition
func (mr *MockEnvironmenterMockRecorder) InspectPosition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectPosition", reflect.TypeOf((*MockEnvironmenter)(nil).InspectPosition), arg0)
}
