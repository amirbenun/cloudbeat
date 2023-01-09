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

// Code generated by mockery v2.15.0. DO NOT EDIT.
package awslib

import (
	aws_sdk_go_v2aws "github.com/aws/aws-sdk-go-v2/aws"
	aws "github.com/elastic/beats/v7/x-pack/libbeat/common/aws"

	context "context"

	logp "github.com/elastic/elastic-agent-libs/logp"

	mock "github.com/stretchr/testify/mock"
)

// MockConfigProviderAPI is an autogenerated mock type for the ConfigProviderAPI type
type MockConfigProviderAPI struct {
	mock.Mock
}

type MockConfigProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConfigProviderAPI) EXPECT() *MockConfigProviderAPI_Expecter {
	return &MockConfigProviderAPI_Expecter{mock: &_m.Mock}
}

// InitializeAWSConfig provides a mock function with given fields: ctx, cfg, log, useDefaultRegion
func (_m *MockConfigProviderAPI) InitializeAWSConfig(ctx context.Context, cfg aws.ConfigAWS, log *logp.Logger, useDefaultRegion bool) (aws_sdk_go_v2aws.Config, error) {
	ret := _m.Called(ctx, cfg, log, useDefaultRegion)

	var r0 aws_sdk_go_v2aws.Config
	if rf, ok := ret.Get(0).(func(context.Context, aws.ConfigAWS, *logp.Logger, bool) aws_sdk_go_v2aws.Config); ok {
		r0 = rf(ctx, cfg, log, useDefaultRegion)
	} else {
		r0 = ret.Get(0).(aws_sdk_go_v2aws.Config)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, aws.ConfigAWS, *logp.Logger, bool) error); ok {
		r1 = rf(ctx, cfg, log, useDefaultRegion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConfigProviderAPI_InitializeAWSConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitializeAWSConfig'
type MockConfigProviderAPI_InitializeAWSConfig_Call struct {
	*mock.Call
}

// InitializeAWSConfig is a helper method to define mock.On call
//   - ctx context.Context
//   - cfg aws.ConfigAWS
//   - log *logp.Logger
//   - useDefaultRegion bool
func (_e *MockConfigProviderAPI_Expecter) InitializeAWSConfig(ctx interface{}, cfg interface{}, log interface{}, useDefaultRegion interface{}) *MockConfigProviderAPI_InitializeAWSConfig_Call {
	return &MockConfigProviderAPI_InitializeAWSConfig_Call{Call: _e.mock.On("InitializeAWSConfig", ctx, cfg, log, useDefaultRegion)}
}

func (_c *MockConfigProviderAPI_InitializeAWSConfig_Call) Run(run func(ctx context.Context, cfg aws.ConfigAWS, log *logp.Logger, useDefaultRegion bool)) *MockConfigProviderAPI_InitializeAWSConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(aws.ConfigAWS), args[2].(*logp.Logger), args[3].(bool))
	})
	return _c
}

func (_c *MockConfigProviderAPI_InitializeAWSConfig_Call) Return(_a0 aws_sdk_go_v2aws.Config, _a1 error) *MockConfigProviderAPI_InitializeAWSConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockConfigProviderAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockConfigProviderAPI creates a new instance of MockConfigProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockConfigProviderAPI(t mockConstructorTestingTNewMockConfigProviderAPI) *MockConfigProviderAPI {
	mock := &MockConfigProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
