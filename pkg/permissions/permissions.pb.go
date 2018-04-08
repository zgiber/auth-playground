// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permissions.proto

/*
Package permissions is a generated protocol buffer package.

It is generated from these files:
	permissions.proto

It has these top-level messages:
	Permission
	ListPermissionsRequest
	ListPermissionsResponse
	RevokePermissionsRequest
	Error
	SimpleResponse
*/
package permissions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Permission struct {
	ClientID string   `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	Resource string   `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	Actions  []string `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
}

func (m *Permission) Reset()                    { *m = Permission{} }
func (m *Permission) String() string            { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()               {}
func (*Permission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Permission) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *Permission) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *Permission) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

type ListPermissionsRequest struct {
	ClientID string `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	Resource string `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
}

func (m *ListPermissionsRequest) Reset()                    { *m = ListPermissionsRequest{} }
func (m *ListPermissionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPermissionsRequest) ProtoMessage()               {}
func (*ListPermissionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListPermissionsRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *ListPermissionsRequest) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type ListPermissionsResponse struct {
	Success     bool          `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Permissions []*Permission `protobuf:"bytes,2,rep,name=permissions" json:"permissions,omitempty"`
	Error       *Error        `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *ListPermissionsResponse) Reset()                    { *m = ListPermissionsResponse{} }
func (m *ListPermissionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPermissionsResponse) ProtoMessage()               {}
func (*ListPermissionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListPermissionsResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ListPermissionsResponse) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *ListPermissionsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type RevokePermissionsRequest struct {
	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *RevokePermissionsRequest) Reset()                    { *m = RevokePermissionsRequest{} }
func (m *RevokePermissionsRequest) String() string            { return proto.CompactTextString(m) }
func (*RevokePermissionsRequest) ProtoMessage()               {}
func (*RevokePermissionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RevokePermissionsRequest) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Error struct {
	ErrorCode   string `protobuf:"bytes,1,opt,name=errorCode" json:"errorCode,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Error) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type SimpleResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   *Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SimpleResponse) Reset()                    { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string            { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()               {}
func (*SimpleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SimpleResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SimpleResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Permission)(nil), "permissions.Permission")
	proto.RegisterType((*ListPermissionsRequest)(nil), "permissions.ListPermissionsRequest")
	proto.RegisterType((*ListPermissionsResponse)(nil), "permissions.ListPermissionsResponse")
	proto.RegisterType((*RevokePermissionsRequest)(nil), "permissions.RevokePermissionsRequest")
	proto.RegisterType((*Error)(nil), "permissions.Error")
	proto.RegisterType((*SimpleResponse)(nil), "permissions.SimpleResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Permissions service

type PermissionsClient interface {
	Grant(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*SimpleResponse, error)
	List(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error)
	Revoke(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*SimpleResponse, error)
	RevokeMulti(ctx context.Context, in *RevokePermissionsRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
}

type permissionsClient struct {
	cc *grpc.ClientConn
}

func NewPermissionsClient(cc *grpc.ClientConn) PermissionsClient {
	return &permissionsClient{cc}
}

func (c *permissionsClient) Grant(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/permissions.Permissions/Grant", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) List(ctx context.Context, in *ListPermissionsRequest, opts ...grpc.CallOption) (*ListPermissionsResponse, error) {
	out := new(ListPermissionsResponse)
	err := grpc.Invoke(ctx, "/permissions.Permissions/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) Revoke(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/permissions.Permissions/Revoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionsClient) RevokeMulti(ctx context.Context, in *RevokePermissionsRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/permissions.Permissions/RevokeMulti", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Permissions service

type PermissionsServer interface {
	Grant(context.Context, *Permission) (*SimpleResponse, error)
	List(context.Context, *ListPermissionsRequest) (*ListPermissionsResponse, error)
	Revoke(context.Context, *Permission) (*SimpleResponse, error)
	RevokeMulti(context.Context, *RevokePermissionsRequest) (*SimpleResponse, error)
}

func RegisterPermissionsServer(s *grpc.Server, srv PermissionsServer) {
	s.RegisterService(&_Permissions_serviceDesc, srv)
}

func _Permissions_Grant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Permission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServer).Grant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.Permissions/Grant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServer).Grant(ctx, req.(*Permission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permissions_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.Permissions/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServer).List(ctx, req.(*ListPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permissions_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Permission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.Permissions/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServer).Revoke(ctx, req.(*Permission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Permissions_RevokeMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionsServer).RevokeMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/permissions.Permissions/RevokeMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionsServer).RevokeMulti(ctx, req.(*RevokePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Permissions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "permissions.Permissions",
	HandlerType: (*PermissionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grant",
			Handler:    _Permissions_Grant_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Permissions_List_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _Permissions_Revoke_Handler,
		},
		{
			MethodName: "RevokeMulti",
			Handler:    _Permissions_RevokeMulti_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permissions.proto",
}

func init() { proto.RegisterFile("permissions.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xb5, 0xad, 0x20, 0x4c, 0x13, 0x13, 0xf7, 0x20, 0x1b, 0xf4, 0xd0, 0xac, 0x9a, 0xf4, 0xc4,
	0x01, 0x4f, 0x9e, 0x34, 0x51, 0x43, 0x4c, 0x34, 0x21, 0x8b, 0x5e, 0x4d, 0xb0, 0xcc, 0x61, 0x23,
	0x74, 0xeb, 0xce, 0xd6, 0x6f, 0xf1, 0x5b, 0xfc, 0x3a, 0xd3, 0x16, 0x4a, 0xab, 0x20, 0x44, 0x8f,
	0x33, 0xf3, 0xfa, 0xe6, 0xbd, 0x37, 0x5d, 0x38, 0x48, 0xd0, 0xcc, 0x14, 0x91, 0xd2, 0x31, 0xf5,
	0x12, 0xa3, 0xad, 0x66, 0x7e, 0xa5, 0x25, 0x9e, 0x01, 0x86, 0x65, 0xc9, 0xba, 0xd0, 0x8a, 0xa6,
	0x0a, 0x63, 0x7b, 0x77, 0xc3, 0x9d, 0xc0, 0x09, 0xdb, 0xb2, 0xac, 0xb3, 0x99, 0x41, 0xd2, 0xa9,
	0x89, 0x90, 0xbb, 0xc5, 0x6c, 0x51, 0x33, 0x0e, 0x7b, 0xe3, 0xc8, 0x66, 0x84, 0xdc, 0x0b, 0xbc,
	0xb0, 0x2d, 0x17, 0xa5, 0x18, 0xc2, 0xe1, 0xbd, 0x22, 0xbb, 0xdc, 0x41, 0x12, 0xdf, 0x52, 0x24,
	0xfb, 0xd7, 0x5d, 0xe2, 0xc3, 0x81, 0xce, 0x0f, 0x4a, 0x4a, 0x74, 0x4c, 0xb9, 0x0e, 0x4a, 0xa3,
	0x08, 0x89, 0x72, 0xca, 0x96, 0x5c, 0x94, 0xec, 0x02, 0xaa, 0xb6, 0xb9, 0x1b, 0x78, 0xa1, 0xdf,
	0xef, 0xf4, 0xaa, 0xe9, 0x2c, 0x09, 0x65, 0x15, 0xcb, 0x42, 0x68, 0xa0, 0x31, 0xda, 0x70, 0x2f,
	0x70, 0x42, 0xbf, 0xcf, 0x6a, 0x1f, 0xdd, 0x66, 0x13, 0x59, 0x00, 0xc4, 0x13, 0x70, 0x89, 0xef,
	0xfa, 0x15, 0x57, 0xd8, 0xfd, 0x26, 0xc0, 0xd9, 0x5e, 0x80, 0x18, 0x40, 0x23, 0x5f, 0xc3, 0x8e,
	0xa1, 0x9d, 0x2f, 0xba, 0xd6, 0x13, 0x9c, 0x67, 0xb6, 0x6c, 0xb0, 0x00, 0xfc, 0x09, 0x52, 0x64,
	0x54, 0x92, 0x45, 0x3f, 0xcf, 0xad, 0xda, 0x12, 0x8f, 0xb0, 0x3f, 0x52, 0xb3, 0x64, 0x8a, 0x5b,
	0x04, 0x56, 0xba, 0x76, 0x37, 0xb8, 0xee, 0x7f, 0xba, 0xe0, 0x57, 0x0c, 0xb3, 0x4b, 0x68, 0x0c,
	0xcc, 0x38, 0xb6, 0x6c, 0x9d, 0xbb, 0xee, 0x51, 0x6d, 0x50, 0x97, 0x24, 0x76, 0xd8, 0x08, 0x76,
	0xb3, 0x03, 0xb3, 0x93, 0x1a, 0x6c, 0xf5, 0x6f, 0xd4, 0x3d, 0xfd, 0x1d, 0x54, 0x92, 0x5e, 0x41,
	0xb3, 0xb8, 0xcd, 0x3f, 0x64, 0xf9, 0x05, 0xc3, 0x43, 0x3a, 0xb5, 0x8a, 0x9d, 0xd5, 0xd0, 0xeb,
	0xee, 0xbe, 0x81, 0xf4, 0xa5, 0x99, 0xbf, 0xc9, 0xf3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad,
	0xda, 0xeb, 0x93, 0xa8, 0x03, 0x00, 0x00,
}