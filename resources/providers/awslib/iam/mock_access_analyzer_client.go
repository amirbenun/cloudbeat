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

package iam

import (
	context "context"

	accessanalyzer "github.com/aws/aws-sdk-go-v2/service/accessanalyzer"

	mock "github.com/stretchr/testify/mock"
)

// MockAccessAnalyzerClient is an autogenerated mock type for the AccessAnalyzerClient type
type MockAccessAnalyzerClient struct {
	mock.Mock
}

type MockAccessAnalyzerClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccessAnalyzerClient) EXPECT() *MockAccessAnalyzerClient_Expecter {
	return &MockAccessAnalyzerClient_Expecter{mock: &_m.Mock}
}

// ListAnalyzers provides a mock function with given fields: ctx, params, optFns
func (_m *MockAccessAnalyzerClient) ListAnalyzers(ctx context.Context, params *accessanalyzer.ListAnalyzersInput, optFns ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *accessanalyzer.ListAnalyzersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *accessanalyzer.ListAnalyzersInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *accessanalyzer.ListAnalyzersInput, ...func(*accessanalyzer.Options)) *accessanalyzer.ListAnalyzersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*accessanalyzer.ListAnalyzersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *accessanalyzer.ListAnalyzersInput, ...func(*accessanalyzer.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessAnalyzerClient_ListAnalyzers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAnalyzers'
type MockAccessAnalyzerClient_ListAnalyzers_Call struct {
	*mock.Call
}

// ListAnalyzers is a helper method to define mock.On call
//   - ctx context.Context
//   - params *accessanalyzer.ListAnalyzersInput
//   - optFns ...func(*accessanalyzer.Options)
func (_e *MockAccessAnalyzerClient_Expecter) ListAnalyzers(ctx interface{}, params interface{}, optFns ...interface{}) *MockAccessAnalyzerClient_ListAnalyzers_Call {
	return &MockAccessAnalyzerClient_ListAnalyzers_Call{Call: _e.mock.On("ListAnalyzers",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockAccessAnalyzerClient_ListAnalyzers_Call) Run(run func(ctx context.Context, params *accessanalyzer.ListAnalyzersInput, optFns ...func(*accessanalyzer.Options))) *MockAccessAnalyzerClient_ListAnalyzers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*accessanalyzer.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*accessanalyzer.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*accessanalyzer.ListAnalyzersInput), variadicArgs...)
	})
	return _c
}

func (_c *MockAccessAnalyzerClient_ListAnalyzers_Call) Return(_a0 *accessanalyzer.ListAnalyzersOutput, _a1 error) *MockAccessAnalyzerClient_ListAnalyzers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccessAnalyzerClient_ListAnalyzers_Call) RunAndReturn(run func(context.Context, *accessanalyzer.ListAnalyzersInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzersOutput, error)) *MockAccessAnalyzerClient_ListAnalyzers_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAccessAnalyzerClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAccessAnalyzerClient creates a new instance of MockAccessAnalyzerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAccessAnalyzerClient(t mockConstructorTestingTNewMockAccessAnalyzerClient) *MockAccessAnalyzerClient {
	mock := &MockAccessAnalyzerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
