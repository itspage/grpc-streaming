// Code generated by protoc-gen-go.
// source: hello.proto
// DO NOT EDIT!

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package hello

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

type HelloRequest struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type HelloResponse struct {
	Reply string `protobuf:"bytes,1,opt,name=reply" json:"reply,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloService service

type HelloServiceClient interface {
	LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_LotsOfRepliesClient, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) LotsOfReplies(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_LotsOfRepliesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_HelloService_serviceDesc.Streams[0], c.cc, "/HelloService/LotsOfReplies", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceLotsOfRepliesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_LotsOfRepliesClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceLotsOfRepliesClient struct {
	grpc.ClientStream
}

func (x *helloServiceLotsOfRepliesClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for HelloService service

type HelloServiceServer interface {
	LotsOfReplies(*HelloRequest, HelloService_LotsOfRepliesServer) error
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_LotsOfReplies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).LotsOfReplies(m, &helloServiceLotsOfRepliesServer{stream})
}

type HelloService_LotsOfRepliesServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type helloServiceLotsOfRepliesServer struct {
	grpc.ServerStream
}

func (x *helloServiceLotsOfRepliesServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LotsOfReplies",
			Handler:       _HelloService_LotsOfReplies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe2, 0xe2, 0xf1, 0x00, 0x71, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0xd2, 0x8b, 0x52, 0x53, 0x4b, 0x32, 0xf3,
	0xd2, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x55, 0x2e, 0x5e, 0xa8, 0xda,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0xa2, 0xd4, 0x82, 0x9c, 0x4a, 0xa8,
	0x4a, 0x08, 0xc7, 0xc8, 0x09, 0x6a, 0x64, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x11,
	0x17, 0xaf, 0x4f, 0x7e, 0x49, 0xb1, 0x7f, 0x5a, 0x50, 0x6a, 0x41, 0x4e, 0x66, 0x6a, 0xb1, 0x10,
	0xaf, 0x1e, 0xb2, 0x95, 0x52, 0x7c, 0x7a, 0x28, 0xa6, 0x2a, 0x31, 0x18, 0x30, 0x26, 0xb1, 0x81,
	0x5d, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x72, 0x09, 0x8e, 0xac, 0x00, 0x00, 0x00,
}
