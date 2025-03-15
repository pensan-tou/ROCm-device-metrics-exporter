//
// Copyright(C) Advanced Micro Devices, Inc. All rights reserved.
//
// You may not use this software and documentation (if any) (collectively,
// the "Materials") except in compliance with the terms and conditions of
// the Software License Agreement included with the Materials or otherwise as
// set forth in writing and signed by you and an authorized signatory of AMD.
// If you do not have a copy of the Software License Agreement, contact your
// AMD representative for a copy.
//
// You agree that you will not reverse engineer or decompile the Materials,
// in whole or in part, except as allowed by applicable law.
//
// THE MATERIALS ARE DISTRIBUTED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OR
// REPRESENTATIONS OF ANY KIND, EITHER EXPRESS OR IMPLIED.
//

//------------------------------------------------------------------------------
///
/// \file
/// protobuf specification for events
///
//------------------------------------------------------------------------------

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: events.proto

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

package amdgpu

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	EventSvc_EventGet_FullMethodName = "/amdgpu.EventSvc/EventGet"
)

// EventSvcClient is the client API for EventSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventSvcClient interface {
	// EventGet() API will return either all events that happened in the
	// system until that moment or only events of interest based on the
	// EventRequest
	// The client is expected to periodically or on-need basis query and
	// get the event information using this API
	EventGet(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type eventSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewEventSvcClient(cc grpc.ClientConnInterface) EventSvcClient {
	return &eventSvcClient{cc}
}

func (c *eventSvcClient) EventGet(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, EventSvc_EventGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventSvcServer is the server API for EventSvc service.
// All implementations must embed UnimplementedEventSvcServer
// for forward compatibility.
type EventSvcServer interface {
	// EventGet() API will return either all events that happened in the
	// system until that moment or only events of interest based on the
	// EventRequest
	// The client is expected to periodically or on-need basis query and
	// get the event information using this API
	EventGet(context.Context, *EventRequest) (*EventResponse, error)
	mustEmbedUnimplementedEventSvcServer()
}

// UnimplementedEventSvcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEventSvcServer struct{}

func (UnimplementedEventSvcServer) EventGet(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventGet not implemented")
}
func (UnimplementedEventSvcServer) mustEmbedUnimplementedEventSvcServer() {}
func (UnimplementedEventSvcServer) testEmbeddedByValue()                  {}

// UnsafeEventSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventSvcServer will
// result in compilation errors.
type UnsafeEventSvcServer interface {
	mustEmbedUnimplementedEventSvcServer()
}

func RegisterEventSvcServer(s grpc.ServiceRegistrar, srv EventSvcServer) {
	// If the following call pancis, it indicates UnimplementedEventSvcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EventSvc_ServiceDesc, srv)
}

func _EventSvc_EventGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSvcServer).EventGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventSvc_EventGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSvcServer).EventGet(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventSvc_ServiceDesc is the grpc.ServiceDesc for EventSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "amdgpu.EventSvc",
	HandlerType: (*EventSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventGet",
			Handler:    _EventSvc_EventGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events.proto",
}

const (
	DebugEventSvc_EventGen_FullMethodName = "/amdgpu.DebugEventSvc/EventGen"
)

// DebugEventSvcClient is the client API for DebugEventSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// experimental debug event APIs, internal debug tools and not for
// external consumption
type DebugEventSvcClient interface {
	// EventGen() API will generate events
	EventGen(ctx context.Context, in *EventGenRequest, opts ...grpc.CallOption) (*EventGenResponse, error)
}

type debugEventSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugEventSvcClient(cc grpc.ClientConnInterface) DebugEventSvcClient {
	return &debugEventSvcClient{cc}
}

func (c *debugEventSvcClient) EventGen(ctx context.Context, in *EventGenRequest, opts ...grpc.CallOption) (*EventGenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EventGenResponse)
	err := c.cc.Invoke(ctx, DebugEventSvc_EventGen_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugEventSvcServer is the server API for DebugEventSvc service.
// All implementations must embed UnimplementedDebugEventSvcServer
// for forward compatibility.
//
// experimental debug event APIs, internal debug tools and not for
// external consumption
type DebugEventSvcServer interface {
	// EventGen() API will generate events
	EventGen(context.Context, *EventGenRequest) (*EventGenResponse, error)
	mustEmbedUnimplementedDebugEventSvcServer()
}

// UnimplementedDebugEventSvcServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDebugEventSvcServer struct{}

func (UnimplementedDebugEventSvcServer) EventGen(context.Context, *EventGenRequest) (*EventGenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventGen not implemented")
}
func (UnimplementedDebugEventSvcServer) mustEmbedUnimplementedDebugEventSvcServer() {}
func (UnimplementedDebugEventSvcServer) testEmbeddedByValue()                       {}

// UnsafeDebugEventSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DebugEventSvcServer will
// result in compilation errors.
type UnsafeDebugEventSvcServer interface {
	mustEmbedUnimplementedDebugEventSvcServer()
}

func RegisterDebugEventSvcServer(s grpc.ServiceRegistrar, srv DebugEventSvcServer) {
	// If the following call pancis, it indicates UnimplementedDebugEventSvcServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DebugEventSvc_ServiceDesc, srv)
}

func _DebugEventSvc_EventGen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventGenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugEventSvcServer).EventGen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebugEventSvc_EventGen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugEventSvcServer).EventGen(ctx, req.(*EventGenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DebugEventSvc_ServiceDesc is the grpc.ServiceDesc for DebugEventSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DebugEventSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "amdgpu.DebugEventSvc",
	HandlerType: (*DebugEventSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventGen",
			Handler:    _DebugEventSvc_EventGen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events.proto",
}
