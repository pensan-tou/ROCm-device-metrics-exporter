/*
Copyright (c) Advanced Micro Devices, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the \"License\");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an \"AS IS\" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: /usr/src/github.com/ROCm/device-metrics-exporter/pkg/amdgpu//gen/amdgpu/gpu_grpc.pb.go
//
// Generated by this command:
//
//	mockgen --destination=/usr/src/github.com/ROCm/device-metrics-exporter/pkg/amdgpu/mock_gen/mock_gpu.go -package=mock_gen --source=/usr/src/github.com/ROCm/device-metrics-exporter/pkg/amdgpu//gen/amdgpu/gpu_grpc.pb.go
//

// Package mock_gen is a generated GoMock package.
package mock_gen

import (
	context "context"
	reflect "reflect"

	amdgpu "github.com/ROCm/device-metrics-exporter/pkg/amdgpu/gen/amdgpu"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockGPUSvcClient is a mock of GPUSvcClient interface.
type MockGPUSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockGPUSvcClientMockRecorder
	isgomock struct{}
}

// MockGPUSvcClientMockRecorder is the mock recorder for MockGPUSvcClient.
type MockGPUSvcClientMockRecorder struct {
	mock *MockGPUSvcClient
}

// NewMockGPUSvcClient creates a new mock instance.
func NewMockGPUSvcClient(ctrl *gomock.Controller) *MockGPUSvcClient {
	mock := &MockGPUSvcClient{ctrl: ctrl}
	mock.recorder = &MockGPUSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGPUSvcClient) EXPECT() *MockGPUSvcClientMockRecorder {
	return m.recorder
}

// GPUComputePartitionGet mocks base method.
func (m *MockGPUSvcClient) GPUComputePartitionGet(ctx context.Context, in *amdgpu.GPUComputePartitionGetRequest, opts ...grpc.CallOption) (*amdgpu.GPUComputePartitionGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GPUComputePartitionGet", varargs...)
	ret0, _ := ret[0].(*amdgpu.GPUComputePartitionGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUComputePartitionGet indicates an expected call of GPUComputePartitionGet.
func (mr *MockGPUSvcClientMockRecorder) GPUComputePartitionGet(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUComputePartitionGet", reflect.TypeOf((*MockGPUSvcClient)(nil).GPUComputePartitionGet), varargs...)
}

// GPUComputePartitionSet mocks base method.
func (m *MockGPUSvcClient) GPUComputePartitionSet(ctx context.Context, in *amdgpu.GPUComputePartitionSetRequest, opts ...grpc.CallOption) (*amdgpu.GPUComputePartitionSetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GPUComputePartitionSet", varargs...)
	ret0, _ := ret[0].(*amdgpu.GPUComputePartitionSetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUComputePartitionSet indicates an expected call of GPUComputePartitionSet.
func (mr *MockGPUSvcClientMockRecorder) GPUComputePartitionSet(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUComputePartitionSet", reflect.TypeOf((*MockGPUSvcClient)(nil).GPUComputePartitionSet), varargs...)
}

// GPUGet mocks base method.
func (m *MockGPUSvcClient) GPUGet(ctx context.Context, in *amdgpu.GPUGetRequest, opts ...grpc.CallOption) (*amdgpu.GPUGetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GPUGet", varargs...)
	ret0, _ := ret[0].(*amdgpu.GPUGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUGet indicates an expected call of GPUGet.
func (mr *MockGPUSvcClientMockRecorder) GPUGet(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUGet", reflect.TypeOf((*MockGPUSvcClient)(nil).GPUGet), varargs...)
}

// GPUReset mocks base method.
func (m *MockGPUSvcClient) GPUReset(ctx context.Context, in *amdgpu.GPUResetRequest, opts ...grpc.CallOption) (*amdgpu.GPUResetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GPUReset", varargs...)
	ret0, _ := ret[0].(*amdgpu.GPUResetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUReset indicates an expected call of GPUReset.
func (mr *MockGPUSvcClientMockRecorder) GPUReset(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUReset", reflect.TypeOf((*MockGPUSvcClient)(nil).GPUReset), varargs...)
}

// GPUUpdate mocks base method.
func (m *MockGPUSvcClient) GPUUpdate(ctx context.Context, in *amdgpu.GPUUpdateRequest, opts ...grpc.CallOption) (*amdgpu.GPUUpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GPUUpdate", varargs...)
	ret0, _ := ret[0].(*amdgpu.GPUUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUUpdate indicates an expected call of GPUUpdate.
func (mr *MockGPUSvcClientMockRecorder) GPUUpdate(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUUpdate", reflect.TypeOf((*MockGPUSvcClient)(nil).GPUUpdate), varargs...)
}

// MockGPUSvcServer is a mock of GPUSvcServer interface.
type MockGPUSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockGPUSvcServerMockRecorder
	isgomock struct{}
}

// MockGPUSvcServerMockRecorder is the mock recorder for MockGPUSvcServer.
type MockGPUSvcServerMockRecorder struct {
	mock *MockGPUSvcServer
}

// NewMockGPUSvcServer creates a new mock instance.
func NewMockGPUSvcServer(ctrl *gomock.Controller) *MockGPUSvcServer {
	mock := &MockGPUSvcServer{ctrl: ctrl}
	mock.recorder = &MockGPUSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGPUSvcServer) EXPECT() *MockGPUSvcServerMockRecorder {
	return m.recorder
}

// GPUComputePartitionGet mocks base method.
func (m *MockGPUSvcServer) GPUComputePartitionGet(arg0 context.Context, arg1 *amdgpu.GPUComputePartitionGetRequest) (*amdgpu.GPUComputePartitionGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GPUComputePartitionGet", arg0, arg1)
	ret0, _ := ret[0].(*amdgpu.GPUComputePartitionGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUComputePartitionGet indicates an expected call of GPUComputePartitionGet.
func (mr *MockGPUSvcServerMockRecorder) GPUComputePartitionGet(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUComputePartitionGet", reflect.TypeOf((*MockGPUSvcServer)(nil).GPUComputePartitionGet), arg0, arg1)
}

// GPUComputePartitionSet mocks base method.
func (m *MockGPUSvcServer) GPUComputePartitionSet(arg0 context.Context, arg1 *amdgpu.GPUComputePartitionSetRequest) (*amdgpu.GPUComputePartitionSetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GPUComputePartitionSet", arg0, arg1)
	ret0, _ := ret[0].(*amdgpu.GPUComputePartitionSetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUComputePartitionSet indicates an expected call of GPUComputePartitionSet.
func (mr *MockGPUSvcServerMockRecorder) GPUComputePartitionSet(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUComputePartitionSet", reflect.TypeOf((*MockGPUSvcServer)(nil).GPUComputePartitionSet), arg0, arg1)
}

// GPUGet mocks base method.
func (m *MockGPUSvcServer) GPUGet(arg0 context.Context, arg1 *amdgpu.GPUGetRequest) (*amdgpu.GPUGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GPUGet", arg0, arg1)
	ret0, _ := ret[0].(*amdgpu.GPUGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUGet indicates an expected call of GPUGet.
func (mr *MockGPUSvcServerMockRecorder) GPUGet(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUGet", reflect.TypeOf((*MockGPUSvcServer)(nil).GPUGet), arg0, arg1)
}

// GPUReset mocks base method.
func (m *MockGPUSvcServer) GPUReset(arg0 context.Context, arg1 *amdgpu.GPUResetRequest) (*amdgpu.GPUResetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GPUReset", arg0, arg1)
	ret0, _ := ret[0].(*amdgpu.GPUResetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUReset indicates an expected call of GPUReset.
func (mr *MockGPUSvcServerMockRecorder) GPUReset(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUReset", reflect.TypeOf((*MockGPUSvcServer)(nil).GPUReset), arg0, arg1)
}

// GPUUpdate mocks base method.
func (m *MockGPUSvcServer) GPUUpdate(arg0 context.Context, arg1 *amdgpu.GPUUpdateRequest) (*amdgpu.GPUUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GPUUpdate", arg0, arg1)
	ret0, _ := ret[0].(*amdgpu.GPUUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUUpdate indicates an expected call of GPUUpdate.
func (mr *MockGPUSvcServerMockRecorder) GPUUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUUpdate", reflect.TypeOf((*MockGPUSvcServer)(nil).GPUUpdate), arg0, arg1)
}

// mustEmbedUnimplementedGPUSvcServer mocks base method.
func (m *MockGPUSvcServer) mustEmbedUnimplementedGPUSvcServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedGPUSvcServer")
}

// mustEmbedUnimplementedGPUSvcServer indicates an expected call of mustEmbedUnimplementedGPUSvcServer.
func (mr *MockGPUSvcServerMockRecorder) mustEmbedUnimplementedGPUSvcServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedGPUSvcServer", reflect.TypeOf((*MockGPUSvcServer)(nil).mustEmbedUnimplementedGPUSvcServer))
}

// MockUnsafeGPUSvcServer is a mock of UnsafeGPUSvcServer interface.
type MockUnsafeGPUSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeGPUSvcServerMockRecorder
	isgomock struct{}
}

// MockUnsafeGPUSvcServerMockRecorder is the mock recorder for MockUnsafeGPUSvcServer.
type MockUnsafeGPUSvcServerMockRecorder struct {
	mock *MockUnsafeGPUSvcServer
}

// NewMockUnsafeGPUSvcServer creates a new mock instance.
func NewMockUnsafeGPUSvcServer(ctrl *gomock.Controller) *MockUnsafeGPUSvcServer {
	mock := &MockUnsafeGPUSvcServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeGPUSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeGPUSvcServer) EXPECT() *MockUnsafeGPUSvcServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedGPUSvcServer mocks base method.
func (m *MockUnsafeGPUSvcServer) mustEmbedUnimplementedGPUSvcServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedGPUSvcServer")
}

// mustEmbedUnimplementedGPUSvcServer indicates an expected call of mustEmbedUnimplementedGPUSvcServer.
func (mr *MockUnsafeGPUSvcServerMockRecorder) mustEmbedUnimplementedGPUSvcServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedGPUSvcServer", reflect.TypeOf((*MockUnsafeGPUSvcServer)(nil).mustEmbedUnimplementedGPUSvcServer))
}

// MockDebugGPUSvcClient is a mock of DebugGPUSvcClient interface.
type MockDebugGPUSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockDebugGPUSvcClientMockRecorder
	isgomock struct{}
}

// MockDebugGPUSvcClientMockRecorder is the mock recorder for MockDebugGPUSvcClient.
type MockDebugGPUSvcClientMockRecorder struct {
	mock *MockDebugGPUSvcClient
}

// NewMockDebugGPUSvcClient creates a new mock instance.
func NewMockDebugGPUSvcClient(ctrl *gomock.Controller) *MockDebugGPUSvcClient {
	mock := &MockDebugGPUSvcClient{ctrl: ctrl}
	mock.recorder = &MockDebugGPUSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDebugGPUSvcClient) EXPECT() *MockDebugGPUSvcClientMockRecorder {
	return m.recorder
}

// GPUBadPageGet mocks base method.
func (m *MockDebugGPUSvcClient) GPUBadPageGet(ctx context.Context, in *amdgpu.GPUBadPageGetRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[amdgpu.GPUBadPageGetResponse], error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GPUBadPageGet", varargs...)
	ret0, _ := ret[0].(grpc.ServerStreamingClient[amdgpu.GPUBadPageGetResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GPUBadPageGet indicates an expected call of GPUBadPageGet.
func (mr *MockDebugGPUSvcClientMockRecorder) GPUBadPageGet(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUBadPageGet", reflect.TypeOf((*MockDebugGPUSvcClient)(nil).GPUBadPageGet), varargs...)
}

// MockDebugGPUSvcServer is a mock of DebugGPUSvcServer interface.
type MockDebugGPUSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockDebugGPUSvcServerMockRecorder
	isgomock struct{}
}

// MockDebugGPUSvcServerMockRecorder is the mock recorder for MockDebugGPUSvcServer.
type MockDebugGPUSvcServerMockRecorder struct {
	mock *MockDebugGPUSvcServer
}

// NewMockDebugGPUSvcServer creates a new mock instance.
func NewMockDebugGPUSvcServer(ctrl *gomock.Controller) *MockDebugGPUSvcServer {
	mock := &MockDebugGPUSvcServer{ctrl: ctrl}
	mock.recorder = &MockDebugGPUSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDebugGPUSvcServer) EXPECT() *MockDebugGPUSvcServerMockRecorder {
	return m.recorder
}

// GPUBadPageGet mocks base method.
func (m *MockDebugGPUSvcServer) GPUBadPageGet(arg0 *amdgpu.GPUBadPageGetRequest, arg1 grpc.ServerStreamingServer[amdgpu.GPUBadPageGetResponse]) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GPUBadPageGet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GPUBadPageGet indicates an expected call of GPUBadPageGet.
func (mr *MockDebugGPUSvcServerMockRecorder) GPUBadPageGet(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GPUBadPageGet", reflect.TypeOf((*MockDebugGPUSvcServer)(nil).GPUBadPageGet), arg0, arg1)
}

// mustEmbedUnimplementedDebugGPUSvcServer mocks base method.
func (m *MockDebugGPUSvcServer) mustEmbedUnimplementedDebugGPUSvcServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDebugGPUSvcServer")
}

// mustEmbedUnimplementedDebugGPUSvcServer indicates an expected call of mustEmbedUnimplementedDebugGPUSvcServer.
func (mr *MockDebugGPUSvcServerMockRecorder) mustEmbedUnimplementedDebugGPUSvcServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDebugGPUSvcServer", reflect.TypeOf((*MockDebugGPUSvcServer)(nil).mustEmbedUnimplementedDebugGPUSvcServer))
}

// MockUnsafeDebugGPUSvcServer is a mock of UnsafeDebugGPUSvcServer interface.
type MockUnsafeDebugGPUSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeDebugGPUSvcServerMockRecorder
	isgomock struct{}
}

// MockUnsafeDebugGPUSvcServerMockRecorder is the mock recorder for MockUnsafeDebugGPUSvcServer.
type MockUnsafeDebugGPUSvcServerMockRecorder struct {
	mock *MockUnsafeDebugGPUSvcServer
}

// NewMockUnsafeDebugGPUSvcServer creates a new mock instance.
func NewMockUnsafeDebugGPUSvcServer(ctrl *gomock.Controller) *MockUnsafeDebugGPUSvcServer {
	mock := &MockUnsafeDebugGPUSvcServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeDebugGPUSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeDebugGPUSvcServer) EXPECT() *MockUnsafeDebugGPUSvcServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedDebugGPUSvcServer mocks base method.
func (m *MockUnsafeDebugGPUSvcServer) mustEmbedUnimplementedDebugGPUSvcServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDebugGPUSvcServer")
}

// mustEmbedUnimplementedDebugGPUSvcServer indicates an expected call of mustEmbedUnimplementedDebugGPUSvcServer.
func (mr *MockUnsafeDebugGPUSvcServerMockRecorder) mustEmbedUnimplementedDebugGPUSvcServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDebugGPUSvcServer", reflect.TypeOf((*MockUnsafeDebugGPUSvcServer)(nil).mustEmbedUnimplementedDebugGPUSvcServer))
}
