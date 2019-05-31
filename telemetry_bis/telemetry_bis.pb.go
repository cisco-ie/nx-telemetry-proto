// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry_bis.proto

package telemetry_bis

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Telemetry struct {
	// Types that are valid to be assigned to NodeId:
	//	*Telemetry_NodeIdStr
	NodeId isTelemetry_NodeId `protobuf_oneof:"node_id"`
	// Types that are valid to be assigned to Subscription:
	//	*Telemetry_SubscriptionIdStr
	Subscription isTelemetry_Subscription `protobuf_oneof:"subscription"`
	// string   sensor_path = 5;               // not produced
	EncodingPath string `protobuf:"bytes,6,opt,name=encoding_path,json=encodingPath,proto3" json:"encoding_path,omitempty"`
	// string   model_version = 7;             // not produced
	CollectionId         uint64             `protobuf:"varint,8,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	CollectionStartTime  uint64             `protobuf:"varint,9,opt,name=collection_start_time,json=collectionStartTime,proto3" json:"collection_start_time,omitempty"`
	MsgTimestamp         uint64             `protobuf:"varint,10,opt,name=msg_timestamp,json=msgTimestamp,proto3" json:"msg_timestamp,omitempty"`
	DataGpbkv            []*TelemetryField  `protobuf:"bytes,11,rep,name=data_gpbkv,json=dataGpbkv,proto3" json:"data_gpbkv,omitempty"`
	DataGpb              *TelemetryGPBTable `protobuf:"bytes,12,opt,name=data_gpb,json=dataGpb,proto3" json:"data_gpb,omitempty"`
	CollectionEndTime    uint64             `protobuf:"varint,13,opt,name=collection_end_time,json=collectionEndTime,proto3" json:"collection_end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Telemetry) Reset()         { *m = Telemetry{} }
func (m *Telemetry) String() string { return proto.CompactTextString(m) }
func (*Telemetry) ProtoMessage()    {}
func (*Telemetry) Descriptor() ([]byte, []int) {
	return fileDescriptor_088ca57c1a32d4b7, []int{0}
}

func (m *Telemetry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry.Unmarshal(m, b)
}
func (m *Telemetry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry.Marshal(b, m, deterministic)
}
func (m *Telemetry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry.Merge(m, src)
}
func (m *Telemetry) XXX_Size() int {
	return xxx_messageInfo_Telemetry.Size(m)
}
func (m *Telemetry) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry proto.InternalMessageInfo

type isTelemetry_NodeId interface {
	isTelemetry_NodeId()
}

type Telemetry_NodeIdStr struct {
	NodeIdStr string `protobuf:"bytes,1,opt,name=node_id_str,json=nodeIdStr,proto3,oneof"`
}

func (*Telemetry_NodeIdStr) isTelemetry_NodeId() {}

func (m *Telemetry) GetNodeId() isTelemetry_NodeId {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Telemetry) GetNodeIdStr() string {
	if x, ok := m.GetNodeId().(*Telemetry_NodeIdStr); ok {
		return x.NodeIdStr
	}
	return ""
}

type isTelemetry_Subscription interface {
	isTelemetry_Subscription()
}

type Telemetry_SubscriptionIdStr struct {
	SubscriptionIdStr string `protobuf:"bytes,3,opt,name=subscription_id_str,json=subscriptionIdStr,proto3,oneof"`
}

func (*Telemetry_SubscriptionIdStr) isTelemetry_Subscription() {}

func (m *Telemetry) GetSubscription() isTelemetry_Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *Telemetry) GetSubscriptionIdStr() string {
	if x, ok := m.GetSubscription().(*Telemetry_SubscriptionIdStr); ok {
		return x.SubscriptionIdStr
	}
	return ""
}

func (m *Telemetry) GetEncodingPath() string {
	if m != nil {
		return m.EncodingPath
	}
	return ""
}

func (m *Telemetry) GetCollectionId() uint64 {
	if m != nil {
		return m.CollectionId
	}
	return 0
}

func (m *Telemetry) GetCollectionStartTime() uint64 {
	if m != nil {
		return m.CollectionStartTime
	}
	return 0
}

func (m *Telemetry) GetMsgTimestamp() uint64 {
	if m != nil {
		return m.MsgTimestamp
	}
	return 0
}

func (m *Telemetry) GetDataGpbkv() []*TelemetryField {
	if m != nil {
		return m.DataGpbkv
	}
	return nil
}

func (m *Telemetry) GetDataGpb() *TelemetryGPBTable {
	if m != nil {
		return m.DataGpb
	}
	return nil
}

func (m *Telemetry) GetCollectionEndTime() uint64 {
	if m != nil {
		return m.CollectionEndTime
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Telemetry) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Telemetry_NodeIdStr)(nil),
		(*Telemetry_SubscriptionIdStr)(nil),
	}
}

type TelemetryField struct {
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ValueByType:
	//	*TelemetryField_BytesValue
	//	*TelemetryField_StringValue
	//	*TelemetryField_BoolValue
	//	*TelemetryField_Uint32Value
	//	*TelemetryField_Uint64Value
	//	*TelemetryField_Sint32Value
	//	*TelemetryField_Sint64Value
	//	*TelemetryField_DoubleValue
	//	*TelemetryField_FloatValue
	ValueByType          isTelemetryField_ValueByType `protobuf_oneof:"value_by_type"`
	Fields               []*TelemetryField            `protobuf:"bytes,15,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *TelemetryField) Reset()         { *m = TelemetryField{} }
func (m *TelemetryField) String() string { return proto.CompactTextString(m) }
func (*TelemetryField) ProtoMessage()    {}
func (*TelemetryField) Descriptor() ([]byte, []int) {
	return fileDescriptor_088ca57c1a32d4b7, []int{1}
}

func (m *TelemetryField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryField.Unmarshal(m, b)
}
func (m *TelemetryField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryField.Marshal(b, m, deterministic)
}
func (m *TelemetryField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryField.Merge(m, src)
}
func (m *TelemetryField) XXX_Size() int {
	return xxx_messageInfo_TelemetryField.Size(m)
}
func (m *TelemetryField) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryField.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryField proto.InternalMessageInfo

func (m *TelemetryField) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TelemetryField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTelemetryField_ValueByType interface {
	isTelemetryField_ValueByType()
}

type TelemetryField_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,4,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type TelemetryField_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type TelemetryField_BoolValue struct {
	BoolValue bool `protobuf:"varint,6,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type TelemetryField_Uint32Value struct {
	Uint32Value uint32 `protobuf:"varint,7,opt,name=uint32_value,json=uint32Value,proto3,oneof"`
}

type TelemetryField_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,8,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type TelemetryField_Sint32Value struct {
	Sint32Value int32 `protobuf:"zigzag32,9,opt,name=sint32_value,json=sint32Value,proto3,oneof"`
}

type TelemetryField_Sint64Value struct {
	Sint64Value int64 `protobuf:"zigzag64,10,opt,name=sint64_value,json=sint64Value,proto3,oneof"`
}

type TelemetryField_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,11,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type TelemetryField_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,12,opt,name=float_value,json=floatValue,proto3,oneof"`
}

func (*TelemetryField_BytesValue) isTelemetryField_ValueByType() {}

func (*TelemetryField_StringValue) isTelemetryField_ValueByType() {}

func (*TelemetryField_BoolValue) isTelemetryField_ValueByType() {}

func (*TelemetryField_Uint32Value) isTelemetryField_ValueByType() {}

func (*TelemetryField_Uint64Value) isTelemetryField_ValueByType() {}

func (*TelemetryField_Sint32Value) isTelemetryField_ValueByType() {}

func (*TelemetryField_Sint64Value) isTelemetryField_ValueByType() {}

func (*TelemetryField_DoubleValue) isTelemetryField_ValueByType() {}

func (*TelemetryField_FloatValue) isTelemetryField_ValueByType() {}

func (m *TelemetryField) GetValueByType() isTelemetryField_ValueByType {
	if m != nil {
		return m.ValueByType
	}
	return nil
}

func (m *TelemetryField) GetBytesValue() []byte {
	if x, ok := m.GetValueByType().(*TelemetryField_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *TelemetryField) GetStringValue() string {
	if x, ok := m.GetValueByType().(*TelemetryField_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *TelemetryField) GetBoolValue() bool {
	if x, ok := m.GetValueByType().(*TelemetryField_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *TelemetryField) GetUint32Value() uint32 {
	if x, ok := m.GetValueByType().(*TelemetryField_Uint32Value); ok {
		return x.Uint32Value
	}
	return 0
}

func (m *TelemetryField) GetUint64Value() uint64 {
	if x, ok := m.GetValueByType().(*TelemetryField_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *TelemetryField) GetSint32Value() int32 {
	if x, ok := m.GetValueByType().(*TelemetryField_Sint32Value); ok {
		return x.Sint32Value
	}
	return 0
}

func (m *TelemetryField) GetSint64Value() int64 {
	if x, ok := m.GetValueByType().(*TelemetryField_Sint64Value); ok {
		return x.Sint64Value
	}
	return 0
}

func (m *TelemetryField) GetDoubleValue() float64 {
	if x, ok := m.GetValueByType().(*TelemetryField_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *TelemetryField) GetFloatValue() float32 {
	if x, ok := m.GetValueByType().(*TelemetryField_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *TelemetryField) GetFields() []*TelemetryField {
	if m != nil {
		return m.Fields
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TelemetryField) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TelemetryField_BytesValue)(nil),
		(*TelemetryField_StringValue)(nil),
		(*TelemetryField_BoolValue)(nil),
		(*TelemetryField_Uint32Value)(nil),
		(*TelemetryField_Uint64Value)(nil),
		(*TelemetryField_Sint32Value)(nil),
		(*TelemetryField_Sint64Value)(nil),
		(*TelemetryField_DoubleValue)(nil),
		(*TelemetryField_FloatValue)(nil),
	}
}

type TelemetryGPBTable struct {
	Row                  []*TelemetryRowGPB `protobuf:"bytes,1,rep,name=row,proto3" json:"row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TelemetryGPBTable) Reset()         { *m = TelemetryGPBTable{} }
func (m *TelemetryGPBTable) String() string { return proto.CompactTextString(m) }
func (*TelemetryGPBTable) ProtoMessage()    {}
func (*TelemetryGPBTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_088ca57c1a32d4b7, []int{2}
}

func (m *TelemetryGPBTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryGPBTable.Unmarshal(m, b)
}
func (m *TelemetryGPBTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryGPBTable.Marshal(b, m, deterministic)
}
func (m *TelemetryGPBTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryGPBTable.Merge(m, src)
}
func (m *TelemetryGPBTable) XXX_Size() int {
	return xxx_messageInfo_TelemetryGPBTable.Size(m)
}
func (m *TelemetryGPBTable) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryGPBTable.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryGPBTable proto.InternalMessageInfo

func (m *TelemetryGPBTable) GetRow() []*TelemetryRowGPB {
	if m != nil {
		return m.Row
	}
	return nil
}

type TelemetryRowGPB struct {
	Timestamp            uint64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Keys                 []byte   `protobuf:"bytes,10,opt,name=keys,proto3" json:"keys,omitempty"`
	Content              []byte   `protobuf:"bytes,11,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetryRowGPB) Reset()         { *m = TelemetryRowGPB{} }
func (m *TelemetryRowGPB) String() string { return proto.CompactTextString(m) }
func (*TelemetryRowGPB) ProtoMessage()    {}
func (*TelemetryRowGPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_088ca57c1a32d4b7, []int{3}
}

func (m *TelemetryRowGPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetryRowGPB.Unmarshal(m, b)
}
func (m *TelemetryRowGPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetryRowGPB.Marshal(b, m, deterministic)
}
func (m *TelemetryRowGPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetryRowGPB.Merge(m, src)
}
func (m *TelemetryRowGPB) XXX_Size() int {
	return xxx_messageInfo_TelemetryRowGPB.Size(m)
}
func (m *TelemetryRowGPB) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetryRowGPB.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetryRowGPB proto.InternalMessageInfo

func (m *TelemetryRowGPB) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TelemetryRowGPB) GetKeys() []byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *TelemetryRowGPB) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Telemetry)(nil), "Telemetry")
	proto.RegisterType((*TelemetryField)(nil), "TelemetryField")
	proto.RegisterType((*TelemetryGPBTable)(nil), "TelemetryGPBTable")
	proto.RegisterType((*TelemetryRowGPB)(nil), "TelemetryRowGPB")
}

func init() { proto.RegisterFile("telemetry_bis.proto", fileDescriptor_088ca57c1a32d4b7) }

var fileDescriptor_088ca57c1a32d4b7 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0xb3, 0xb1, 0x6b, 0x5b, 0x23, 0x39, 0xae, 0xd7, 0x14, 0xf6, 0x50, 0xa8, 0xe2, 0x1c,
	0xaa, 0x4b, 0x4d, 0x71, 0x42, 0x7a, 0x17, 0xb4, 0x4e, 0x6e, 0x66, 0x63, 0x7a, 0x28, 0x14, 0x21,
	0x59, 0x1b, 0x47, 0x44, 0xd2, 0x0a, 0xed, 0x3a, 0xc1, 0xef, 0xd4, 0x97, 0xe8, 0x9b, 0x95, 0x59,
	0x49, 0x91, 0xdc, 0x16, 0x72, 0x93, 0xfe, 0xf9, 0xf6, 0x67, 0x76, 0xfe, 0x61, 0x61, 0xa6, 0x45,
	0x2a, 0x32, 0xa1, 0xcb, 0x43, 0x10, 0x25, 0x6a, 0x51, 0x94, 0x52, 0xcb, 0xf9, 0xaf, 0x1e, 0x58,
	0x9b, 0x46, 0xa7, 0x2e, 0xd8, 0xb9, 0x8c, 0x45, 0x90, 0xc4, 0x81, 0xd2, 0x25, 0x23, 0x2e, 0xf1,
	0xac, 0x9b, 0x13, 0x6e, 0xa1, 0x78, 0x1b, 0xdf, 0xe9, 0x92, 0x7e, 0x86, 0x99, 0xda, 0x47, 0x6a,
	0x5b, 0x26, 0x85, 0x4e, 0x64, 0xde, 0x90, 0x3d, 0x43, 0x12, 0x3e, 0xed, 0x16, 0xab, 0x13, 0x17,
	0x30, 0x16, 0xf9, 0x56, 0xc6, 0x49, 0xbe, 0x0b, 0x8a, 0x50, 0x3f, 0xb0, 0x01, 0xb2, 0xdc, 0x69,
	0xc4, 0x75, 0xa8, 0x1f, 0x10, 0xda, 0xca, 0x34, 0x15, 0xdb, 0xda, 0x94, 0x8d, 0x5c, 0xe2, 0xf5,
	0xb9, 0xd3, 0x8a, 0xb7, 0x31, 0x5d, 0xc2, 0xbb, 0x0e, 0xa4, 0x74, 0x58, 0xea, 0x40, 0x27, 0x99,
	0x60, 0x96, 0x81, 0x67, 0x6d, 0xf1, 0x0e, 0x6b, 0x9b, 0x24, 0x13, 0x68, 0x9c, 0xa9, 0x9d, 0xc1,
	0x94, 0x0e, 0xb3, 0x82, 0x41, 0x65, 0x9c, 0xa9, 0xdd, 0xa6, 0xd1, 0xe8, 0x02, 0x20, 0x0e, 0x75,
	0x18, 0xec, 0x8a, 0xe8, 0xf1, 0x89, 0xd9, 0x6e, 0xcf, 0xb3, 0x97, 0x93, 0xc5, 0xcb, 0x58, 0xbe,
	0x25, 0x22, 0x8d, 0xb9, 0x85, 0xc8, 0x0a, 0x09, 0xfa, 0x09, 0x46, 0x0d, 0xcf, 0x1c, 0x97, 0x78,
	0xf6, 0x92, 0xb6, 0xf4, 0x6a, 0xed, 0x6f, 0xc2, 0x28, 0x15, 0x7c, 0x58, 0x1f, 0xa0, 0x0b, 0xe8,
	0xb4, 0x16, 0x88, 0x3c, 0xae, 0xba, 0x1e, 0x9b, 0x4e, 0xa6, 0x6d, 0xe9, 0x6b, 0x1e, 0x63, 0x4f,
	0xbe, 0x05, 0xc3, 0x3a, 0x05, 0xff, 0x0c, 0x9c, 0xee, 0x44, 0xe7, 0xbf, 0x7b, 0x70, 0x76, 0xdc,
	0x17, 0x7d, 0x0f, 0x56, 0x7b, 0x3b, 0x62, 0x3c, 0x5b, 0x81, 0x52, 0xe8, 0xe7, 0x61, 0x26, 0xd8,
	0xa9, 0x19, 0xba, 0xf9, 0xa6, 0xe7, 0x60, 0x47, 0x07, 0x2d, 0x54, 0xf0, 0x14, 0xa6, 0x7b, 0xc1,
	0xfa, 0x2e, 0xf1, 0x9c, 0x9b, 0x13, 0x0e, 0x46, 0xfc, 0x8e, 0x1a, 0xbd, 0x00, 0x47, 0xe9, 0x12,
	0x23, 0xab, 0x98, 0x37, 0xf5, 0x26, 0xd8, 0x95, 0x5a, 0x41, 0x1f, 0x00, 0x22, 0x29, 0xd3, 0x1a,
	0xc1, 0x58, 0x47, 0xb8, 0x2c, 0xa8, 0xbd, 0xb8, 0xec, 0x93, 0x5c, 0x5f, 0x2e, 0x6b, 0x64, 0xe8,
	0x12, 0x6f, 0x8c, 0x2e, 0x95, 0x7a, 0x04, 0x5d, 0x5f, 0xd5, 0x90, 0x49, 0xbe, 0x81, 0xae, 0xaf,
	0xda, 0x7e, 0xba, 0x4e, 0x98, 0xf8, 0xd4, 0xf4, 0x73, 0xec, 0xa4, 0xba, 0x4e, 0x18, 0x35, 0x6d,
	0xa0, 0x8e, 0x53, 0x2c, 0xf7, 0x51, 0x2a, 0x6a, 0xc8, 0x76, 0x89, 0x47, 0x10, 0xaa, 0xd4, 0x0a,
	0x3a, 0x07, 0xfb, 0x3e, 0x95, 0xa1, 0xae, 0x19, 0xcc, 0xf8, 0x14, 0x27, 0x64, 0xc4, 0x0a, 0xf9,
	0x08, 0x83, 0x7b, 0x9c, 0xbf, 0x62, 0x93, 0xff, 0xef, 0x4b, 0x5d, 0xf6, 0x27, 0x30, 0x36, 0x2e,
	0x41, 0x74, 0x08, 0xf4, 0xa1, 0x10, 0xf3, 0x2f, 0x30, 0xfd, 0x67, 0x59, 0xe8, 0x1c, 0x7a, 0xa5,
	0x7c, 0x66, 0xc4, 0x78, 0xbd, 0x6d, 0xbd, 0xb8, 0x7c, 0x5e, 0xad, 0x7d, 0x8e, 0xc5, 0xf9, 0x4f,
	0x98, 0xfc, 0xa5, 0xbf, 0x1e, 0xfe, 0xa3, 0x38, 0x28, 0x33, 0x08, 0x87, 0x9b, 0x6f, 0xca, 0x60,
	0xb8, 0x95, 0xb9, 0x16, 0xb9, 0x36, 0x57, 0x77, 0x78, 0xf3, 0xeb, 0x4f, 0x7e, 0x8c, 0x8f, 0x5e,
	0x88, 0x68, 0x60, 0x9e, 0x88, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x64, 0xd1, 0x66,
	0x39, 0x04, 0x00, 0x00,
}
