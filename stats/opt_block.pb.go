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
	ListOptBlocksRequest_Q_POOL_ID      ListOptBlocksRequest_QueryField = 0
	ListOptBlocksRequest_Q_HEIGHT       ListOptBlocksRequest_QueryField = 1
	ListOptBlocksRequest_Q_TIMESTAMP    ListOptBlocksRequest_QueryField = 2
	ListOptBlocksRequest_Q_MEDIAN_AT_TS ListOptBlocksRequest_QueryField = 3
)

var ListOptBlocksRequest_QueryField_name = map[int32]string{
	0: "Q_POOL_ID",
	1: "Q_HEIGHT",
	2: "Q_TIMESTAMP",
	3: "Q_MEDIAN_AT_TS",
}

var ListOptBlocksRequest_QueryField_value = map[string]int32{
	"Q_POOL_ID":      0,
	"Q_HEIGHT":       1,
	"Q_TIMESTAMP":    2,
	"Q_MEDIAN_AT_TS": 3,
}

func (x ListOptBlocksRequest_QueryField) String() string {
	return proto.EnumName(ListOptBlocksRequest_QueryField_name, int32(x))
}

func (ListOptBlocksRequest_QueryField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 0}
}

type ListOptBlocksRequest_SortField int32

const (
	ListOptBlocksRequest_S_HEIGHT   ListOptBlocksRequest_SortField = 0
	ListOptBlocksRequest_S_SIZE     ListOptBlocksRequest_SortField = 1
	ListOptBlocksRequest_S_TX_COUNT ListOptBlocksRequest_SortField = 2
)

var ListOptBlocksRequest_SortField_name = map[int32]string{
	0: "S_HEIGHT",
	1: "S_SIZE",
	2: "S_TX_COUNT",
}

var ListOptBlocksRequest_SortField_value = map[string]int32{
	"S_HEIGHT":   0,
	"S_SIZE":     1,
	"S_TX_COUNT": 2,
}

func (x ListOptBlocksRequest_SortField) String() string {
	return proto.EnumName(ListOptBlocksRequest_SortField_name, int32(x))
}

func (ListOptBlocksRequest_SortField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 1}
}

type ListOptBlocksRequest_SortDirection int32

const (
	ListOptBlocksRequest_ASC  ListOptBlocksRequest_SortDirection = 0
	ListOptBlocksRequest_DESC ListOptBlocksRequest_SortDirection = 1
)

var ListOptBlocksRequest_SortDirection_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}

var ListOptBlocksRequest_SortDirection_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x ListOptBlocksRequest_SortDirection) String() string {
	return proto.EnumName(ListOptBlocksRequest_SortDirection_name, int32(x))
}

func (ListOptBlocksRequest_SortDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 2}
}

type ListOptBlocksRequest_QueryOperator int32

const (
	ListOptBlocksRequest_EQ  ListOptBlocksRequest_QueryOperator = 0
	ListOptBlocksRequest_NE  ListOptBlocksRequest_QueryOperator = 1
	ListOptBlocksRequest_GTE ListOptBlocksRequest_QueryOperator = 2
	ListOptBlocksRequest_LTE ListOptBlocksRequest_QueryOperator = 3
	ListOptBlocksRequest_LE  ListOptBlocksRequest_QueryOperator = 4
	ListOptBlocksRequest_GE  ListOptBlocksRequest_QueryOperator = 5
)

var ListOptBlocksRequest_QueryOperator_name = map[int32]string{
	0: "EQ",
	1: "NE",
	2: "GTE",
	3: "LTE",
	4: "LE",
	5: "GE",
}

var ListOptBlocksRequest_QueryOperator_value = map[string]int32{
	"EQ":  0,
	"NE":  1,
	"GTE": 2,
	"LTE": 3,
	"LE":  4,
	"GE":  5,
}

func (x ListOptBlocksRequest_QueryOperator) String() string {
	return proto.EnumName(ListOptBlocksRequest_QueryOperator_name, int32(x))
}

func (ListOptBlocksRequest_QueryOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 3}
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
	Sorts                []*ListOptBlocksRequest_Sort  `protobuf:"bytes,2,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Limit                int32                         `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Skip                 int32                         `protobuf:"varint,4,opt,name=skip,proto3" json:"skip,omitempty"`
	OnlyTotal            bool                          `protobuf:"varint,5,opt,name=onlyTotal,proto3" json:"onlyTotal,omitempty"`
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

func (m *ListOptBlocksRequest) GetSorts() []*ListOptBlocksRequest_Sort {
	if m != nil {
		return m.Sorts
	}
	return nil
}

func (m *ListOptBlocksRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListOptBlocksRequest) GetSkip() int32 {
	if m != nil {
		return m.Skip
	}
	return 0
}

func (m *ListOptBlocksRequest) GetOnlyTotal() bool {
	if m != nil {
		return m.OnlyTotal
	}
	return false
}

type ListOptBlocksRequest_QueryInt32Value struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOptBlocksRequest_QueryInt32Value) Reset()         { *m = ListOptBlocksRequest_QueryInt32Value{} }
func (m *ListOptBlocksRequest_QueryInt32Value) String() string { return proto.CompactTextString(m) }
func (*ListOptBlocksRequest_QueryInt32Value) ProtoMessage()    {}
func (*ListOptBlocksRequest_QueryInt32Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 0}
}

func (m *ListOptBlocksRequest_QueryInt32Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOptBlocksRequest_QueryInt32Value.Unmarshal(m, b)
}
func (m *ListOptBlocksRequest_QueryInt32Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOptBlocksRequest_QueryInt32Value.Marshal(b, m, deterministic)
}
func (m *ListOptBlocksRequest_QueryInt32Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOptBlocksRequest_QueryInt32Value.Merge(m, src)
}
func (m *ListOptBlocksRequest_QueryInt32Value) XXX_Size() int {
	return xxx_messageInfo_ListOptBlocksRequest_QueryInt32Value.Size(m)
}
func (m *ListOptBlocksRequest_QueryInt32Value) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOptBlocksRequest_QueryInt32Value.DiscardUnknown(m)
}

var xxx_messageInfo_ListOptBlocksRequest_QueryInt32Value proto.InternalMessageInfo

func (m *ListOptBlocksRequest_QueryInt32Value) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ListOptBlocksRequest_QueryDoubleValue struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOptBlocksRequest_QueryDoubleValue) Reset()         { *m = ListOptBlocksRequest_QueryDoubleValue{} }
func (m *ListOptBlocksRequest_QueryDoubleValue) String() string { return proto.CompactTextString(m) }
func (*ListOptBlocksRequest_QueryDoubleValue) ProtoMessage()    {}
func (*ListOptBlocksRequest_QueryDoubleValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 1}
}

func (m *ListOptBlocksRequest_QueryDoubleValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOptBlocksRequest_QueryDoubleValue.Unmarshal(m, b)
}
func (m *ListOptBlocksRequest_QueryDoubleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOptBlocksRequest_QueryDoubleValue.Marshal(b, m, deterministic)
}
func (m *ListOptBlocksRequest_QueryDoubleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOptBlocksRequest_QueryDoubleValue.Merge(m, src)
}
func (m *ListOptBlocksRequest_QueryDoubleValue) XXX_Size() int {
	return xxx_messageInfo_ListOptBlocksRequest_QueryDoubleValue.Size(m)
}
func (m *ListOptBlocksRequest_QueryDoubleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOptBlocksRequest_QueryDoubleValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListOptBlocksRequest_QueryDoubleValue proto.InternalMessageInfo

func (m *ListOptBlocksRequest_QueryDoubleValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ListOptBlocksRequest_QueryStringValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOptBlocksRequest_QueryStringValue) Reset()         { *m = ListOptBlocksRequest_QueryStringValue{} }
func (m *ListOptBlocksRequest_QueryStringValue) String() string { return proto.CompactTextString(m) }
func (*ListOptBlocksRequest_QueryStringValue) ProtoMessage()    {}
func (*ListOptBlocksRequest_QueryStringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 2}
}

func (m *ListOptBlocksRequest_QueryStringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOptBlocksRequest_QueryStringValue.Unmarshal(m, b)
}
func (m *ListOptBlocksRequest_QueryStringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOptBlocksRequest_QueryStringValue.Marshal(b, m, deterministic)
}
func (m *ListOptBlocksRequest_QueryStringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOptBlocksRequest_QueryStringValue.Merge(m, src)
}
func (m *ListOptBlocksRequest_QueryStringValue) XXX_Size() int {
	return xxx_messageInfo_ListOptBlocksRequest_QueryStringValue.Size(m)
}
func (m *ListOptBlocksRequest_QueryStringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOptBlocksRequest_QueryStringValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListOptBlocksRequest_QueryStringValue proto.InternalMessageInfo

func (m *ListOptBlocksRequest_QueryStringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
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
	return fileDescriptor_52510bcac66c2e66, []int{1, 3}
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
	return ListOptBlocksRequest_Q_POOL_ID
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

type ListOptBlocksRequest_Sort struct {
	Field                ListOptBlocksRequest_SortField     `protobuf:"varint,1,opt,name=field,proto3,enum=stats.ListOptBlocksRequest_SortField" json:"field,omitempty"`
	Direction            ListOptBlocksRequest_SortDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=stats.ListOptBlocksRequest_SortDirection" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ListOptBlocksRequest_Sort) Reset()         { *m = ListOptBlocksRequest_Sort{} }
func (m *ListOptBlocksRequest_Sort) String() string { return proto.CompactTextString(m) }
func (*ListOptBlocksRequest_Sort) ProtoMessage()    {}
func (*ListOptBlocksRequest_Sort) Descriptor() ([]byte, []int) {
	return fileDescriptor_52510bcac66c2e66, []int{1, 4}
}

func (m *ListOptBlocksRequest_Sort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOptBlocksRequest_Sort.Unmarshal(m, b)
}
func (m *ListOptBlocksRequest_Sort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOptBlocksRequest_Sort.Marshal(b, m, deterministic)
}
func (m *ListOptBlocksRequest_Sort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOptBlocksRequest_Sort.Merge(m, src)
}
func (m *ListOptBlocksRequest_Sort) XXX_Size() int {
	return xxx_messageInfo_ListOptBlocksRequest_Sort.Size(m)
}
func (m *ListOptBlocksRequest_Sort) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOptBlocksRequest_Sort.DiscardUnknown(m)
}

var xxx_messageInfo_ListOptBlocksRequest_Sort proto.InternalMessageInfo

func (m *ListOptBlocksRequest_Sort) GetField() ListOptBlocksRequest_SortField {
	if m != nil {
		return m.Field
	}
	return ListOptBlocksRequest_S_HEIGHT
}

func (m *ListOptBlocksRequest_Sort) GetDirection() ListOptBlocksRequest_SortDirection {
	if m != nil {
		return m.Direction
	}
	return ListOptBlocksRequest_ASC
}

type ListOptBlocksResponse struct {
	OptBlocks            []*OptBlock `protobuf:"bytes,1,rep,name=optBlocks,proto3" json:"optBlocks,omitempty"`
	Total                int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
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

func (m *ListOptBlocksResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterEnum("stats.ListOptBlocksRequest_QueryField", ListOptBlocksRequest_QueryField_name, ListOptBlocksRequest_QueryField_value)
	proto.RegisterEnum("stats.ListOptBlocksRequest_SortField", ListOptBlocksRequest_SortField_name, ListOptBlocksRequest_SortField_value)
	proto.RegisterEnum("stats.ListOptBlocksRequest_SortDirection", ListOptBlocksRequest_SortDirection_name, ListOptBlocksRequest_SortDirection_value)
	proto.RegisterEnum("stats.ListOptBlocksRequest_QueryOperator", ListOptBlocksRequest_QueryOperator_name, ListOptBlocksRequest_QueryOperator_value)
	proto.RegisterType((*OptBlock)(nil), "stats.OptBlock")
	proto.RegisterType((*ListOptBlocksRequest)(nil), "stats.ListOptBlocksRequest")
	proto.RegisterType((*ListOptBlocksRequest_QueryInt32Value)(nil), "stats.ListOptBlocksRequest.QueryInt32Value")
	proto.RegisterType((*ListOptBlocksRequest_QueryDoubleValue)(nil), "stats.ListOptBlocksRequest.QueryDoubleValue")
	proto.RegisterType((*ListOptBlocksRequest_QueryStringValue)(nil), "stats.ListOptBlocksRequest.QueryStringValue")
	proto.RegisterType((*ListOptBlocksRequest_Query)(nil), "stats.ListOptBlocksRequest.Query")
	proto.RegisterType((*ListOptBlocksRequest_Sort)(nil), "stats.ListOptBlocksRequest.Sort")
	proto.RegisterType((*ListOptBlocksResponse)(nil), "stats.ListOptBlocksResponse")
}

func init() { proto.RegisterFile("opt_block.proto", fileDescriptor_52510bcac66c2e66) }

var fileDescriptor_52510bcac66c2e66 = []byte{
	// 883 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xe3, 0xfc, 0xcf, 0xc9, 0x3f, 0x73, 0x54, 0x58, 0xb7, 0xac, 0x50, 0xd6, 0x68, 0xd9,
	0x80, 0x20, 0x2b, 0x65, 0x05, 0x17, 0x2c, 0x37, 0xa1, 0xf1, 0xb6, 0x91, 0xd2, 0xa6, 0xb1, 0x0d,
	0x42, 0x08, 0xc9, 0x72, 0xe2, 0x49, 0x3b, 0xaa, 0xe3, 0xf1, 0xda, 0xe3, 0xb2, 0xe1, 0x09, 0xb8,
	0xe0, 0x81, 0x78, 0x0e, 0x9e, 0x08, 0xcd, 0x38, 0x93, 0x6c, 0x56, 0x55, 0xbb, 0x57, 0x99, 0xf3,
	0xcd, 0xef, 0x9b, 0x39, 0x99, 0xf9, 0xc6, 0xd0, 0x65, 0x31, 0xf7, 0x16, 0x21, 0x5b, 0xde, 0x0e,
	0xe2, 0x84, 0x71, 0x86, 0x95, 0x94, 0xfb, 0x3c, 0x3d, 0x39, 0xbe, 0x66, 0xec, 0x3a, 0x24, 0x2f,
	0xa5, 0xb8, 0xc8, 0x56, 0x2f, 0xfd, 0x68, 0x93, 0x13, 0xe6, 0x7f, 0x65, 0xa8, 0xcf, 0x62, 0xfe,
	0xb3, 0x30, 0x61, 0x07, 0x8a, 0x34, 0x30, 0xb4, 0x9e, 0xd6, 0x6f, 0xd8, 0x45, 0x1a, 0x20, 0x42,
	0xf9, 0xc6, 0x4f, 0x6f, 0x8c, 0xa2, 0x54, 0xe4, 0x18, 0x3f, 0x83, 0xea, 0x0d, 0xa1, 0xd7, 0x37,
	0xdc, 0x28, 0xf5, 0xb4, 0x7e, 0xc5, 0xde, 0x56, 0xf8, 0x25, 0xb4, 0xe3, 0x84, 0xdc, 0x51, 0x96,
	0xa5, 0x9e, 0x34, 0x95, 0xa5, 0xa9, 0xa5, 0xc4, 0x73, 0x61, 0x46, 0x28, 0xa7, 0xf4, 0x2f, 0x62,
	0x54, 0xa4, 0x55, 0x8e, 0xf1, 0x0b, 0x80, 0x80, 0xae, 0x56, 0x74, 0x99, 0x85, 0x7c, 0x63, 0x54,
	0x7b, 0x5a, 0x5f, 0xb3, 0xdf, 0x53, 0xf0, 0x09, 0xd4, 0x62, 0xc6, 0x42, 0x8f, 0x06, 0x46, 0x2d,
	0xdf, 0x51, 0x94, 0x93, 0x00, 0xbf, 0x05, 0xa4, 0x01, 0x89, 0x38, 0x5d, 0x6d, 0x3c, 0x49, 0xac,
	0x59, 0x40, 0x8c, 0xba, 0xdc, 0x56, 0x57, 0x33, 0x57, 0x8c, 0x85, 0x17, 0x2c, 0x20, 0xf8, 0x1c,
	0x3a, 0x4b, 0x46, 0xa3, 0x85, 0x9f, 0x12, 0xef, 0xce, 0x0f, 0x33, 0x62, 0x34, 0x7a, 0x5a, 0xbf,
	0x64, 0xb7, 0x95, 0xfa, 0xab, 0x10, 0xf1, 0x19, 0xb4, 0x76, 0xd8, 0x8a, 0x10, 0x03, 0x24, 0xd4,
	0x54, 0xda, 0x1b, 0x42, 0xf0, 0x05, 0x74, 0x77, 0x48, 0x42, 0xfe, 0xf4, 0x93, 0xc0, 0x68, 0x4a,
	0x6a, 0xb7, 0x81, 0x2d, 0x55, 0xfc, 0x11, 0x8e, 0x77, 0x20, 0x49, 0x39, 0x5d, 0xfb, 0x9c, 0x04,
	0xca, 0xd2, 0x92, 0x96, 0x27, 0x0a, 0xb0, 0xd4, 0xfc, 0xd6, 0x6b, 0x40, 0xed, 0x8e, 0x24, 0x29,
	0x65, 0x91, 0xd1, 0x96, 0xff, 0x5a, 0x95, 0x78, 0x0c, 0x75, 0xfe, 0xce, 0x5b, 0xb2, 0x2c, 0xe2,
	0x46, 0x27, 0x9f, 0xe2, 0xef, 0x4e, 0x45, 0x89, 0x4f, 0xa1, 0xc1, 0xe9, 0x9a, 0xa4, 0xdc, 0x5f,
	0xc7, 0x46, 0x57, 0xce, 0xed, 0x05, 0x34, 0xa1, 0xbd, 0x4c, 0x88, 0xec, 0xc1, 0xe7, 0x1e, 0x4f,
	0x0d, 0x5d, 0x12, 0xcd, 0xad, 0x38, 0xe2, 0x6e, 0x2a, 0x98, 0x2c, 0x0e, 0xde, 0x63, 0x3e, 0xc9,
	0x99, 0xad, 0x28, 0x99, 0x1e, 0xb4, 0xd6, 0x24, 0xa0, 0x7e, 0xb4, 0x45, 0x50, 0x22, 0x90, 0x6b,
	0x82, 0x30, 0xff, 0xae, 0xc1, 0xd1, 0x94, 0xa6, 0x5c, 0x05, 0x2b, 0xb5, 0xc9, 0xdb, 0x8c, 0xa4,
	0x1c, 0x5f, 0x43, 0xed, 0x6d, 0x46, 0x12, 0x4a, 0x52, 0x43, 0xeb, 0x95, 0xfa, 0xcd, 0xe1, 0xb3,
	0x81, 0x4c, 0xe8, 0xe0, 0x3e, 0x7a, 0x30, 0xcf, 0x48, 0xb2, 0xb1, 0x95, 0x03, 0x7f, 0x80, 0x4a,
	0xca, 0x12, 0x9e, 0x1a, 0x45, 0x69, 0xed, 0x3d, 0x64, 0x75, 0x58, 0xc2, 0xed, 0x1c, 0xc7, 0x23,
	0xa8, 0x84, 0x74, 0x4d, 0x55, 0x60, 0xf3, 0x42, 0x46, 0xf1, 0x96, 0xc6, 0x32, 0xa6, 0x22, 0x8a,
	0xb7, 0x34, 0x16, 0xe7, 0xc7, 0xa2, 0x70, 0xe3, 0x32, 0xee, 0x87, 0x32, 0xa3, 0x75, 0x7b, 0x2f,
	0x9c, 0xbc, 0x80, 0xae, 0xec, 0x68, 0x12, 0xf1, 0x57, 0xc3, 0x3c, 0x2d, 0x47, 0x50, 0xc9, 0xb3,
	0xa4, 0xe5, 0x4b, 0xcb, 0xe2, 0xa4, 0x0f, 0xba, 0x04, 0xc7, 0x2c, 0x5b, 0x84, 0xe4, 0x1e, 0x52,
	0xfb, 0x90, 0x74, 0x78, 0x42, 0xa3, 0xeb, 0x7b, 0xc8, 0x86, 0x22, 0xff, 0xd5, 0xa0, 0x22, 0x51,
	0xfc, 0x09, 0x2a, 0x2b, 0x4a, 0xc2, 0xfc, 0x9d, 0x76, 0x86, 0x5f, 0x3d, 0x7a, 0x82, 0x6f, 0x04,
	0x6d, 0xe7, 0x26, 0xb4, 0xa0, 0xce, 0x62, 0x92, 0xf8, 0x9c, 0x25, 0xf2, 0x59, 0x77, 0x86, 0x5f,
	0x3f, 0xba, 0xc0, 0x6c, 0x6b, 0xb0, 0x77, 0x56, 0xfc, 0x46, 0x35, 0x29, 0xce, 0xb4, 0x39, 0x3c,
	0x1a, 0xe4, 0x5f, 0x98, 0x81, 0xfa, 0xc2, 0x0c, 0x46, 0xd1, 0x46, 0xb5, 0xfe, 0x8f, 0x06, 0x65,
	0x71, 0x1f, 0xf8, 0xfa, 0xb0, 0xf3, 0xe7, 0x8f, 0x5d, 0xe0, 0x41, 0xe3, 0x67, 0xd0, 0x08, 0x68,
	0x42, 0x96, 0x5c, 0x3c, 0x89, 0x8f, 0xe8, 0x5c, 0x2c, 0x30, 0x56, 0x06, 0x7b, 0xef, 0x35, 0x2f,
	0x01, 0xf6, 0xc7, 0x82, 0x6d, 0x68, 0xcc, 0xbd, 0xab, 0xd9, 0x6c, 0xea, 0x4d, 0xc6, 0x7a, 0x01,
	0x5b, 0x50, 0x9f, 0x7b, 0xe7, 0xd6, 0xe4, 0xec, 0xdc, 0xd5, 0x35, 0xec, 0x42, 0x73, 0xee, 0xb9,
	0x93, 0x0b, 0xcb, 0x71, 0x47, 0x17, 0x57, 0x7a, 0x11, 0x11, 0x3a, 0x73, 0xef, 0xc2, 0x1a, 0x4f,
	0x46, 0x97, 0xde, 0xc8, 0xf5, 0x5c, 0x47, 0x2f, 0x99, 0xdf, 0x43, 0x63, 0xd7, 0xac, 0xf0, 0x3b,
	0xca, 0x5f, 0x40, 0x80, 0xaa, 0xe3, 0x39, 0x93, 0xdf, 0x2d, 0x5d, 0xc3, 0x0e, 0x80, 0xe3, 0xb9,
	0xbf, 0x79, 0xa7, 0xb3, 0x5f, 0x2e, 0x5d, 0xbd, 0x68, 0x9a, 0xd0, 0x3e, 0x68, 0x11, 0x6b, 0x50,
	0x1a, 0x39, 0xa7, 0x7a, 0x01, 0xeb, 0x50, 0x1e, 0x5b, 0xce, 0xa9, 0xae, 0x99, 0x23, 0x68, 0x1f,
	0x5c, 0x00, 0x56, 0xa1, 0x68, 0xcd, 0xf5, 0x82, 0xf8, 0xbd, 0x14, 0x8b, 0xd6, 0xa0, 0x74, 0xe6,
	0x5a, 0x7a, 0x51, 0x0c, 0xa6, 0xae, 0xa5, 0x97, 0xc4, 0xcc, 0xd4, 0xd2, 0xcb, 0xe2, 0xf7, 0xcc,
	0xd2, 0x2b, 0xe6, 0x1f, 0xf0, 0xe9, 0x07, 0xc7, 0x93, 0xc6, 0x2c, 0x4a, 0x09, 0x7e, 0x07, 0x0d,
	0xa6, 0xc4, 0xed, 0x63, 0xec, 0x6e, 0xcf, 0x53, 0xc1, 0xf6, 0x9e, 0x10, 0xa9, 0xe4, 0xf2, 0x59,
	0x14, 0xf3, 0xa4, 0xcb, 0x62, 0xe8, 0x41, 0x57, 0xc1, 0x0e, 0x49, 0xee, 0xe8, 0x92, 0xe0, 0x14,
	0xda, 0x07, 0x1b, 0xe2, 0xe7, 0x0f, 0xdc, 0xd2, 0xc9, 0xd3, 0xfb, 0x27, 0xf3, 0x1e, 0xcd, 0xc2,
	0xa2, 0x2a, 0x03, 0xf5, 0xea, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x91, 0x50, 0x51, 0xda,
	0x06, 0x00, 0x00,
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
