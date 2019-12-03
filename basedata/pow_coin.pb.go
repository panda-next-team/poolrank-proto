// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pow_coin.proto

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

type POW_Coin struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EnName               string   `protobuf:"bytes,3,opt,name=en_name,json=enName,proto3" json:"en_name,omitempty"`
	EnTag                string   `protobuf:"bytes,4,opt,name=en_tag,json=enTag,proto3" json:"en_tag,omitempty"`
	MaxSupply            int32    `protobuf:"varint,5,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty"`
	AlgorithmId          int32    `protobuf:"varint,6,opt,name=algorithm_id,json=algorithmId,proto3" json:"algorithm_id,omitempty"`
	BlockTime            string   `protobuf:"bytes,7,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	WebsiteUrl           string   `protobuf:"bytes,8,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
	GithubUrl            string   `protobuf:"bytes,9,opt,name=github_url,json=githubUrl,proto3" json:"github_url,omitempty"`
	Icon                 string   `protobuf:"bytes,10,opt,name=icon,proto3" json:"icon,omitempty"`
	Intro                string   `protobuf:"bytes,11,opt,name=intro,proto3" json:"intro,omitempty"`
	Status               int32    `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	ListOrder            int32    `protobuf:"varint,13,opt,name=list_order,json=listOrder,proto3" json:"list_order,omitempty"`
	CreatedAtTs          int32    `protobuf:"varint,14,opt,name=created_at_ts,json=createdAtTs,proto3" json:"created_at_ts,omitempty"`
	UpdatedAtTs          int32    `protobuf:"varint,15,opt,name=updated_at_ts,json=updatedAtTs,proto3" json:"updated_at_ts,omitempty"`
	CreatedAt            string   `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,17,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *POW_Coin) Reset()         { *m = POW_Coin{} }
func (m *POW_Coin) String() string { return proto.CompactTextString(m) }
func (*POW_Coin) ProtoMessage()    {}
func (*POW_Coin) Descriptor() ([]byte, []int) {
	return fileDescriptor_04940dfa92f54e82, []int{0}
}

func (m *POW_Coin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POW_Coin.Unmarshal(m, b)
}
func (m *POW_Coin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POW_Coin.Marshal(b, m, deterministic)
}
func (m *POW_Coin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POW_Coin.Merge(m, src)
}
func (m *POW_Coin) XXX_Size() int {
	return xxx_messageInfo_POW_Coin.Size(m)
}
func (m *POW_Coin) XXX_DiscardUnknown() {
	xxx_messageInfo_POW_Coin.DiscardUnknown(m)
}

var xxx_messageInfo_POW_Coin proto.InternalMessageInfo

func (m *POW_Coin) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *POW_Coin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *POW_Coin) GetEnName() string {
	if m != nil {
		return m.EnName
	}
	return ""
}

func (m *POW_Coin) GetEnTag() string {
	if m != nil {
		return m.EnTag
	}
	return ""
}

func (m *POW_Coin) GetMaxSupply() int32 {
	if m != nil {
		return m.MaxSupply
	}
	return 0
}

func (m *POW_Coin) GetAlgorithmId() int32 {
	if m != nil {
		return m.AlgorithmId
	}
	return 0
}

func (m *POW_Coin) GetBlockTime() string {
	if m != nil {
		return m.BlockTime
	}
	return ""
}

func (m *POW_Coin) GetWebsiteUrl() string {
	if m != nil {
		return m.WebsiteUrl
	}
	return ""
}

func (m *POW_Coin) GetGithubUrl() string {
	if m != nil {
		return m.GithubUrl
	}
	return ""
}

func (m *POW_Coin) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *POW_Coin) GetIntro() string {
	if m != nil {
		return m.Intro
	}
	return ""
}

func (m *POW_Coin) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *POW_Coin) GetListOrder() int32 {
	if m != nil {
		return m.ListOrder
	}
	return 0
}

func (m *POW_Coin) GetCreatedAtTs() int32 {
	if m != nil {
		return m.CreatedAtTs
	}
	return 0
}

func (m *POW_Coin) GetUpdatedAtTs() int32 {
	if m != nil {
		return m.UpdatedAtTs
	}
	return 0
}

func (m *POW_Coin) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *POW_Coin) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CountResponse struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResponse) Reset()         { *m = CountResponse{} }
func (m *CountResponse) String() string { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()    {}
func (*CountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04940dfa92f54e82, []int{1}
}

func (m *CountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountResponse.Unmarshal(m, b)
}
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountResponse.Marshal(b, m, deterministic)
}
func (m *CountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResponse.Merge(m, src)
}
func (m *CountResponse) XXX_Size() int {
	return xxx_messageInfo_CountResponse.Size(m)
}
func (m *CountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountResponse proto.InternalMessageInfo

func (m *CountResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type GetIDRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIDRequest) Reset()         { *m = GetIDRequest{} }
func (m *GetIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetIDRequest) ProtoMessage()    {}
func (*GetIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04940dfa92f54e82, []int{2}
}

func (m *GetIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDRequest.Unmarshal(m, b)
}
func (m *GetIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDRequest.Marshal(b, m, deterministic)
}
func (m *GetIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDRequest.Merge(m, src)
}
func (m *GetIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetIDRequest.Size(m)
}
func (m *GetIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDRequest proto.InternalMessageInfo

func (m *GetIDRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetIDResponse struct {
	Coin                 *POW_Coin `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetIDResponse) Reset()         { *m = GetIDResponse{} }
func (m *GetIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetIDResponse) ProtoMessage()    {}
func (*GetIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04940dfa92f54e82, []int{3}
}

func (m *GetIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIDResponse.Unmarshal(m, b)
}
func (m *GetIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIDResponse.Marshal(b, m, deterministic)
}
func (m *GetIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIDResponse.Merge(m, src)
}
func (m *GetIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetIDResponse.Size(m)
}
func (m *GetIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIDResponse proto.InternalMessageInfo

func (m *GetIDResponse) GetCoin() *POW_Coin {
	if m != nil {
		return m.Coin
	}
	return nil
}

func init() {
	proto.RegisterType((*POW_Coin)(nil), "basedata.POW_Coin")
	proto.RegisterType((*CountResponse)(nil), "basedata.CountResponse")
	proto.RegisterType((*GetIDRequest)(nil), "basedata.GetIDRequest")
	proto.RegisterType((*GetIDResponse)(nil), "basedata.GetIDResponse")
}

func init() { proto.RegisterFile("pow_coin.proto", fileDescriptor_04940dfa92f54e82) }

var fileDescriptor_04940dfa92f54e82 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x9b, 0x34, 0x49, 0x93, 0xc9, 0x1f, 0x60, 0x05, 0xe9, 0xaa, 0xa8, 0x50, 0x2c, 0x81,
	0x7a, 0x72, 0xa5, 0x72, 0x40, 0xe2, 0x86, 0x0a, 0x42, 0xbd, 0x90, 0x2a, 0x0d, 0xea, 0x71, 0xb5,
	0xb6, 0x07, 0x77, 0x85, 0xbd, 0x6b, 0xbc, 0x63, 0xda, 0xbe, 0x01, 0x4f, 0xca, 0x73, 0xa0, 0xdd,
	0x75, 0xfe, 0x94, 0x9b, 0xe7, 0xf7, 0x7d, 0xf3, 0xed, 0x8c, 0x3c, 0x30, 0xab, 0xcc, 0x9d, 0x48,
	0x8d, 0xd2, 0x71, 0x55, 0x1b, 0x32, 0x6c, 0x98, 0x48, 0x8b, 0x99, 0x24, 0x79, 0xf4, 0x32, 0x37,
	0x26, 0x2f, 0xf0, 0xcc, 0xf3, 0xa4, 0xf9, 0x71, 0x86, 0x65, 0x45, 0x0f, 0xc1, 0x16, 0xfd, 0xdd,
	0x87, 0xe1, 0xd5, 0xe2, 0x46, 0x5c, 0x18, 0xa5, 0xd9, 0x0c, 0xba, 0x2a, 0xe3, 0x9d, 0x93, 0xce,
	0x69, 0x7f, 0xd9, 0x55, 0x19, 0x63, 0xd0, 0xd3, 0xb2, 0x44, 0xde, 0x3d, 0xe9, 0x9c, 0x8e, 0x96,
	0xfe, 0x9b, 0x1d, 0xc2, 0x01, 0x6a, 0xe1, 0xf1, 0xbe, 0xc7, 0x03, 0xd4, 0xdf, 0x9c, 0xf0, 0x02,
	0x06, 0xa8, 0x05, 0xc9, 0x9c, 0xf7, 0x3c, 0xef, 0xa3, 0x5e, 0xc9, 0x9c, 0x1d, 0x03, 0x94, 0xf2,
	0x5e, 0xd8, 0xa6, 0xaa, 0x8a, 0x07, 0xde, 0xf7, 0xd9, 0xa3, 0x52, 0xde, 0x5f, 0x7b, 0xc0, 0xde,
	0xc0, 0x44, 0x16, 0xb9, 0xa9, 0x15, 0xdd, 0x96, 0x42, 0x65, 0x7c, 0xe0, 0x0d, 0xe3, 0x0d, 0xbb,
	0xcc, 0x5c, 0x42, 0x52, 0x98, 0xf4, 0xa7, 0x20, 0x55, 0x22, 0x3f, 0xf0, 0xe1, 0x23, 0x4f, 0x56,
	0xaa, 0x44, 0xf6, 0x1a, 0xc6, 0x77, 0x98, 0x58, 0x45, 0x28, 0x9a, 0xba, 0xe0, 0x43, 0xaf, 0x43,
	0x8b, 0xbe, 0xd7, 0x85, 0xeb, 0xcf, 0x15, 0xdd, 0x36, 0x89, 0xd7, 0x47, 0xa1, 0x3f, 0x10, 0x27,
	0x33, 0xe8, 0xa9, 0xd4, 0x68, 0x0e, 0x61, 0x49, 0xf7, 0xcd, 0x9e, 0x43, 0x5f, 0x69, 0xaa, 0x0d,
	0x1f, 0x87, 0x55, 0x7c, 0xc1, 0xe6, 0x30, 0xb0, 0x24, 0xa9, 0xb1, 0x7c, 0xe2, 0xa7, 0x6c, 0x2b,
	0xf7, 0x40, 0xa1, 0x2c, 0x09, 0x53, 0x67, 0x58, 0xf3, 0x69, 0x58, 0xd1, 0x91, 0x85, 0x03, 0x2c,
	0x82, 0x69, 0x5a, 0xa3, 0x24, 0xcc, 0x84, 0x24, 0x41, 0x96, 0xcf, 0xc2, 0x8e, 0x2d, 0xfc, 0x44,
	0x2b, 0xeb, 0x3c, 0x4d, 0x95, 0xed, 0x78, 0x9e, 0x04, 0x4f, 0x0b, 0xbd, 0xe7, 0x18, 0x60, 0x9b,
	0xc3, 0x9f, 0x86, 0x3d, 0x36, 0x21, 0x4e, 0xde, 0x46, 0xf0, 0x67, 0x41, 0xde, 0xf4, 0x47, 0x6f,
	0x61, 0x7a, 0x61, 0x1a, 0x4d, 0x4b, 0xb4, 0x95, 0xd1, 0x16, 0xdd, 0x8e, 0x64, 0x48, 0x16, 0xed,
	0xff, 0x0e, 0x45, 0xf4, 0x0a, 0x26, 0x5f, 0x91, 0x2e, 0x3f, 0x2f, 0xf1, 0x57, 0x83, 0x96, 0xfe,
	0x3f, 0x89, 0xe8, 0x03, 0x4c, 0x5b, 0xbd, 0x8d, 0x79, 0x07, 0x3d, 0x77, 0x75, 0xde, 0x32, 0x3e,
	0x67, 0xf1, 0xfa, 0xec, 0xe2, 0xf5, 0x55, 0x2d, 0xbd, 0x7e, 0xfe, 0xa7, 0x03, 0xb3, 0xab, 0xc5,
	0x8d, 0x23, 0xd7, 0x58, 0xff, 0x56, 0x29, 0xb2, 0x8f, 0xd0, 0xf7, 0x23, 0xb1, 0x79, 0x1c, 0x4e,
	0x34, 0x5e, 0x9f, 0x68, 0xfc, 0xc5, 0x9d, 0xe8, 0xd1, 0xe1, 0x36, 0xed, 0xd1, 0xec, 0xd1, 0x9e,
	0xeb, 0xf5, 0x73, 0xb0, 0xf9, 0xd6, 0xb3, 0x3b, 0xf8, 0x6e, 0xef, 0xa3, 0x81, 0xa3, 0xbd, 0x64,
	0xe0, 0x9f, 0x79, 0xff, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xe8, 0xd8, 0xfd, 0x33, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// POWCoinServiceClient is the client API for POWCoinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type POWCoinServiceClient interface {
	Count(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountResponse, error)
	GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error)
}

type pOWCoinServiceClient struct {
	cc *grpc.ClientConn
}

func NewPOWCoinServiceClient(cc *grpc.ClientConn) POWCoinServiceClient {
	return &pOWCoinServiceClient{cc}
}

func (c *pOWCoinServiceClient) Count(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/basedata.POWCoinService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pOWCoinServiceClient) GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error) {
	out := new(GetIDResponse)
	err := c.cc.Invoke(ctx, "/basedata.POWCoinService/GetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// POWCoinServiceServer is the server API for POWCoinService service.
type POWCoinServiceServer interface {
	Count(context.Context, *empty.Empty) (*CountResponse, error)
	GetID(context.Context, *GetIDRequest) (*GetIDResponse, error)
}

// UnimplementedPOWCoinServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPOWCoinServiceServer struct {
}

func (*UnimplementedPOWCoinServiceServer) Count(ctx context.Context, req *empty.Empty) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (*UnimplementedPOWCoinServiceServer) GetID(ctx context.Context, req *GetIDRequest) (*GetIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetID not implemented")
}

func RegisterPOWCoinServiceServer(s *grpc.Server, srv POWCoinServiceServer) {
	s.RegisterService(&_POWCoinService_serviceDesc, srv)
}

func _POWCoinService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(POWCoinServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basedata.POWCoinService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(POWCoinServiceServer).Count(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _POWCoinService_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(POWCoinServiceServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basedata.POWCoinService/GetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(POWCoinServiceServer).GetID(ctx, req.(*GetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _POWCoinService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "basedata.POWCoinService",
	HandlerType: (*POWCoinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Count",
			Handler:    _POWCoinService_Count_Handler,
		},
		{
			MethodName: "GetID",
			Handler:    _POWCoinService_GetID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pow_coin.proto",
}
