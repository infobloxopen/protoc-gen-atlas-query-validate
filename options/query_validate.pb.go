// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: options/query_validate.proto

package options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryValidate_FilterOperator int32

const (
	QueryValidate_EQ    QueryValidate_FilterOperator = 0
	QueryValidate_MATCH QueryValidate_FilterOperator = 1
	QueryValidate_GT    QueryValidate_FilterOperator = 2
	QueryValidate_GE    QueryValidate_FilterOperator = 3
	QueryValidate_LT    QueryValidate_FilterOperator = 4
	QueryValidate_LE    QueryValidate_FilterOperator = 5
	QueryValidate_ALL   QueryValidate_FilterOperator = 6
	QueryValidate_IEQ   QueryValidate_FilterOperator = 7
	QueryValidate_IN    QueryValidate_FilterOperator = 8
)

// Enum value maps for QueryValidate_FilterOperator.
var (
	QueryValidate_FilterOperator_name = map[int32]string{
		0: "EQ",
		1: "MATCH",
		2: "GT",
		3: "GE",
		4: "LT",
		5: "LE",
		6: "ALL",
		7: "IEQ",
		8: "IN",
	}
	QueryValidate_FilterOperator_value = map[string]int32{
		"EQ":    0,
		"MATCH": 1,
		"GT":    2,
		"GE":    3,
		"LT":    4,
		"LE":    5,
		"ALL":   6,
		"IEQ":   7,
		"IN":    8,
	}
)

func (x QueryValidate_FilterOperator) Enum() *QueryValidate_FilterOperator {
	p := new(QueryValidate_FilterOperator)
	*p = x
	return p
}

func (x QueryValidate_FilterOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryValidate_FilterOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_options_query_validate_proto_enumTypes[0].Descriptor()
}

func (QueryValidate_FilterOperator) Type() protoreflect.EnumType {
	return &file_options_query_validate_proto_enumTypes[0]
}

func (x QueryValidate_FilterOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryValidate_FilterOperator.Descriptor instead.
func (QueryValidate_FilterOperator) EnumDescriptor() ([]byte, []int) {
	return file_options_query_validate_proto_rawDescGZIP(), []int{0, 0}
}

type QueryValidate_ValueType int32

const (
	QueryValidate_DEFAULT QueryValidate_ValueType = 0
	QueryValidate_STRING  QueryValidate_ValueType = 1
	QueryValidate_NUMBER  QueryValidate_ValueType = 2
	QueryValidate_BOOL    QueryValidate_ValueType = 3
)

// Enum value maps for QueryValidate_ValueType.
var (
	QueryValidate_ValueType_name = map[int32]string{
		0: "DEFAULT",
		1: "STRING",
		2: "NUMBER",
		3: "BOOL",
	}
	QueryValidate_ValueType_value = map[string]int32{
		"DEFAULT": 0,
		"STRING":  1,
		"NUMBER":  2,
		"BOOL":    3,
	}
)

func (x QueryValidate_ValueType) Enum() *QueryValidate_ValueType {
	p := new(QueryValidate_ValueType)
	*p = x
	return p
}

func (x QueryValidate_ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryValidate_ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_options_query_validate_proto_enumTypes[1].Descriptor()
}

func (QueryValidate_ValueType) Type() protoreflect.EnumType {
	return &file_options_query_validate_proto_enumTypes[1]
}

func (x QueryValidate_ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryValidate_ValueType.Descriptor instead.
func (QueryValidate_ValueType) EnumDescriptor() ([]byte, []int) {
	return file_options_query_validate_proto_rawDescGZIP(), []int{0, 1}
}

type QueryValidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filtering          *QueryValidate_Filtering      `protobuf:"bytes,1,opt,name=filtering,proto3" json:"filtering,omitempty"`
	Sorting            *QueryValidate_Sorting        `protobuf:"bytes,2,opt,name=sorting,proto3" json:"sorting,omitempty"`
	FieldSelection     *QueryValidate_FieldSelection `protobuf:"bytes,3,opt,name=field_selection,json=fieldSelection,proto3" json:"field_selection,omitempty"`
	ValueType          QueryValidate_ValueType       `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=atlas.query.QueryValidate_ValueType" json:"value_type,omitempty"`
	ValueTypeUrl       string                        `protobuf:"bytes,5,opt,name=value_type_url,json=valueTypeUrl,proto3" json:"value_type_url,omitempty"`
	EnableNestedFields bool                          `protobuf:"varint,6,opt,name=enable_nested_fields,json=enableNestedFields,proto3" json:"enable_nested_fields,omitempty"`
	NestedFields       []string                      `protobuf:"bytes,7,rep,name=nested_fields,json=nestedFields,proto3" json:"nested_fields,omitempty"`
}

func (x *QueryValidate) Reset() {
	*x = QueryValidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_query_validate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidate) ProtoMessage() {}

func (x *QueryValidate) ProtoReflect() protoreflect.Message {
	mi := &file_options_query_validate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryValidate.ProtoReflect.Descriptor instead.
func (*QueryValidate) Descriptor() ([]byte, []int) {
	return file_options_query_validate_proto_rawDescGZIP(), []int{0}
}

func (x *QueryValidate) GetFiltering() *QueryValidate_Filtering {
	if x != nil {
		return x.Filtering
	}
	return nil
}

func (x *QueryValidate) GetSorting() *QueryValidate_Sorting {
	if x != nil {
		return x.Sorting
	}
	return nil
}

func (x *QueryValidate) GetFieldSelection() *QueryValidate_FieldSelection {
	if x != nil {
		return x.FieldSelection
	}
	return nil
}

func (x *QueryValidate) GetValueType() QueryValidate_ValueType {
	if x != nil {
		return x.ValueType
	}
	return QueryValidate_DEFAULT
}

func (x *QueryValidate) GetValueTypeUrl() string {
	if x != nil {
		return x.ValueTypeUrl
	}
	return ""
}

func (x *QueryValidate) GetEnableNestedFields() bool {
	if x != nil {
		return x.EnableNestedFields
	}
	return false
}

func (x *QueryValidate) GetNestedFields() []string {
	if x != nil {
		return x.NestedFields
	}
	return nil
}

type MessageQueryValidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validate              []*MessageQueryValidate_QueryValidateEntry `protobuf:"bytes,1,rep,name=validate,proto3" json:"validate,omitempty"`
	NestedFieldDepthLimit int32                                      `protobuf:"varint,2,opt,name=nested_field_depth_limit,json=nestedFieldDepthLimit,proto3" json:"nested_field_depth_limit,omitempty"`
	EnableNestedFields    bool                                       `protobuf:"varint,3,opt,name=enable_nested_fields,json=enableNestedFields,proto3" json:"enable_nested_fields,omitempty"`
}

func (x *MessageQueryValidate) Reset() {
	*x = MessageQueryValidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_query_validate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageQueryValidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageQueryValidate) ProtoMessage() {}

func (x *MessageQueryValidate) ProtoReflect() protoreflect.Message {
	mi := &file_options_query_validate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageQueryValidate.ProtoReflect.Descriptor instead.
func (*MessageQueryValidate) Descriptor() ([]byte, []int) {
	return file_options_query_validate_proto_rawDescGZIP(), []int{1}
}

func (x *MessageQueryValidate) GetValidate() []*MessageQueryValidate_QueryValidateEntry {
	if x != nil {
		return x.Validate
	}
	return nil
}

func (x *MessageQueryValidate) GetNestedFieldDepthLimit() int32 {
	if x != nil {
		return x.NestedFieldDepthLimit
	}
	return 0
}

func (x *MessageQueryValidate) GetEnableNestedFields() bool {
	if x != nil {
		return x.EnableNestedFields
	}
	return false
}

type QueryValidate_Filtering struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allow []QueryValidate_FilterOperator `protobuf:"varint,1,rep,packed,name=allow,proto3,enum=atlas.query.QueryValidate_FilterOperator" json:"allow,omitempty"`
	Deny  []QueryValidate_FilterOperator `protobuf:"varint,2,rep,packed,name=deny,proto3,enum=atlas.query.QueryValidate_FilterOperator" json:"deny,omitempty"`
}

func (x *QueryValidate_Filtering) Reset() {
	*x = QueryValidate_Filtering{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_query_validate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidate_Filtering) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidate_Filtering) ProtoMessage() {}

func (x *QueryValidate_Filtering) ProtoReflect() protoreflect.Message {
	mi := &file_options_query_validate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryValidate_Filtering.ProtoReflect.Descriptor instead.
func (*QueryValidate_Filtering) Descriptor() ([]byte, []int) {
	return file_options_query_validate_proto_rawDescGZIP(), []int{0, 0}
}

func (x *QueryValidate_Filtering) GetAllow() []QueryValidate_FilterOperator {
	if x != nil {
		return x.Allow
	}
	return nil
}

func (x *QueryValidate_Filtering) GetDeny() []QueryValidate_FilterOperator {
	if x != nil {
		return x.Deny
	}
	return nil
}

type QueryValidate_Sorting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disable bool `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (x *QueryValidate_Sorting) Reset() {
	*x = QueryValidate_Sorting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_query_validate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidate_Sorting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidate_Sorting) ProtoMessage() {}

func (x *QueryValidate_Sorting) ProtoReflect() protoreflect.Message {
	mi := &file_options_query_validate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryValidate_Sorting.ProtoReflect.Descriptor instead.
func (*QueryValidate_Sorting) Descriptor() ([]byte, []int) {
	return file_options_query_validate_proto_rawDescGZIP(), []int{0, 1}
}

func (x *QueryValidate_Sorting) GetDisable() bool {
	if x != nil {
		return x.Disable
	}
	return false
}

type QueryValidate_FieldSelection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disable bool `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (x *QueryValidate_FieldSelection) Reset() {
	*x = QueryValidate_FieldSelection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_query_validate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryValidate_FieldSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryValidate_FieldSelection) ProtoMessage() {}

func (x *QueryValidate_FieldSelection) ProtoReflect() protoreflect.Message {
	mi := &file_options_query_validate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryValidate_FieldSelection.ProtoReflect.Descriptor instead.
func (*QueryValidate_FieldSelection) Descriptor() ([]byte, []int) {
	return file_options_query_validate_proto_rawDescGZIP(), []int{0, 2}
}

func (x *QueryValidate_FieldSelection) GetDisable() bool {
	if x != nil {
		return x.Disable
	}
	return false
}

type MessageQueryValidate_QueryValidateEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *QueryValidate `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MessageQueryValidate_QueryValidateEntry) Reset() {
	*x = MessageQueryValidate_QueryValidateEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_options_query_validate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageQueryValidate_QueryValidateEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageQueryValidate_QueryValidateEntry) ProtoMessage() {}

func (x *MessageQueryValidate_QueryValidateEntry) ProtoReflect() protoreflect.Message {
	mi := &file_options_query_validate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageQueryValidate_QueryValidateEntry.ProtoReflect.Descriptor instead.
func (*MessageQueryValidate_QueryValidateEntry) Descriptor() ([]byte, []int) {
	return file_options_query_validate_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MessageQueryValidate_QueryValidateEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MessageQueryValidate_QueryValidateEntry) GetValue() *QueryValidate {
	if x != nil {
		return x.Value
	}
	return nil
}

var file_options_query_validate_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*QueryValidate)(nil),
		Field:         52121,
		Name:          "atlas.query.validate",
		Tag:           "bytes,52121,opt,name=validate",
		Filename:      "options/query_validate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageQueryValidate)(nil),
		Field:         52121,
		Name:          "atlas.query.message",
		Tag:           "bytes,52121,opt,name=message",
		Filename:      "options/query_validate.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional atlas.query.QueryValidate validate = 52121;
	E_Validate = &file_options_query_validate_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional atlas.query.MessageQueryValidate message = 52121;
	E_Message = &file_options_query_validate_proto_extTypes[1]
)

var File_options_query_validate_proto protoreflect.FileDescriptor

var file_options_query_validate_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x06,
	0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x42, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x52, 0x0a, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61, 0x74, 0x6c, 0x61,
	0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x8b, 0x01, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52,
	0x04, 0x64, 0x65, 0x6e, 0x79, 0x1a, 0x23, 0x0a, 0x07, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x2a, 0x0a, 0x0e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x5d, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x51, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x47,
	0x54, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x45, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x4c,
	0x54, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x41,
	0x4c, 0x4c, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x45, 0x51, 0x10, 0x07, 0x12, 0x06, 0x0a,
	0x02, 0x49, 0x4e, 0x10, 0x08, 0x22, 0x3a, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e,
	0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10,
	0x03, 0x22, 0xaf, 0x02, 0x0a, 0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x08, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x61,
	0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x18,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x65, 0x70,
	0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x70, 0x74, 0x68,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x12, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x5a, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x57, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x99,
	0x97, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x5e, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x99, 0x97, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x49, 0x5a, 0x47,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62,
	0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_options_query_validate_proto_rawDescOnce sync.Once
	file_options_query_validate_proto_rawDescData = file_options_query_validate_proto_rawDesc
)

func file_options_query_validate_proto_rawDescGZIP() []byte {
	file_options_query_validate_proto_rawDescOnce.Do(func() {
		file_options_query_validate_proto_rawDescData = protoimpl.X.CompressGZIP(file_options_query_validate_proto_rawDescData)
	})
	return file_options_query_validate_proto_rawDescData
}

var file_options_query_validate_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_options_query_validate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_options_query_validate_proto_goTypes = []interface{}{
	(QueryValidate_FilterOperator)(0),               // 0: atlas.query.QueryValidate.FilterOperator
	(QueryValidate_ValueType)(0),                    // 1: atlas.query.QueryValidate.ValueType
	(*QueryValidate)(nil),                           // 2: atlas.query.QueryValidate
	(*MessageQueryValidate)(nil),                    // 3: atlas.query.MessageQueryValidate
	(*QueryValidate_Filtering)(nil),                 // 4: atlas.query.QueryValidate.Filtering
	(*QueryValidate_Sorting)(nil),                   // 5: atlas.query.QueryValidate.Sorting
	(*QueryValidate_FieldSelection)(nil),            // 6: atlas.query.QueryValidate.FieldSelection
	(*MessageQueryValidate_QueryValidateEntry)(nil), // 7: atlas.query.MessageQueryValidate.QueryValidateEntry
	(*descriptorpb.FieldOptions)(nil),               // 8: google.protobuf.FieldOptions
	(*descriptorpb.MessageOptions)(nil),             // 9: google.protobuf.MessageOptions
}
var file_options_query_validate_proto_depIdxs = []int32{
	4,  // 0: atlas.query.QueryValidate.filtering:type_name -> atlas.query.QueryValidate.Filtering
	5,  // 1: atlas.query.QueryValidate.sorting:type_name -> atlas.query.QueryValidate.Sorting
	6,  // 2: atlas.query.QueryValidate.field_selection:type_name -> atlas.query.QueryValidate.FieldSelection
	1,  // 3: atlas.query.QueryValidate.value_type:type_name -> atlas.query.QueryValidate.ValueType
	7,  // 4: atlas.query.MessageQueryValidate.validate:type_name -> atlas.query.MessageQueryValidate.QueryValidateEntry
	0,  // 5: atlas.query.QueryValidate.Filtering.allow:type_name -> atlas.query.QueryValidate.FilterOperator
	0,  // 6: atlas.query.QueryValidate.Filtering.deny:type_name -> atlas.query.QueryValidate.FilterOperator
	2,  // 7: atlas.query.MessageQueryValidate.QueryValidateEntry.value:type_name -> atlas.query.QueryValidate
	8,  // 8: atlas.query.validate:extendee -> google.protobuf.FieldOptions
	9,  // 9: atlas.query.message:extendee -> google.protobuf.MessageOptions
	2,  // 10: atlas.query.validate:type_name -> atlas.query.QueryValidate
	3,  // 11: atlas.query.message:type_name -> atlas.query.MessageQueryValidate
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	10, // [10:12] is the sub-list for extension type_name
	8,  // [8:10] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_options_query_validate_proto_init() }
func file_options_query_validate_proto_init() {
	if File_options_query_validate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_options_query_validate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidate); i {
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
		file_options_query_validate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageQueryValidate); i {
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
		file_options_query_validate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidate_Filtering); i {
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
		file_options_query_validate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidate_Sorting); i {
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
		file_options_query_validate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryValidate_FieldSelection); i {
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
		file_options_query_validate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageQueryValidate_QueryValidateEntry); i {
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
			RawDescriptor: file_options_query_validate_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_options_query_validate_proto_goTypes,
		DependencyIndexes: file_options_query_validate_proto_depIdxs,
		EnumInfos:         file_options_query_validate_proto_enumTypes,
		MessageInfos:      file_options_query_validate_proto_msgTypes,
		ExtensionInfos:    file_options_query_validate_proto_extTypes,
	}.Build()
	File_options_query_validate_proto = out.File
	file_options_query_validate_proto_rawDesc = nil
	file_options_query_validate_proto_goTypes = nil
	file_options_query_validate_proto_depIdxs = nil
}
