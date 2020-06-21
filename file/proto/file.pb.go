// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package file

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FileShare struct {
	//The uuid of the file share.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The createdAt representing the server time when the file share was created.
	CreatedAt string `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// The updatedAt representing the server time when the file share was updated.
	UpdatedAt string `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// The name of the file share.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the file share.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// The uuid of the project that the file share belongs to.
	TenantId string `protobuf:"bytes,6,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// The uuid of the user that the file share belongs to.
	UserId string `protobuf:"bytes,7,opt,name=userId,proto3" json:"userId,omitempty"`
	// The uuid of the backend that the file share belongs to.
	BackendId string `protobuf:"bytes,8,opt,name=backendId,proto3" json:"backendId,omitempty"`
	// The name of the backend that the file share belongs to.
	Backend string `protobuf:"bytes,9,opt,name=backend,proto3" json:"backend,omitempty"`
	// The size of the file share requested by the user.
	Size int64 `protobuf:"varint,10,opt,name=size,proto3" json:"size,omitempty"`
	// The type of the file share.
	Type string `protobuf:"bytes,11,opt,name=type,proto3" json:"type,omitempty"`
	// The location that file share belongs to.
	Region string `protobuf:"bytes,12,opt,name=region,proto3" json:"region,omitempty"`
	// The locality that file share belongs to.
	AvailabilityZone string `protobuf:"bytes,13,opt,name=availabilityZone,proto3" json:"availabilityZone,omitempty"`
	// The status of the file share.
	Status string `protobuf:"bytes,14,opt,name=status,proto3" json:"status,omitempty"`
	// The uuid of the snapshot from which the file share is created
	SnapshotId string `protobuf:"bytes,15,opt,name=snapshotId,proto3" json:"snapshotId,omitempty"`
	// The protocol of the fileshare. e.g NFS, SMB etc.
	Protocols []string `protobuf:"bytes,16,rep,name=protocols,proto3" json:"protocols,omitempty"`
	// Any tags assigned to the file share.
	Tags []*Tag `protobuf:"bytes,17,rep,name=tags,proto3" json:"tags,omitempty"`
	// Indicates whether the file share is encrypted.
	Encrypted bool `protobuf:"varint,18,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	// EncryptionSettings that was used to protect the file share encryption.
	EncryptionSettings map[string]string `protobuf:"bytes,19,rep,name=encryptionSettings,proto3" json:"encryptionSettings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata should be kept until the semantics between file share and backend storage resource.
	Metadata             *_struct.Struct `protobuf:"bytes,20,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FileShare) Reset()         { *m = FileShare{} }
func (m *FileShare) String() string { return proto.CompactTextString(m) }
func (*FileShare) ProtoMessage()    {}
func (*FileShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *FileShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileShare.Unmarshal(m, b)
}
func (m *FileShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileShare.Marshal(b, m, deterministic)
}
func (m *FileShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileShare.Merge(m, src)
}
func (m *FileShare) XXX_Size() int {
	return xxx_messageInfo_FileShare.Size(m)
}
func (m *FileShare) XXX_DiscardUnknown() {
	xxx_messageInfo_FileShare.DiscardUnknown(m)
}

var xxx_messageInfo_FileShare proto.InternalMessageInfo

func (m *FileShare) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FileShare) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *FileShare) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *FileShare) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileShare) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FileShare) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *FileShare) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *FileShare) GetBackendId() string {
	if m != nil {
		return m.BackendId
	}
	return ""
}

func (m *FileShare) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func (m *FileShare) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileShare) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FileShare) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *FileShare) GetAvailabilityZone() string {
	if m != nil {
		return m.AvailabilityZone
	}
	return ""
}

func (m *FileShare) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FileShare) GetSnapshotId() string {
	if m != nil {
		return m.SnapshotId
	}
	return ""
}

func (m *FileShare) GetProtocols() []string {
	if m != nil {
		return m.Protocols
	}
	return nil
}

func (m *FileShare) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FileShare) GetEncrypted() bool {
	if m != nil {
		return m.Encrypted
	}
	return false
}

func (m *FileShare) GetEncryptionSettings() map[string]string {
	if m != nil {
		return m.EncryptionSettings
	}
	return nil
}

func (m *FileShare) GetMetadata() *_struct.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Tag struct {
	// The key of the tag.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value of the tag.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{1}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ListFileShareRequest struct {
	Limit                int32             `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32             `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SortKeys             []string          `protobuf:"bytes,3,rep,name=sortKeys,proto3" json:"sortKeys,omitempty"`
	SortDirs             []string          `protobuf:"bytes,4,rep,name=sortDirs,proto3" json:"sortDirs,omitempty"`
	Filter               map[string]string `protobuf:"bytes,5,rep,name=Filter,proto3" json:"Filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListFileShareRequest) Reset()         { *m = ListFileShareRequest{} }
func (m *ListFileShareRequest) String() string { return proto.CompactTextString(m) }
func (*ListFileShareRequest) ProtoMessage()    {}
func (*ListFileShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{2}
}

func (m *ListFileShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFileShareRequest.Unmarshal(m, b)
}
func (m *ListFileShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFileShareRequest.Marshal(b, m, deterministic)
}
func (m *ListFileShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFileShareRequest.Merge(m, src)
}
func (m *ListFileShareRequest) XXX_Size() int {
	return xxx_messageInfo_ListFileShareRequest.Size(m)
}
func (m *ListFileShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFileShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFileShareRequest proto.InternalMessageInfo

func (m *ListFileShareRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListFileShareRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListFileShareRequest) GetSortKeys() []string {
	if m != nil {
		return m.SortKeys
	}
	return nil
}

func (m *ListFileShareRequest) GetSortDirs() []string {
	if m != nil {
		return m.SortDirs
	}
	return nil
}

func (m *ListFileShareRequest) GetFilter() map[string]string {
	if m != nil {
		return m.Filter
	}
	return nil
}

type ListFileShareResponse struct {
	Fileshares           []*FileShare `protobuf:"bytes,1,rep,name=fileshares,proto3" json:"fileshares,omitempty"`
	Next                 int32        `protobuf:"varint,2,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListFileShareResponse) Reset()         { *m = ListFileShareResponse{} }
func (m *ListFileShareResponse) String() string { return proto.CompactTextString(m) }
func (*ListFileShareResponse) ProtoMessage()    {}
func (*ListFileShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{3}
}

func (m *ListFileShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFileShareResponse.Unmarshal(m, b)
}
func (m *ListFileShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFileShareResponse.Marshal(b, m, deterministic)
}
func (m *ListFileShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFileShareResponse.Merge(m, src)
}
func (m *ListFileShareResponse) XXX_Size() int {
	return xxx_messageInfo_ListFileShareResponse.Size(m)
}
func (m *ListFileShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFileShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFileShareResponse proto.InternalMessageInfo

func (m *ListFileShareResponse) GetFileshares() []*FileShare {
	if m != nil {
		return m.Fileshares
	}
	return nil
}

func (m *ListFileShareResponse) GetNext() int32 {
	if m != nil {
		return m.Next
	}
	return 0
}

type GetFileShareRequest struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fileshare            *FileShare `protobuf:"bytes,2,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetFileShareRequest) Reset()         { *m = GetFileShareRequest{} }
func (m *GetFileShareRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileShareRequest) ProtoMessage()    {}
func (*GetFileShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{4}
}

func (m *GetFileShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileShareRequest.Unmarshal(m, b)
}
func (m *GetFileShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileShareRequest.Marshal(b, m, deterministic)
}
func (m *GetFileShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileShareRequest.Merge(m, src)
}
func (m *GetFileShareRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileShareRequest.Size(m)
}
func (m *GetFileShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileShareRequest proto.InternalMessageInfo

func (m *GetFileShareRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetFileShareRequest) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type GetFileShareResponse struct {
	Fileshare            *FileShare `protobuf:"bytes,1,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetFileShareResponse) Reset()         { *m = GetFileShareResponse{} }
func (m *GetFileShareResponse) String() string { return proto.CompactTextString(m) }
func (*GetFileShareResponse) ProtoMessage()    {}
func (*GetFileShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{5}
}

func (m *GetFileShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileShareResponse.Unmarshal(m, b)
}
func (m *GetFileShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileShareResponse.Marshal(b, m, deterministic)
}
func (m *GetFileShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileShareResponse.Merge(m, src)
}
func (m *GetFileShareResponse) XXX_Size() int {
	return xxx_messageInfo_GetFileShareResponse.Size(m)
}
func (m *GetFileShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileShareResponse proto.InternalMessageInfo

func (m *GetFileShareResponse) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type CreateFileShareRequest struct {
	Fileshare            *FileShare `protobuf:"bytes,1,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateFileShareRequest) Reset()         { *m = CreateFileShareRequest{} }
func (m *CreateFileShareRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFileShareRequest) ProtoMessage()    {}
func (*CreateFileShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{6}
}

func (m *CreateFileShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileShareRequest.Unmarshal(m, b)
}
func (m *CreateFileShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileShareRequest.Marshal(b, m, deterministic)
}
func (m *CreateFileShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileShareRequest.Merge(m, src)
}
func (m *CreateFileShareRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFileShareRequest.Size(m)
}
func (m *CreateFileShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileShareRequest proto.InternalMessageInfo

func (m *CreateFileShareRequest) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

type CreateFileShareResponse struct {
	Fileshare            *FileShare `protobuf:"bytes,1,opt,name=fileshare,proto3" json:"fileshare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateFileShareResponse) Reset()         { *m = CreateFileShareResponse{} }
func (m *CreateFileShareResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFileShareResponse) ProtoMessage()    {}
func (*CreateFileShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{7}
}

func (m *CreateFileShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileShareResponse.Unmarshal(m, b)
}
func (m *CreateFileShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileShareResponse.Marshal(b, m, deterministic)
}
func (m *CreateFileShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileShareResponse.Merge(m, src)
}
func (m *CreateFileShareResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFileShareResponse.Size(m)
}
func (m *CreateFileShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileShareResponse proto.InternalMessageInfo

func (m *CreateFileShareResponse) GetFileshare() *FileShare {
	if m != nil {
		return m.Fileshare
	}
	return nil
}

func init() {
	proto.RegisterType((*FileShare)(nil), "FileShare")
	proto.RegisterMapType((map[string]string)(nil), "FileShare.EncryptionSettingsEntry")
	proto.RegisterType((*Tag)(nil), "Tag")
	proto.RegisterType((*ListFileShareRequest)(nil), "ListFileShareRequest")
	proto.RegisterMapType((map[string]string)(nil), "ListFileShareRequest.FilterEntry")
	proto.RegisterType((*ListFileShareResponse)(nil), "ListFileShareResponse")
	proto.RegisterType((*GetFileShareRequest)(nil), "GetFileShareRequest")
	proto.RegisterType((*GetFileShareResponse)(nil), "GetFileShareResponse")
	proto.RegisterType((*CreateFileShareRequest)(nil), "CreateFileShareRequest")
	proto.RegisterType((*CreateFileShareResponse)(nil), "CreateFileShareResponse")
}

func init() {
	proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162)
}

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x4e, 0xdb, 0x3c,
	0x18, 0x26, 0xfd, 0xa3, 0x7d, 0xcb, 0x4f, 0x3f, 0x53, 0xa8, 0x55, 0xa1, 0x4f, 0x5d, 0x8e, 0x2a,
	0xa4, 0x19, 0x09, 0x4e, 0xc6, 0xa4, 0x4d, 0x30, 0x06, 0x13, 0xda, 0xa4, 0x4d, 0x01, 0x69, 0xd2,
	0xce, 0xdc, 0xe6, 0x6d, 0xb1, 0x08, 0x49, 0x16, 0x3b, 0x68, 0xdd, 0xed, 0xed, 0x16, 0x76, 0x25,
	0xbb, 0x82, 0xc9, 0x76, 0x9a, 0xb6, 0x34, 0x4c, 0xb0, 0x9d, 0xf9, 0x79, 0x1e, 0xfb, 0xf1, 0xfb,
	0x67, 0x03, 0x8c, 0x44, 0x80, 0x2c, 0x4e, 0x22, 0x15, 0x75, 0x77, 0xc7, 0x51, 0x34, 0x0e, 0x70,
	0xdf, 0xa0, 0x41, 0x3a, 0xda, 0x97, 0x2a, 0x49, 0x87, 0xca, 0xaa, 0xee, 0x8f, 0x2a, 0x34, 0xce,
	0x45, 0x80, 0x97, 0xd7, 0x3c, 0x41, 0xb2, 0x01, 0x25, 0xe1, 0x53, 0xa7, 0xe7, 0xf4, 0x1b, 0x5e,
	0x49, 0xf8, 0x64, 0x17, 0x1a, 0xc3, 0x04, 0xb9, 0x42, 0xff, 0x44, 0xd1, 0x92, 0xa1, 0x67, 0x84,
	0x56, 0xd3, 0xd8, 0xcf, 0xd4, 0xb2, 0x55, 0x73, 0x82, 0x10, 0xa8, 0x84, 0xfc, 0x16, 0x69, 0xc5,
	0x08, 0x66, 0x4d, 0x7a, 0xd0, 0xf4, 0x51, 0x0e, 0x13, 0x11, 0x2b, 0x11, 0x85, 0xb4, 0x6a, 0xa4,
	0x79, 0x8a, 0x74, 0xa1, 0xae, 0x30, 0xe4, 0xa1, 0xba, 0xf0, 0x69, 0xcd, 0xc8, 0x39, 0x26, 0x3b,
	0x50, 0x4b, 0x25, 0x26, 0x17, 0x3e, 0x5d, 0x35, 0x4a, 0x86, 0x74, 0x1c, 0x03, 0x3e, 0xbc, 0xc1,
	0xd0, 0xbf, 0xf0, 0x69, 0xdd, 0xc6, 0x91, 0x13, 0x84, 0xc2, 0x6a, 0x06, 0x68, 0xc3, 0x68, 0x53,
	0xa8, 0x23, 0x94, 0xe2, 0x3b, 0x52, 0xe8, 0x39, 0xfd, 0xb2, 0x67, 0xd6, 0x9a, 0x53, 0x93, 0x18,
	0x69, 0xd3, 0x46, 0xad, 0xd7, 0xfa, 0xde, 0x04, 0xc7, 0x3a, 0xe0, 0x35, 0x7b, 0xaf, 0x45, 0x64,
	0x0f, 0x5a, 0xfc, 0x8e, 0x8b, 0x80, 0x0f, 0x44, 0x20, 0xd4, 0xe4, 0x4b, 0x14, 0x22, 0x5d, 0x37,
	0x3b, 0x96, 0x78, 0xed, 0x21, 0x15, 0x57, 0xa9, 0xa4, 0x1b, 0xd6, 0xc3, 0x22, 0xf2, 0x3f, 0x80,
	0x0c, 0x79, 0x2c, 0xaf, 0x23, 0x9d, 0xf1, 0xa6, 0xd1, 0xe6, 0x18, 0x9d, 0x9b, 0x69, 0xd4, 0x30,
	0x0a, 0x24, 0x6d, 0xf5, 0xca, 0x3a, 0xb7, 0x9c, 0x20, 0x14, 0x2a, 0x8a, 0x8f, 0x25, 0xfd, 0xaf,
	0x57, 0xee, 0x37, 0x0f, 0x2a, 0xec, 0x8a, 0x8f, 0x3d, 0xc3, 0xe8, 0x73, 0x18, 0x0e, 0x93, 0x49,
	0xac, 0xd0, 0xa7, 0xa4, 0xe7, 0xf4, 0xeb, 0xde, 0x8c, 0x20, 0x1e, 0x90, 0x0c, 0x88, 0x28, 0xbc,
	0x44, 0xa5, 0x44, 0x38, 0x96, 0x74, 0xcb, 0xb8, 0xb8, 0x2c, 0x9f, 0x07, 0x76, 0xb6, 0xb4, 0xe9,
	0x2c, 0x54, 0xc9, 0xc4, 0x2b, 0x38, 0x4d, 0x0e, 0xa1, 0x7e, 0x8b, 0x8a, 0xfb, 0x5c, 0x71, 0xda,
	0xee, 0x39, 0xfd, 0xe6, 0x41, 0x87, 0xd9, 0xd1, 0x63, 0xd3, 0xd1, 0x63, 0x97, 0x66, 0xf4, 0xbc,
	0x7c, 0x63, 0xf7, 0x0c, 0x3a, 0x0f, 0xdc, 0x41, 0x5a, 0x50, 0xbe, 0xc1, 0x49, 0x36, 0x8c, 0x7a,
	0x49, 0xda, 0x50, 0xbd, 0xe3, 0x41, 0x8a, 0xd9, 0x24, 0x5a, 0xf0, 0xb2, 0xf4, 0xc2, 0x71, 0x9f,
	0x43, 0xf9, 0x8a, 0x8f, 0x1f, 0x7b, 0xc4, 0xfd, 0xe5, 0x40, 0xfb, 0x83, 0x90, 0x2a, 0x4f, 0xd4,
	0xc3, 0xaf, 0x29, 0x4a, 0xa5, 0xb7, 0x07, 0xe2, 0x56, 0x28, 0x63, 0x51, 0xf5, 0x2c, 0xd0, 0xbd,
	0x8b, 0x46, 0x23, 0x89, 0xf6, 0x09, 0x54, 0xbd, 0x0c, 0xe9, 0x59, 0x95, 0x51, 0xa2, 0xde, 0xe3,
	0x44, 0xd2, 0xb2, 0x69, 0x4d, 0x8e, 0xa7, 0xda, 0x5b, 0x91, 0x48, 0x5a, 0x99, 0x69, 0x1a, 0x93,
	0x23, 0xa8, 0x9d, 0x8b, 0x40, 0x61, 0x42, 0xab, 0xa6, 0xe2, 0xcf, 0x58, 0x51, 0x30, 0xcc, 0xee,
	0xb1, 0x05, 0xcf, 0x0e, 0x74, 0x8f, 0xa0, 0x39, 0x47, 0x3f, 0xa9, 0x46, 0x9f, 0x61, 0xfb, 0xde,
	0x35, 0x32, 0x8e, 0x42, 0x89, 0x64, 0xcf, 0x7e, 0x17, 0x52, 0x93, 0x92, 0x3a, 0x26, 0x24, 0x98,
	0x0d, 0x81, 0x37, 0xa7, 0x9a, 0x47, 0x8d, 0xdf, 0xa6, 0x85, 0x30, 0x6b, 0xf7, 0x23, 0x6c, 0xbd,
	0xc3, 0xe5, 0x5a, 0xde, 0xff, 0x4b, 0xfa, 0xd0, 0xc8, 0x8d, 0xcc, 0xf9, 0xc5, 0x5b, 0x66, 0xa2,
	0x7b, 0x0c, 0xed, 0x45, 0xc3, 0x2c, 0xd0, 0x05, 0x07, 0xe7, 0x4f, 0x0e, 0x6f, 0x60, 0xe7, 0xd4,
	0x7c, 0x53, 0x4b, 0x51, 0x3d, 0xde, 0xe3, 0x14, 0x3a, 0x4b, 0x1e, 0x4f, 0x0d, 0xe4, 0xe0, 0x67,
	0x09, 0x2a, 0x5a, 0x20, 0xc7, 0xb0, 0xbe, 0x50, 0x7d, 0xb2, 0x5d, 0xd8, 0xf4, 0xee, 0x0e, 0x2b,
	0x6c, 0x92, 0xbb, 0x42, 0x5e, 0xc1, 0xda, 0x7c, 0x55, 0x48, 0x9b, 0x15, 0x54, 0xbd, 0xbb, 0xcd,
	0x8a, 0x4a, 0xe7, 0xae, 0x90, 0x73, 0xd8, 0xbc, 0x97, 0x0e, 0xe9, 0xb0, 0xe2, 0x22, 0x75, 0x29,
	0x7b, 0x20, 0x73, 0x77, 0x85, 0x9c, 0x42, 0xeb, 0x53, 0x1a, 0x04, 0x27, 0x41, 0xf0, 0x0f, 0xb9,
	0xbc, 0x86, 0x75, 0x6d, 0xf2, 0xb7, 0xc9, 0x0c, 0x6a, 0xe6, 0x47, 0x39, 0xfc, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xeb, 0xa6, 0xa7, 0x7c, 0xe8, 0x06, 0x00, 0x00,
}
