// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: gateway/internal/test.proto

package gateway_test

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_internal_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_internal_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_gateway_internal_test_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_internal_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_internal_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_gateway_internal_test_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserWithPtr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PtrValue *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=ptr_value,json=ptrValue,proto3" json:"ptr_value,omitempty"`
}

func (x *UserWithPtr) Reset() {
	*x = UserWithPtr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_internal_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserWithPtr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserWithPtr) ProtoMessage() {}

func (x *UserWithPtr) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_internal_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserWithPtr.ProtoReflect.Descriptor instead.
func (*UserWithPtr) Descriptor() ([]byte, []int) {
	return file_gateway_internal_test_proto_rawDescGZIP(), []int{2}
}

func (x *UserWithPtr) GetPtrValue() *wrapperspb.Int64Value {
	if x != nil {
		return x.PtrValue
	}
	return nil
}

type UserWithPtrResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *UserWithPtr `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UserWithPtrResult) Reset() {
	*x = UserWithPtrResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_internal_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserWithPtrResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserWithPtrResult) ProtoMessage() {}

func (x *UserWithPtrResult) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_internal_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserWithPtrResult.ProtoReflect.Descriptor instead.
func (*UserWithPtrResult) Descriptor() ([]byte, []int) {
	return file_gateway_internal_test_proto_rawDescGZIP(), []int{3}
}

func (x *UserWithPtrResult) GetResult() *UserWithPtr {
	if x != nil {
		return x.Result
	}
	return nil
}

type BadResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success []*User `protobuf:"bytes,1,rep,name=success,proto3" json:"success,omitempty"`
}

func (x *BadResult) Reset() {
	*x = BadResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_internal_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadResult) ProtoMessage() {}

func (x *BadResult) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_internal_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadResult.ProtoReflect.Descriptor instead.
func (*BadResult) Descriptor() ([]byte, []int) {
	return file_gateway_internal_test_proto_rawDescGZIP(), []int{4}
}

func (x *BadResult) GetSuccess() []*User {
	if x != nil {
		return x.Success
	}
	return nil
}

var File_gateway_internal_test_proto protoreflect.FileDescriptor

var file_gateway_internal_test_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x69,
	0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x47, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x50, 0x74, 0x72, 0x12, 0x38,
	0x0a, 0x09, 0x70, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08,
	0x70, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72,
	0x57, 0x69, 0x74, 0x68, 0x50, 0x74, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x50, 0x74, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x3a, 0x0a, 0x09, 0x42, 0x61, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f,
	0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2d, 0x61,
	0x70, 0x70, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_internal_test_proto_rawDescOnce sync.Once
	file_gateway_internal_test_proto_rawDescData = file_gateway_internal_test_proto_rawDesc
)

func file_gateway_internal_test_proto_rawDescGZIP() []byte {
	file_gateway_internal_test_proto_rawDescOnce.Do(func() {
		file_gateway_internal_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_internal_test_proto_rawDescData)
	})
	return file_gateway_internal_test_proto_rawDescData
}

var file_gateway_internal_test_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gateway_internal_test_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: infoblox.test.User
	(*Result)(nil),                // 1: infoblox.test.Result
	(*UserWithPtr)(nil),           // 2: infoblox.test.UserWithPtr
	(*UserWithPtrResult)(nil),     // 3: infoblox.test.UserWithPtrResult
	(*BadResult)(nil),             // 4: infoblox.test.BadResult
	(*wrapperspb.Int64Value)(nil), // 5: google.protobuf.Int64Value
}
var file_gateway_internal_test_proto_depIdxs = []int32{
	0, // 0: infoblox.test.Result.users:type_name -> infoblox.test.User
	5, // 1: infoblox.test.UserWithPtr.ptr_value:type_name -> google.protobuf.Int64Value
	2, // 2: infoblox.test.UserWithPtrResult.result:type_name -> infoblox.test.UserWithPtr
	0, // 3: infoblox.test.BadResult.success:type_name -> infoblox.test.User
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gateway_internal_test_proto_init() }
func file_gateway_internal_test_proto_init() {
	if File_gateway_internal_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_internal_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_gateway_internal_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_gateway_internal_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserWithPtr); i {
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
		file_gateway_internal_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserWithPtrResult); i {
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
		file_gateway_internal_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadResult); i {
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
			RawDescriptor: file_gateway_internal_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_internal_test_proto_goTypes,
		DependencyIndexes: file_gateway_internal_test_proto_depIdxs,
		MessageInfos:      file_gateway_internal_test_proto_msgTypes,
	}.Build()
	File_gateway_internal_test_proto = out.File
	file_gateway_internal_test_proto_rawDesc = nil
	file_gateway_internal_test_proto_goTypes = nil
	file_gateway_internal_test_proto_depIdxs = nil
}