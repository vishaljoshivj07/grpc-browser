// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heurist.proto

package pb

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

type CheckUsernameRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckUsernameRequest) Reset()         { *m = CheckUsernameRequest{} }
func (m *CheckUsernameRequest) String() string { return proto.CompactTextString(m) }
func (*CheckUsernameRequest) ProtoMessage()    {}
func (*CheckUsernameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_heurist_7f57832cc076c67f, []int{0}
}
func (m *CheckUsernameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUsernameRequest.Unmarshal(m, b)
}
func (m *CheckUsernameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUsernameRequest.Marshal(b, m, deterministic)
}
func (dst *CheckUsernameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUsernameRequest.Merge(dst, src)
}
func (m *CheckUsernameRequest) XXX_Size() int {
	return xxx_messageInfo_CheckUsernameRequest.Size(m)
}
func (m *CheckUsernameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUsernameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUsernameRequest proto.InternalMessageInfo

func (m *CheckUsernameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type CheckUsernameResponse struct {
	IsValid              bool     `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckUsernameResponse) Reset()         { *m = CheckUsernameResponse{} }
func (m *CheckUsernameResponse) String() string { return proto.CompactTextString(m) }
func (*CheckUsernameResponse) ProtoMessage()    {}
func (*CheckUsernameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_heurist_7f57832cc076c67f, []int{1}
}
func (m *CheckUsernameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckUsernameResponse.Unmarshal(m, b)
}
func (m *CheckUsernameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckUsernameResponse.Marshal(b, m, deterministic)
}
func (dst *CheckUsernameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckUsernameResponse.Merge(dst, src)
}
func (m *CheckUsernameResponse) XXX_Size() int {
	return xxx_messageInfo_CheckUsernameResponse.Size(m)
}
func (m *CheckUsernameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckUsernameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckUsernameResponse proto.InternalMessageInfo

func (m *CheckUsernameResponse) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

type GetUserDetailsRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserDetailsRequest) Reset()         { *m = GetUserDetailsRequest{} }
func (m *GetUserDetailsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserDetailsRequest) ProtoMessage()    {}
func (*GetUserDetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_heurist_7f57832cc076c67f, []int{2}
}
func (m *GetUserDetailsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserDetailsRequest.Unmarshal(m, b)
}
func (m *GetUserDetailsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserDetailsRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserDetailsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserDetailsRequest.Merge(dst, src)
}
func (m *GetUserDetailsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserDetailsRequest.Size(m)
}
func (m *GetUserDetailsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserDetailsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserDetailsRequest proto.InternalMessageInfo

func (m *GetUserDetailsRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUserDetailsResponse struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	ImageUrl             string   `protobuf:"bytes,2,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	About                string   `protobuf:"bytes,4,opt,name=about,proto3" json:"about,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserDetailsResponse) Reset()         { *m = GetUserDetailsResponse{} }
func (m *GetUserDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserDetailsResponse) ProtoMessage()    {}
func (*GetUserDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_heurist_7f57832cc076c67f, []int{3}
}
func (m *GetUserDetailsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserDetailsResponse.Unmarshal(m, b)
}
func (m *GetUserDetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserDetailsResponse.Marshal(b, m, deterministic)
}
func (dst *GetUserDetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserDetailsResponse.Merge(dst, src)
}
func (m *GetUserDetailsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserDetailsResponse.Size(m)
}
func (m *GetUserDetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserDetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserDetailsResponse proto.InternalMessageInfo

func (m *GetUserDetailsResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserDetailsResponse) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *GetUserDetailsResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserDetailsResponse) GetAbout() string {
	if m != nil {
		return m.About
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckUsernameRequest)(nil), "pb.CheckUsernameRequest")
	proto.RegisterType((*CheckUsernameResponse)(nil), "pb.CheckUsernameResponse")
	proto.RegisterType((*GetUserDetailsRequest)(nil), "pb.GetUserDetailsRequest")
	proto.RegisterType((*GetUserDetailsResponse)(nil), "pb.GetUserDetailsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeuristGrpcClient is the client API for HeuristGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeuristGrpcClient interface {
	Check(ctx context.Context, in *CheckUsernameRequest, opts ...grpc.CallOption) (*CheckUsernameResponse, error)
	GetUser(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error)
}

type heuristGrpcClient struct {
	cc *grpc.ClientConn
}

func NewHeuristGrpcClient(cc *grpc.ClientConn) HeuristGrpcClient {
	return &heuristGrpcClient{cc}
}

func (c *heuristGrpcClient) Check(ctx context.Context, in *CheckUsernameRequest, opts ...grpc.CallOption) (*CheckUsernameResponse, error) {
	out := new(CheckUsernameResponse)
	err := c.cc.Invoke(ctx, "/pb.HeuristGrpc/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heuristGrpcClient) GetUser(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error) {
	out := new(GetUserDetailsResponse)
	err := c.cc.Invoke(ctx, "/pb.HeuristGrpc/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeuristGrpcServer is the server API for HeuristGrpc service.
type HeuristGrpcServer interface {
	Check(context.Context, *CheckUsernameRequest) (*CheckUsernameResponse, error)
	GetUser(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error)
}

func RegisterHeuristGrpcServer(s *grpc.Server, srv HeuristGrpcServer) {
	s.RegisterService(&_HeuristGrpc_serviceDesc, srv)
}

func _HeuristGrpc_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeuristGrpcServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HeuristGrpc/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeuristGrpcServer).Check(ctx, req.(*CheckUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeuristGrpc_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeuristGrpcServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HeuristGrpc/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeuristGrpcServer).GetUser(ctx, req.(*GetUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeuristGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HeuristGrpc",
	HandlerType: (*HeuristGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _HeuristGrpc_Check_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _HeuristGrpc_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heurist.proto",
}

func init() { proto.RegisterFile("heurist.proto", fileDescriptor_heurist_7f57832cc076c67f) }

var fileDescriptor_heurist_7f57832cc076c67f = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x49, 0x69, 0x69, 0x39, 0xc4, 0x72, 0x6a, 0x91, 0xc9, 0x84, 0x32, 0x31, 0x45, 0xa2,
	0xdd, 0x19, 0x00, 0xa9, 0xcc, 0x91, 0xca, 0x6e, 0x97, 0x13, 0xb5, 0x48, 0x6b, 0xe3, 0x3f, 0x0b,
	0xdf, 0x82, 0x6f, 0x8c, 0x7a, 0x4e, 0x18, 0x22, 0x0b, 0x75, 0xcb, 0x7b, 0x2f, 0x3f, 0x9f, 0xef,
	0x19, 0xae, 0x77, 0x14, 0x9d, 0xf6, 0xa1, 0xb6, 0xce, 0x04, 0x83, 0x23, 0xab, 0xaa, 0x25, 0xcc,
	0x9f, 0x77, 0xb4, 0xfd, 0xdc, 0x78, 0x72, 0x07, 0xb9, 0xa7, 0x86, 0xbe, 0x22, 0xf9, 0x80, 0x25,
	0xcc, 0x62, 0x67, 0x89, 0xe2, 0xae, 0xb8, 0xbf, 0x6c, 0xfe, 0x74, 0xf5, 0x00, 0x8b, 0x01, 0xe3,
	0xad, 0x39, 0x78, 0x42, 0x01, 0x53, 0xed, 0xdf, 0x64, 0xab, 0xdf, 0x99, 0x99, 0x35, 0xbd, 0xac,
	0x56, 0xb0, 0x58, 0x53, 0x38, 0x02, 0x2f, 0x14, 0xa4, 0x6e, 0xfd, 0x29, 0x73, 0xbe, 0xe1, 0x66,
	0x08, 0x75, 0x83, 0xfe, 0xa1, 0x8e, 0x99, 0xde, 0xcb, 0x0f, 0xda, 0xb8, 0x56, 0x8c, 0x52, 0xd6,
	0x6b, 0x44, 0x18, 0x33, 0x73, 0xce, 0x3e, 0x7f, 0xe3, 0x1c, 0x26, 0x52, 0x99, 0x18, 0xc4, 0x98,
	0xcd, 0x24, 0x96, 0x3f, 0x05, 0x5c, 0xbd, 0xa6, 0xb6, 0xd6, 0xce, 0x6e, 0xf1, 0x11, 0x26, 0xbc,
	0x33, 0x8a, 0xda, 0xaa, 0x3a, 0x57, 0x59, 0x79, 0x9b, 0x49, 0xd2, 0x7d, 0xab, 0x33, 0x7c, 0x82,
	0x69, 0xb7, 0x0b, 0xf2, 0x7f, 0xd9, 0x36, 0xca, 0x32, 0x17, 0xf5, 0x67, 0xa8, 0x0b, 0x7e, 0xb6,
	0xd5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x58, 0x27, 0xd8, 0xc7, 0x01, 0x00, 0x00,
}