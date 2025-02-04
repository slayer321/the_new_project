// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: services/user_profile/user_service/pb/user.proto

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

type GetUserProfileRes_UserRole int32

const (
	GetUserProfileRes_USER_ADMIN    GetUserProfileRes_UserRole = 0
	GetUserProfileRes_USER_NORMAL   GetUserProfileRes_UserRole = 1
	GetUserProfileRes_USER_SUPPORT  GetUserProfileRes_UserRole = 2
	GetUserProfileRes_USER_SECURITY GetUserProfileRes_UserRole = 3
)

// Enum value maps for GetUserProfileRes_UserRole.
var (
	GetUserProfileRes_UserRole_name = map[int32]string{
		0: "USER_ADMIN",
		1: "USER_NORMAL",
		2: "USER_SUPPORT",
		3: "USER_SECURITY",
	}
	GetUserProfileRes_UserRole_value = map[string]int32{
		"USER_ADMIN":    0,
		"USER_NORMAL":   1,
		"USER_SUPPORT":  2,
		"USER_SECURITY": 3,
	}
)

func (x GetUserProfileRes_UserRole) Enum() *GetUserProfileRes_UserRole {
	p := new(GetUserProfileRes_UserRole)
	*p = x
	return p
}

func (x GetUserProfileRes_UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetUserProfileRes_UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_services_user_profile_user_service_pb_user_proto_enumTypes[0].Descriptor()
}

func (GetUserProfileRes_UserRole) Type() protoreflect.EnumType {
	return &file_services_user_profile_user_service_pb_user_proto_enumTypes[0]
}

func (x GetUserProfileRes_UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetUserProfileRes_UserRole.Descriptor instead.
func (GetUserProfileRes_UserRole) EnumDescriptor() ([]byte, []int) {
	return file_services_user_profile_user_service_pb_user_proto_rawDescGZIP(), []int{1, 0}
}

type GetUserProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserProfileReq) Reset() {
	*x = GetUserProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_profile_user_service_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileReq) ProtoMessage() {}

func (x *GetUserProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_profile_user_service_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileReq.ProtoReflect.Descriptor instead.
func (*GetUserProfileReq) Descriptor() ([]byte, []int) {
	return file_services_user_profile_user_service_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserProfileReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserProfileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName  string                     `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName   string                     `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email      string                     `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	ProfilePic []byte                     `protobuf:"bytes,5,opt,name=profilePic,proto3" json:"profilePic,omitempty"`
	CreateTime *timestamppb.Timestamp     `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime *timestamppb.Timestamp     `protobuf:"bytes,7,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	IsActive   bool                       `protobuf:"varint,8,opt,name=isActive,proto3" json:"isActive,omitempty"`
	Role       GetUserProfileRes_UserRole `protobuf:"varint,9,opt,name=role,proto3,enum=auth.GetUserProfileRes_UserRole" json:"role,omitempty"`
	LastLogin  *timestamppb.Timestamp     `protobuf:"bytes,10,opt,name=lastLogin,proto3" json:"lastLogin,omitempty"`
}

func (x *GetUserProfileRes) Reset() {
	*x = GetUserProfileRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_profile_user_service_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserProfileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileRes) ProtoMessage() {}

func (x *GetUserProfileRes) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_profile_user_service_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileRes.ProtoReflect.Descriptor instead.
func (*GetUserProfileRes) Descriptor() ([]byte, []int) {
	return file_services_user_profile_user_service_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserProfileRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserProfileRes) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetUserProfileRes) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetUserProfileRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserProfileRes) GetProfilePic() []byte {
	if x != nil {
		return x.ProfilePic
	}
	return nil
}

func (x *GetUserProfileRes) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GetUserProfileRes) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *GetUserProfileRes) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *GetUserProfileRes) GetRole() GetUserProfileRes_UserRole {
	if x != nil {
		return x.Role
	}
	return GetUserProfileRes_USER_ADMIN
}

func (x *GetUserProfileRes) GetLastLogin() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLogin
	}
	return nil
}

var File_services_user_profile_user_service_pb_user_proto protoreflect.FileDescriptor

var file_services_user_profile_user_service_pb_user_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xe9,
	0x03, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x69, 0x63, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x50, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x52,
	0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x55,
	0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x10, 0x03, 0x32, 0x53, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42,
	0x29, 0x5a, 0x27, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_services_user_profile_user_service_pb_user_proto_rawDescOnce sync.Once
	file_services_user_profile_user_service_pb_user_proto_rawDescData = file_services_user_profile_user_service_pb_user_proto_rawDesc
)

func file_services_user_profile_user_service_pb_user_proto_rawDescGZIP() []byte {
	file_services_user_profile_user_service_pb_user_proto_rawDescOnce.Do(func() {
		file_services_user_profile_user_service_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_user_profile_user_service_pb_user_proto_rawDescData)
	})
	return file_services_user_profile_user_service_pb_user_proto_rawDescData
}

var file_services_user_profile_user_service_pb_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_user_profile_user_service_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_services_user_profile_user_service_pb_user_proto_goTypes = []interface{}{
	(GetUserProfileRes_UserRole)(0), // 0: auth.GetUserProfileRes.UserRole
	(*GetUserProfileReq)(nil),       // 1: auth.GetUserProfileReq
	(*GetUserProfileRes)(nil),       // 2: auth.GetUserProfileRes
	(*timestamppb.Timestamp)(nil),   // 3: google.protobuf.Timestamp
}
var file_services_user_profile_user_service_pb_user_proto_depIdxs = []int32{
	3, // 0: auth.GetUserProfileRes.createTime:type_name -> google.protobuf.Timestamp
	3, // 1: auth.GetUserProfileRes.updateTime:type_name -> google.protobuf.Timestamp
	0, // 2: auth.GetUserProfileRes.role:type_name -> auth.GetUserProfileRes.UserRole
	3, // 3: auth.GetUserProfileRes.lastLogin:type_name -> google.protobuf.Timestamp
	1, // 4: auth.UserService.GetUserProfile:input_type -> auth.GetUserProfileReq
	2, // 5: auth.UserService.GetUserProfile:output_type -> auth.GetUserProfileRes
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_user_profile_user_service_pb_user_proto_init() }
func file_services_user_profile_user_service_pb_user_proto_init() {
	if File_services_user_profile_user_service_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_user_profile_user_service_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserProfileReq); i {
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
		file_services_user_profile_user_service_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserProfileRes); i {
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
			RawDescriptor: file_services_user_profile_user_service_pb_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_user_profile_user_service_pb_user_proto_goTypes,
		DependencyIndexes: file_services_user_profile_user_service_pb_user_proto_depIdxs,
		EnumInfos:         file_services_user_profile_user_service_pb_user_proto_enumTypes,
		MessageInfos:      file_services_user_profile_user_service_pb_user_proto_msgTypes,
	}.Build()
	File_services_user_profile_user_service_pb_user_proto = out.File
	file_services_user_profile_user_service_pb_user_proto_rawDesc = nil
	file_services_user_profile_user_service_pb_user_proto_goTypes = nil
	file_services_user_profile_user_service_pb_user_proto_depIdxs = nil
}
