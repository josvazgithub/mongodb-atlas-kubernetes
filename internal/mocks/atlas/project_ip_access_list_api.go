// Code generated by mockery. DO NOT EDIT.

package atlas

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20231115004/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// ProjectIPAccessListApiMock is an autogenerated mock type for the ProjectIPAccessListApi type
type ProjectIPAccessListApiMock struct {
	mock.Mock
}

type ProjectIPAccessListApiMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectIPAccessListApiMock) EXPECT() *ProjectIPAccessListApiMock_Expecter {
	return &ProjectIPAccessListApiMock_Expecter{mock: &_m.Mock}
}

// CreateProjectIpAccessList provides a mock function with given fields: ctx, groupId, networkPermissionEntry
func (_m *ProjectIPAccessListApiMock) CreateProjectIpAccessList(ctx context.Context, groupId string, networkPermissionEntry *[]admin.NetworkPermissionEntry) admin.CreateProjectIpAccessListApiRequest {
	ret := _m.Called(ctx, groupId, networkPermissionEntry)

	if len(ret) == 0 {
		panic("no return value specified for CreateProjectIpAccessList")
	}

	var r0 admin.CreateProjectIpAccessListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *[]admin.NetworkPermissionEntry) admin.CreateProjectIpAccessListApiRequest); ok {
		r0 = rf(ctx, groupId, networkPermissionEntry)
	} else {
		r0 = ret.Get(0).(admin.CreateProjectIpAccessListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProjectIpAccessList'
type ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call struct {
	*mock.Call
}

// CreateProjectIpAccessList is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - networkPermissionEntry *[]admin.NetworkPermissionEntry
func (_e *ProjectIPAccessListApiMock_Expecter) CreateProjectIpAccessList(ctx interface{}, groupId interface{}, networkPermissionEntry interface{}) *ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call {
	return &ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call{Call: _e.mock.On("CreateProjectIpAccessList", ctx, groupId, networkPermissionEntry)}
}

func (_c *ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call) Run(run func(ctx context.Context, groupId string, networkPermissionEntry *[]admin.NetworkPermissionEntry)) *ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*[]admin.NetworkPermissionEntry))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call) Return(_a0 admin.CreateProjectIpAccessListApiRequest) *ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call) RunAndReturn(run func(context.Context, string, *[]admin.NetworkPermissionEntry) admin.CreateProjectIpAccessListApiRequest) *ProjectIPAccessListApiMock_CreateProjectIpAccessList_Call {
	_c.Call.Return(run)
	return _c
}

// CreateProjectIpAccessListExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApiMock) CreateProjectIpAccessListExecute(r admin.CreateProjectIpAccessListApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateProjectIpAccessListExecute")
	}

	var r0 *admin.PaginatedNetworkAccess
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateProjectIpAccessListApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateProjectIpAccessListApiRequest) *admin.PaginatedNetworkAccess); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedNetworkAccess)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateProjectIpAccessListApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateProjectIpAccessListApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProjectIpAccessListExecute'
type ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call struct {
	*mock.Call
}

// CreateProjectIpAccessListExecute is a helper method to define mock.On call
//   - r admin.CreateProjectIpAccessListApiRequest
func (_e *ProjectIPAccessListApiMock_Expecter) CreateProjectIpAccessListExecute(r interface{}) *ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call {
	return &ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call{Call: _e.mock.On("CreateProjectIpAccessListExecute", r)}
}

func (_c *ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call) Run(run func(r admin.CreateProjectIpAccessListApiRequest)) *ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateProjectIpAccessListApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call) Return(_a0 *admin.PaginatedNetworkAccess, _a1 *http.Response, _a2 error) *ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call) RunAndReturn(run func(admin.CreateProjectIpAccessListApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error)) *ProjectIPAccessListApiMock_CreateProjectIpAccessListExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateProjectIpAccessListWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApiMock) CreateProjectIpAccessListWithParams(ctx context.Context, args *admin.CreateProjectIpAccessListApiParams) admin.CreateProjectIpAccessListApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateProjectIpAccessListWithParams")
	}

	var r0 admin.CreateProjectIpAccessListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateProjectIpAccessListApiParams) admin.CreateProjectIpAccessListApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateProjectIpAccessListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProjectIpAccessListWithParams'
type ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call struct {
	*mock.Call
}

// CreateProjectIpAccessListWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateProjectIpAccessListApiParams
func (_e *ProjectIPAccessListApiMock_Expecter) CreateProjectIpAccessListWithParams(ctx interface{}, args interface{}) *ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call {
	return &ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call{Call: _e.mock.On("CreateProjectIpAccessListWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateProjectIpAccessListApiParams)) *ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateProjectIpAccessListApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call) Return(_a0 admin.CreateProjectIpAccessListApiRequest) *ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateProjectIpAccessListApiParams) admin.CreateProjectIpAccessListApiRequest) *ProjectIPAccessListApiMock_CreateProjectIpAccessListWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProjectIpAccessList provides a mock function with given fields: ctx, groupId, entryValue
func (_m *ProjectIPAccessListApiMock) DeleteProjectIpAccessList(ctx context.Context, groupId string, entryValue string) admin.DeleteProjectIpAccessListApiRequest {
	ret := _m.Called(ctx, groupId, entryValue)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProjectIpAccessList")
	}

	var r0 admin.DeleteProjectIpAccessListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.DeleteProjectIpAccessListApiRequest); ok {
		r0 = rf(ctx, groupId, entryValue)
	} else {
		r0 = ret.Get(0).(admin.DeleteProjectIpAccessListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProjectIpAccessList'
type ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call struct {
	*mock.Call
}

// DeleteProjectIpAccessList is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - entryValue string
func (_e *ProjectIPAccessListApiMock_Expecter) DeleteProjectIpAccessList(ctx interface{}, groupId interface{}, entryValue interface{}) *ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call {
	return &ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call{Call: _e.mock.On("DeleteProjectIpAccessList", ctx, groupId, entryValue)}
}

func (_c *ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call) Run(run func(ctx context.Context, groupId string, entryValue string)) *ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call) Return(_a0 admin.DeleteProjectIpAccessListApiRequest) *ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call) RunAndReturn(run func(context.Context, string, string) admin.DeleteProjectIpAccessListApiRequest) *ProjectIPAccessListApiMock_DeleteProjectIpAccessList_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProjectIpAccessListExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApiMock) DeleteProjectIpAccessListExecute(r admin.DeleteProjectIpAccessListApiRequest) (map[string]interface{}, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProjectIpAccessListExecute")
	}

	var r0 map[string]interface{}
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DeleteProjectIpAccessListApiRequest) (map[string]interface{}, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeleteProjectIpAccessListApiRequest) map[string]interface{}); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeleteProjectIpAccessListApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DeleteProjectIpAccessListApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProjectIpAccessListExecute'
type ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call struct {
	*mock.Call
}

// DeleteProjectIpAccessListExecute is a helper method to define mock.On call
//   - r admin.DeleteProjectIpAccessListApiRequest
func (_e *ProjectIPAccessListApiMock_Expecter) DeleteProjectIpAccessListExecute(r interface{}) *ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call {
	return &ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call{Call: _e.mock.On("DeleteProjectIpAccessListExecute", r)}
}

func (_c *ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call) Run(run func(r admin.DeleteProjectIpAccessListApiRequest)) *ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeleteProjectIpAccessListApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call) Return(_a0 map[string]interface{}, _a1 *http.Response, _a2 error) *ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call) RunAndReturn(run func(admin.DeleteProjectIpAccessListApiRequest) (map[string]interface{}, *http.Response, error)) *ProjectIPAccessListApiMock_DeleteProjectIpAccessListExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProjectIpAccessListWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApiMock) DeleteProjectIpAccessListWithParams(ctx context.Context, args *admin.DeleteProjectIpAccessListApiParams) admin.DeleteProjectIpAccessListApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProjectIpAccessListWithParams")
	}

	var r0 admin.DeleteProjectIpAccessListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeleteProjectIpAccessListApiParams) admin.DeleteProjectIpAccessListApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeleteProjectIpAccessListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProjectIpAccessListWithParams'
type ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call struct {
	*mock.Call
}

// DeleteProjectIpAccessListWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeleteProjectIpAccessListApiParams
func (_e *ProjectIPAccessListApiMock_Expecter) DeleteProjectIpAccessListWithParams(ctx interface{}, args interface{}) *ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call {
	return &ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call{Call: _e.mock.On("DeleteProjectIpAccessListWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call) Run(run func(ctx context.Context, args *admin.DeleteProjectIpAccessListApiParams)) *ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeleteProjectIpAccessListApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call) Return(_a0 admin.DeleteProjectIpAccessListApiRequest) *ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeleteProjectIpAccessListApiParams) admin.DeleteProjectIpAccessListApiRequest) *ProjectIPAccessListApiMock_DeleteProjectIpAccessListWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpAccessListStatus provides a mock function with given fields: ctx, groupId, entryValue
func (_m *ProjectIPAccessListApiMock) GetProjectIpAccessListStatus(ctx context.Context, groupId string, entryValue string) admin.GetProjectIpAccessListStatusApiRequest {
	ret := _m.Called(ctx, groupId, entryValue)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpAccessListStatus")
	}

	var r0 admin.GetProjectIpAccessListStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetProjectIpAccessListStatusApiRequest); ok {
		r0 = rf(ctx, groupId, entryValue)
	} else {
		r0 = ret.Get(0).(admin.GetProjectIpAccessListStatusApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpAccessListStatus'
type ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call struct {
	*mock.Call
}

// GetProjectIpAccessListStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - entryValue string
func (_e *ProjectIPAccessListApiMock_Expecter) GetProjectIpAccessListStatus(ctx interface{}, groupId interface{}, entryValue interface{}) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call {
	return &ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call{Call: _e.mock.On("GetProjectIpAccessListStatus", ctx, groupId, entryValue)}
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call) Run(run func(ctx context.Context, groupId string, entryValue string)) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call) Return(_a0 admin.GetProjectIpAccessListStatusApiRequest) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call) RunAndReturn(run func(context.Context, string, string) admin.GetProjectIpAccessListStatusApiRequest) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpAccessListStatusExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApiMock) GetProjectIpAccessListStatusExecute(r admin.GetProjectIpAccessListStatusApiRequest) (*admin.NetworkPermissionEntryStatus, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpAccessListStatusExecute")
	}

	var r0 *admin.NetworkPermissionEntryStatus
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetProjectIpAccessListStatusApiRequest) (*admin.NetworkPermissionEntryStatus, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetProjectIpAccessListStatusApiRequest) *admin.NetworkPermissionEntryStatus); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.NetworkPermissionEntryStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetProjectIpAccessListStatusApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetProjectIpAccessListStatusApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpAccessListStatusExecute'
type ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call struct {
	*mock.Call
}

// GetProjectIpAccessListStatusExecute is a helper method to define mock.On call
//   - r admin.GetProjectIpAccessListStatusApiRequest
func (_e *ProjectIPAccessListApiMock_Expecter) GetProjectIpAccessListStatusExecute(r interface{}) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call {
	return &ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call{Call: _e.mock.On("GetProjectIpAccessListStatusExecute", r)}
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call) Run(run func(r admin.GetProjectIpAccessListStatusApiRequest)) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetProjectIpAccessListStatusApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call) Return(_a0 *admin.NetworkPermissionEntryStatus, _a1 *http.Response, _a2 error) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call) RunAndReturn(run func(admin.GetProjectIpAccessListStatusApiRequest) (*admin.NetworkPermissionEntryStatus, *http.Response, error)) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpAccessListStatusWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApiMock) GetProjectIpAccessListStatusWithParams(ctx context.Context, args *admin.GetProjectIpAccessListStatusApiParams) admin.GetProjectIpAccessListStatusApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpAccessListStatusWithParams")
	}

	var r0 admin.GetProjectIpAccessListStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetProjectIpAccessListStatusApiParams) admin.GetProjectIpAccessListStatusApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetProjectIpAccessListStatusApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpAccessListStatusWithParams'
type ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call struct {
	*mock.Call
}

// GetProjectIpAccessListStatusWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetProjectIpAccessListStatusApiParams
func (_e *ProjectIPAccessListApiMock_Expecter) GetProjectIpAccessListStatusWithParams(ctx interface{}, args interface{}) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call {
	return &ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call{Call: _e.mock.On("GetProjectIpAccessListStatusWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call) Run(run func(ctx context.Context, args *admin.GetProjectIpAccessListStatusApiParams)) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetProjectIpAccessListStatusApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call) Return(_a0 admin.GetProjectIpAccessListStatusApiRequest) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetProjectIpAccessListStatusApiParams) admin.GetProjectIpAccessListStatusApiRequest) *ProjectIPAccessListApiMock_GetProjectIpAccessListStatusWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpList provides a mock function with given fields: ctx, groupId, entryValue
func (_m *ProjectIPAccessListApiMock) GetProjectIpList(ctx context.Context, groupId string, entryValue string) admin.GetProjectIpListApiRequest {
	ret := _m.Called(ctx, groupId, entryValue)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpList")
	}

	var r0 admin.GetProjectIpListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetProjectIpListApiRequest); ok {
		r0 = rf(ctx, groupId, entryValue)
	} else {
		r0 = ret.Get(0).(admin.GetProjectIpListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_GetProjectIpList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpList'
type ProjectIPAccessListApiMock_GetProjectIpList_Call struct {
	*mock.Call
}

// GetProjectIpList is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - entryValue string
func (_e *ProjectIPAccessListApiMock_Expecter) GetProjectIpList(ctx interface{}, groupId interface{}, entryValue interface{}) *ProjectIPAccessListApiMock_GetProjectIpList_Call {
	return &ProjectIPAccessListApiMock_GetProjectIpList_Call{Call: _e.mock.On("GetProjectIpList", ctx, groupId, entryValue)}
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpList_Call) Run(run func(ctx context.Context, groupId string, entryValue string)) *ProjectIPAccessListApiMock_GetProjectIpList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpList_Call) Return(_a0 admin.GetProjectIpListApiRequest) *ProjectIPAccessListApiMock_GetProjectIpList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpList_Call) RunAndReturn(run func(context.Context, string, string) admin.GetProjectIpListApiRequest) *ProjectIPAccessListApiMock_GetProjectIpList_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpListExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApiMock) GetProjectIpListExecute(r admin.GetProjectIpListApiRequest) (*admin.NetworkPermissionEntry, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpListExecute")
	}

	var r0 *admin.NetworkPermissionEntry
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetProjectIpListApiRequest) (*admin.NetworkPermissionEntry, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetProjectIpListApiRequest) *admin.NetworkPermissionEntry); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.NetworkPermissionEntry)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetProjectIpListApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetProjectIpListApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApiMock_GetProjectIpListExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpListExecute'
type ProjectIPAccessListApiMock_GetProjectIpListExecute_Call struct {
	*mock.Call
}

// GetProjectIpListExecute is a helper method to define mock.On call
//   - r admin.GetProjectIpListApiRequest
func (_e *ProjectIPAccessListApiMock_Expecter) GetProjectIpListExecute(r interface{}) *ProjectIPAccessListApiMock_GetProjectIpListExecute_Call {
	return &ProjectIPAccessListApiMock_GetProjectIpListExecute_Call{Call: _e.mock.On("GetProjectIpListExecute", r)}
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpListExecute_Call) Run(run func(r admin.GetProjectIpListApiRequest)) *ProjectIPAccessListApiMock_GetProjectIpListExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetProjectIpListApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpListExecute_Call) Return(_a0 *admin.NetworkPermissionEntry, _a1 *http.Response, _a2 error) *ProjectIPAccessListApiMock_GetProjectIpListExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpListExecute_Call) RunAndReturn(run func(admin.GetProjectIpListApiRequest) (*admin.NetworkPermissionEntry, *http.Response, error)) *ProjectIPAccessListApiMock_GetProjectIpListExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpListWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApiMock) GetProjectIpListWithParams(ctx context.Context, args *admin.GetProjectIpListApiParams) admin.GetProjectIpListApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpListWithParams")
	}

	var r0 admin.GetProjectIpListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetProjectIpListApiParams) admin.GetProjectIpListApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetProjectIpListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpListWithParams'
type ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call struct {
	*mock.Call
}

// GetProjectIpListWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetProjectIpListApiParams
func (_e *ProjectIPAccessListApiMock_Expecter) GetProjectIpListWithParams(ctx interface{}, args interface{}) *ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call {
	return &ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call{Call: _e.mock.On("GetProjectIpListWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call) Run(run func(ctx context.Context, args *admin.GetProjectIpListApiParams)) *ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetProjectIpListApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call) Return(_a0 admin.GetProjectIpListApiRequest) *ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetProjectIpListApiParams) admin.GetProjectIpListApiRequest) *ProjectIPAccessListApiMock_GetProjectIpListWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectIpAccessLists provides a mock function with given fields: ctx, groupId
func (_m *ProjectIPAccessListApiMock) ListProjectIpAccessLists(ctx context.Context, groupId string) admin.ListProjectIpAccessListsApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectIpAccessLists")
	}

	var r0 admin.ListProjectIpAccessListsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListProjectIpAccessListsApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ListProjectIpAccessListsApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectIpAccessLists'
type ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call struct {
	*mock.Call
}

// ListProjectIpAccessLists is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *ProjectIPAccessListApiMock_Expecter) ListProjectIpAccessLists(ctx interface{}, groupId interface{}) *ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call {
	return &ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call{Call: _e.mock.On("ListProjectIpAccessLists", ctx, groupId)}
}

func (_c *ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call) Run(run func(ctx context.Context, groupId string)) *ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call) Return(_a0 admin.ListProjectIpAccessListsApiRequest) *ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call) RunAndReturn(run func(context.Context, string) admin.ListProjectIpAccessListsApiRequest) *ProjectIPAccessListApiMock_ListProjectIpAccessLists_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectIpAccessListsExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApiMock) ListProjectIpAccessListsExecute(r admin.ListProjectIpAccessListsApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectIpAccessListsExecute")
	}

	var r0 *admin.PaginatedNetworkAccess
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListProjectIpAccessListsApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListProjectIpAccessListsApiRequest) *admin.PaginatedNetworkAccess); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedNetworkAccess)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListProjectIpAccessListsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListProjectIpAccessListsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectIpAccessListsExecute'
type ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call struct {
	*mock.Call
}

// ListProjectIpAccessListsExecute is a helper method to define mock.On call
//   - r admin.ListProjectIpAccessListsApiRequest
func (_e *ProjectIPAccessListApiMock_Expecter) ListProjectIpAccessListsExecute(r interface{}) *ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call {
	return &ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call{Call: _e.mock.On("ListProjectIpAccessListsExecute", r)}
}

func (_c *ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call) Run(run func(r admin.ListProjectIpAccessListsApiRequest)) *ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListProjectIpAccessListsApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call) Return(_a0 *admin.PaginatedNetworkAccess, _a1 *http.Response, _a2 error) *ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call) RunAndReturn(run func(admin.ListProjectIpAccessListsApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error)) *ProjectIPAccessListApiMock_ListProjectIpAccessListsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectIpAccessListsWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApiMock) ListProjectIpAccessListsWithParams(ctx context.Context, args *admin.ListProjectIpAccessListsApiParams) admin.ListProjectIpAccessListsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectIpAccessListsWithParams")
	}

	var r0 admin.ListProjectIpAccessListsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListProjectIpAccessListsApiParams) admin.ListProjectIpAccessListsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListProjectIpAccessListsApiRequest)
	}

	return r0
}

// ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectIpAccessListsWithParams'
type ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call struct {
	*mock.Call
}

// ListProjectIpAccessListsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListProjectIpAccessListsApiParams
func (_e *ProjectIPAccessListApiMock_Expecter) ListProjectIpAccessListsWithParams(ctx interface{}, args interface{}) *ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call {
	return &ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call{Call: _e.mock.On("ListProjectIpAccessListsWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call) Run(run func(ctx context.Context, args *admin.ListProjectIpAccessListsApiParams)) *ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListProjectIpAccessListsApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call) Return(_a0 admin.ListProjectIpAccessListsApiRequest) *ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListProjectIpAccessListsApiParams) admin.ListProjectIpAccessListsApiRequest) *ProjectIPAccessListApiMock_ListProjectIpAccessListsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectIPAccessListApiMock creates a new instance of ProjectIPAccessListApiMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectIPAccessListApiMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectIPAccessListApiMock {
	mock := &ProjectIPAccessListApiMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
