// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	v1alpha1 "github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// BaseWorkflowWithStatus is an autogenerated mock type for the BaseWorkflowWithStatus type
type BaseWorkflowWithStatus struct {
	mock.Mock
}

type BaseWorkflowWithStatus_FromNode struct {
	*mock.Call
}

func (_m BaseWorkflowWithStatus_FromNode) Return(_a0 []string, _a1 error) *BaseWorkflowWithStatus_FromNode {
	return &BaseWorkflowWithStatus_FromNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *BaseWorkflowWithStatus) OnFromNode(name string) *BaseWorkflowWithStatus_FromNode {
	c := _m.On("FromNode", name)
	return &BaseWorkflowWithStatus_FromNode{Call: c}
}

func (_m *BaseWorkflowWithStatus) OnFromNodeMatch(matchers ...interface{}) *BaseWorkflowWithStatus_FromNode {
	c := _m.On("FromNode", matchers...)
	return &BaseWorkflowWithStatus_FromNode{Call: c}
}

// FromNode provides a mock function with given fields: name
func (_m *BaseWorkflowWithStatus) FromNode(name string) ([]string, error) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type BaseWorkflowWithStatus_GetID struct {
	*mock.Call
}

func (_m BaseWorkflowWithStatus_GetID) Return(_a0 string) *BaseWorkflowWithStatus_GetID {
	return &BaseWorkflowWithStatus_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *BaseWorkflowWithStatus) OnGetID() *BaseWorkflowWithStatus_GetID {
	c := _m.On("GetID")
	return &BaseWorkflowWithStatus_GetID{Call: c}
}

func (_m *BaseWorkflowWithStatus) OnGetIDMatch(matchers ...interface{}) *BaseWorkflowWithStatus_GetID {
	c := _m.On("GetID", matchers...)
	return &BaseWorkflowWithStatus_GetID{Call: c}
}

// GetID provides a mock function with given fields:
func (_m *BaseWorkflowWithStatus) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type BaseWorkflowWithStatus_GetNode struct {
	*mock.Call
}

func (_m BaseWorkflowWithStatus_GetNode) Return(_a0 v1alpha1.ExecutableNode, _a1 bool) *BaseWorkflowWithStatus_GetNode {
	return &BaseWorkflowWithStatus_GetNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *BaseWorkflowWithStatus) OnGetNode(nodeID string) *BaseWorkflowWithStatus_GetNode {
	c := _m.On("GetNode", nodeID)
	return &BaseWorkflowWithStatus_GetNode{Call: c}
}

func (_m *BaseWorkflowWithStatus) OnGetNodeMatch(matchers ...interface{}) *BaseWorkflowWithStatus_GetNode {
	c := _m.On("GetNode", matchers...)
	return &BaseWorkflowWithStatus_GetNode{Call: c}
}

// GetNode provides a mock function with given fields: nodeID
func (_m *BaseWorkflowWithStatus) GetNode(nodeID string) (v1alpha1.ExecutableNode, bool) {
	ret := _m.Called(nodeID)

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableNode); ok {
		r0 = rf(nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

type BaseWorkflowWithStatus_GetNodeExecutionStatus struct {
	*mock.Call
}

func (_m BaseWorkflowWithStatus_GetNodeExecutionStatus) Return(_a0 v1alpha1.ExecutableNodeStatus) *BaseWorkflowWithStatus_GetNodeExecutionStatus {
	return &BaseWorkflowWithStatus_GetNodeExecutionStatus{Call: _m.Call.Return(_a0)}
}

func (_m *BaseWorkflowWithStatus) OnGetNodeExecutionStatus(ctx context.Context, id string) *BaseWorkflowWithStatus_GetNodeExecutionStatus {
	c := _m.On("GetNodeExecutionStatus", ctx, id)
	return &BaseWorkflowWithStatus_GetNodeExecutionStatus{Call: c}
}

func (_m *BaseWorkflowWithStatus) OnGetNodeExecutionStatusMatch(matchers ...interface{}) *BaseWorkflowWithStatus_GetNodeExecutionStatus {
	c := _m.On("GetNodeExecutionStatus", matchers...)
	return &BaseWorkflowWithStatus_GetNodeExecutionStatus{Call: c}
}

// GetNodeExecutionStatus provides a mock function with given fields: ctx, id
func (_m *BaseWorkflowWithStatus) GetNodeExecutionStatus(ctx context.Context, id string) v1alpha1.ExecutableNodeStatus {
	ret := _m.Called(ctx, id)

	var r0 v1alpha1.ExecutableNodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, string) v1alpha1.ExecutableNodeStatus); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNodeStatus)
		}
	}

	return r0
}

type BaseWorkflowWithStatus_StartNode struct {
	*mock.Call
}

func (_m BaseWorkflowWithStatus_StartNode) Return(_a0 v1alpha1.ExecutableNode) *BaseWorkflowWithStatus_StartNode {
	return &BaseWorkflowWithStatus_StartNode{Call: _m.Call.Return(_a0)}
}

func (_m *BaseWorkflowWithStatus) OnStartNode() *BaseWorkflowWithStatus_StartNode {
	c := _m.On("StartNode")
	return &BaseWorkflowWithStatus_StartNode{Call: c}
}

func (_m *BaseWorkflowWithStatus) OnStartNodeMatch(matchers ...interface{}) *BaseWorkflowWithStatus_StartNode {
	c := _m.On("StartNode", matchers...)
	return &BaseWorkflowWithStatus_StartNode{Call: c}
}

// StartNode provides a mock function with given fields:
func (_m *BaseWorkflowWithStatus) StartNode() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}

type BaseWorkflowWithStatus_ToNode struct {
	*mock.Call
}

func (_m BaseWorkflowWithStatus_ToNode) Return(_a0 []string, _a1 error) *BaseWorkflowWithStatus_ToNode {
	return &BaseWorkflowWithStatus_ToNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *BaseWorkflowWithStatus) OnToNode(name string) *BaseWorkflowWithStatus_ToNode {
	c := _m.On("ToNode", name)
	return &BaseWorkflowWithStatus_ToNode{Call: c}
}

func (_m *BaseWorkflowWithStatus) OnToNodeMatch(matchers ...interface{}) *BaseWorkflowWithStatus_ToNode {
	c := _m.On("ToNode", matchers...)
	return &BaseWorkflowWithStatus_ToNode{Call: c}
}

// ToNode provides a mock function with given fields: name
func (_m *BaseWorkflowWithStatus) ToNode(name string) ([]string, error) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
