// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: laptop_message.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Laptop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram      *Memory    `protobuf:"bytes,5,opt,name=Ram,proto3" json:"Ram,omitempty"`
	Gpu      []*GPU     `protobuf:"bytes,6,rep,name=gpu,proto3" json:"gpu,omitempty"`
	Storage  []*Storage `protobuf:"bytes,7,rep,name=storage,proto3" json:"storage,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are assignable to Weight:
	//	*Laptop_WightKg
	//	*Laptop_WightLb
	Weight      isLaptop_Weight        `protobuf_oneof:"weight"`
	PriceUsd    float64                `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear uint32                 `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdateAt    *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *Laptop) Reset() {
	*x = Laptop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_laptop_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Laptop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Laptop) ProtoMessage() {}

func (x *Laptop) ProtoReflect() protoreflect.Message {
	mi := &file_laptop_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Laptop.ProtoReflect.Descriptor instead.
func (*Laptop) Descriptor() ([]byte, []int) {
	return file_laptop_message_proto_rawDescGZIP(), []int{0}
}

func (x *Laptop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Laptop) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Laptop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Laptop) GetCpu() *CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *Laptop) GetRam() *Memory {
	if x != nil {
		return x.Ram
	}
	return nil
}

func (x *Laptop) GetGpu() []*GPU {
	if x != nil {
		return x.Gpu
	}
	return nil
}

func (x *Laptop) GetStorage() []*Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *Laptop) GetScreen() *Screen {
	if x != nil {
		return x.Screen
	}
	return nil
}

func (x *Laptop) GetKeyboard() *Keyboard {
	if x != nil {
		return x.Keyboard
	}
	return nil
}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (x *Laptop) GetWightKg() float64 {
	if x, ok := x.GetWeight().(*Laptop_WightKg); ok {
		return x.WightKg
	}
	return 0
}

func (x *Laptop) GetWightLb() float64 {
	if x, ok := x.GetWeight().(*Laptop_WightLb); ok {
		return x.WightLb
	}
	return 0
}

func (x *Laptop) GetPriceUsd() float64 {
	if x != nil {
		return x.PriceUsd
	}
	return 0
}

func (x *Laptop) GetReleaseYear() uint32 {
	if x != nil {
		return x.ReleaseYear
	}
	return 0
}

func (x *Laptop) GetUpdateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateAt
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WightKg struct {
	WightKg float64 `protobuf:"fixed64,10,opt,name=wight_kg,json=wightKg,proto3,oneof"`
}

type Laptop_WightLb struct {
	WightLb float64 `protobuf:"fixed64,11,opt,name=wight_lb,json=wightLb,proto3,oneof"`
}

func (*Laptop_WightKg) isLaptop_Weight() {}

func (*Laptop_WightLb) isLaptop_Weight() {}

var File_laptop_message_proto protoreflect.FileDescriptor

var file_laptop_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x03, 0x0a, 0x06,
	0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e,
	0x43, 0x50, 0x55, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x19, 0x0a, 0x03, 0x52, 0x61, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x03,
	0x52, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x03, 0x67, 0x70, 0x75, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x04, 0x2e, 0x47, 0x50, 0x55, 0x52, 0x03, 0x67, 0x70, 0x75, 0x12, 0x22, 0x0a, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x06, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x12, 0x25, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x08, 0x6b,
	0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x77, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x6b, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x07, 0x77, 0x69, 0x67,
	0x68, 0x74, 0x4b, 0x67, 0x12, 0x1b, 0x0a, 0x08, 0x77, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6c, 0x62,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x07, 0x77, 0x69, 0x67, 0x68, 0x74, 0x4c,
	0x62, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x55, 0x73, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x37, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_laptop_message_proto_rawDescOnce sync.Once
	file_laptop_message_proto_rawDescData = file_laptop_message_proto_rawDesc
)

func file_laptop_message_proto_rawDescGZIP() []byte {
	file_laptop_message_proto_rawDescOnce.Do(func() {
		file_laptop_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_laptop_message_proto_rawDescData)
	})
	return file_laptop_message_proto_rawDescData
}

var file_laptop_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_laptop_message_proto_goTypes = []interface{}{
	(*Laptop)(nil),                // 0: Laptop
	(*CPU)(nil),                   // 1: CPU
	(*Memory)(nil),                // 2: Memory
	(*GPU)(nil),                   // 3: GPU
	(*Storage)(nil),               // 4: Storage
	(*Screen)(nil),                // 5: Screen
	(*Keyboard)(nil),              // 6: Keyboard
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_laptop_message_proto_depIdxs = []int32{
	1, // 0: Laptop.cpu:type_name -> CPU
	2, // 1: Laptop.Ram:type_name -> Memory
	3, // 2: Laptop.gpu:type_name -> GPU
	4, // 3: Laptop.storage:type_name -> Storage
	5, // 4: Laptop.screen:type_name -> Screen
	6, // 5: Laptop.keyboard:type_name -> Keyboard
	7, // 6: Laptop.update_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_laptop_message_proto_init() }
func file_laptop_message_proto_init() {
	if File_laptop_message_proto != nil {
		return
	}
	file_processor_message_proto_init()
	file_memory_message_proto_init()
	file_storage_message_proto_init()
	file_screen_message_proto_init()
	file_keyboard_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_laptop_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Laptop); i {
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
	file_laptop_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Laptop_WightKg)(nil),
		(*Laptop_WightLb)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_laptop_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_laptop_message_proto_goTypes,
		DependencyIndexes: file_laptop_message_proto_depIdxs,
		MessageInfos:      file_laptop_message_proto_msgTypes,
	}.Build()
	File_laptop_message_proto = out.File
	file_laptop_message_proto_rawDesc = nil
	file_laptop_message_proto_goTypes = nil
	file_laptop_message_proto_depIdxs = nil
}
