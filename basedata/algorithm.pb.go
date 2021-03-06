// Code generated by protoc-gen-go. DO NOT EDIT.
// source: algorithm.proto

package basedata

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Algorithm struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CratedAtTs           int64    `protobuf:"varint,3,opt,name=crated_at_ts,json=cratedAtTs,proto3" json:"crated_at_ts,omitempty"`
	UpdatedAtTs          int64    `protobuf:"varint,4,opt,name=updated_at_ts,json=updatedAtTs,proto3" json:"updated_at_ts,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Algorithm) Reset()         { *m = Algorithm{} }
func (m *Algorithm) String() string { return proto.CompactTextString(m) }
func (*Algorithm) ProtoMessage()    {}
func (*Algorithm) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{0}
}

func (m *Algorithm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Algorithm.Unmarshal(m, b)
}
func (m *Algorithm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Algorithm.Marshal(b, m, deterministic)
}
func (m *Algorithm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Algorithm.Merge(m, src)
}
func (m *Algorithm) XXX_Size() int {
	return xxx_messageInfo_Algorithm.Size(m)
}
func (m *Algorithm) XXX_DiscardUnknown() {
	xxx_messageInfo_Algorithm.DiscardUnknown(m)
}

var xxx_messageInfo_Algorithm proto.InternalMessageInfo

func (m *Algorithm) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Algorithm) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Algorithm) GetCratedAtTs() int64 {
	if m != nil {
		return m.CratedAtTs
	}
	return 0
}

func (m *Algorithm) GetUpdatedAtTs() int64 {
	if m != nil {
		return m.UpdatedAtTs
	}
	return 0
}

func (m *Algorithm) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Algorithm) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CountAlgorithmsResponse struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountAlgorithmsResponse) Reset()         { *m = CountAlgorithmsResponse{} }
func (m *CountAlgorithmsResponse) String() string { return proto.CompactTextString(m) }
func (*CountAlgorithmsResponse) ProtoMessage()    {}
func (*CountAlgorithmsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{1}
}

func (m *CountAlgorithmsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountAlgorithmsResponse.Unmarshal(m, b)
}
func (m *CountAlgorithmsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountAlgorithmsResponse.Marshal(b, m, deterministic)
}
func (m *CountAlgorithmsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountAlgorithmsResponse.Merge(m, src)
}
func (m *CountAlgorithmsResponse) XXX_Size() int {
	return xxx_messageInfo_CountAlgorithmsResponse.Size(m)
}
func (m *CountAlgorithmsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountAlgorithmsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountAlgorithmsResponse proto.InternalMessageInfo

func (m *CountAlgorithmsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetAlgorithmRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAlgorithmRequest) Reset()         { *m = GetAlgorithmRequest{} }
func (m *GetAlgorithmRequest) String() string { return proto.CompactTextString(m) }
func (*GetAlgorithmRequest) ProtoMessage()    {}
func (*GetAlgorithmRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{2}
}

func (m *GetAlgorithmRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlgorithmRequest.Unmarshal(m, b)
}
func (m *GetAlgorithmRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlgorithmRequest.Marshal(b, m, deterministic)
}
func (m *GetAlgorithmRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlgorithmRequest.Merge(m, src)
}
func (m *GetAlgorithmRequest) XXX_Size() int {
	return xxx_messageInfo_GetAlgorithmRequest.Size(m)
}
func (m *GetAlgorithmRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlgorithmRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlgorithmRequest proto.InternalMessageInfo

func (m *GetAlgorithmRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAlgorithmResponse struct {
	Algorithm            *Algorithm `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAlgorithmResponse) Reset()         { *m = GetAlgorithmResponse{} }
func (m *GetAlgorithmResponse) String() string { return proto.CompactTextString(m) }
func (*GetAlgorithmResponse) ProtoMessage()    {}
func (*GetAlgorithmResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{3}
}

func (m *GetAlgorithmResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlgorithmResponse.Unmarshal(m, b)
}
func (m *GetAlgorithmResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlgorithmResponse.Marshal(b, m, deterministic)
}
func (m *GetAlgorithmResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlgorithmResponse.Merge(m, src)
}
func (m *GetAlgorithmResponse) XXX_Size() int {
	return xxx_messageInfo_GetAlgorithmResponse.Size(m)
}
func (m *GetAlgorithmResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlgorithmResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlgorithmResponse proto.InternalMessageInfo

func (m *GetAlgorithmResponse) GetAlgorithm() *Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

func init() {
	proto.RegisterType((*Algorithm)(nil), "basedata.Algorithm")
	proto.RegisterType((*CountAlgorithmsResponse)(nil), "basedata.CountAlgorithmsResponse")
	proto.RegisterType((*GetAlgorithmRequest)(nil), "basedata.GetAlgorithmRequest")
	proto.RegisterType((*GetAlgorithmResponse)(nil), "basedata.GetAlgorithmResponse")
}

func init() { proto.RegisterFile("algorithm.proto", fileDescriptor_8c381a4f1e580eed) }

var fileDescriptor_8c381a4f1e580eed = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xdd, 0x4a, 0xfb, 0x40,
	0x10, 0xc5, 0xff, 0xdb, 0x2f, 0xfe, 0x99, 0x56, 0x2b, 0xdb, 0xa2, 0xa1, 0x52, 0x89, 0x0b, 0x42,
	0xaf, 0xb6, 0x58, 0x9f, 0xa0, 0x88, 0x88, 0x37, 0x0a, 0xd1, 0xfb, 0xb2, 0x4d, 0xc6, 0x1a, 0x68,
	0xb2, 0x31, 0xbb, 0x11, 0x7c, 0xae, 0xbe, 0xa0, 0x64, 0xf3, 0xb1, 0xb5, 0xd8, 0xbb, 0x64, 0xce,
	0x6f, 0x0e, 0xe7, 0xec, 0xc0, 0x50, 0x6c, 0x37, 0x32, 0x8b, 0xf4, 0x47, 0xcc, 0xd3, 0x4c, 0x6a,
	0x49, 0xff, 0xaf, 0x85, 0xc2, 0x50, 0x68, 0x31, 0xb9, 0xdc, 0x48, 0xb9, 0xd9, 0xe2, 0xdc, 0xcc,
	0xd7, 0xf9, 0xfb, 0x1c, 0xe3, 0x54, 0x7f, 0x97, 0x18, 0xdb, 0x11, 0x70, 0x96, 0xf5, 0x2a, 0x3d,
	0x85, 0x56, 0x14, 0xba, 0xc4, 0x23, 0xb3, 0xae, 0xdf, 0x8a, 0x42, 0x4a, 0xa1, 0x93, 0x88, 0x18,
	0xdd, 0x96, 0x47, 0x66, 0x8e, 0x6f, 0xbe, 0xa9, 0x07, 0x83, 0x20, 0x13, 0x1a, 0xc3, 0x95, 0xd0,
	0x2b, 0xad, 0xdc, 0xb6, 0x47, 0x66, 0x6d, 0x1f, 0xca, 0xd9, 0x52, 0xbf, 0x29, 0xca, 0xe0, 0x24,
	0x4f, 0xc3, 0x3d, 0xa4, 0x63, 0x90, 0x7e, 0x35, 0x34, 0xcc, 0x14, 0x20, 0xc8, 0xb0, 0x62, 0xdc,
	0xae, 0xf1, 0x77, 0xaa, 0xc9, 0x52, 0x17, 0xb2, 0xb5, 0x70, 0x7b, 0xa5, 0xdc, 0xec, 0xb3, 0x39,
	0x5c, 0xdc, 0xcb, 0x3c, 0xd1, 0x4d, 0x72, 0xe5, 0xa3, 0x4a, 0x65, 0xa2, 0x90, 0x8e, 0xa1, 0x1b,
	0x14, 0x52, 0xd5, 0xa2, 0xfc, 0x61, 0x37, 0x30, 0x7a, 0x44, 0x8b, 0xfb, 0xf8, 0x99, 0xa3, 0xd2,
	0x87, 0x7d, 0xd9, 0x13, 0x8c, 0x7f, 0x63, 0x95, 0xe9, 0x2d, 0x38, 0xcd, 0xfb, 0x1a, 0xbc, 0xbf,
	0x18, 0xf1, 0xfa, 0x81, 0xb9, 0xe5, 0x2d, 0xb5, 0xd8, 0x11, 0x38, 0x6b, 0x84, 0x57, 0xcc, 0xbe,
	0xa2, 0x00, 0xe9, 0x33, 0x0c, 0x0f, 0x72, 0xd3, 0x73, 0x5e, 0x9e, 0x87, 0xd7, 0xe7, 0xe1, 0x0f,
	0xc5, 0x79, 0x26, 0xd7, 0xd6, 0xff, 0x48, 0x55, 0xf6, 0x8f, 0xbe, 0xc0, 0x60, 0x3f, 0x2f, 0x9d,
	0xda, 0xa5, 0x3f, 0xea, 0x4e, 0xae, 0x8e, 0xc9, 0xb5, 0xe1, 0xba, 0x67, 0x52, 0xdc, 0xfd, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x90, 0x61, 0x26, 0x69, 0x4f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlgorithmServiceClient is the client API for AlgorithmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlgorithmServiceClient interface {
	CountAlgorithms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountAlgorithmsResponse, error)
	GetAlgorithm(ctx context.Context, in *GetAlgorithmRequest, opts ...grpc.CallOption) (*GetAlgorithmResponse, error)
}

type algorithmServiceClient struct {
	cc *grpc.ClientConn
}

func NewAlgorithmServiceClient(cc *grpc.ClientConn) AlgorithmServiceClient {
	return &algorithmServiceClient{cc}
}

func (c *algorithmServiceClient) CountAlgorithms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountAlgorithmsResponse, error) {
	out := new(CountAlgorithmsResponse)
	err := c.cc.Invoke(ctx, "/basedata.AlgorithmService/CountAlgorithms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algorithmServiceClient) GetAlgorithm(ctx context.Context, in *GetAlgorithmRequest, opts ...grpc.CallOption) (*GetAlgorithmResponse, error) {
	out := new(GetAlgorithmResponse)
	err := c.cc.Invoke(ctx, "/basedata.AlgorithmService/GetAlgorithm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlgorithmServiceServer is the server API for AlgorithmService service.
type AlgorithmServiceServer interface {
	CountAlgorithms(context.Context, *empty.Empty) (*CountAlgorithmsResponse, error)
	GetAlgorithm(context.Context, *GetAlgorithmRequest) (*GetAlgorithmResponse, error)
}

// UnimplementedAlgorithmServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlgorithmServiceServer struct {
}

func (*UnimplementedAlgorithmServiceServer) CountAlgorithms(ctx context.Context, req *empty.Empty) (*CountAlgorithmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAlgorithms not implemented")
}
func (*UnimplementedAlgorithmServiceServer) GetAlgorithm(ctx context.Context, req *GetAlgorithmRequest) (*GetAlgorithmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlgorithm not implemented")
}

func RegisterAlgorithmServiceServer(s *grpc.Server, srv AlgorithmServiceServer) {
	s.RegisterService(&_AlgorithmService_serviceDesc, srv)
}

func _AlgorithmService_CountAlgorithms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmServiceServer).CountAlgorithms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basedata.AlgorithmService/CountAlgorithms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmServiceServer).CountAlgorithms(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlgorithmService_GetAlgorithm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlgorithmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmServiceServer).GetAlgorithm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basedata.AlgorithmService/GetAlgorithm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmServiceServer).GetAlgorithm(ctx, req.(*GetAlgorithmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlgorithmService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "basedata.AlgorithmService",
	HandlerType: (*AlgorithmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountAlgorithms",
			Handler:    _AlgorithmService_CountAlgorithms_Handler,
		},
		{
			MethodName: "GetAlgorithm",
			Handler:    _AlgorithmService_GetAlgorithm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "algorithm.proto",
}
