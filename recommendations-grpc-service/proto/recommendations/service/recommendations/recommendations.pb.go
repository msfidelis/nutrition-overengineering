// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: recommendations.proto

package recommendations

import (
	context "context"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weight   float64 `protobuf:"fixed64,1,opt,name=weight,proto3" json:"weight,omitempty"`
	Height   float64 `protobuf:"fixed64,2,opt,name=height,proto3" json:"height,omitempty"`
	Calories float64 `protobuf:"fixed64,3,opt,name=calories,proto3" json:"calories,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_recommendations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_recommendations_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Message) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Message) GetCalories() float64 {
	if x != nil {
		return x.Calories
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WaterValue         float64 `protobuf:"fixed64,1,opt,name=waterValue,proto3" json:"waterValue,omitempty"`
	WaterUnit          string  `protobuf:"bytes,2,opt,name=waterUnit,proto3" json:"waterUnit,omitempty"`
	ProteinsUnit       string  `protobuf:"bytes,3,opt,name=proteinsUnit,proto3" json:"proteinsUnit,omitempty"`
	ProteinsValue      int64   `protobuf:"varint,4,opt,name=proteinsValue,proto3" json:"proteinsValue,omitempty"`
	CaloriesToLoss     float64 `protobuf:"fixed64,5,opt,name=caloriesToLoss,proto3" json:"caloriesToLoss,omitempty"`
	CaloriesToMaintein float64 `protobuf:"fixed64,6,opt,name=caloriesToMaintein,proto3" json:"caloriesToMaintein,omitempty"`
	CaloriesToGain     float64 `protobuf:"fixed64,7,opt,name=caloriesToGain,proto3" json:"caloriesToGain,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_recommendations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_recommendations_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetWaterValue() float64 {
	if x != nil {
		return x.WaterValue
	}
	return 0
}

func (x *Response) GetWaterUnit() string {
	if x != nil {
		return x.WaterUnit
	}
	return ""
}

func (x *Response) GetProteinsUnit() string {
	if x != nil {
		return x.ProteinsUnit
	}
	return ""
}

func (x *Response) GetProteinsValue() int64 {
	if x != nil {
		return x.ProteinsValue
	}
	return 0
}

func (x *Response) GetCaloriesToLoss() float64 {
	if x != nil {
		return x.CaloriesToLoss
	}
	return 0
}

func (x *Response) GetCaloriesToMaintein() float64 {
	if x != nil {
		return x.CaloriesToMaintein
	}
	return 0
}

func (x *Response) GetCaloriesToGain() float64 {
	if x != nil {
		return x.CaloriesToGain
	}
	return 0
}

var File_recommendations_proto protoreflect.FileDescriptor

var file_recommendations_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x55, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22,
	0x92, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x77, 0x61, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x77, 0x61, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x77, 0x61, 0x74, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x54, 0x6f, 0x4c, 0x6f, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x61,
	0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x12,
	0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x54, 0x6f, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0e,
	0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x54, 0x6f, 0x47, 0x61, 0x69, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x63, 0x61, 0x6c, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x54, 0x6f,
	0x47, 0x61, 0x69, 0x6e, 0x32, 0x5a, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x19, 0x5a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_recommendations_proto_rawDescOnce sync.Once
	file_recommendations_proto_rawDescData = file_recommendations_proto_rawDesc
)

func file_recommendations_proto_rawDescGZIP() []byte {
	file_recommendations_proto_rawDescOnce.Do(func() {
		file_recommendations_proto_rawDescData = protoimpl.X.CompressGZIP(file_recommendations_proto_rawDescData)
	})
	return file_recommendations_proto_rawDescData
}

var file_recommendations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_recommendations_proto_goTypes = []interface{}{
	(*Message)(nil),  // 0: recommendations.Message
	(*Response)(nil), // 1: recommendations.Response
}
var file_recommendations_proto_depIdxs = []int32{
	0, // 0: recommendations.RecomendationsService.SayHello:input_type -> recommendations.Message
	1, // 1: recommendations.RecomendationsService.SayHello:output_type -> recommendations.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_recommendations_proto_init() }
func file_recommendations_proto_init() {
	if File_recommendations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recommendations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_recommendations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_recommendations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recommendations_proto_goTypes,
		DependencyIndexes: file_recommendations_proto_depIdxs,
		MessageInfos:      file_recommendations_proto_msgTypes,
	}.Build()
	File_recommendations_proto = out.File
	file_recommendations_proto_rawDesc = nil
	file_recommendations_proto_goTypes = nil
	file_recommendations_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RecomendationsServiceClient is the client API for RecomendationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecomendationsServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error)
}

type recomendationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecomendationsServiceClient(cc grpc.ClientConnInterface) RecomendationsServiceClient {
	return &recomendationsServiceClient{cc}
}

func (c *recomendationsServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/recommendations.RecomendationsService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecomendationsServiceServer is the server API for RecomendationsService service.
type RecomendationsServiceServer interface {
	SayHello(context.Context, *Message) (*Response, error)
}

// UnimplementedRecomendationsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRecomendationsServiceServer struct {
}

func (*UnimplementedRecomendationsServiceServer) SayHello(context.Context, *Message) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterRecomendationsServiceServer(s *grpc.Server, srv RecomendationsServiceServer) {
	s.RegisterService(&_RecomendationsService_serviceDesc, srv)
}

func _RecomendationsService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecomendationsServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recommendations.RecomendationsService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecomendationsServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecomendationsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recommendations.RecomendationsService",
	HandlerType: (*RecomendationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _RecomendationsService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommendations.proto",
}
