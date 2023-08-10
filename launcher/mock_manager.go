// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.24.0. DO NOT EDIT.

package launcher

import (
	client "github.com/elastic/elastic-agent-client/v7/pkg/client"
	config "github.com/elastic/elastic-agent-libs/config"

	management "github.com/elastic/beats/v7/libbeat/management"

	mock "github.com/stretchr/testify/mock"
)

// MockManager is an autogenerated mock type for the Manager type
type MockManager struct {
	mock.Mock
}

type MockManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockManager) EXPECT() *MockManager_Expecter {
	return &MockManager_Expecter{mock: &_m.Mock}
}

// CheckRawConfig provides a mock function with given fields: cfg
func (_m *MockManager) CheckRawConfig(cfg *config.C) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*config.C) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockManager_CheckRawConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckRawConfig'
type MockManager_CheckRawConfig_Call struct {
	*mock.Call
}

// CheckRawConfig is a helper method to define mock.On call
//   - cfg *config.C
func (_e *MockManager_Expecter) CheckRawConfig(cfg interface{}) *MockManager_CheckRawConfig_Call {
	return &MockManager_CheckRawConfig_Call{Call: _e.mock.On("CheckRawConfig", cfg)}
}

func (_c *MockManager_CheckRawConfig_Call) Run(run func(cfg *config.C)) *MockManager_CheckRawConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*config.C))
	})
	return _c
}

func (_c *MockManager_CheckRawConfig_Call) Return(_a0 error) *MockManager_CheckRawConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManager_CheckRawConfig_Call) RunAndReturn(run func(*config.C) error) *MockManager_CheckRawConfig_Call {
	_c.Call.Return(run)
	return _c
}

// Enabled provides a mock function with given fields:
func (_m *MockManager) Enabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockManager_Enabled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Enabled'
type MockManager_Enabled_Call struct {
	*mock.Call
}

// Enabled is a helper method to define mock.On call
func (_e *MockManager_Expecter) Enabled() *MockManager_Enabled_Call {
	return &MockManager_Enabled_Call{Call: _e.mock.On("Enabled")}
}

func (_c *MockManager_Enabled_Call) Run(run func()) *MockManager_Enabled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockManager_Enabled_Call) Return(_a0 bool) *MockManager_Enabled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManager_Enabled_Call) RunAndReturn(run func() bool) *MockManager_Enabled_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterAction provides a mock function with given fields: action
func (_m *MockManager) RegisterAction(action client.Action) {
	_m.Called(action)
}

// MockManager_RegisterAction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterAction'
type MockManager_RegisterAction_Call struct {
	*mock.Call
}

// RegisterAction is a helper method to define mock.On call
//   - action client.Action
func (_e *MockManager_Expecter) RegisterAction(action interface{}) *MockManager_RegisterAction_Call {
	return &MockManager_RegisterAction_Call{Call: _e.mock.On("RegisterAction", action)}
}

func (_c *MockManager_RegisterAction_Call) Run(run func(action client.Action)) *MockManager_RegisterAction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.Action))
	})
	return _c
}

func (_c *MockManager_RegisterAction_Call) Return() *MockManager_RegisterAction_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_RegisterAction_Call) RunAndReturn(run func(client.Action)) *MockManager_RegisterAction_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterDiagnosticHook provides a mock function with given fields: name, description, filename, contentType, hook
func (_m *MockManager) RegisterDiagnosticHook(name string, description string, filename string, contentType string, hook client.DiagnosticHook) {
	_m.Called(name, description, filename, contentType, hook)
}

// MockManager_RegisterDiagnosticHook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterDiagnosticHook'
type MockManager_RegisterDiagnosticHook_Call struct {
	*mock.Call
}

// RegisterDiagnosticHook is a helper method to define mock.On call
//   - name string
//   - description string
//   - filename string
//   - contentType string
//   - hook client.DiagnosticHook
func (_e *MockManager_Expecter) RegisterDiagnosticHook(name interface{}, description interface{}, filename interface{}, contentType interface{}, hook interface{}) *MockManager_RegisterDiagnosticHook_Call {
	return &MockManager_RegisterDiagnosticHook_Call{Call: _e.mock.On("RegisterDiagnosticHook", name, description, filename, contentType, hook)}
}

func (_c *MockManager_RegisterDiagnosticHook_Call) Run(run func(name string, description string, filename string, contentType string, hook client.DiagnosticHook)) *MockManager_RegisterDiagnosticHook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(client.DiagnosticHook))
	})
	return _c
}

func (_c *MockManager_RegisterDiagnosticHook_Call) Return() *MockManager_RegisterDiagnosticHook_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_RegisterDiagnosticHook_Call) RunAndReturn(run func(string, string, string, string, client.DiagnosticHook)) *MockManager_RegisterDiagnosticHook_Call {
	_c.Call.Return(run)
	return _c
}

// SetPayload provides a mock function with given fields: _a0
func (_m *MockManager) SetPayload(_a0 map[string]interface{}) {
	_m.Called(_a0)
}

// MockManager_SetPayload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPayload'
type MockManager_SetPayload_Call struct {
	*mock.Call
}

// SetPayload is a helper method to define mock.On call
//   - _a0 map[string]interface{}
func (_e *MockManager_Expecter) SetPayload(_a0 interface{}) *MockManager_SetPayload_Call {
	return &MockManager_SetPayload_Call{Call: _e.mock.On("SetPayload", _a0)}
}

func (_c *MockManager_SetPayload_Call) Run(run func(_a0 map[string]interface{})) *MockManager_SetPayload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]interface{}))
	})
	return _c
}

func (_c *MockManager_SetPayload_Call) Return() *MockManager_SetPayload_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_SetPayload_Call) RunAndReturn(run func(map[string]interface{})) *MockManager_SetPayload_Call {
	_c.Call.Return(run)
	return _c
}

// SetStopCallback provides a mock function with given fields: f
func (_m *MockManager) SetStopCallback(f func()) {
	_m.Called(f)
}

// MockManager_SetStopCallback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStopCallback'
type MockManager_SetStopCallback_Call struct {
	*mock.Call
}

// SetStopCallback is a helper method to define mock.On call
//   - f func()
func (_e *MockManager_Expecter) SetStopCallback(f interface{}) *MockManager_SetStopCallback_Call {
	return &MockManager_SetStopCallback_Call{Call: _e.mock.On("SetStopCallback", f)}
}

func (_c *MockManager_SetStopCallback_Call) Run(run func(f func())) *MockManager_SetStopCallback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func()))
	})
	return _c
}

func (_c *MockManager_SetStopCallback_Call) Return() *MockManager_SetStopCallback_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_SetStopCallback_Call) RunAndReturn(run func(func())) *MockManager_SetStopCallback_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockManager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockManager_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockManager_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockManager_Expecter) Start() *MockManager_Start_Call {
	return &MockManager_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockManager_Start_Call) Run(run func()) *MockManager_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockManager_Start_Call) Return(_a0 error) *MockManager_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManager_Start_Call) RunAndReturn(run func() error) *MockManager_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockManager) Stop() {
	_m.Called()
}

// MockManager_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockManager_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockManager_Expecter) Stop() *MockManager_Stop_Call {
	return &MockManager_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockManager_Stop_Call) Run(run func()) *MockManager_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockManager_Stop_Call) Return() *MockManager_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_Stop_Call) RunAndReturn(run func()) *MockManager_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// UnregisterAction provides a mock function with given fields: action
func (_m *MockManager) UnregisterAction(action client.Action) {
	_m.Called(action)
}

// MockManager_UnregisterAction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnregisterAction'
type MockManager_UnregisterAction_Call struct {
	*mock.Call
}

// UnregisterAction is a helper method to define mock.On call
//   - action client.Action
func (_e *MockManager_Expecter) UnregisterAction(action interface{}) *MockManager_UnregisterAction_Call {
	return &MockManager_UnregisterAction_Call{Call: _e.mock.On("UnregisterAction", action)}
}

func (_c *MockManager_UnregisterAction_Call) Run(run func(action client.Action)) *MockManager_UnregisterAction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.Action))
	})
	return _c
}

func (_c *MockManager_UnregisterAction_Call) Return() *MockManager_UnregisterAction_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_UnregisterAction_Call) RunAndReturn(run func(client.Action)) *MockManager_UnregisterAction_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatus provides a mock function with given fields: status, msg
func (_m *MockManager) UpdateStatus(status management.Status, msg string) {
	_m.Called(status, msg)
}

// MockManager_UpdateStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatus'
type MockManager_UpdateStatus_Call struct {
	*mock.Call
}

// UpdateStatus is a helper method to define mock.On call
//   - status management.Status
//   - msg string
func (_e *MockManager_Expecter) UpdateStatus(status interface{}, msg interface{}) *MockManager_UpdateStatus_Call {
	return &MockManager_UpdateStatus_Call{Call: _e.mock.On("UpdateStatus", status, msg)}
}

func (_c *MockManager_UpdateStatus_Call) Run(run func(status management.Status, msg string)) *MockManager_UpdateStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(management.Status), args[1].(string))
	})
	return _c
}

func (_c *MockManager_UpdateStatus_Call) Return() *MockManager_UpdateStatus_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManager_UpdateStatus_Call) RunAndReturn(run func(management.Status, string)) *MockManager_UpdateStatus_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockManager creates a new instance of MockManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockManager(t mockConstructorTestingTNewMockManager) *MockManager {
	mock := &MockManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
