// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: nested.proto

package testingpb

import (
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SearchResponse_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{0}
}

func (x *SearchResponse) GetResults() []*SearchResponse_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type SomeOtherMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *SearchResponse_Result     `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Ques   *SomeOtherMessage_Question `protobuf:"bytes,2,opt,name=ques,proto3" json:"ques,omitempty"`
}

func (x *SomeOtherMessage) Reset() {
	*x = SomeOtherMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeOtherMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeOtherMessage) ProtoMessage() {}

func (x *SomeOtherMessage) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeOtherMessage.ProtoReflect.Descriptor instead.
func (*SomeOtherMessage) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{1}
}

func (x *SomeOtherMessage) GetResult() *SearchResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *SomeOtherMessage) GetQues() *SomeOtherMessage_Question {
	if x != nil {
		return x.Ques
	}
	return nil
}

type SearchResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets []string `protobuf:"bytes,3,rep,name=snippets,proto3" json:"snippets,omitempty"`
}

func (x *SearchResponse_Result) Reset() {
	*x = SearchResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse_Result) ProtoMessage() {}

func (x *SearchResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse_Result.ProtoReflect.Descriptor instead.
func (*SearchResponse_Result) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SearchResponse_Result) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SearchResponse_Result) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SearchResponse_Result) GetSnippets() []string {
	if x != nil {
		return x.Snippets
	}
	return nil
}

type SomeOtherMessage_Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Results map[int64]*SearchResponse_Result `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SomeOtherMessage_Question) Reset() {
	*x = SomeOtherMessage_Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SomeOtherMessage_Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeOtherMessage_Question) ProtoMessage() {}

func (x *SomeOtherMessage_Question) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeOtherMessage_Question.ProtoReflect.Descriptor instead.
func (*SomeOtherMessage_Question) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SomeOtherMessage_Question) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SomeOtherMessage_Question) GetResults() map[int64]*SearchResponse_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_nested_proto protoreflect.FileDescriptor

var file_nested_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x1a, 0x4c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73,
	0x22, 0xe2, 0x02, 0x0a, 0x10, 0x53, 0x6f, 0x6d, 0x65, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x71, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x71, 0x75, 0x65,
	0x73, 0x1a, 0xd1, 0x01, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x1a, 0x60, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65,
	0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_nested_proto_rawDescOnce sync.Once
	file_nested_proto_rawDescData = file_nested_proto_rawDesc
)

func file_nested_proto_rawDescGZIP() []byte {
	file_nested_proto_rawDescOnce.Do(func() {
		file_nested_proto_rawDescData = protoimpl.X.CompressGZIP(file_nested_proto_rawDescData)
	})
	return file_nested_proto_rawDescData
}

var file_nested_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nested_proto_goTypes = []interface{}{
	(*SearchResponse)(nil),            // 0: nestedservice.SearchResponse
	(*SomeOtherMessage)(nil),          // 1: nestedservice.SomeOtherMessage
	(*SearchResponse_Result)(nil),     // 2: nestedservice.SearchResponse.Result
	(*SomeOtherMessage_Question)(nil), // 3: nestedservice.SomeOtherMessage.Question
	nil,                               // 4: nestedservice.SomeOtherMessage.Question.ResultsEntry
}
var file_nested_proto_depIdxs = []int32{
	2, // 0: nestedservice.SearchResponse.results:type_name -> nestedservice.SearchResponse.Result
	2, // 1: nestedservice.SomeOtherMessage.result:type_name -> nestedservice.SearchResponse.Result
	3, // 2: nestedservice.SomeOtherMessage.ques:type_name -> nestedservice.SomeOtherMessage.Question
	4, // 3: nestedservice.SomeOtherMessage.Question.results:type_name -> nestedservice.SomeOtherMessage.Question.ResultsEntry
	2, // 4: nestedservice.SomeOtherMessage.Question.ResultsEntry.value:type_name -> nestedservice.SearchResponse.Result
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_nested_proto_init() }
func file_nested_proto_init() {
	if File_nested_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nested_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_nested_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeOtherMessage); i {
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
		file_nested_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse_Result); i {
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
		file_nested_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SomeOtherMessage_Question); i {
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
			RawDescriptor: file_nested_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nested_proto_goTypes,
		DependencyIndexes: file_nested_proto_depIdxs,
		MessageInfos:      file_nested_proto_msgTypes,
	}.Build()
	File_nested_proto = out.File
	file_nested_proto_rawDesc = nil
	file_nested_proto_goTypes = nil
	file_nested_proto_depIdxs = nil
}
