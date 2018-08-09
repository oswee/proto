// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application.proto

package app

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

type GetApplicationRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetApplicationRequest) Reset()         { *m = GetApplicationRequest{} }
func (m *GetApplicationRequest) String() string { return proto.CompactTextString(m) }
func (*GetApplicationRequest) ProtoMessage()    {}
func (*GetApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_96596db0568e8f62, []int{0}
}
func (m *GetApplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetApplicationRequest.Unmarshal(m, b)
}
func (m *GetApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetApplicationRequest.Marshal(b, m, deterministic)
}
func (dst *GetApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetApplicationRequest.Merge(dst, src)
}
func (m *GetApplicationRequest) XXX_Size() int {
	return xxx_messageInfo_GetApplicationRequest.Size(m)
}
func (m *GetApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetApplicationRequest proto.InternalMessageInfo

func (m *GetApplicationRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetApplicationResponse struct {
	Application          *Application `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetApplicationResponse) Reset()         { *m = GetApplicationResponse{} }
func (m *GetApplicationResponse) String() string { return proto.CompactTextString(m) }
func (*GetApplicationResponse) ProtoMessage()    {}
func (*GetApplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_96596db0568e8f62, []int{1}
}
func (m *GetApplicationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetApplicationResponse.Unmarshal(m, b)
}
func (m *GetApplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetApplicationResponse.Marshal(b, m, deterministic)
}
func (dst *GetApplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetApplicationResponse.Merge(dst, src)
}
func (m *GetApplicationResponse) XXX_Size() int {
	return xxx_messageInfo_GetApplicationResponse.Size(m)
}
func (m *GetApplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetApplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetApplicationResponse proto.InternalMessageInfo

func (m *GetApplicationResponse) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

type ListApplicationsRequest struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	PageNumber           int32    `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32    `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListApplicationsRequest) Reset()         { *m = ListApplicationsRequest{} }
func (m *ListApplicationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListApplicationsRequest) ProtoMessage()    {}
func (*ListApplicationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_96596db0568e8f62, []int{2}
}
func (m *ListApplicationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListApplicationsRequest.Unmarshal(m, b)
}
func (m *ListApplicationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListApplicationsRequest.Marshal(b, m, deterministic)
}
func (dst *ListApplicationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListApplicationsRequest.Merge(dst, src)
}
func (m *ListApplicationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListApplicationsRequest.Size(m)
}
func (m *ListApplicationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListApplicationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListApplicationsRequest proto.InternalMessageInfo

func (m *ListApplicationsRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ListApplicationsRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *ListApplicationsRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

type ListApplicationsResponse struct {
	Applications         []*Application `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListApplicationsResponse) Reset()         { *m = ListApplicationsResponse{} }
func (m *ListApplicationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListApplicationsResponse) ProtoMessage()    {}
func (*ListApplicationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_96596db0568e8f62, []int{3}
}
func (m *ListApplicationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListApplicationsResponse.Unmarshal(m, b)
}
func (m *ListApplicationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListApplicationsResponse.Marshal(b, m, deterministic)
}
func (dst *ListApplicationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListApplicationsResponse.Merge(dst, src)
}
func (m *ListApplicationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListApplicationsResponse.Size(m)
}
func (m *ListApplicationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListApplicationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListApplicationsResponse proto.InternalMessageInfo

func (m *ListApplicationsResponse) GetApplications() []*Application {
	if m != nil {
		return m.Applications
	}
	return nil
}

type Application struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId             int32    `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	FullName             string   `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Permalink            string   `protobuf:"bytes,5,opt,name=permalink,proto3" json:"permalink,omitempty"`
	Type                 int32    `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_96596db0568e8f62, []int{4}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (dst *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(dst, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Application) GetParentId() int32 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Application) GetPermalink() string {
	if m != nil {
		return m.Permalink
	}
	return ""
}

func (m *Application) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*GetApplicationRequest)(nil), "GetApplicationRequest")
	proto.RegisterType((*GetApplicationResponse)(nil), "GetApplicationResponse")
	proto.RegisterType((*ListApplicationsRequest)(nil), "ListApplicationsRequest")
	proto.RegisterType((*ListApplicationsResponse)(nil), "ListApplicationsResponse")
	proto.RegisterType((*Application)(nil), "Application")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationServiceClient is the client API for ApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationServiceClient interface {
	GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
	ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error)
}

type applicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServiceClient(cc *grpc.ClientConn) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	out := new(GetApplicationResponse)
	err := c.cc.Invoke(ctx, "/ApplicationService/GetApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServiceClient) ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error) {
	out := new(ListApplicationsResponse)
	err := c.cc.Invoke(ctx, "/ApplicationService/ListApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServiceServer is the server API for ApplicationService service.
type ApplicationServiceServer interface {
	GetApplication(context.Context, *GetApplicationRequest) (*GetApplicationResponse, error)
	ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error)
}

func RegisterApplicationServiceServer(s *grpc.Server, srv ApplicationServiceServer) {
	s.RegisterService(&_ApplicationService_serviceDesc, srv)
}

func _ApplicationService_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApplicationService/GetApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).GetApplication(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationService_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApplicationService/ListApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ListApplications(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApplication",
			Handler:    _ApplicationService_GetApplication_Handler,
		},
		{
			MethodName: "ListApplications",
			Handler:    _ApplicationService_ListApplications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application.proto",
}

func init() { proto.RegisterFile("application.proto", fileDescriptor_application_96596db0568e8f62) }

var fileDescriptor_application_96596db0568e8f62 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4f, 0x4f, 0x3a, 0x31,
	0x10, 0xcd, 0xf2, 0x2f, 0x30, 0xcb, 0x8f, 0x9f, 0x36, 0x11, 0x2a, 0x9a, 0x48, 0xf6, 0xa0, 0x9e,
	0x1a, 0x83, 0x9f, 0x00, 0x2f, 0x68, 0x42, 0x08, 0xc1, 0x9b, 0x97, 0x4d, 0x61, 0x47, 0xd2, 0xb8,
	0xec, 0xd6, 0xb6, 0x4b, 0xe2, 0x87, 0x31, 0xf1, 0xa3, 0x9a, 0x16, 0x94, 0x02, 0x72, 0x6b, 0xde,
	0x7b, 0x9d, 0x79, 0xf3, 0x66, 0xe0, 0x94, 0x4b, 0x99, 0x8a, 0x39, 0x37, 0x22, 0xcf, 0x98, 0x54,
	0xb9, 0xc9, 0xa3, 0x1b, 0x38, 0x1b, 0xa2, 0x19, 0x6c, 0xf1, 0x29, 0xbe, 0x17, 0xa8, 0x0d, 0x69,
	0x41, 0x49, 0x24, 0x34, 0xe8, 0x05, 0xb7, 0xd5, 0x69, 0x49, 0x24, 0xd1, 0x23, 0xb4, 0xf7, 0x85,
	0x5a, 0xe6, 0x99, 0x46, 0xc2, 0x20, 0xf4, 0xea, 0xba, 0x2f, 0x61, 0xbf, 0xc9, 0x7c, 0xa9, 0x2f,
	0x88, 0x56, 0xd0, 0x19, 0x09, 0xed, 0x97, 0xd2, 0x3f, 0x4d, 0x09, 0x54, 0xcc, 0x87, 0xc4, 0x4d,
	0x5b, 0xf7, 0x26, 0x57, 0x10, 0x4a, 0xbe, 0xc0, 0x38, 0x2b, 0x96, 0x33, 0x54, 0xb4, 0xe4, 0x28,
	0xb0, 0xd0, 0xd8, 0x21, 0xe4, 0x1a, 0xfe, 0x2b, 0xd4, 0x45, 0x6a, 0x62, 0x89, 0x2a, 0xb6, 0x04,
	0x2d, 0x3b, 0xd1, 0xbf, 0x35, 0x3c, 0x41, 0x35, 0xe1, 0x0b, 0x8c, 0x46, 0x40, 0x0f, 0xfb, 0x6e,
	0x66, 0xb8, 0x83, 0xa6, 0x67, 0x51, 0xd3, 0xa0, 0x57, 0x3e, 0x18, 0x62, 0x47, 0x11, 0x7d, 0x06,
	0x10, 0x7a, 0xec, 0x7e, 0x5e, 0xe4, 0x02, 0x1a, 0x92, 0x2b, 0xcc, 0x4c, 0x2c, 0x92, 0x8d, 0xe9,
	0xfa, 0x1a, 0x78, 0x4a, 0xec, 0x9c, 0x19, 0x5f, 0xae, 0x7d, 0x36, 0xa6, 0xee, 0x6d, 0x3f, 0xbc,
	0x16, 0x69, 0x1a, 0x3b, 0xa2, 0xe2, 0x88, 0xba, 0x05, 0xc6, 0x96, 0xbc, 0x84, 0x86, 0x44, 0xb5,
	0xe4, 0xa9, 0xc8, 0xde, 0x68, 0xd5, 0x91, 0x5b, 0xe0, 0x37, 0xb6, 0xda, 0x36, 0xb6, 0xfe, 0x57,
	0x00, 0xc4, 0xf3, 0xf7, 0x8c, 0x6a, 0x25, 0xe6, 0x48, 0x06, 0xd0, 0xda, 0x5d, 0x23, 0x69, 0xb3,
	0x3f, 0x0f, 0xa0, 0xdb, 0x61, 0x47, 0xf6, 0x3d, 0x84, 0x93, 0xfd, 0x1c, 0x09, 0x65, 0x47, 0x56,
	0xda, 0x3d, 0x67, 0xc7, 0x42, 0x7f, 0xa8, 0xbe, 0x94, 0xb9, 0x94, 0xb3, 0x9a, 0xbb, 0xc4, 0xfb,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0xe8, 0xdf, 0x0f, 0x9e, 0x02, 0x00, 0x00,
}
