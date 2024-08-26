// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: InverseMatrix.proto

package Inverse

import (
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

type RealRows struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []float64 `protobuf:"fixed64,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *RealRows) Reset() {
	*x = RealRows{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InverseMatrix_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealRows) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealRows) ProtoMessage() {}

func (x *RealRows) ProtoReflect() protoreflect.Message {
	mi := &file_InverseMatrix_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealRows.ProtoReflect.Descriptor instead.
func (*RealRows) Descriptor() ([]byte, []int) {
	return file_InverseMatrix_proto_rawDescGZIP(), []int{0}
}

func (x *RealRows) GetVal() []float64 {
	if x != nil {
		return x.Val
	}
	return nil
}

type InverseMatrixRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*RealRows `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *InverseMatrixRequest) Reset() {
	*x = InverseMatrixRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InverseMatrix_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InverseMatrixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InverseMatrixRequest) ProtoMessage() {}

func (x *InverseMatrixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_InverseMatrix_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InverseMatrixRequest.ProtoReflect.Descriptor instead.
func (*InverseMatrixRequest) Descriptor() ([]byte, []int) {
	return file_InverseMatrix_proto_rawDescGZIP(), []int{1}
}

func (x *InverseMatrixRequest) GetRows() []*RealRows {
	if x != nil {
		return x.Rows
	}
	return nil
}

type InverseMatrixResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*RealRows `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *InverseMatrixResponse) Reset() {
	*x = InverseMatrixResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InverseMatrix_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InverseMatrixResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InverseMatrixResponse) ProtoMessage() {}

func (x *InverseMatrixResponse) ProtoReflect() protoreflect.Message {
	mi := &file_InverseMatrix_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InverseMatrixResponse.ProtoReflect.Descriptor instead.
func (*InverseMatrixResponse) Descriptor() ([]byte, []int) {
	return file_InverseMatrix_proto_rawDescGZIP(), []int{2}
}

func (x *InverseMatrixResponse) GetRows() []*RealRows {
	if x != nil {
		return x.Rows
	}
	return nil
}

var File_InverseMatrix_proto protoreflect.FileDescriptor

var file_InverseMatrix_proto_rawDesc = []byte{
	0x0a, 0x13, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x6c, 0x52, 0x6f, 0x77,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x35, 0x0a, 0x14, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x61, 0x6c,
	0x52, 0x6f, 0x77, 0x73, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x36, 0x0a, 0x15, 0x49, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x04, 0x72, 0x6f,
	0x77, 0x73, 0x32, 0x56, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x12, 0x45, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x12, 0x15, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x49,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InverseMatrix_proto_rawDescOnce sync.Once
	file_InverseMatrix_proto_rawDescData = file_InverseMatrix_proto_rawDesc
)

func file_InverseMatrix_proto_rawDescGZIP() []byte {
	file_InverseMatrix_proto_rawDescOnce.Do(func() {
		file_InverseMatrix_proto_rawDescData = protoimpl.X.CompressGZIP(file_InverseMatrix_proto_rawDescData)
	})
	return file_InverseMatrix_proto_rawDescData
}

var file_InverseMatrix_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_InverseMatrix_proto_goTypes = []any{
	(*RealRows)(nil),              // 0: RealRows
	(*InverseMatrixRequest)(nil),  // 1: InverseMatrixRequest
	(*InverseMatrixResponse)(nil), // 2: InverseMatrixResponse
}
var file_InverseMatrix_proto_depIdxs = []int32{
	0, // 0: InverseMatrixRequest.rows:type_name -> RealRows
	0, // 1: InverseMatrixResponse.rows:type_name -> RealRows
	1, // 2: InverseMatrix.GetInverseMatrix:input_type -> InverseMatrixRequest
	2, // 3: InverseMatrix.GetInverseMatrix:output_type -> InverseMatrixResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_InverseMatrix_proto_init() }
func file_InverseMatrix_proto_init() {
	if File_InverseMatrix_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_InverseMatrix_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RealRows); i {
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
		file_InverseMatrix_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InverseMatrixRequest); i {
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
		file_InverseMatrix_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InverseMatrixResponse); i {
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
			RawDescriptor: file_InverseMatrix_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_InverseMatrix_proto_goTypes,
		DependencyIndexes: file_InverseMatrix_proto_depIdxs,
		MessageInfos:      file_InverseMatrix_proto_msgTypes,
	}.Build()
	File_InverseMatrix_proto = out.File
	file_InverseMatrix_proto_rawDesc = nil
	file_InverseMatrix_proto_goTypes = nil
	file_InverseMatrix_proto_depIdxs = nil
}