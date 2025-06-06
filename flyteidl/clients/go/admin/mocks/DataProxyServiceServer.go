// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/service"
	mock "github.com/stretchr/testify/mock"
)

// DataProxyServiceServer is an autogenerated mock type for the DataProxyServiceServer type
type DataProxyServiceServer struct {
	mock.Mock
}

type DataProxyServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *DataProxyServiceServer) EXPECT() *DataProxyServiceServer_Expecter {
	return &DataProxyServiceServer_Expecter{mock: &_m.Mock}
}

// CreateDownloadLink provides a mock function with given fields: _a0, _a1
func (_m *DataProxyServiceServer) CreateDownloadLink(_a0 context.Context, _a1 *service.CreateDownloadLinkRequest) (*service.CreateDownloadLinkResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateDownloadLink")
	}

	var r0 *service.CreateDownloadLinkResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateDownloadLinkRequest) (*service.CreateDownloadLinkResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateDownloadLinkRequest) *service.CreateDownloadLinkResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateDownloadLinkResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateDownloadLinkRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataProxyServiceServer_CreateDownloadLink_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDownloadLink'
type DataProxyServiceServer_CreateDownloadLink_Call struct {
	*mock.Call
}

// CreateDownloadLink is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *service.CreateDownloadLinkRequest
func (_e *DataProxyServiceServer_Expecter) CreateDownloadLink(_a0 interface{}, _a1 interface{}) *DataProxyServiceServer_CreateDownloadLink_Call {
	return &DataProxyServiceServer_CreateDownloadLink_Call{Call: _e.mock.On("CreateDownloadLink", _a0, _a1)}
}

func (_c *DataProxyServiceServer_CreateDownloadLink_Call) Run(run func(_a0 context.Context, _a1 *service.CreateDownloadLinkRequest)) *DataProxyServiceServer_CreateDownloadLink_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*service.CreateDownloadLinkRequest))
	})
	return _c
}

func (_c *DataProxyServiceServer_CreateDownloadLink_Call) Return(_a0 *service.CreateDownloadLinkResponse, _a1 error) *DataProxyServiceServer_CreateDownloadLink_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataProxyServiceServer_CreateDownloadLink_Call) RunAndReturn(run func(context.Context, *service.CreateDownloadLinkRequest) (*service.CreateDownloadLinkResponse, error)) *DataProxyServiceServer_CreateDownloadLink_Call {
	_c.Call.Return(run)
	return _c
}

// CreateDownloadLocation provides a mock function with given fields: _a0, _a1
func (_m *DataProxyServiceServer) CreateDownloadLocation(_a0 context.Context, _a1 *service.CreateDownloadLocationRequest) (*service.CreateDownloadLocationResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateDownloadLocation")
	}

	var r0 *service.CreateDownloadLocationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateDownloadLocationRequest) (*service.CreateDownloadLocationResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateDownloadLocationRequest) *service.CreateDownloadLocationResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateDownloadLocationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateDownloadLocationRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataProxyServiceServer_CreateDownloadLocation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDownloadLocation'
type DataProxyServiceServer_CreateDownloadLocation_Call struct {
	*mock.Call
}

// CreateDownloadLocation is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *service.CreateDownloadLocationRequest
func (_e *DataProxyServiceServer_Expecter) CreateDownloadLocation(_a0 interface{}, _a1 interface{}) *DataProxyServiceServer_CreateDownloadLocation_Call {
	return &DataProxyServiceServer_CreateDownloadLocation_Call{Call: _e.mock.On("CreateDownloadLocation", _a0, _a1)}
}

func (_c *DataProxyServiceServer_CreateDownloadLocation_Call) Run(run func(_a0 context.Context, _a1 *service.CreateDownloadLocationRequest)) *DataProxyServiceServer_CreateDownloadLocation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*service.CreateDownloadLocationRequest))
	})
	return _c
}

func (_c *DataProxyServiceServer_CreateDownloadLocation_Call) Return(_a0 *service.CreateDownloadLocationResponse, _a1 error) *DataProxyServiceServer_CreateDownloadLocation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataProxyServiceServer_CreateDownloadLocation_Call) RunAndReturn(run func(context.Context, *service.CreateDownloadLocationRequest) (*service.CreateDownloadLocationResponse, error)) *DataProxyServiceServer_CreateDownloadLocation_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUploadLocation provides a mock function with given fields: _a0, _a1
func (_m *DataProxyServiceServer) CreateUploadLocation(_a0 context.Context, _a1 *service.CreateUploadLocationRequest) (*service.CreateUploadLocationResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateUploadLocation")
	}

	var r0 *service.CreateUploadLocationResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateUploadLocationRequest) (*service.CreateUploadLocationResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateUploadLocationRequest) *service.CreateUploadLocationResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateUploadLocationResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateUploadLocationRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataProxyServiceServer_CreateUploadLocation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUploadLocation'
type DataProxyServiceServer_CreateUploadLocation_Call struct {
	*mock.Call
}

// CreateUploadLocation is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *service.CreateUploadLocationRequest
func (_e *DataProxyServiceServer_Expecter) CreateUploadLocation(_a0 interface{}, _a1 interface{}) *DataProxyServiceServer_CreateUploadLocation_Call {
	return &DataProxyServiceServer_CreateUploadLocation_Call{Call: _e.mock.On("CreateUploadLocation", _a0, _a1)}
}

func (_c *DataProxyServiceServer_CreateUploadLocation_Call) Run(run func(_a0 context.Context, _a1 *service.CreateUploadLocationRequest)) *DataProxyServiceServer_CreateUploadLocation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*service.CreateUploadLocationRequest))
	})
	return _c
}

func (_c *DataProxyServiceServer_CreateUploadLocation_Call) Return(_a0 *service.CreateUploadLocationResponse, _a1 error) *DataProxyServiceServer_CreateUploadLocation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataProxyServiceServer_CreateUploadLocation_Call) RunAndReturn(run func(context.Context, *service.CreateUploadLocationRequest) (*service.CreateUploadLocationResponse, error)) *DataProxyServiceServer_CreateUploadLocation_Call {
	_c.Call.Return(run)
	return _c
}

// GetData provides a mock function with given fields: _a0, _a1
func (_m *DataProxyServiceServer) GetData(_a0 context.Context, _a1 *service.GetDataRequest) (*service.GetDataResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetData")
	}

	var r0 *service.GetDataResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *service.GetDataRequest) (*service.GetDataResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *service.GetDataRequest) *service.GetDataResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.GetDataResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *service.GetDataRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataProxyServiceServer_GetData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetData'
type DataProxyServiceServer_GetData_Call struct {
	*mock.Call
}

// GetData is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *service.GetDataRequest
func (_e *DataProxyServiceServer_Expecter) GetData(_a0 interface{}, _a1 interface{}) *DataProxyServiceServer_GetData_Call {
	return &DataProxyServiceServer_GetData_Call{Call: _e.mock.On("GetData", _a0, _a1)}
}

func (_c *DataProxyServiceServer_GetData_Call) Run(run func(_a0 context.Context, _a1 *service.GetDataRequest)) *DataProxyServiceServer_GetData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*service.GetDataRequest))
	})
	return _c
}

func (_c *DataProxyServiceServer_GetData_Call) Return(_a0 *service.GetDataResponse, _a1 error) *DataProxyServiceServer_GetData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DataProxyServiceServer_GetData_Call) RunAndReturn(run func(context.Context, *service.GetDataRequest) (*service.GetDataResponse, error)) *DataProxyServiceServer_GetData_Call {
	_c.Call.Return(run)
	return _c
}

// NewDataProxyServiceServer creates a new instance of DataProxyServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataProxyServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataProxyServiceServer {
	mock := &DataProxyServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
