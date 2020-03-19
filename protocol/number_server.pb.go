// Code generated by protoc-gen-go. DO NOT EDIT.
// source: number_server.proto

package grpc_protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NumberMessage struct {
	Number               string   `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberMessage) Reset()         { *m = NumberMessage{} }
func (m *NumberMessage) String() string { return proto.CompactTextString(m) }
func (*NumberMessage) ProtoMessage()    {}
func (*NumberMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8034ca23551f4c, []int{0}
}

func (m *NumberMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberMessage.Unmarshal(m, b)
}
func (m *NumberMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberMessage.Marshal(b, m, deterministic)
}
func (m *NumberMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberMessage.Merge(m, src)
}
func (m *NumberMessage) XXX_Size() int {
	return xxx_messageInfo_NumberMessage.Size(m)
}
func (m *NumberMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NumberMessage proto.InternalMessageInfo

func (m *NumberMessage) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8034ca23551f4c, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NumberMessage)(nil), "grpc_protocol.NumberMessage")
	proto.RegisterType((*Empty)(nil), "grpc_protocol.Empty")
}

func init() {
	proto.RegisterFile("number_server.proto", fileDescriptor_2a8034ca23551f4c)
}

var fileDescriptor_2a8034ca23551f4c = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x2b, 0xcd, 0x4d,
	0x4a, 0x2d, 0x8a, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x4d, 0x2f, 0x2a, 0x48, 0x8e, 0x07, 0xb3, 0x93, 0xf3, 0x73, 0x94, 0xd4, 0xb9, 0x78, 0xfd,
	0xc0, 0xaa, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x20, 0xda, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x25, 0x76, 0x2e, 0x56, 0xd7, 0xdc, 0x82, 0x92,
	0x4a, 0xa3, 0x60, 0x2e, 0x1e, 0x88, 0x8e, 0x60, 0xb0, 0xb1, 0x42, 0xce, 0x5c, 0xdc, 0xc1, 0xa9,
	0x79, 0x29, 0x30, 0xfd, 0x32, 0x7a, 0x28, 0x16, 0xe8, 0xa1, 0x98, 0x2e, 0x25, 0x82, 0x26, 0x0b,
	0x36, 0x52, 0x89, 0x21, 0x89, 0x0d, 0x2c, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x61, 0xec,
	0x77, 0xd8, 0xb3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NumberServerClient is the client API for NumberServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NumberServerClient interface {
	SendMessage(ctx context.Context, in *NumberMessage, opts ...grpc.CallOption) (*Empty, error)
}

type numberServerClient struct {
	cc grpc.ClientConnInterface
}

func NewNumberServerClient(cc grpc.ClientConnInterface) NumberServerClient {
	return &numberServerClient{cc}
}

func (c *numberServerClient) SendMessage(ctx context.Context, in *NumberMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc_protocol.NumberServer/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NumberServerServer is the server API for NumberServer service.
type NumberServerServer interface {
	SendMessage(context.Context, *NumberMessage) (*Empty, error)
}

// UnimplementedNumberServerServer can be embedded to have forward compatible implementations.
type UnimplementedNumberServerServer struct {
}

func (*UnimplementedNumberServerServer) SendMessage(ctx context.Context, req *NumberMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterNumberServerServer(s *grpc.Server, srv NumberServerServer) {
	s.RegisterService(&_NumberServer_serviceDesc, srv)
}

func _NumberServer_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberServerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_protocol.NumberServer/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberServerServer).SendMessage(ctx, req.(*NumberMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NumberServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_protocol.NumberServer",
	HandlerType: (*NumberServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _NumberServer_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "number_server.proto",
}