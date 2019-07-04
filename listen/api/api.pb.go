// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Nope struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nope) Reset()         { *m = Nope{} }
func (m *Nope) String() string { return proto.CompactTextString(m) }
func (*Nope) ProtoMessage()    {}
func (*Nope) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *Nope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nope.Unmarshal(m, b)
}
func (m *Nope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nope.Marshal(b, m, deterministic)
}
func (m *Nope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nope.Merge(m, src)
}
func (m *Nope) XXX_Size() int {
	return xxx_messageInfo_Nope.Size(m)
}
func (m *Nope) XXX_DiscardUnknown() {
	xxx_messageInfo_Nope.DiscardUnknown(m)
}

var xxx_messageInfo_Nope proto.InternalMessageInfo

// TopicName contains Watermill topic name
type TopicName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicName) Reset()         { *m = TopicName{} }
func (m *TopicName) String() string { return proto.CompactTextString(m) }
func (*TopicName) ProtoMessage()    {}
func (*TopicName) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *TopicName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicName.Unmarshal(m, b)
}
func (m *TopicName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicName.Marshal(b, m, deterministic)
}
func (m *TopicName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicName.Merge(m, src)
}
func (m *TopicName) XXX_Size() int {
	return xxx_messageInfo_TopicName.Size(m)
}
func (m *TopicName) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicName.DiscardUnknown(m)
}

var xxx_messageInfo_TopicName proto.InternalMessageInfo

func (m *TopicName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// TopicName contains Watermill topic names.
type TopicNames struct {
	Name                 []string `protobuf:"bytes,1,rep,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicNames) Reset()         { *m = TopicNames{} }
func (m *TopicNames) String() string { return proto.CompactTextString(m) }
func (*TopicNames) ProtoMessage()    {}
func (*TopicNames) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{2}
}

func (m *TopicNames) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicNames.Unmarshal(m, b)
}
func (m *TopicNames) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicNames.Marshal(b, m, deterministic)
}
func (m *TopicNames) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicNames.Merge(m, src)
}
func (m *TopicNames) XXX_Size() int {
	return xxx_messageInfo_TopicNames.Size(m)
}
func (m *TopicNames) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicNames.DiscardUnknown(m)
}

var xxx_messageInfo_TopicNames proto.InternalMessageInfo

func (m *TopicNames) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

// TopicContent contains list of Key=Value pairs of Metadata field from received messages.
type TopicContent struct {
	KV                   []string `protobuf:"bytes,1,rep,name=KV,proto3" json:"KV,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicContent) Reset()         { *m = TopicContent{} }
func (m *TopicContent) String() string { return proto.CompactTextString(m) }
func (*TopicContent) ProtoMessage()    {}
func (*TopicContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{3}
}

func (m *TopicContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicContent.Unmarshal(m, b)
}
func (m *TopicContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicContent.Marshal(b, m, deterministic)
}
func (m *TopicContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicContent.Merge(m, src)
}
func (m *TopicContent) XXX_Size() int {
	return xxx_messageInfo_TopicContent.Size(m)
}
func (m *TopicContent) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicContent.DiscardUnknown(m)
}

var xxx_messageInfo_TopicContent proto.InternalMessageInfo

func (m *TopicContent) GetKV() []string {
	if m != nil {
		return m.KV
	}
	return nil
}

func init() {
	proto.RegisterType((*Nope)(nil), "api.Nope")
	proto.RegisterType((*TopicName)(nil), "api.TopicName")
	proto.RegisterType((*TopicNames)(nil), "api.TopicNames")
	proto.RegisterType((*TopicContent)(nil), "api.TopicContent")
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784) }

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92,
	0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x4b, 0xe5, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64,
	0xe6, 0xe7, 0x15, 0x43, 0x94, 0x28, 0xb1, 0x71, 0xb1, 0xf8, 0xe5, 0x17, 0xa4, 0x2a, 0xc9, 0x73,
	0x71, 0x86, 0xe4, 0x17, 0x64, 0x26, 0xfb, 0x25, 0xe6, 0xa6, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25,
	0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x0a, 0x5c, 0x5c, 0x70,
	0x05, 0xc5, 0x20, 0x15, 0x7e, 0x10, 0x15, 0xcc, 0x20, 0x15, 0x20, 0xb6, 0x92, 0x1c, 0x17, 0x0f,
	0x58, 0x85, 0x73, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x89, 0x10, 0x1f, 0x17, 0x93, 0x77, 0x18, 0x54,
	0x05, 0x93, 0x77, 0x98, 0x51, 0x1b, 0x23, 0x17, 0x67, 0x40, 0x69, 0x52, 0x4e, 0x66, 0x71, 0x46,
	0x6a, 0x91, 0x90, 0x05, 0x17, 0x1b, 0x58, 0x75, 0xb1, 0x10, 0xa7, 0x1e, 0xc8, 0xc5, 0x20, 0x57,
	0x48, 0xf1, 0x83, 0x99, 0x08, 0x7b, 0x94, 0x84, 0x9a, 0x2e, 0x3f, 0x99, 0xcc, 0xc4, 0xa3, 0xc4,
	0xae, 0x5f, 0x02, 0x56, 0x6c, 0xc5, 0xa8, 0x25, 0x64, 0xcf, 0xc5, 0x0a, 0x56, 0x21, 0xc4, 0x87,
	0xaa, 0x5a, 0x4a, 0x10, 0xc1, 0x87, 0xba, 0x41, 0x49, 0x10, 0xac, 0x9f, 0x5b, 0x89, 0x0d, 0xa2,
	0xdf, 0x8a, 0x51, 0xcb, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x75, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x82, 0x9b, 0xf0, 0x51, 0x2e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PublisherClient is the client API for Publisher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublisherClient interface {
	// Topics returns all registered topics.
	Topics(ctx context.Context, in *Nope, opts ...grpc.CallOption) (*TopicNames, error)
	// Topic returns content for provided topic name.
	Topic(ctx context.Context, in *TopicName, opts ...grpc.CallOption) (Publisher_TopicClient, error)
}

type publisherClient struct {
	cc *grpc.ClientConn
}

func NewPublisherClient(cc *grpc.ClientConn) PublisherClient {
	return &publisherClient{cc}
}

func (c *publisherClient) Topics(ctx context.Context, in *Nope, opts ...grpc.CallOption) (*TopicNames, error) {
	out := new(TopicNames)
	err := c.cc.Invoke(ctx, "/api.Publisher/Topics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherClient) Topic(ctx context.Context, in *TopicName, opts ...grpc.CallOption) (Publisher_TopicClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Publisher_serviceDesc.Streams[0], "/api.Publisher/Topic", opts...)
	if err != nil {
		return nil, err
	}
	x := &publisherTopicClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Publisher_TopicClient interface {
	Recv() (*TopicContent, error)
	grpc.ClientStream
}

type publisherTopicClient struct {
	grpc.ClientStream
}

func (x *publisherTopicClient) Recv() (*TopicContent, error) {
	m := new(TopicContent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PublisherServer is the server API for Publisher service.
type PublisherServer interface {
	// Topics returns all registered topics.
	Topics(context.Context, *Nope) (*TopicNames, error)
	// Topic returns content for provided topic name.
	Topic(*TopicName, Publisher_TopicServer) error
}

// UnimplementedPublisherServer can be embedded to have forward compatible implementations.
type UnimplementedPublisherServer struct {
}

func (*UnimplementedPublisherServer) Topics(ctx context.Context, req *Nope) (*TopicNames, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Topics not implemented")
}
func (*UnimplementedPublisherServer) Topic(req *TopicName, srv Publisher_TopicServer) error {
	return status.Errorf(codes.Unimplemented, "method Topic not implemented")
}

func RegisterPublisherServer(s *grpc.Server, srv PublisherServer) {
	s.RegisterService(&_Publisher_serviceDesc, srv)
}

func _Publisher_Topics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).Topics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Publisher/Topics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).Topics(ctx, req.(*Nope))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publisher_Topic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TopicName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PublisherServer).Topic(m, &publisherTopicServer{stream})
}

type Publisher_TopicServer interface {
	Send(*TopicContent) error
	grpc.ServerStream
}

type publisherTopicServer struct {
	grpc.ServerStream
}

func (x *publisherTopicServer) Send(m *TopicContent) error {
	return x.ServerStream.SendMsg(m)
}

var _Publisher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Publisher",
	HandlerType: (*PublisherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Topics",
			Handler:    _Publisher_Topics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Topic",
			Handler:       _Publisher_Topic_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/api.proto",
}
