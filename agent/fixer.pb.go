// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fixer.proto

package agent

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

type Rate struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Base                 string   `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	CurrencyRate         float64  `protobuf:"fixed64,4,opt,name=currency_rate,json=currencyRate,proto3" json:"currency_rate,omitempty"`
	CratedAtTs           int64    `protobuf:"varint,5,opt,name=crated_at_ts,json=cratedAtTs,proto3" json:"crated_at_ts,omitempty"`
	UpdatedAtTs          int64    `protobuf:"varint,6,opt,name=updated_at_ts,json=updatedAtTs,proto3" json:"updated_at_ts,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rate) Reset()         { *m = Rate{} }
func (m *Rate) String() string { return proto.CompactTextString(m) }
func (*Rate) ProtoMessage()    {}
func (*Rate) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f614da6dc767a1, []int{0}
}

func (m *Rate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rate.Unmarshal(m, b)
}
func (m *Rate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rate.Marshal(b, m, deterministic)
}
func (m *Rate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rate.Merge(m, src)
}
func (m *Rate) XXX_Size() int {
	return xxx_messageInfo_Rate.Size(m)
}
func (m *Rate) XXX_DiscardUnknown() {
	xxx_messageInfo_Rate.DiscardUnknown(m)
}

var xxx_messageInfo_Rate proto.InternalMessageInfo

func (m *Rate) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Rate) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Rate) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Rate) GetCurrencyRate() float64 {
	if m != nil {
		return m.CurrencyRate
	}
	return 0
}

func (m *Rate) GetCratedAtTs() int64 {
	if m != nil {
		return m.CratedAtTs
	}
	return 0
}

func (m *Rate) GetUpdatedAtTs() int64 {
	if m != nil {
		return m.UpdatedAtTs
	}
	return 0
}

func (m *Rate) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Rate) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type GetRateRequest struct {
	Base                 string   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Currency             string   `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRateRequest) Reset()         { *m = GetRateRequest{} }
func (m *GetRateRequest) String() string { return proto.CompactTextString(m) }
func (*GetRateRequest) ProtoMessage()    {}
func (*GetRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f614da6dc767a1, []int{1}
}

func (m *GetRateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRateRequest.Unmarshal(m, b)
}
func (m *GetRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRateRequest.Marshal(b, m, deterministic)
}
func (m *GetRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRateRequest.Merge(m, src)
}
func (m *GetRateRequest) XXX_Size() int {
	return xxx_messageInfo_GetRateRequest.Size(m)
}
func (m *GetRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRateRequest proto.InternalMessageInfo

func (m *GetRateRequest) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *GetRateRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type GetRateResponse struct {
	Rate                 *Rate    `protobuf:"bytes,1,opt,name=rate,proto3" json:"rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRateResponse) Reset()         { *m = GetRateResponse{} }
func (m *GetRateResponse) String() string { return proto.CompactTextString(m) }
func (*GetRateResponse) ProtoMessage()    {}
func (*GetRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3f614da6dc767a1, []int{2}
}

func (m *GetRateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRateResponse.Unmarshal(m, b)
}
func (m *GetRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRateResponse.Marshal(b, m, deterministic)
}
func (m *GetRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRateResponse.Merge(m, src)
}
func (m *GetRateResponse) XXX_Size() int {
	return xxx_messageInfo_GetRateResponse.Size(m)
}
func (m *GetRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRateResponse proto.InternalMessageInfo

func (m *GetRateResponse) GetRate() *Rate {
	if m != nil {
		return m.Rate
	}
	return nil
}

func init() {
	proto.RegisterType((*Rate)(nil), "agent.Rate")
	proto.RegisterType((*GetRateRequest)(nil), "agent.GetRateRequest")
	proto.RegisterType((*GetRateResponse)(nil), "agent.GetRateResponse")
}

func init() { proto.RegisterFile("fixer.proto", fileDescriptor_f3f614da6dc767a1) }

var fileDescriptor_f3f614da6dc767a1 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x3d, 0x4f, 0x84, 0x40,
	0x10, 0x75, 0x39, 0xb8, 0x8f, 0x81, 0x3b, 0x93, 0x4d, 0x34, 0x9b, 0x4b, 0x8c, 0x04, 0x1b, 0x2a,
	0x0a, 0xec, 0xac, 0xbc, 0x46, 0x13, 0xcb, 0xd5, 0x9e, 0x70, 0x30, 0x1a, 0x1a, 0xc0, 0xdd, 0xc1,
	0xe8, 0xaf, 0xf6, 0x2f, 0x18, 0x06, 0x0e, 0x3d, 0x93, 0xeb, 0x86, 0xf7, 0x1e, 0x6f, 0xde, 0x9b,
	0x05, 0xff, 0xb5, 0xfa, 0x44, 0x93, 0xb4, 0xa6, 0xa1, 0x46, 0x7a, 0xf9, 0x1b, 0xd6, 0x14, 0x7d,
	0x0b, 0x70, 0x75, 0x4e, 0x28, 0x37, 0xe0, 0x54, 0xa5, 0x12, 0xa1, 0x88, 0x3d, 0xed, 0x54, 0xa5,
	0x94, 0xe0, 0xee, 0x73, 0x8b, 0xca, 0x09, 0x45, 0xbc, 0xd2, 0x3c, 0xcb, 0x2d, 0x2c, 0x8b, 0xce,
	0x18, 0xac, 0x8b, 0x2f, 0x35, 0x63, 0x7c, 0xfa, 0x96, 0x37, 0xb0, 0x3e, 0xcc, 0x99, 0xc9, 0x09,
	0x95, 0x1b, 0x8a, 0x58, 0xe8, 0xe0, 0x00, 0xf2, 0x92, 0x10, 0x82, 0xa2, 0x27, 0xcb, 0x2c, 0xa7,
	0x8c, 0xac, 0xf2, 0x42, 0x11, 0xcf, 0x34, 0x0c, 0xd8, 0x8e, 0x5e, 0xac, 0x8c, 0x60, 0xdd, 0xb5,
	0xe5, 0x1f, 0xc9, 0x9c, 0x25, 0xfe, 0x08, 0xb2, 0xe6, 0x0a, 0xa0, 0x30, 0x38, 0x6a, 0xd4, 0x82,
	0x83, 0xac, 0x46, 0x64, 0x47, 0x3d, 0xfd, 0x6b, 0xa1, 0x96, 0x03, 0x3d, 0xfd, 0x1f, 0xdd, 0xc3,
	0xe6, 0x11, 0xa9, 0x8f, 0xa3, 0xf1, 0xbd, 0x43, 0x4b, 0x53, 0x55, 0x71, 0xa2, 0xaa, 0x73, 0x5c,
	0x35, 0x4a, 0xe1, 0x7c, 0x72, 0xb0, 0x6d, 0x53, 0x5b, 0x94, 0xd7, 0xe0, 0x72, 0xe9, 0xde, 0xc2,
	0x4f, 0xfd, 0x84, 0x8f, 0x9b, 0xb0, 0x84, 0x89, 0xf4, 0x09, 0x82, 0x87, 0xfe, 0xfa, 0xcf, 0x68,
	0x3e, 0xaa, 0x02, 0xe5, 0x1d, 0x2c, 0x46, 0x0f, 0x79, 0x31, 0xaa, 0x8f, 0x53, 0x6d, 0x2f, 0xff,
	0xc3, 0xc3, 0xaa, 0xe8, 0x6c, 0x3f, 0xe7, 0x17, 0xbc, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xac,
	0x80, 0xb3, 0x9b, 0xd0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FixerServiceClient is the client API for FixerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FixerServiceClient interface {
	GetRate(ctx context.Context, in *GetRateRequest, opts ...grpc.CallOption) (*GetRateResponse, error)
}

type fixerServiceClient struct {
	cc *grpc.ClientConn
}

func NewFixerServiceClient(cc *grpc.ClientConn) FixerServiceClient {
	return &fixerServiceClient{cc}
}

func (c *fixerServiceClient) GetRate(ctx context.Context, in *GetRateRequest, opts ...grpc.CallOption) (*GetRateResponse, error) {
	out := new(GetRateResponse)
	err := c.cc.Invoke(ctx, "/agent.FixerService/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FixerServiceServer is the server API for FixerService service.
type FixerServiceServer interface {
	GetRate(context.Context, *GetRateRequest) (*GetRateResponse, error)
}

// UnimplementedFixerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFixerServiceServer struct {
}

func (*UnimplementedFixerServiceServer) GetRate(ctx context.Context, req *GetRateRequest) (*GetRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}

func RegisterFixerServiceServer(s *grpc.Server, srv FixerServiceServer) {
	s.RegisterService(&_FixerService_serviceDesc, srv)
}

func _FixerService_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixerServiceServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.FixerService/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixerServiceServer).GetRate(ctx, req.(*GetRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FixerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.FixerService",
	HandlerType: (*FixerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _FixerService_GetRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fixer.proto",
}
