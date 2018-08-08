// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sfapi.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type GetFilmRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFilmRequest) Reset()         { *m = GetFilmRequest{} }
func (m *GetFilmRequest) String() string { return proto.CompactTextString(m) }
func (*GetFilmRequest) ProtoMessage()    {}
func (*GetFilmRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sfapi_68aba19ab3b5b8b1, []int{0}
}
func (m *GetFilmRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFilmRequest.Unmarshal(m, b)
}
func (m *GetFilmRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFilmRequest.Marshal(b, m, deterministic)
}
func (dst *GetFilmRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFilmRequest.Merge(dst, src)
}
func (m *GetFilmRequest) XXX_Size() int {
	return xxx_messageInfo_GetFilmRequest.Size(m)
}
func (m *GetFilmRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFilmRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFilmRequest proto.InternalMessageInfo

func (m *GetFilmRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetFilmResponse struct {
	Film                 *Film    `protobuf:"bytes,1,opt,name=film,proto3" json:"film,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFilmResponse) Reset()         { *m = GetFilmResponse{} }
func (m *GetFilmResponse) String() string { return proto.CompactTextString(m) }
func (*GetFilmResponse) ProtoMessage()    {}
func (*GetFilmResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sfapi_68aba19ab3b5b8b1, []int{1}
}
func (m *GetFilmResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFilmResponse.Unmarshal(m, b)
}
func (m *GetFilmResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFilmResponse.Marshal(b, m, deterministic)
}
func (dst *GetFilmResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFilmResponse.Merge(dst, src)
}
func (m *GetFilmResponse) XXX_Size() int {
	return xxx_messageInfo_GetFilmResponse.Size(m)
}
func (m *GetFilmResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFilmResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFilmResponse proto.InternalMessageInfo

func (m *GetFilmResponse) GetFilm() *Film {
	if m != nil {
		return m.Film
	}
	return nil
}

type ListFilmsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFilmsRequest) Reset()         { *m = ListFilmsRequest{} }
func (m *ListFilmsRequest) String() string { return proto.CompactTextString(m) }
func (*ListFilmsRequest) ProtoMessage()    {}
func (*ListFilmsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sfapi_68aba19ab3b5b8b1, []int{2}
}
func (m *ListFilmsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilmsRequest.Unmarshal(m, b)
}
func (m *ListFilmsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilmsRequest.Marshal(b, m, deterministic)
}
func (dst *ListFilmsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilmsRequest.Merge(dst, src)
}
func (m *ListFilmsRequest) XXX_Size() int {
	return xxx_messageInfo_ListFilmsRequest.Size(m)
}
func (m *ListFilmsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilmsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilmsRequest proto.InternalMessageInfo

type ListFilmsResponse struct {
	Films                []*Film  `protobuf:"bytes,1,rep,name=films,proto3" json:"films,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFilmsResponse) Reset()         { *m = ListFilmsResponse{} }
func (m *ListFilmsResponse) String() string { return proto.CompactTextString(m) }
func (*ListFilmsResponse) ProtoMessage()    {}
func (*ListFilmsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_sfapi_68aba19ab3b5b8b1, []int{3}
}
func (m *ListFilmsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilmsResponse.Unmarshal(m, b)
}
func (m *ListFilmsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilmsResponse.Marshal(b, m, deterministic)
}
func (dst *ListFilmsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilmsResponse.Merge(dst, src)
}
func (m *ListFilmsResponse) XXX_Size() int {
	return xxx_messageInfo_ListFilmsResponse.Size(m)
}
func (m *ListFilmsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilmsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilmsResponse proto.InternalMessageInfo

func (m *ListFilmsResponse) GetFilms() []*Film {
	if m != nil {
		return m.Films
	}
	return nil
}

type Film struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Director             string               `protobuf:"bytes,3,opt,name=director,proto3" json:"director,omitempty"`
	Producer             string               `protobuf:"bytes,4,opt,name=producer,proto3" json:"producer,omitempty"`
	ReleaseDate          *timestamp.Timestamp `protobuf:"bytes,5,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Film) Reset()         { *m = Film{} }
func (m *Film) String() string { return proto.CompactTextString(m) }
func (*Film) ProtoMessage()    {}
func (*Film) Descriptor() ([]byte, []int) {
	return fileDescriptor_sfapi_68aba19ab3b5b8b1, []int{4}
}
func (m *Film) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Film.Unmarshal(m, b)
}
func (m *Film) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Film.Marshal(b, m, deterministic)
}
func (dst *Film) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Film.Merge(dst, src)
}
func (m *Film) XXX_Size() int {
	return xxx_messageInfo_Film.Size(m)
}
func (m *Film) XXX_DiscardUnknown() {
	xxx_messageInfo_Film.DiscardUnknown(m)
}

var xxx_messageInfo_Film proto.InternalMessageInfo

func (m *Film) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Film) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Film) GetDirector() string {
	if m != nil {
		return m.Director
	}
	return ""
}

func (m *Film) GetProducer() string {
	if m != nil {
		return m.Producer
	}
	return ""
}

func (m *Film) GetReleaseDate() *timestamp.Timestamp {
	if m != nil {
		return m.ReleaseDate
	}
	return nil
}

func init() {
	proto.RegisterType((*GetFilmRequest)(nil), "GetFilmRequest")
	proto.RegisterType((*GetFilmResponse)(nil), "GetFilmResponse")
	proto.RegisterType((*ListFilmsRequest)(nil), "ListFilmsRequest")
	proto.RegisterType((*ListFilmsResponse)(nil), "ListFilmsResponse")
	proto.RegisterType((*Film)(nil), "Film")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StarfriendsClient is the client API for Starfriends service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StarfriendsClient interface {
	// Get a single Film by unique ID
	GetFilm(ctx context.Context, in *GetFilmRequest, opts ...grpc.CallOption) (*GetFilmResponse, error)
	// Get a list of all Films
	ListFilms(ctx context.Context, in *ListFilmsRequest, opts ...grpc.CallOption) (*ListFilmsResponse, error)
}

type starfriendsClient struct {
	cc *grpc.ClientConn
}

func NewStarfriendsClient(cc *grpc.ClientConn) StarfriendsClient {
	return &starfriendsClient{cc}
}

func (c *starfriendsClient) GetFilm(ctx context.Context, in *GetFilmRequest, opts ...grpc.CallOption) (*GetFilmResponse, error) {
	out := new(GetFilmResponse)
	err := c.cc.Invoke(ctx, "/Starfriends/GetFilm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starfriendsClient) ListFilms(ctx context.Context, in *ListFilmsRequest, opts ...grpc.CallOption) (*ListFilmsResponse, error) {
	out := new(ListFilmsResponse)
	err := c.cc.Invoke(ctx, "/Starfriends/ListFilms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StarfriendsServer is the server API for Starfriends service.
type StarfriendsServer interface {
	// Get a single Film by unique ID
	GetFilm(context.Context, *GetFilmRequest) (*GetFilmResponse, error)
	// Get a list of all Films
	ListFilms(context.Context, *ListFilmsRequest) (*ListFilmsResponse, error)
}

func RegisterStarfriendsServer(s *grpc.Server, srv StarfriendsServer) {
	s.RegisterService(&_Starfriends_serviceDesc, srv)
}

func _Starfriends_GetFilm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarfriendsServer).GetFilm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Starfriends/GetFilm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarfriendsServer).GetFilm(ctx, req.(*GetFilmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starfriends_ListFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarfriendsServer).ListFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Starfriends/ListFilms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarfriendsServer).ListFilms(ctx, req.(*ListFilmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Starfriends_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Starfriends",
	HandlerType: (*StarfriendsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFilm",
			Handler:    _Starfriends_GetFilm_Handler,
		},
		{
			MethodName: "ListFilms",
			Handler:    _Starfriends_ListFilms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sfapi.proto",
}

func init() { proto.RegisterFile("sfapi.proto", fileDescriptor_sfapi_68aba19ab3b5b8b1) }

var fileDescriptor_sfapi_68aba19ab3b5b8b1 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x95, 0xb6, 0xa1, 0xf4, 0x04, 0xf5, 0x62, 0x31, 0x98, 0x30, 0x10, 0x65, 0x62, 0xa8,
	0x5c, 0x14, 0x66, 0x16, 0x84, 0x60, 0x61, 0x0a, 0x4c, 0x2c, 0xc8, 0xd4, 0x27, 0x95, 0xa5, 0xa4,
	0x0e, 0xb6, 0xf3, 0x44, 0xbc, 0x28, 0x8a, 0x9d, 0x04, 0x08, 0x53, 0xf4, 0x5f, 0xf2, 0xeb, 0xf8,
	0x83, 0xc8, 0x14, 0xbc, 0x96, 0xac, 0xd6, 0xca, 0xaa, 0xf8, 0xea, 0xa0, 0xd4, 0xa1, 0xc4, 0x9d,
	0x53, 0x1f, 0x4d, 0xb1, 0xb3, 0xb2, 0x42, 0x63, 0x79, 0x55, 0xfb, 0x42, 0x9a, 0xc0, 0xf2, 0x09,
	0xed, 0xa3, 0x2c, 0xab, 0x1c, 0x3f, 0x1b, 0x34, 0x96, 0x2c, 0x61, 0x22, 0x05, 0x0d, 0x92, 0xe0,
	0x7a, 0x91, 0x4f, 0xa4, 0x48, 0xb7, 0xb0, 0x1a, 0x1a, 0xa6, 0x56, 0x47, 0x83, 0xe4, 0x02, 0x66,
	0x85, 0x2c, 0x2b, 0x57, 0x8a, 0xb2, 0x90, 0xb9, 0xd0, 0x59, 0x29, 0x81, 0xf5, 0xb3, 0x34, 0xae,
	0x6e, 0xba, 0xc5, 0xf4, 0x06, 0x36, 0xbf, 0xbc, 0x6e, 0xe3, 0x12, 0xc2, 0xf6, 0x07, 0x43, 0x83,
	0x64, 0xfa, 0x33, 0xe2, 0xbd, 0xf4, 0x2b, 0x80, 0x59, 0xab, 0xc7, 0xc7, 0x90, 0x73, 0x08, 0xad,
	0xb4, 0x25, 0xd2, 0x89, 0xb3, 0xbc, 0x20, 0x31, 0x9c, 0x0a, 0xa9, 0x71, 0x6f, 0x95, 0xa6, 0x53,
	0x17, 0x0c, 0xba, 0xcd, 0x6a, 0xad, 0x44, 0xb3, 0x47, 0x4d, 0x67, 0x3e, 0xeb, 0x35, 0xb9, 0x83,
	0x33, 0x8d, 0x25, 0x72, 0x83, 0xef, 0x82, 0x5b, 0xa4, 0xa1, 0x7b, 0x4f, 0xcc, 0x3c, 0x34, 0xd6,
	0x43, 0x63, 0xaf, 0x3d, 0xb4, 0x3c, 0xea, 0xfa, 0x0f, 0xdc, 0x62, 0xa6, 0x20, 0x7a, 0xb1, 0x5c,
	0x17, 0x5a, 0xe2, 0x51, 0x18, 0xb2, 0x85, 0x79, 0x07, 0x8a, 0xac, 0xd8, 0x5f, 0xa8, 0xf1, 0x9a,
	0x8d, 0x19, 0x66, 0xb0, 0x18, 0xa0, 0x90, 0x0d, 0x1b, 0x43, 0x8b, 0x09, 0xfb, 0xc7, 0xec, 0x7e,
	0xfe, 0x16, 0xfa, 0x9b, 0x4e, 0xdc, 0xe7, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x0f, 0x3e,
	0x7c, 0xec, 0x01, 0x00, 0x00,
}