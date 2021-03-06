// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: subscription.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sports  []string `protobuf:"bytes,1,rep,name=Sports,proto3" json:"Sports,omitempty"`
	Seconds int32    `protobuf:"varint,2,opt,name=Seconds,proto3" json:"Seconds,omitempty"`
}

func (x *SubRequest) Reset() {
	*x = SubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubRequest) ProtoMessage() {}

func (x *SubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubRequest.ProtoReflect.Descriptor instead.
func (*SubRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *SubRequest) GetSports() []string {
	if x != nil {
		return x.Sports
	}
	return nil
}

func (x *SubRequest) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type Line struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sport string `protobuf:"bytes,1,opt,name=Sport,proto3" json:"Sport,omitempty"`
	Coef  string `protobuf:"bytes,2,opt,name=Coef,proto3" json:"Coef,omitempty"`
}

func (x *Line) Reset() {
	*x = Line{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Line) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Line) ProtoMessage() {}

func (x *Line) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Line.ProtoReflect.Descriptor instead.
func (*Line) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *Line) GetSport() string {
	if x != nil {
		return x.Sport
	}
	return ""
}

func (x *Line) GetCoef() string {
	if x != nil {
		return x.Coef
	}
	return ""
}

type SubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines []*Line `protobuf:"bytes,1,rep,name=Lines,proto3" json:"Lines,omitempty"`
}

func (x *SubResponse) Reset() {
	*x = SubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubResponse) ProtoMessage() {}

func (x *SubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubResponse.ProtoReflect.Descriptor instead.
func (*SubResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *SubResponse) GetLines() []*Line {
	if x != nil {
		return x.Lines
	}
	return nil
}

var File_subscription_proto protoreflect.FileDescriptor

var file_subscription_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0a, 0x53,
	0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x53, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x30, 0x0a, 0x04, 0x4c,
	0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x65,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x65, 0x66, 0x22, 0x30, 0x0a,
	0x0b, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05,
	0x4c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x05, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x32,
	0x51, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x6e, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscription_proto_rawDescOnce sync.Once
	file_subscription_proto_rawDescData = file_subscription_proto_rawDesc
)

func file_subscription_proto_rawDescGZIP() []byte {
	file_subscription_proto_rawDescOnce.Do(func() {
		file_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscription_proto_rawDescData)
	})
	return file_subscription_proto_rawDescData
}

var file_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_subscription_proto_goTypes = []interface{}{
	(*SubRequest)(nil),  // 0: proto.SubRequest
	(*Line)(nil),        // 1: proto.Line
	(*SubResponse)(nil), // 2: proto.SubResponse
}
var file_subscription_proto_depIdxs = []int32{
	1, // 0: proto.SubResponse.Lines:type_name -> proto.Line
	0, // 1: proto.SubService.SubscribeOnSportsLines:input_type -> proto.SubRequest
	2, // 2: proto.SubService.SubscribeOnSportsLines:output_type -> proto.SubResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_subscription_proto_init() }
func file_subscription_proto_init() {
	if File_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Line); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscription_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscription_proto_goTypes,
		DependencyIndexes: file_subscription_proto_depIdxs,
		MessageInfos:      file_subscription_proto_msgTypes,
	}.Build()
	File_subscription_proto = out.File
	file_subscription_proto_rawDesc = nil
	file_subscription_proto_goTypes = nil
	file_subscription_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SubServiceClient is the client API for SubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubServiceClient interface {
	SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (SubService_SubscribeOnSportsLinesClient, error)
}

type subServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubServiceClient(cc grpc.ClientConnInterface) SubServiceClient {
	return &subServiceClient{cc}
}

func (c *subServiceClient) SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (SubService_SubscribeOnSportsLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SubService_serviceDesc.Streams[0], "/proto.SubService/SubscribeOnSportsLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &subServiceSubscribeOnSportsLinesClient{stream}
	return x, nil
}

type SubService_SubscribeOnSportsLinesClient interface {
	Send(*SubRequest) error
	Recv() (*SubResponse, error)
	grpc.ClientStream
}

type subServiceSubscribeOnSportsLinesClient struct {
	grpc.ClientStream
}

func (x *subServiceSubscribeOnSportsLinesClient) Send(m *SubRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *subServiceSubscribeOnSportsLinesClient) Recv() (*SubResponse, error) {
	m := new(SubResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubServiceServer is the server API for SubService service.
type SubServiceServer interface {
	SubscribeOnSportsLines(SubService_SubscribeOnSportsLinesServer) error
}

// UnimplementedSubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSubServiceServer struct {
}

func (*UnimplementedSubServiceServer) SubscribeOnSportsLines(SubService_SubscribeOnSportsLinesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeOnSportsLines not implemented")
}

func RegisterSubServiceServer(s *grpc.Server, srv SubServiceServer) {
	s.RegisterService(&_SubService_serviceDesc, srv)
}

func _SubService_SubscribeOnSportsLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SubServiceServer).SubscribeOnSportsLines(&subServiceSubscribeOnSportsLinesServer{stream})
}

type SubService_SubscribeOnSportsLinesServer interface {
	Send(*SubResponse) error
	Recv() (*SubRequest, error)
	grpc.ServerStream
}

type subServiceSubscribeOnSportsLinesServer struct {
	grpc.ServerStream
}

func (x *subServiceSubscribeOnSportsLinesServer) Send(m *SubResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *subServiceSubscribeOnSportsLinesServer) Recv() (*SubRequest, error) {
	m := new(SubRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SubService",
	HandlerType: (*SubServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeOnSportsLines",
			Handler:       _SubService_SubscribeOnSportsLines_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "subscription.proto",
}
