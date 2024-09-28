// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: evidence/evidence.proto

package evidencepb

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

type Evidence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CaseId      string `protobuf:"bytes,2,opt,name=case_id,json=caseId,proto3" json:"case_id,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Location    string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Evidence) Reset() {
	*x = Evidence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evidence_evidence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Evidence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Evidence) ProtoMessage() {}

func (x *Evidence) ProtoReflect() protoreflect.Message {
	mi := &file_evidence_evidence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Evidence.ProtoReflect.Descriptor instead.
func (*Evidence) Descriptor() ([]byte, []int) {
	return file_evidence_evidence_proto_rawDescGZIP(), []int{0}
}

func (x *Evidence) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Evidence) GetCaseId() string {
	if x != nil {
		return x.CaseId
	}
	return ""
}

func (x *Evidence) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Evidence) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type ListEvidenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaseId string `protobuf:"bytes,1,opt,name=case_id,json=caseId,proto3" json:"case_id,omitempty"`
}

func (x *ListEvidenceRequest) Reset() {
	*x = ListEvidenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evidence_evidence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEvidenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEvidenceRequest) ProtoMessage() {}

func (x *ListEvidenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evidence_evidence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEvidenceRequest.ProtoReflect.Descriptor instead.
func (*ListEvidenceRequest) Descriptor() ([]byte, []int) {
	return file_evidence_evidence_proto_rawDescGZIP(), []int{1}
}

func (x *ListEvidenceRequest) GetCaseId() string {
	if x != nil {
		return x.CaseId
	}
	return ""
}

type EvidenceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evidence []*Evidence `protobuf:"bytes,1,rep,name=evidence,proto3" json:"evidence,omitempty"`
}

func (x *EvidenceList) Reset() {
	*x = EvidenceList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evidence_evidence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvidenceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvidenceList) ProtoMessage() {}

func (x *EvidenceList) ProtoReflect() protoreflect.Message {
	mi := &file_evidence_evidence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvidenceList.ProtoReflect.Descriptor instead.
func (*EvidenceList) Descriptor() ([]byte, []int) {
	return file_evidence_evidence_proto_rawDescGZIP(), []int{2}
}

func (x *EvidenceList) GetEvidence() []*Evidence {
	if x != nil {
		return x.Evidence
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evidence_evidence_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_evidence_evidence_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_evidence_evidence_proto_rawDescGZIP(), []int{3}
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaseId string `protobuf:"bytes,1,opt,name=case_id,json=caseId,proto3" json:"case_id,omitempty"`
}

func (x *ListLocationsRequest) Reset() {
	*x = ListLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evidence_evidence_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocationsRequest) ProtoMessage() {}

func (x *ListLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evidence_evidence_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocationsRequest.ProtoReflect.Descriptor instead.
func (*ListLocationsRequest) Descriptor() ([]byte, []int) {
	return file_evidence_evidence_proto_rawDescGZIP(), []int{4}
}

func (x *ListLocationsRequest) GetCaseId() string {
	if x != nil {
		return x.CaseId
	}
	return ""
}

type LocationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *LocationList) Reset() {
	*x = LocationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evidence_evidence_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationList) ProtoMessage() {}

func (x *LocationList) ProtoReflect() protoreflect.Message {
	mi := &file_evidence_evidence_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationList.ProtoReflect.Descriptor instead.
func (*LocationList) Descriptor() ([]byte, []int) {
	return file_evidence_evidence_proto_rawDescGZIP(), []int{5}
}

func (x *LocationList) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

var File_evidence_evidence_proto protoreflect.FileDescriptor

var file_evidence_evidence_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x65, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x70, 0x62, 0x22, 0x71, 0x0a, 0x08, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x0c, 0x45, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x65, 0x76, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x08, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0c, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32,
	0xa9, 0x01, 0x0a, 0x0f, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x70,
	0x62, 0x2e, 0x45, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x20, 0x2e, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x70, 0x62, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x23, 0x5a, 0x21, 0x6f,
	0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x3b, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evidence_evidence_proto_rawDescOnce sync.Once
	file_evidence_evidence_proto_rawDescData = file_evidence_evidence_proto_rawDesc
)

func file_evidence_evidence_proto_rawDescGZIP() []byte {
	file_evidence_evidence_proto_rawDescOnce.Do(func() {
		file_evidence_evidence_proto_rawDescData = protoimpl.X.CompressGZIP(file_evidence_evidence_proto_rawDescData)
	})
	return file_evidence_evidence_proto_rawDescData
}

var file_evidence_evidence_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_evidence_evidence_proto_goTypes = []any{
	(*Evidence)(nil),             // 0: evidencepb.Evidence
	(*ListEvidenceRequest)(nil),  // 1: evidencepb.ListEvidenceRequest
	(*EvidenceList)(nil),         // 2: evidencepb.EvidenceList
	(*Location)(nil),             // 3: evidencepb.Location
	(*ListLocationsRequest)(nil), // 4: evidencepb.ListLocationsRequest
	(*LocationList)(nil),         // 5: evidencepb.LocationList
}
var file_evidence_evidence_proto_depIdxs = []int32{
	0, // 0: evidencepb.EvidenceList.evidence:type_name -> evidencepb.Evidence
	3, // 1: evidencepb.LocationList.locations:type_name -> evidencepb.Location
	1, // 2: evidencepb.EvidenceService.ListEvidence:input_type -> evidencepb.ListEvidenceRequest
	4, // 3: evidencepb.EvidenceService.ListLocations:input_type -> evidencepb.ListLocationsRequest
	2, // 4: evidencepb.EvidenceService.ListEvidence:output_type -> evidencepb.EvidenceList
	5, // 5: evidencepb.EvidenceService.ListLocations:output_type -> evidencepb.LocationList
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_evidence_evidence_proto_init() }
func file_evidence_evidence_proto_init() {
	if File_evidence_evidence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_evidence_evidence_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Evidence); i {
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
		file_evidence_evidence_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListEvidenceRequest); i {
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
		file_evidence_evidence_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EvidenceList); i {
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
		file_evidence_evidence_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Location); i {
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
		file_evidence_evidence_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListLocationsRequest); i {
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
		file_evidence_evidence_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LocationList); i {
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
			RawDescriptor: file_evidence_evidence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_evidence_evidence_proto_goTypes,
		DependencyIndexes: file_evidence_evidence_proto_depIdxs,
		MessageInfos:      file_evidence_evidence_proto_msgTypes,
	}.Build()
	File_evidence_evidence_proto = out.File
	file_evidence_evidence_proto_rawDesc = nil
	file_evidence_evidence_proto_goTypes = nil
	file_evidence_evidence_proto_depIdxs = nil
}
