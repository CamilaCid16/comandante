// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: comandante.proto

package __

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

// El comando a ejecutar
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector   string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base     string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	ClientId string `protobuf:"bytes,3,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comandante_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_comandante_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_comandante_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *Command) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *Command) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

// La respuesta del servidor
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acknowledgement string  `protobuf:"bytes,1,opt,name=acknowledgement,proto3" json:"acknowledgement,omitempty"`
	FulcrumServer   string  `protobuf:"bytes,2,opt,name=FulcrumServer,proto3" json:"FulcrumServer,omitempty"`
	VectorClock     []int32 `protobuf:"varint,3,rep,packed,name=VectorClock,proto3" json:"VectorClock,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comandante_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_comandante_proto_msgTypes[1]
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
	return file_comandante_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetAcknowledgement() string {
	if x != nil {
		return x.Acknowledgement
	}
	return ""
}

func (x *Response) GetFulcrumServer() string {
	if x != nil {
		return x.FulcrumServer
	}
	return ""
}

func (x *Response) GetVectorClock() []int32 {
	if x != nil {
		return x.VectorClock
	}
	return nil
}

var File_comandante_proto protoreflect.FileDescriptor

var file_comandante_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x51, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x7c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x32, 0x43,
	0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x45, 0x6e, 0x65, 0x6d, 0x69, 0x67, 0x6f, 0x73, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x13,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_comandante_proto_rawDescOnce sync.Once
	file_comandante_proto_rawDescData = file_comandante_proto_rawDesc
)

func file_comandante_proto_rawDescGZIP() []byte {
	file_comandante_proto_rawDescOnce.Do(func() {
		file_comandante_proto_rawDescData = protoimpl.X.CompressGZIP(file_comandante_proto_rawDescData)
	})
	return file_comandante_proto_rawDescData
}

var file_comandante_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_comandante_proto_goTypes = []interface{}{
	(*Command)(nil),  // 0: commander.Command
	(*Response)(nil), // 1: commander.Response
}
var file_comandante_proto_depIdxs = []int32{
	0, // 0: commander.Commander.GetEnemigos:input_type -> commander.Command
	1, // 1: commander.Commander.GetEnemigos:output_type -> commander.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comandante_proto_init() }
func file_comandante_proto_init() {
	if File_comandante_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comandante_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_comandante_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_comandante_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comandante_proto_goTypes,
		DependencyIndexes: file_comandante_proto_depIdxs,
		MessageInfos:      file_comandante_proto_msgTypes,
	}.Build()
	File_comandante_proto = out.File
	file_comandante_proto_rawDesc = nil
	file_comandante_proto_goTypes = nil
	file_comandante_proto_depIdxs = nil
}
