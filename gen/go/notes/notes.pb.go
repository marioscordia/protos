// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.27.1
// source: notes/notes.proto

package notesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	OwnerId       int64                  `protobuf:"varint,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Collaborators []int64                `protobuf:"varint,4,rep,packed,name=collaborators,proto3" json:"collaborators,omitempty"`
	Tags          []string               `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	IsPublic      bool                   `protobuf:"varint,6,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_notes_notes_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *CreateRequest) GetCollaborators() []int64 {
	if x != nil {
		return x.Collaborators
	}
	return nil
}

func (x *CreateRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NoteId        int64                  `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_notes_notes_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

type GetByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NoteId        int64                  `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetByIdRequest) Reset() {
	*x = GetByIdRequest{}
	mi := &file_notes_notes_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdRequest) ProtoMessage() {}

func (x *GetByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdRequest.ProtoReflect.Descriptor instead.
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{2}
}

func (x *GetByIdRequest) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

type GetByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NoteId        int64                  `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	OwnerId       int64                  `protobuf:"varint,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Collaborators []int64                `protobuf:"varint,5,rep,packed,name=collaborators,proto3" json:"collaborators,omitempty"`
	Tags          []string               `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	IsPublic      bool                   `protobuf:"varint,7,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetByIdResponse) Reset() {
	*x = GetByIdResponse{}
	mi := &file_notes_notes_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdResponse) ProtoMessage() {}

func (x *GetByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdResponse.ProtoReflect.Descriptor instead.
func (*GetByIdResponse) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{3}
}

func (x *GetByIdResponse) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

func (x *GetByIdResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetByIdResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetByIdResponse) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *GetByIdResponse) GetCollaborators() []int64 {
	if x != nil {
		return x.Collaborators
	}
	return nil
}

func (x *GetByIdResponse) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetByIdResponse) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *GetByIdResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetByIdResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	NoteId        int64                   `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	Title         *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Tags          []string                `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_notes_notes_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

func (x *UpdateRequest) GetTitle() *wrapperspb.StringValue {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *UpdateRequest) GetContent() *wrapperspb.StringValue {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *UpdateRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NoteId        int64                  `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_notes_notes_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRequest) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

type AddCollaboratorsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserIds       []int64                `protobuf:"varint,1,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddCollaboratorsRequest) Reset() {
	*x = AddCollaboratorsRequest{}
	mi := &file_notes_notes_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCollaboratorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCollaboratorsRequest) ProtoMessage() {}

func (x *AddCollaboratorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notes_notes_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCollaboratorsRequest.ProtoReflect.Descriptor instead.
func (*AddCollaboratorsRequest) Descriptor() ([]byte, []int) {
	return file_notes_notes_proto_rawDescGZIP(), []int{6}
}

func (x *AddCollaboratorsRequest) GetUserIds() []int64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

var File_notes_notes_proto protoreflect.FileDescriptor

var file_notes_notes_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x29, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65,
	0x49, 0x64, 0x22, 0xc2, 0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x22, 0x28, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x17,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x32, 0xb4, 0x02, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1a, 0x5a, 0x18, 0x73, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x61, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x3b, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_notes_notes_proto_rawDescOnce sync.Once
	file_notes_notes_proto_rawDescData []byte
)

func file_notes_notes_proto_rawDescGZIP() []byte {
	file_notes_notes_proto_rawDescOnce.Do(func() {
		file_notes_notes_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_notes_notes_proto_rawDesc), len(file_notes_notes_proto_rawDesc)))
	})
	return file_notes_notes_proto_rawDescData
}

var file_notes_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_notes_notes_proto_goTypes = []any{
	(*CreateRequest)(nil),           // 0: notes.CreateRequest
	(*CreateResponse)(nil),          // 1: notes.CreateResponse
	(*GetByIdRequest)(nil),          // 2: notes.GetByIdRequest
	(*GetByIdResponse)(nil),         // 3: notes.GetByIdResponse
	(*UpdateRequest)(nil),           // 4: notes.UpdateRequest
	(*DeleteRequest)(nil),           // 5: notes.DeleteRequest
	(*AddCollaboratorsRequest)(nil), // 6: notes.AddCollaboratorsRequest
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil),  // 8: google.protobuf.StringValue
	(*emptypb.Empty)(nil),           // 9: google.protobuf.Empty
}
var file_notes_notes_proto_depIdxs = []int32{
	7, // 0: notes.GetByIdResponse.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: notes.GetByIdResponse.updated_at:type_name -> google.protobuf.Timestamp
	8, // 2: notes.UpdateRequest.title:type_name -> google.protobuf.StringValue
	8, // 3: notes.UpdateRequest.content:type_name -> google.protobuf.StringValue
	0, // 4: notes.Notes.Create:input_type -> notes.CreateRequest
	2, // 5: notes.Notes.GetById:input_type -> notes.GetByIdRequest
	4, // 6: notes.Notes.Update:input_type -> notes.UpdateRequest
	5, // 7: notes.Notes.Delete:input_type -> notes.DeleteRequest
	6, // 8: notes.Notes.AddCollaborators:input_type -> notes.AddCollaboratorsRequest
	1, // 9: notes.Notes.Create:output_type -> notes.CreateResponse
	3, // 10: notes.Notes.GetById:output_type -> notes.GetByIdResponse
	9, // 11: notes.Notes.Update:output_type -> google.protobuf.Empty
	9, // 12: notes.Notes.Delete:output_type -> google.protobuf.Empty
	9, // 13: notes.Notes.AddCollaborators:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_notes_notes_proto_init() }
func file_notes_notes_proto_init() {
	if File_notes_notes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_notes_notes_proto_rawDesc), len(file_notes_notes_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notes_notes_proto_goTypes,
		DependencyIndexes: file_notes_notes_proto_depIdxs,
		MessageInfos:      file_notes_notes_proto_msgTypes,
	}.Build()
	File_notes_notes_proto = out.File
	file_notes_notes_proto_goTypes = nil
	file_notes_notes_proto_depIdxs = nil
}
