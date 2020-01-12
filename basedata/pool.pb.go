// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pool.proto

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

type AddressType int32

const (
	AddressType_ALL      AddressType = 0
	AddressType_COINBASE AddressType = 1
	AddressType_PAYMENT  AddressType = 2
	AddressType_TRANSIT  AddressType = 3
	AddressType_UNKNOWN  AddressType = 4
)

var AddressType_name = map[int32]string{
	0: "ALL",
	1: "COINBASE",
	2: "PAYMENT",
	3: "TRANSIT",
	4: "UNKNOWN",
}

var AddressType_value = map[string]int32{
	"ALL":      0,
	"COINBASE": 1,
	"PAYMENT":  2,
	"TRANSIT":  3,
	"UNKNOWN":  4,
}

func (x AddressType) String() string {
	return proto.EnumName(AddressType_name, int32(x))
}

func (AddressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{0}
}

type Pool struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	WebsiteUrl           string   `protobuf:"bytes,3,opt,name=website_url,json=websiteUrl,proto3" json:"website_url,omitempty"`
	Icon                 string   `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Status               int32    `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	ListOrder            int32    `protobuf:"varint,6,opt,name=list_order,json=listOrder,proto3" json:"list_order,omitempty"`
	CreatedAtTs          int32    `protobuf:"varint,7,opt,name=created_at_ts,json=createdAtTs,proto3" json:"created_at_ts,omitempty"`
	UpdatedAtTs          int32    `protobuf:"varint,8,opt,name=updated_at_ts,json=updatedAtTs,proto3" json:"updated_at_ts,omitempty"`
	CreatedAt            string   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{0}
}

func (m *Pool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pool.Unmarshal(m, b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return xxx_messageInfo_Pool.Size(m)
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pool) GetWebsiteUrl() string {
	if m != nil {
		return m.WebsiteUrl
	}
	return ""
}

func (m *Pool) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Pool) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Pool) GetListOrder() int32 {
	if m != nil {
		return m.ListOrder
	}
	return 0
}

func (m *Pool) GetCreatedAtTs() int32 {
	if m != nil {
		return m.CreatedAtTs
	}
	return 0
}

func (m *Pool) GetUpdatedAtTs() int32 {
	if m != nil {
		return m.UpdatedAtTs
	}
	return 0
}

func (m *Pool) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Pool) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type PoolAddress struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PoolId               int32    `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	CoinId               int32    `protobuf:"varint,4,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	Type                 int32    `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAtTs          int32    `protobuf:"varint,6,opt,name=created_at_ts,json=createdAtTs,proto3" json:"created_at_ts,omitempty"`
	UpdatedAtTs          int32    `protobuf:"varint,7,opt,name=updated_at_ts,json=updatedAtTs,proto3" json:"updated_at_ts,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolAddress) Reset()         { *m = PoolAddress{} }
func (m *PoolAddress) String() string { return proto.CompactTextString(m) }
func (*PoolAddress) ProtoMessage()    {}
func (*PoolAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{1}
}

func (m *PoolAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolAddress.Unmarshal(m, b)
}
func (m *PoolAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolAddress.Marshal(b, m, deterministic)
}
func (m *PoolAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolAddress.Merge(m, src)
}
func (m *PoolAddress) XXX_Size() int {
	return xxx_messageInfo_PoolAddress.Size(m)
}
func (m *PoolAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PoolAddress proto.InternalMessageInfo

func (m *PoolAddress) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PoolAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PoolAddress) GetPoolId() int32 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *PoolAddress) GetCoinId() int32 {
	if m != nil {
		return m.CoinId
	}
	return 0
}

func (m *PoolAddress) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *PoolAddress) GetCreatedAtTs() int32 {
	if m != nil {
		return m.CreatedAtTs
	}
	return 0
}

func (m *PoolAddress) GetUpdatedAtTs() int32 {
	if m != nil {
		return m.UpdatedAtTs
	}
	return 0
}

func (m *PoolAddress) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *PoolAddress) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type PoolCoinbaseTag struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PoolId               int32    `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	CreatedAtTs          int32    `protobuf:"varint,4,opt,name=created_at_ts,json=createdAtTs,proto3" json:"created_at_ts,omitempty"`
	UpdatedAtTs          int32    `protobuf:"varint,5,opt,name=updated_at_ts,json=updatedAtTs,proto3" json:"updated_at_ts,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolCoinbaseTag) Reset()         { *m = PoolCoinbaseTag{} }
func (m *PoolCoinbaseTag) String() string { return proto.CompactTextString(m) }
func (*PoolCoinbaseTag) ProtoMessage()    {}
func (*PoolCoinbaseTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{2}
}

func (m *PoolCoinbaseTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolCoinbaseTag.Unmarshal(m, b)
}
func (m *PoolCoinbaseTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolCoinbaseTag.Marshal(b, m, deterministic)
}
func (m *PoolCoinbaseTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolCoinbaseTag.Merge(m, src)
}
func (m *PoolCoinbaseTag) XXX_Size() int {
	return xxx_messageInfo_PoolCoinbaseTag.Size(m)
}
func (m *PoolCoinbaseTag) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolCoinbaseTag.DiscardUnknown(m)
}

var xxx_messageInfo_PoolCoinbaseTag proto.InternalMessageInfo

func (m *PoolCoinbaseTag) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PoolCoinbaseTag) GetPoolId() int32 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *PoolCoinbaseTag) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *PoolCoinbaseTag) GetCreatedAtTs() int32 {
	if m != nil {
		return m.CreatedAtTs
	}
	return 0
}

func (m *PoolCoinbaseTag) GetUpdatedAtTs() int32 {
	if m != nil {
		return m.UpdatedAtTs
	}
	return 0
}

func (m *PoolCoinbaseTag) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *PoolCoinbaseTag) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type CountPoolsResponse struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountPoolsResponse) Reset()         { *m = CountPoolsResponse{} }
func (m *CountPoolsResponse) String() string { return proto.CompactTextString(m) }
func (*CountPoolsResponse) ProtoMessage()    {}
func (*CountPoolsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{3}
}

func (m *CountPoolsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountPoolsResponse.Unmarshal(m, b)
}
func (m *CountPoolsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountPoolsResponse.Marshal(b, m, deterministic)
}
func (m *CountPoolsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountPoolsResponse.Merge(m, src)
}
func (m *CountPoolsResponse) XXX_Size() int {
	return xxx_messageInfo_CountPoolsResponse.Size(m)
}
func (m *CountPoolsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountPoolsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountPoolsResponse proto.InternalMessageInfo

func (m *CountPoolsResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetPoolRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPoolRequest) Reset()         { *m = GetPoolRequest{} }
func (m *GetPoolRequest) String() string { return proto.CompactTextString(m) }
func (*GetPoolRequest) ProtoMessage()    {}
func (*GetPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{4}
}

func (m *GetPoolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolRequest.Unmarshal(m, b)
}
func (m *GetPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolRequest.Marshal(b, m, deterministic)
}
func (m *GetPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolRequest.Merge(m, src)
}
func (m *GetPoolRequest) XXX_Size() int {
	return xxx_messageInfo_GetPoolRequest.Size(m)
}
func (m *GetPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolRequest proto.InternalMessageInfo

func (m *GetPoolRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetPoolResponse struct {
	Pool                 *Pool    `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPoolResponse) Reset()         { *m = GetPoolResponse{} }
func (m *GetPoolResponse) String() string { return proto.CompactTextString(m) }
func (*GetPoolResponse) ProtoMessage()    {}
func (*GetPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{5}
}

func (m *GetPoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolResponse.Unmarshal(m, b)
}
func (m *GetPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolResponse.Marshal(b, m, deterministic)
}
func (m *GetPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolResponse.Merge(m, src)
}
func (m *GetPoolResponse) XXX_Size() int {
	return xxx_messageInfo_GetPoolResponse.Size(m)
}
func (m *GetPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolResponse proto.InternalMessageInfo

func (m *GetPoolResponse) GetPool() *Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

type GetPoolAddressesRequest struct {
	PoolId               int32       `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	CoinId               int32       `protobuf:"varint,2,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	Type                 AddressType `protobuf:"varint,3,opt,name=type,proto3,enum=basedata.AddressType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetPoolAddressesRequest) Reset()         { *m = GetPoolAddressesRequest{} }
func (m *GetPoolAddressesRequest) String() string { return proto.CompactTextString(m) }
func (*GetPoolAddressesRequest) ProtoMessage()    {}
func (*GetPoolAddressesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{6}
}

func (m *GetPoolAddressesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolAddressesRequest.Unmarshal(m, b)
}
func (m *GetPoolAddressesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolAddressesRequest.Marshal(b, m, deterministic)
}
func (m *GetPoolAddressesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolAddressesRequest.Merge(m, src)
}
func (m *GetPoolAddressesRequest) XXX_Size() int {
	return xxx_messageInfo_GetPoolAddressesRequest.Size(m)
}
func (m *GetPoolAddressesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolAddressesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolAddressesRequest proto.InternalMessageInfo

func (m *GetPoolAddressesRequest) GetPoolId() int32 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *GetPoolAddressesRequest) GetCoinId() int32 {
	if m != nil {
		return m.CoinId
	}
	return 0
}

func (m *GetPoolAddressesRequest) GetType() AddressType {
	if m != nil {
		return m.Type
	}
	return AddressType_ALL
}

type GetPoolAddressesResponse struct {
	Addresses            []*PoolAddress `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetPoolAddressesResponse) Reset()         { *m = GetPoolAddressesResponse{} }
func (m *GetPoolAddressesResponse) String() string { return proto.CompactTextString(m) }
func (*GetPoolAddressesResponse) ProtoMessage()    {}
func (*GetPoolAddressesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{7}
}

func (m *GetPoolAddressesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolAddressesResponse.Unmarshal(m, b)
}
func (m *GetPoolAddressesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolAddressesResponse.Marshal(b, m, deterministic)
}
func (m *GetPoolAddressesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolAddressesResponse.Merge(m, src)
}
func (m *GetPoolAddressesResponse) XXX_Size() int {
	return xxx_messageInfo_GetPoolAddressesResponse.Size(m)
}
func (m *GetPoolAddressesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolAddressesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolAddressesResponse proto.InternalMessageInfo

func (m *GetPoolAddressesResponse) GetAddresses() []*PoolAddress {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type GetPoolCoinbaseTagsRequest struct {
	PoolId               int32    `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPoolCoinbaseTagsRequest) Reset()         { *m = GetPoolCoinbaseTagsRequest{} }
func (m *GetPoolCoinbaseTagsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPoolCoinbaseTagsRequest) ProtoMessage()    {}
func (*GetPoolCoinbaseTagsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{8}
}

func (m *GetPoolCoinbaseTagsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolCoinbaseTagsRequest.Unmarshal(m, b)
}
func (m *GetPoolCoinbaseTagsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolCoinbaseTagsRequest.Marshal(b, m, deterministic)
}
func (m *GetPoolCoinbaseTagsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolCoinbaseTagsRequest.Merge(m, src)
}
func (m *GetPoolCoinbaseTagsRequest) XXX_Size() int {
	return xxx_messageInfo_GetPoolCoinbaseTagsRequest.Size(m)
}
func (m *GetPoolCoinbaseTagsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolCoinbaseTagsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolCoinbaseTagsRequest proto.InternalMessageInfo

func (m *GetPoolCoinbaseTagsRequest) GetPoolId() int32 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

type GetPoolCoinbaseTagsResponse struct {
	Tags                 []*PoolCoinbaseTag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetPoolCoinbaseTagsResponse) Reset()         { *m = GetPoolCoinbaseTagsResponse{} }
func (m *GetPoolCoinbaseTagsResponse) String() string { return proto.CompactTextString(m) }
func (*GetPoolCoinbaseTagsResponse) ProtoMessage()    {}
func (*GetPoolCoinbaseTagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a14d8612184524f, []int{9}
}

func (m *GetPoolCoinbaseTagsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPoolCoinbaseTagsResponse.Unmarshal(m, b)
}
func (m *GetPoolCoinbaseTagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPoolCoinbaseTagsResponse.Marshal(b, m, deterministic)
}
func (m *GetPoolCoinbaseTagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPoolCoinbaseTagsResponse.Merge(m, src)
}
func (m *GetPoolCoinbaseTagsResponse) XXX_Size() int {
	return xxx_messageInfo_GetPoolCoinbaseTagsResponse.Size(m)
}
func (m *GetPoolCoinbaseTagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPoolCoinbaseTagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPoolCoinbaseTagsResponse proto.InternalMessageInfo

func (m *GetPoolCoinbaseTagsResponse) GetTags() []*PoolCoinbaseTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("basedata.AddressType", AddressType_name, AddressType_value)
	proto.RegisterType((*Pool)(nil), "basedata.Pool")
	proto.RegisterType((*PoolAddress)(nil), "basedata.PoolAddress")
	proto.RegisterType((*PoolCoinbaseTag)(nil), "basedata.PoolCoinbaseTag")
	proto.RegisterType((*CountPoolsResponse)(nil), "basedata.CountPoolsResponse")
	proto.RegisterType((*GetPoolRequest)(nil), "basedata.GetPoolRequest")
	proto.RegisterType((*GetPoolResponse)(nil), "basedata.GetPoolResponse")
	proto.RegisterType((*GetPoolAddressesRequest)(nil), "basedata.GetPoolAddressesRequest")
	proto.RegisterType((*GetPoolAddressesResponse)(nil), "basedata.GetPoolAddressesResponse")
	proto.RegisterType((*GetPoolCoinbaseTagsRequest)(nil), "basedata.GetPoolCoinbaseTagsRequest")
	proto.RegisterType((*GetPoolCoinbaseTagsResponse)(nil), "basedata.GetPoolCoinbaseTagsResponse")
}

func init() { proto.RegisterFile("pool.proto", fileDescriptor_8a14d8612184524f) }

var fileDescriptor_8a14d8612184524f = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0x5d, 0xd2, 0xb4, 0x69, 0x6f, 0xbe, 0xaf, 0x8b, 0x0c, 0x6c, 0x5e, 0xc7, 0x44, 0x89, 0x40,
	0x1a, 0x93, 0xe8, 0xa4, 0x4d, 0x7b, 0xa7, 0x8c, 0x09, 0x55, 0x1b, 0xed, 0x94, 0x75, 0x42, 0x88,
	0x87, 0xca, 0x6d, 0x4c, 0x15, 0xa9, 0xab, 0x43, 0xec, 0x0e, 0xed, 0x99, 0xbf, 0xc1, 0xaf, 0xe2,
	0x89, 0x9f, 0x83, 0xec, 0x38, 0x4d, 0xda, 0x94, 0x76, 0x6f, 0xbe, 0xf7, 0x9e, 0x7b, 0x7d, 0xcf,
	0xf1, 0x49, 0x00, 0x22, 0xc6, 0x26, 0xad, 0x28, 0x66, 0x82, 0xa1, 0xea, 0x90, 0x70, 0x1a, 0x10,
	0x41, 0x1a, 0xfb, 0x63, 0xc6, 0xc6, 0x13, 0x7a, 0xac, 0xf2, 0xc3, 0xd9, 0xb7, 0x63, 0x7a, 0x17,
	0x89, 0x87, 0x04, 0xe6, 0xfd, 0x32, 0xc1, 0xba, 0x66, 0x6c, 0x82, 0xea, 0x60, 0x86, 0x01, 0x36,
	0x9a, 0xc6, 0x61, 0xd9, 0x37, 0xc3, 0x00, 0x21, 0xb0, 0xa6, 0xe4, 0x8e, 0x62, 0xb3, 0x69, 0x1c,
	0xd6, 0x7c, 0x75, 0x46, 0x2f, 0xc0, 0xf9, 0x41, 0x87, 0x3c, 0x14, 0x74, 0x30, 0x8b, 0x27, 0xb8,
	0xa4, 0x4a, 0xa0, 0x53, 0xb7, 0xf1, 0x44, 0x36, 0x85, 0x23, 0x36, 0xc5, 0x56, 0xd2, 0x24, 0xcf,
	0x68, 0x07, 0x2a, 0x5c, 0x10, 0x31, 0xe3, 0xb8, 0xac, 0x86, 0xeb, 0x08, 0x1d, 0x00, 0x4c, 0x42,
	0x2e, 0x06, 0x2c, 0x0e, 0x68, 0x8c, 0x2b, 0xaa, 0x56, 0x93, 0x99, 0x9e, 0x4c, 0x20, 0x0f, 0xfe,
	0x1f, 0xc5, 0x94, 0x08, 0x1a, 0x0c, 0x88, 0x18, 0x08, 0x8e, 0x6d, 0x85, 0x70, 0x74, 0xb2, 0x2d,
	0xfa, 0x5c, 0x62, 0x66, 0x51, 0x90, 0xc3, 0x54, 0x13, 0x8c, 0x4e, 0x2a, 0xcc, 0x01, 0x40, 0x36,
	0x07, 0xd7, 0xd4, 0x62, 0xb5, 0xf9, 0x10, 0x59, 0xce, 0x46, 0x60, 0x48, 0xca, 0xf3, 0x7e, 0xef,
	0xa7, 0x09, 0x8e, 0x94, 0xa7, 0x1d, 0x04, 0x31, 0xe5, 0xbc, 0xa0, 0x12, 0x06, 0x9b, 0x24, 0x25,
	0x2d, 0x54, 0x1a, 0xa2, 0x5d, 0xb0, 0xe5, 0x6b, 0x0c, 0xc2, 0x40, 0xe9, 0x54, 0xf6, 0x2b, 0x32,
	0xec, 0x04, 0xb2, 0x30, 0x62, 0xe1, 0x54, 0x16, 0xac, 0xa4, 0x20, 0xc3, 0x8e, 0x52, 0x5c, 0x3c,
	0x44, 0x54, 0xcb, 0xa4, 0xce, 0x45, 0x15, 0x2a, 0x8f, 0x50, 0xc1, 0xde, 0xa4, 0x42, 0x75, 0xbd,
	0x0a, 0xb5, 0x65, 0x15, 0x7e, 0x1b, 0xb0, 0x2d, 0x55, 0x38, 0x67, 0xe1, 0x54, 0xda, 0xaa, 0x4f,
	0xc6, 0x05, 0x25, 0x72, 0x7c, 0xcd, 0x05, 0xbe, 0x2e, 0x94, 0x04, 0x19, 0x6b, 0xb3, 0xc8, 0x63,
	0x91, 0x94, 0xf5, 0x08, 0x52, 0xe5, 0x4d, 0xa4, 0x2a, 0xeb, 0x49, 0xd9, 0xcb, 0xa4, 0x8e, 0x00,
	0x9d, 0xb3, 0xd9, 0x54, 0x48, 0x62, 0xdc, 0xa7, 0x3c, 0x62, 0x53, 0x4e, 0xd1, 0x53, 0x28, 0x8f,
	0x64, 0x56, 0x33, 0x4b, 0x02, 0xaf, 0x09, 0xf5, 0x8f, 0x54, 0x21, 0x7d, 0xfa, 0x7d, 0x46, 0xb9,
	0x58, 0xa6, 0xef, 0x9d, 0xc1, 0xf6, 0x1c, 0xa1, 0x47, 0x79, 0x60, 0x49, 0x09, 0x14, 0xc8, 0x39,
	0xa9, 0xb7, 0xd2, 0x0f, 0xb2, 0xa5, 0x50, 0xaa, 0xe6, 0xdd, 0xc3, 0xae, 0x6e, 0xd3, 0x0e, 0xa3,
	0x3c, 0xbd, 0x21, 0x27, 0xa8, 0xf1, 0x2f, 0x03, 0x99, 0x0b, 0x06, 0x7a, 0xa3, 0x0d, 0x24, 0xa5,
	0xae, 0x9f, 0x3c, 0xcb, 0x2e, 0xd4, 0xb3, 0xfb, 0x0f, 0x11, 0x4d, 0x7c, 0xe5, 0xf5, 0x00, 0x17,
	0xef, 0xd5, 0x7b, 0x9f, 0x42, 0x8d, 0xa4, 0x49, 0x6c, 0x34, 0x4b, 0x87, 0x4e, 0x7e, 0x56, 0xae,
	0xc7, 0xcf, 0x70, 0xde, 0x19, 0x34, 0xf4, 0xc0, 0x9c, 0x49, 0x36, 0x72, 0xf1, 0xae, 0x60, 0x7f,
	0x65, 0x9b, 0x5e, 0xe5, 0x2d, 0x58, 0x82, 0x8c, 0xd3, 0x2d, 0xf6, 0x16, 0xb7, 0xc8, 0x75, 0xf8,
	0x0a, 0x76, 0x74, 0x09, 0x4e, 0x8e, 0x2a, 0xb2, 0xa1, 0xd4, 0xbe, 0xba, 0x72, 0xb7, 0xd0, 0x7f,
	0x50, 0x3d, 0xef, 0x75, 0xba, 0xef, 0xdb, 0x37, 0x17, 0xae, 0x81, 0x1c, 0xb0, 0xaf, 0xdb, 0x5f,
	0x3e, 0x5d, 0x74, 0xfb, 0xae, 0x29, 0x83, 0xbe, 0xdf, 0xee, 0xde, 0x74, 0xfa, 0x6e, 0x49, 0x06,
	0xb7, 0xdd, 0xcb, 0x6e, 0xef, 0x73, 0xd7, 0xb5, 0x4e, 0xfe, 0xe8, 0x4f, 0xff, 0x86, 0xc6, 0xf7,
	0xe1, 0x88, 0xa2, 0x0f, 0x00, 0x99, 0x5f, 0xd0, 0x4e, 0x2b, 0xf9, 0xab, 0xb6, 0xd2, 0xbf, 0x6a,
	0xeb, 0x42, 0xfe, 0x55, 0x1b, 0xcf, 0xb3, 0x1d, 0x8b, 0xee, 0xf2, 0xb6, 0xd0, 0x3b, 0xb0, 0x35,
	0x61, 0x84, 0x33, 0xe8, 0xa2, 0xb9, 0x1a, 0x7b, 0x2b, 0x2a, 0xf3, 0x09, 0x5f, 0xc1, 0x5d, 0x7e,
	0x3a, 0xf4, 0xb2, 0xd0, 0xb0, 0x6c, 0xa7, 0x86, 0xb7, 0x0e, 0x32, 0x1f, 0x1e, 0xc0, 0x93, 0x15,
	0xef, 0x81, 0x5e, 0x15, 0x9a, 0x57, 0xbc, 0x72, 0xe3, 0xf5, 0x06, 0x54, 0x7a, 0xcb, 0xb0, 0xa2,
	0x44, 0x3b, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x38, 0xd2, 0xcf, 0xb0, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PoolServiceClient is the client API for PoolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PoolServiceClient interface {
	CountPools(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountPoolsResponse, error)
	GetPool(ctx context.Context, in *GetPoolRequest, opts ...grpc.CallOption) (*GetPoolResponse, error)
	GetPoolAddresses(ctx context.Context, in *GetPoolAddressesRequest, opts ...grpc.CallOption) (*GetPoolAddressesResponse, error)
	GetPoolCoinbaseTags(ctx context.Context, in *GetPoolCoinbaseTagsRequest, opts ...grpc.CallOption) (*GetPoolCoinbaseTagsResponse, error)
}

type poolServiceClient struct {
	cc *grpc.ClientConn
}

func NewPoolServiceClient(cc *grpc.ClientConn) PoolServiceClient {
	return &poolServiceClient{cc}
}

func (c *poolServiceClient) CountPools(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountPoolsResponse, error) {
	out := new(CountPoolsResponse)
	err := c.cc.Invoke(ctx, "/basedata.PoolService/CountPools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolServiceClient) GetPool(ctx context.Context, in *GetPoolRequest, opts ...grpc.CallOption) (*GetPoolResponse, error) {
	out := new(GetPoolResponse)
	err := c.cc.Invoke(ctx, "/basedata.PoolService/GetPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolServiceClient) GetPoolAddresses(ctx context.Context, in *GetPoolAddressesRequest, opts ...grpc.CallOption) (*GetPoolAddressesResponse, error) {
	out := new(GetPoolAddressesResponse)
	err := c.cc.Invoke(ctx, "/basedata.PoolService/GetPoolAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolServiceClient) GetPoolCoinbaseTags(ctx context.Context, in *GetPoolCoinbaseTagsRequest, opts ...grpc.CallOption) (*GetPoolCoinbaseTagsResponse, error) {
	out := new(GetPoolCoinbaseTagsResponse)
	err := c.cc.Invoke(ctx, "/basedata.PoolService/GetPoolCoinbaseTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoolServiceServer is the server API for PoolService service.
type PoolServiceServer interface {
	CountPools(context.Context, *empty.Empty) (*CountPoolsResponse, error)
	GetPool(context.Context, *GetPoolRequest) (*GetPoolResponse, error)
	GetPoolAddresses(context.Context, *GetPoolAddressesRequest) (*GetPoolAddressesResponse, error)
	GetPoolCoinbaseTags(context.Context, *GetPoolCoinbaseTagsRequest) (*GetPoolCoinbaseTagsResponse, error)
}

// UnimplementedPoolServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPoolServiceServer struct {
}

func (*UnimplementedPoolServiceServer) CountPools(ctx context.Context, req *empty.Empty) (*CountPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountPools not implemented")
}
func (*UnimplementedPoolServiceServer) GetPool(ctx context.Context, req *GetPoolRequest) (*GetPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPool not implemented")
}
func (*UnimplementedPoolServiceServer) GetPoolAddresses(ctx context.Context, req *GetPoolAddressesRequest) (*GetPoolAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoolAddresses not implemented")
}
func (*UnimplementedPoolServiceServer) GetPoolCoinbaseTags(ctx context.Context, req *GetPoolCoinbaseTagsRequest) (*GetPoolCoinbaseTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoolCoinbaseTags not implemented")
}

func RegisterPoolServiceServer(s *grpc.Server, srv PoolServiceServer) {
	s.RegisterService(&_PoolService_serviceDesc, srv)
}

func _PoolService_CountPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServiceServer).CountPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basedata.PoolService/CountPools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServiceServer).CountPools(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolService_GetPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServiceServer).GetPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basedata.PoolService/GetPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServiceServer).GetPool(ctx, req.(*GetPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolService_GetPoolAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServiceServer).GetPoolAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basedata.PoolService/GetPoolAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServiceServer).GetPoolAddresses(ctx, req.(*GetPoolAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolService_GetPoolCoinbaseTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolCoinbaseTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServiceServer).GetPoolCoinbaseTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/basedata.PoolService/GetPoolCoinbaseTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServiceServer).GetPoolCoinbaseTags(ctx, req.(*GetPoolCoinbaseTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PoolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "basedata.PoolService",
	HandlerType: (*PoolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountPools",
			Handler:    _PoolService_CountPools_Handler,
		},
		{
			MethodName: "GetPool",
			Handler:    _PoolService_GetPool_Handler,
		},
		{
			MethodName: "GetPoolAddresses",
			Handler:    _PoolService_GetPoolAddresses_Handler,
		},
		{
			MethodName: "GetPoolCoinbaseTags",
			Handler:    _PoolService_GetPoolCoinbaseTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pool.proto",
}
