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

package providers

import (
	config "github.com/elastic/cloudbeat/config"
	kubernetes "k8s.io/client-go/kubernetes"

	mock "github.com/stretchr/testify/mock"
)

// MockKubernetesClusterNameProviderApi is an autogenerated mock type for the KubernetesClusterNameProviderApi type
type MockKubernetesClusterNameProviderApi struct {
	mock.Mock
}

type MockKubernetesClusterNameProviderApi_Expecter struct {
	mock *mock.Mock
}

func (_m *MockKubernetesClusterNameProviderApi) EXPECT() *MockKubernetesClusterNameProviderApi_Expecter {
	return &MockKubernetesClusterNameProviderApi_Expecter{mock: &_m.Mock}
}

// GetClusterName provides a mock function with given fields: cfg, client
func (_m *MockKubernetesClusterNameProviderApi) GetClusterName(cfg *config.Config, client kubernetes.Interface) (string, error) {
	ret := _m.Called(cfg, client)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*config.Config, kubernetes.Interface) (string, error)); ok {
		return rf(cfg, client)
	}
	if rf, ok := ret.Get(0).(func(*config.Config, kubernetes.Interface) string); ok {
		r0 = rf(cfg, client)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*config.Config, kubernetes.Interface) error); ok {
		r1 = rf(cfg, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockKubernetesClusterNameProviderApi_GetClusterName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClusterName'
type MockKubernetesClusterNameProviderApi_GetClusterName_Call struct {
	*mock.Call
}

// GetClusterName is a helper method to define mock.On call
//   - cfg *config.Config
//   - client kubernetes.Interface
func (_e *MockKubernetesClusterNameProviderApi_Expecter) GetClusterName(cfg interface{}, client interface{}) *MockKubernetesClusterNameProviderApi_GetClusterName_Call {
	return &MockKubernetesClusterNameProviderApi_GetClusterName_Call{Call: _e.mock.On("GetClusterName", cfg, client)}
}

func (_c *MockKubernetesClusterNameProviderApi_GetClusterName_Call) Run(run func(cfg *config.Config, client kubernetes.Interface)) *MockKubernetesClusterNameProviderApi_GetClusterName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*config.Config), args[1].(kubernetes.Interface))
	})
	return _c
}

func (_c *MockKubernetesClusterNameProviderApi_GetClusterName_Call) Return(_a0 string, _a1 error) *MockKubernetesClusterNameProviderApi_GetClusterName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockKubernetesClusterNameProviderApi_GetClusterName_Call) RunAndReturn(run func(*config.Config, kubernetes.Interface) (string, error)) *MockKubernetesClusterNameProviderApi_GetClusterName_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockKubernetesClusterNameProviderApi interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockKubernetesClusterNameProviderApi creates a new instance of MockKubernetesClusterNameProviderApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockKubernetesClusterNameProviderApi(t mockConstructorTestingTNewMockKubernetesClusterNameProviderApi) *MockKubernetesClusterNameProviderApi {
	mock := &MockKubernetesClusterNameProviderApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
