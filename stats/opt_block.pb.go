// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opt_block.proto

package stats

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type ListOptBlocksRequest_QueryField int32

const (
	ListOptBlocksRequest_POOL_ID      ListOptBlocksRequest_QueryField = 0
	ListOptBlocksRequest_HEIGHT       ListOptBlocksRequest_QueryField = 1
	ListOptBlocksRequest_TIMESTAMP    ListOptBlocksRequest_QueryField = 2
	ListOptBlocksRequest_MEDIAN_AT_TS ListOptBlocksRequest_QueryField = 3
)

var ListOptBlocksRequest_QueryField_name = map[int32]string{
	0: "POOL_ID",
	1: "HEIGHT",
	2: "TIMESTAMP",
	3: "MEDIAN_AT_TS",
}

var ListOptBlocksRequest_QueryField_value = map[string]int32{
	"POOL_ID":      0,
	"HEIGHT":       1,
	"TIMESTAMP":    2,
	"MEDIAN_AT_TS": 3,
}

func (x ListOptBlocksRequest_QueryField) String() string {
	return proto.EnumName(ListOptBlocksRequest_QueryField_name, int32(x))
}

func (ListOptBlocksRequest_QueryField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 0}
}

type ListOptBlocksRequest_QueryOperator int32

const (
	ListOptBlocksRequest_EQ ListOptBlocksRequest_QueryOperator = 0
	ListOptBlocksRequest_NE ListOptBlocksRequest_QueryOperator = 1
	ListOptBlocksRequest_GT ListOptBlocksRequest_QueryOperator = 2
	ListOptBlocksRequest_LT ListOptBlocksRequest_QueryOperator = 3
	ListOptBlocksRequest_LE ListOptBlocksRequest_QueryOperator = 4
	ListOptBlocksRequest_GE ListOptBlocksRequest_QueryOperator = 5
)

var ListOptBlocksRequest_QueryOperator_name = map[int32]string{
	0: "EQ",
	1: "NE",
	2: "GT",
	3: "LT",
	4: "LE",
	5: "GE",
}

var ListOptBlocksRequest_QueryOperator_value = map[string]int32{
	"EQ": 0,
	"NE": 1,
	"GT": 2,
	"LT": 3,
	"LE": 4,
	"GE": 5,
}

func (x ListOptBlocksRequest_QueryOperator) String() string {
	return proto.EnumName(ListOptBlocksRequest_QueryOperator_name, int32(x))
}

func (ListOptBlocksRequest_QueryOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 1}
}

type OptBlock struct {
	Id                      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hash                    string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Height                  int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	PreviousHash            string   `protobuf:"bytes,4,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	Size                    int32    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Difficulty              float64  `protobuf:"fixed64,6,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	PoolId                  int32    `protobuf:"varint,7,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	IdentifyPoolMode        string   `protobuf:"bytes,8,opt,name=identify_pool_mode,json=identifyPoolMode,proto3" json:"identify_pool_mode,omitempty"`
	CoinbaseValue           int64    `protobuf:"varint,9,opt,name=coinbase_value,json=coinbaseValue,proto3" json:"coinbase_value,omitempty"`
	CoinbaseFee             int64    `protobuf:"varint,10,opt,name=coinbase_fee,json=coinbaseFee,proto3" json:"coinbase_fee,omitempty"`
	CoinbaseReward          int64    `protobuf:"varint,11,opt,name=coinbase_reward,json=coinbaseReward,proto3" json:"coinbase_reward,omitempty"`
	CoinbaseEstimatedReward int64    `protobuf:"varint,12,opt,name=coinbase_estimated_reward,json=coinbaseEstimatedReward,proto3" json:"coinbase_estimated_reward,omitempty"`
	Version                 int32    `protobuf:"varint,13,opt,name=version,proto3" json:"version,omitempty"`
	TxCount                 int32    `protobuf:"varint,14,opt,name=tx_count,json=txCount,proto3" json:"tx_count,omitempty"`
	Timestamp               int32    `protobuf:"varint,15,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CreatedAtTs             int32    `protobuf:"varint,16,opt,name=created_at_ts,json=createdAtTs,proto3" json:"created_at_ts,omitempty"`
	UpdatedAtTs             int32    `protobuf:"varint,17,opt,name=updated_at_ts,json=updatedAtTs,proto3" json:"updated_at_ts,omitempty"`
	MedianAtTs              int32    `protobuf:"varint,18,opt,name=median_at_ts,json=medianAtTs,proto3" json:"median_at_ts,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *OptBlock) Reset()         { *m = OptBlock{} }
func (m *OptBlock) String() string { return proto.CompactTextString(m) }
func (*OptBlock) ProtoMessage()    {}
func (*OptBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{0}
}

func (m *OptBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptBlock.Unmarshal(m, b)
}
func (m *OptBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptBlock.Marshal(b, m, deterministic)
}
func (m *OptBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptBlock.Merge(m, src)
}
func (m *OptBlock) XXX_Size() int {
	return xxx_messageInfo_OptBlock.Size(m)
}
func (m *OptBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_OptBlock.DiscardUnknown(m)
}

var xxx_messageInfo_OptBlock proto.InternalMessageInfo

func (m *OptBlock) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OptBlock) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *OptBlock) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *OptBlock) GetPreviousHash() string {
	if m != nil {
		return m.PreviousHash
	}
	return ""
}

func (m *OptBlock) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *OptBlock) GetDifficulty() float64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *OptBlock) GetPoolId() int32 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *OptBlock) GetIdentifyPoolMode() string {
	if m != nil {
		return m.IdentifyPoolMode
	}
	return ""
}

func (m *OptBlock) GetCoinbaseValue() int64 {
	if m != nil {
		return m.CoinbaseValue
	}
	return 0
}

func (m *OptBlock) GetCoinbaseFee() int64 {
	if m != nil {
		return m.CoinbaseFee
	}
	return 0
}

func (m *OptBlock) GetCoinbaseReward() int64 {
	if m != nil {
		return m.CoinbaseReward
	}
	return 0
}

func (m *OptBlock) GetCoinbaseEstimatedReward() int64 {
	if m != nil {
		return m.CoinbaseEstimatedReward
	}
	return 0
}

func (m *OptBlock) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *OptBlock) GetTxCount() int32 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

func (m *OptBlock) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *OptBlock) GetCreatedAtTs() int32 {
	if m != nil {
		return m.CreatedAtTs
	}
	return 0
}

func (m *OptBlock) GetUpdatedAtTs() int32 {
	if m != nil {
		return m.UpdatedAtTs
	}
	return 0
}

func (m *OptBlock) GetMedianAtTs() int32 {
	if m != nil {
		return m.MedianAtTs
	}
	return 0
}

type ListOptBlocksRequest struct {
	Queries              []*ListOptBlocksRequest_Query `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ListOptBlocksRequest) Reset()         { *m = ListOptBlocksRequest{} }
func (m *ListOptBlocksRequest) String() string { return proto.CompactTextString(m) }
func (*ListOptBlocksRequest) ProtoMessage()    {}
func (*ListOptBlocksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1}
}

func (m *ListOptBlocksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOptBlocksRequest.Unmarshal(m, b)
}
func (m *ListOptBlocksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOptBlocksRequest.Marshal(b, m, deterministic)
}
func (m *ListOptBlocksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOptBlocksRequest.Merge(m, src)
}
func (m *ListOptBlocksRequest) XXX_Size() int {
	return xxx_messageInfo_ListOptBlocksRequest.Size(m)
}
func (m *ListOptBlocksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOptBlocksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListOptBlocksRequest proto.InternalMessageInfo

func (m *ListOptBlocksRequest) GetQueries() []*ListOptBlocksRequest_Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type ListOptBlocksRequest_Query struct {
	Field                ListOptBlocksRequest_QueryField    `protobuf:"varint,1,opt,name=field,proto3,enum=stats.ListOptBlocksRequest_QueryField" json:"field,omitempty"`
	Operator             ListOptBlocksRequest_QueryOperator `protobuf:"varint,2,opt,name=operator,proto3,enum=stats.ListOptBlocksRequest_QueryOperator" json:"operator,omitempty"`
	Value                *any.Any                           `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ListOptBlocksRequest_Query) Reset()         { *m = ListOptBlocksRequest_Query{} }
func (m *ListOptBlocksRequest_Query) String() string { return proto.CompactTextString(m) }
func (*ListOptBlocksRequest_Query) ProtoMessage()    {}
func (*ListOptBlocksRequest_Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 0}
}

func (m *ListOptBlocksRequest_Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOptBlocksRequest_Query.Unmarshal(m, b)
}
func (m *ListOptBlocksRequest_Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOptBlocksRequest_Query.Marshal(b, m, deterministic)
}
func (m *ListOptBlocksRequest_Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOptBlocksRequest_Query.Merge(m, src)
}
func (m *ListOptBlocksRequest_Query) XXX_Size() int {
	return xxx_messageInfo_ListOptBlocksRequest_Query.Size(m)
}
func (m *ListOptBlocksRequest_Query) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOptBlocksRequest_Query.DiscardUnknown(m)
}

var xxx_messageInfo_ListOptBlocksRequest_Query proto.InternalMessageInfo

func (m *ListOptBlocksRequest_Query) GetField() ListOptBlocksRequest_QueryField {
	if m != nil {
		return m.Field
	}
	return ListOptBlocksRequest_POOL_ID
}

func (m *ListOptBlocksRequest_Query) GetOperator() ListOptBlocksRequest_QueryOperator {
	if m != nil {
		return m.Operator
	}
	return ListOptBlocksRequest_EQ
}

func (m *ListOptBlocksRequest_Query) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type ListOptBlocksResponse struct {
	OptBlocks            []*OptBlock `protobuf:"bytes,1,rep,name=optBlocks,proto3" json:"optBlocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListOptBlocksResponse) Reset()         { *m = ListOptBlocksResponse{} }
func (m *ListOptBlocksResponse) String() string { return proto.CompactTextString(m) }
func (*ListOptBlocksResponse) ProtoMessage()    {}
func (*ListOptBlocksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{2}
}

func (m *ListOptBlocksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOptBlocksResponse.Unmarshal(m, b)
}
func (m *ListOptBlocksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOptBlocksResponse.Marshal(b, m, deterministic)
}
func (m *ListOptBlocksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOptBlocksResponse.Merge(m, src)
}
func (m *ListOptBlocksResponse) XXX_Size() int {
	return xxx_messageInfo_ListOptBlocksResponse.Size(m)
}
func (m *ListOptBlocksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOptBlocksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOptBlocksResponse proto.InternalMessageInfo

func (m *ListOptBlocksResponse) GetOptBlocks() []*OptBlock {
	if m != nil {
		return m.OptBlocks
	}
	return nil
}

func init() {
	proto.RegisterEnum("stats.ListOptBlocksRequest_QueryField", ListOptBlocksRequest_QueryField_name, ListOptBlocksRequest_QueryField_value)
	proto.RegisterEnum("stats.ListOptBlocksRequest_QueryOperator", ListOptBlocksRequest_QueryOperator_name, ListOptBlocksRequest_QueryOperator_value)
	proto.RegisterType((*OptBlock)(nil), "stats.OptBlock")
	proto.RegisterType((*ListOptBlocksRequest)(nil), "stats.ListOptBlocksRequest")
	proto.RegisterType((*ListOptBlocksRequest_Query)(nil), "stats.ListOptBlocksRequest.Query")
	proto.RegisterType((*ListOptBlocksResponse)(nil), "stats.ListOptBlocksResponse")
}

func init() { proto.RegisterFile("opt_block.proto", fileDescriptor_52510bcac66c2e66) }

var fileDescriptor_52510bcac66c2e66 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x6d, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x97, 0x3e, 0xf7, 0xfa, 0x14, 0x4e, 0x83, 0x79, 0x63, 0x42, 0x5d, 0x11, 0x50, 0x10,
	0x74, 0x52, 0x79, 0x07, 0x48, 0xa8, 0xb0, 0x74, 0xab, 0xd4, 0xae, 0x5b, 0x16, 0xf1, 0x36, 0x4a,
	0x1b, 0x77, 0xb5, 0x68, 0xe3, 0x2c, 0x76, 0xca, 0xca, 0x17, 0xe1, 0xab, 0xf0, 0x39, 0xf8, 0x44,
	0x28, 0x4e, 0xd2, 0x6d, 0x68, 0x62, 0xaf, 0x2e, 0xf7, 0xbf, 0xdf, 0xdf, 0x3e, 0x3b, 0x67, 0x68,
	0x70, 0x5f, 0xda, 0x93, 0x05, 0x9f, 0x7e, 0xef, 0xf8, 0x01, 0x97, 0x1c, 0xf3, 0x42, 0x3a, 0x52,
	0xec, 0xed, 0x5e, 0x72, 0x7e, 0xb9, 0xa0, 0x87, 0x4a, 0x9c, 0x84, 0xb3, 0x43, 0xc7, 0x5b, 0xc7,
	0x44, 0xeb, 0x4f, 0x0e, 0x4a, 0x63, 0x5f, 0x7e, 0x89, 0x4c, 0x58, 0x87, 0x0c, 0x73, 0x89, 0xd6,
	0xd4, 0xda, 0x65, 0x33, 0xc3, 0x5c, 0x44, 0xc8, 0xcd, 0x1d, 0x31, 0x27, 0x19, 0xa5, 0xa8, 0x6f,
	0x7c, 0x02, 0x85, 0x39, 0x65, 0x97, 0x73, 0x49, 0xb2, 0x4d, 0xad, 0x9d, 0x37, 0x93, 0x0c, 0x9f,
	0x43, 0xcd, 0x0f, 0xe8, 0x8a, 0xf1, 0x50, 0xd8, 0xca, 0x94, 0x53, 0xa6, 0x6a, 0x2a, 0x9e, 0x44,
	0x66, 0x84, 0x9c, 0x60, 0x3f, 0x29, 0xc9, 0x2b, 0xab, 0xfa, 0xc6, 0x67, 0x00, 0x2e, 0x9b, 0xcd,
	0xd8, 0x34, 0x5c, 0xc8, 0x35, 0x29, 0x34, 0xb5, 0xb6, 0x66, 0xde, 0x52, 0x70, 0x07, 0x8a, 0x3e,
	0xe7, 0x0b, 0x9b, 0xb9, 0xa4, 0x18, 0xef, 0x18, 0xa5, 0x03, 0x17, 0xdf, 0x02, 0x32, 0x97, 0x7a,
	0x92, 0xcd, 0xd6, 0xb6, 0x22, 0x96, 0xdc, 0xa5, 0xa4, 0xa4, 0xb6, 0xd5, 0xd3, 0xca, 0x19, 0xe7,
	0x8b, 0x11, 0x77, 0x29, 0xbe, 0x80, 0xfa, 0x94, 0x33, 0x6f, 0xe2, 0x08, 0x6a, 0xaf, 0x9c, 0x45,
	0x48, 0x49, 0xb9, 0xa9, 0xb5, 0xb3, 0x66, 0x2d, 0x55, 0xbf, 0x45, 0x22, 0x1e, 0x40, 0x75, 0x83,
	0xcd, 0x28, 0x25, 0xa0, 0xa0, 0x4a, 0xaa, 0xf5, 0x29, 0xc5, 0x57, 0xd0, 0xd8, 0x20, 0x01, 0xfd,
	0xe1, 0x04, 0x2e, 0xa9, 0x28, 0x6a, 0xb3, 0x81, 0xa9, 0x54, 0xfc, 0x00, 0xbb, 0x1b, 0x90, 0x0a,
	0xc9, 0x96, 0x8e, 0xa4, 0x6e, 0x6a, 0xa9, 0x2a, 0xcb, 0x4e, 0x0a, 0x18, 0x69, 0x3d, 0xf1, 0x12,
	0x28, 0xae, 0x68, 0x20, 0x18, 0xf7, 0x48, 0x4d, 0x9d, 0x3a, 0x4d, 0x71, 0x17, 0x4a, 0xf2, 0xda,
	0x9e, 0xf2, 0xd0, 0x93, 0xa4, 0x1e, 0x97, 0xe4, 0xf5, 0xd7, 0x28, 0xc5, 0x7d, 0x28, 0x4b, 0xb6,
	0xa4, 0x42, 0x3a, 0x4b, 0x9f, 0x34, 0x54, 0xed, 0x46, 0xc0, 0x16, 0xd4, 0xa6, 0x01, 0x55, 0x3d,
	0x38, 0xd2, 0x96, 0x82, 0xe8, 0x8a, 0xa8, 0x24, 0x62, 0x4f, 0x5a, 0x22, 0x62, 0x42, 0xdf, 0xbd,
	0xc5, 0x3c, 0x8a, 0x99, 0x44, 0x54, 0x4c, 0x13, 0xaa, 0x4b, 0xea, 0x32, 0xc7, 0x4b, 0x10, 0x54,
	0x08, 0xc4, 0x5a, 0x44, 0xb4, 0x7e, 0x65, 0x61, 0x7b, 0xc8, 0x84, 0x4c, 0x07, 0x4b, 0x98, 0xf4,
	0x2a, 0xa4, 0x42, 0xe2, 0x47, 0x28, 0x5e, 0x85, 0x34, 0x60, 0x54, 0x10, 0xad, 0x99, 0x6d, 0x57,
	0xba, 0x07, 0x1d, 0x35, 0xa1, 0x9d, 0xfb, 0xe8, 0xce, 0x79, 0x48, 0x83, 0xb5, 0x99, 0x3a, 0xf6,
	0x7e, 0x6b, 0x90, 0x57, 0x12, 0x7e, 0x82, 0xfc, 0x8c, 0xd1, 0x45, 0x3c, 0xaa, 0xf5, 0xee, 0xcb,
	0x07, 0x17, 0xe9, 0x47, 0xb4, 0x19, 0x9b, 0xd0, 0x80, 0x12, 0xf7, 0x69, 0xe0, 0x48, 0x1e, 0xa8,
	0xc9, 0xae, 0x77, 0x5f, 0x3f, 0xb8, 0xc0, 0x38, 0x31, 0x98, 0x1b, 0x2b, 0xbe, 0x81, 0x7c, 0x3c,
	0x47, 0xd1, 0x3b, 0xa8, 0x74, 0xb7, 0x3b, 0xf1, 0x23, 0xeb, 0xa4, 0x8f, 0xac, 0xd3, 0xf3, 0xd6,
	0x66, 0x8c, 0xb4, 0xfa, 0x00, 0x37, 0x7d, 0x60, 0x05, 0x8a, 0x67, 0xe3, 0xf1, 0xd0, 0x1e, 0x1c,
	0xe9, 0x5b, 0x08, 0x50, 0x38, 0x31, 0x06, 0xc7, 0x27, 0x96, 0xae, 0x61, 0x0d, 0xca, 0xd6, 0x60,
	0x64, 0x5c, 0x58, 0xbd, 0xd1, 0x99, 0x9e, 0x41, 0x1d, 0xaa, 0x23, 0xe3, 0x68, 0xd0, 0x3b, 0xb5,
	0x7b, 0x96, 0x6d, 0x5d, 0xe8, 0xd9, 0xd6, 0x67, 0xa8, 0xdd, 0x69, 0x07, 0x0b, 0x90, 0x31, 0xce,
	0xf5, 0xad, 0x28, 0x9e, 0x1a, 0xba, 0x16, 0xc5, 0x63, 0x4b, 0xcf, 0x44, 0x71, 0x68, 0xe9, 0x59,
	0x15, 0x0d, 0x3d, 0xa7, 0x74, 0x43, 0xcf, 0xb7, 0xfa, 0xf0, 0xf8, 0x9f, 0x43, 0x0a, 0x9f, 0x7b,
	0x82, 0xe2, 0x3b, 0x28, 0xf3, 0x54, 0x4c, 0xfe, 0x4d, 0x23, 0xb9, 0x95, 0x14, 0x36, 0x6f, 0x88,
	0xae, 0x0d, 0x8d, 0x54, 0xbe, 0xa0, 0xc1, 0x8a, 0x4d, 0x29, 0x0e, 0xa1, 0x76, 0x67, 0x69, 0x7c,
	0xfa, 0x9f, 0x5b, 0xdd, 0xdb, 0xbf, 0xbf, 0x18, 0x77, 0xd3, 0xda, 0x9a, 0x14, 0xd4, 0x35, 0xbe,
	0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x59, 0x96, 0x60, 0xd3, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OptBlockServiceClient is the client API for OptBlockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OptBlockServiceClient interface {
	ListOptBlocks(ctx context.Context, in *ListOptBlocksRequest, opts ...grpc.CallOption) (*ListOptBlocksResponse, error)
}

type optBlockServiceClient struct {
	cc *grpc.ClientConn
}

func NewOptBlockServiceClient(cc *grpc.ClientConn) OptBlockServiceClient {
	return &optBlockServiceClient{cc}
}

func (c *optBlockServiceClient) ListOptBlocks(ctx context.Context, in *ListOptBlocksRequest, opts ...grpc.CallOption) (*ListOptBlocksResponse, error) {
	out := new(ListOptBlocksResponse)
	err := c.cc.Invoke(ctx, "/stats.OptBlockService/ListOptBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OptBlockServiceServer is the server API for OptBlockService service.
type OptBlockServiceServer interface {
	ListOptBlocks(context.Context, *ListOptBlocksRequest) (*ListOptBlocksResponse, error)
}

// UnimplementedOptBlockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOptBlockServiceServer struct {
}

func (*UnimplementedOptBlockServiceServer) ListOptBlocks(ctx context.Context, req *ListOptBlocksRequest) (*ListOptBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOptBlocks not implemented")
}

func RegisterOptBlockServiceServer(s *grpc.Server, srv OptBlockServiceServer) {
	s.RegisterService(&_OptBlockService_serviceDesc, srv)
}

func _OptBlockService_ListOptBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOptBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptBlockServiceServer).ListOptBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.OptBlockService/ListOptBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptBlockServiceServer).ListOptBlocks(ctx, req.(*ListOptBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OptBlockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stats.OptBlockService",
	HandlerType: (*OptBlockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOptBlocks",
			Handler:    _OptBlockService_ListOptBlocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "opt_block.proto",
}
