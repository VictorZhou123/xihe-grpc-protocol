// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: competition.proto

package protocol

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

type SubmissionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompetitionId string  `protobuf:"bytes,2,opt,name=competition_id,json=competitionId,proto3" json:"competition_id,omitempty"`
	Phase         string  `protobuf:"bytes,3,opt,name=phase,proto3" json:"phase,omitempty"`
	Status        string  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Score         float32 `protobuf:"fixed32,5,opt,name=score,proto3" json:"score,omitempty"`
	PlayerId      string  `protobuf:"bytes,6,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *SubmissionInfo) Reset() {
	*x = SubmissionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmissionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionInfo) ProtoMessage() {}

func (x *SubmissionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionInfo.ProtoReflect.Descriptor instead.
func (*SubmissionInfo) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{0}
}

func (x *SubmissionInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubmissionInfo) GetCompetitionId() string {
	if x != nil {
		return x.CompetitionId
	}
	return ""
}

func (x *SubmissionInfo) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *SubmissionInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SubmissionInfo) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *SubmissionInfo) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type SubmissionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmissionResult) Reset() {
	*x = SubmissionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmissionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionResult) ProtoMessage() {}

func (x *SubmissionResult) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionResult.ProtoReflect.Descriptor instead.
func (*SubmissionResult) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{1}
}

type TeamMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (x *TeamMembersRequest) Reset() {
	*x = TeamMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamMembersRequest) ProtoMessage() {}

func (x *TeamMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamMembersRequest.ProtoReflect.Descriptor instead.
func (*TeamMembersRequest) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{2}
}

func (x *TeamMembersRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

type TeamMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members      []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	ErrorMessage string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *TeamMembersResponse) Reset() {
	*x = TeamMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_competition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamMembersResponse) ProtoMessage() {}

func (x *TeamMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_competition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamMembersResponse.ProtoReflect.Descriptor instead.
func (*TeamMembersResponse) Descriptor() ([]byte, []int) {
	return file_competition_proto_rawDescGZIP(), []int{3}
}

func (x *TeamMembersResponse) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *TeamMembersResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_competition_proto protoreflect.FileDescriptor

var file_competition_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xa8, 0x01, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x28, 0x0a, 0x12, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x22, 0x54, 0x0a, 0x13, 0x54, 0x65, 0x61,
	0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0xb3, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4f, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_competition_proto_rawDescOnce sync.Once
	file_competition_proto_rawDescData = file_competition_proto_rawDesc
)

func file_competition_proto_rawDescGZIP() []byte {
	file_competition_proto_rawDescOnce.Do(func() {
		file_competition_proto_rawDescData = protoimpl.X.CompressGZIP(file_competition_proto_rawDescData)
	})
	return file_competition_proto_rawDescData
}

var file_competition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_competition_proto_goTypes = []interface{}{
	(*SubmissionInfo)(nil),      // 0: competition.SubmissionInfo
	(*SubmissionResult)(nil),    // 1: competition.SubmissionResult
	(*TeamMembersRequest)(nil),  // 2: competition.TeamMembersRequest
	(*TeamMembersResponse)(nil), // 3: competition.TeamMembersResponse
}
var file_competition_proto_depIdxs = []int32{
	0, // 0: competition.Competition.SetSubmissionInfo:input_type -> competition.SubmissionInfo
	2, // 1: competition.Competition.GetTeamMembers:input_type -> competition.TeamMembersRequest
	1, // 2: competition.Competition.SetSubmissionInfo:output_type -> competition.SubmissionResult
	3, // 3: competition.Competition.GetTeamMembers:output_type -> competition.TeamMembersResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_competition_proto_init() }
func file_competition_proto_init() {
	if File_competition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_competition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmissionInfo); i {
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
		file_competition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmissionResult); i {
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
		file_competition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamMembersRequest); i {
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
		file_competition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamMembersResponse); i {
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
			RawDescriptor: file_competition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_competition_proto_goTypes,
		DependencyIndexes: file_competition_proto_depIdxs,
		MessageInfos:      file_competition_proto_msgTypes,
	}.Build()
	File_competition_proto = out.File
	file_competition_proto_rawDesc = nil
	file_competition_proto_goTypes = nil
	file_competition_proto_depIdxs = nil
}
