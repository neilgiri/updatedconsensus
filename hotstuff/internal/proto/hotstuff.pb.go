// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: internal/proto/hotstuff.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/relab/gorums"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BlockHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *BlockHash) Reset() {
	*x = BlockHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_hotstuff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHash) ProtoMessage() {}

func (x *BlockHash) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_hotstuff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHash.ProtoReflect.Descriptor instead.
func (*BlockHash) Descriptor() ([]byte, []int) {
	return file_internal_proto_hotstuff_proto_rawDescGZIP(), []int{0}
}

func (x *BlockHash) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent   []byte      `protobuf:"bytes,1,opt,name=Parent,proto3" json:"Parent,omitempty"`
	QC       *QuorumCert `protobuf:"bytes,2,opt,name=QC,proto3" json:"QC,omitempty"`
	View     uint64      `protobuf:"varint,3,opt,name=View,proto3" json:"View,omitempty"`
	Command  []byte      `protobuf:"bytes,4,opt,name=Command,proto3" json:"Command,omitempty"`
	Proposer uint32      `protobuf:"varint,5,opt,name=Proposer,proto3" json:"Proposer,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_hotstuff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_hotstuff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_internal_proto_hotstuff_proto_rawDescGZIP(), []int{1}
}

func (x *Block) GetParent() []byte {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Block) GetQC() *QuorumCert {
	if x != nil {
		return x.QC
	}
	return nil
}

func (x *Block) GetView() uint64 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *Block) GetCommand() []byte {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *Block) GetProposer() uint32 {
	if x != nil {
		return x.Proposer
	}
	return 0
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplicaID uint32 `protobuf:"varint,1,opt,name=ReplicaID,proto3" json:"ReplicaID,omitempty"`
	R         []byte `protobuf:"bytes,2,opt,name=R,proto3" json:"R,omitempty"`
	S         []byte `protobuf:"bytes,3,opt,name=S,proto3" json:"S,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_hotstuff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_hotstuff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_internal_proto_hotstuff_proto_rawDescGZIP(), []int{2}
}

func (x *Signature) GetReplicaID() uint32 {
	if x != nil {
		return x.ReplicaID
	}
	return 0
}

func (x *Signature) GetR() []byte {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *Signature) GetS() []byte {
	if x != nil {
		return x.S
	}
	return nil
}

type PartialCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sig  *Signature `protobuf:"bytes,1,opt,name=Sig,proto3" json:"Sig,omitempty"`
	Hash []byte     `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *PartialCert) Reset() {
	*x = PartialCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_hotstuff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialCert) ProtoMessage() {}

func (x *PartialCert) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_hotstuff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialCert.ProtoReflect.Descriptor instead.
func (*PartialCert) Descriptor() ([]byte, []int) {
	return file_internal_proto_hotstuff_proto_rawDescGZIP(), []int{3}
}

func (x *PartialCert) GetSig() *Signature {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *PartialCert) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type QuorumCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sigs map[uint32]*Signature `protobuf:"bytes,1,rep,name=Sigs,proto3" json:"Sigs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hash []byte                `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
}

func (x *QuorumCert) Reset() {
	*x = QuorumCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_hotstuff_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuorumCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuorumCert) ProtoMessage() {}

func (x *QuorumCert) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_hotstuff_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuorumCert.ProtoReflect.Descriptor instead.
func (*QuorumCert) Descriptor() ([]byte, []int) {
	return file_internal_proto_hotstuff_proto_rawDescGZIP(), []int{4}
}

func (x *QuorumCert) GetSigs() map[uint32]*Signature {
	if x != nil {
		return x.Sigs
	}
	return nil
}

func (x *QuorumCert) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type NewViewMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View uint64      `protobuf:"varint,1,opt,name=View,proto3" json:"View,omitempty"`
	QC   *QuorumCert `protobuf:"bytes,2,opt,name=QC,proto3" json:"QC,omitempty"`
}

func (x *NewViewMsg) Reset() {
	*x = NewViewMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_hotstuff_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewViewMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewViewMsg) ProtoMessage() {}

func (x *NewViewMsg) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_hotstuff_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewViewMsg.ProtoReflect.Descriptor instead.
func (*NewViewMsg) Descriptor() ([]byte, []int) {
	return file_internal_proto_hotstuff_proto_rawDescGZIP(), []int{5}
}

func (x *NewViewMsg) GetView() uint64 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *NewViewMsg) GetQC() *QuorumCert {
	if x != nil {
		return x.QC
	}
	return nil
}

var File_internal_proto_hotstuff_proto protoreflect.FileDescriptor

var file_internal_proto_hotstuff_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x68, 0x6f, 0x74, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6f, 0x72, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1f, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61,
	0x73, 0x68, 0x22, 0x8c, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x02, 0x51, 0x43, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43,
	0x65, 0x72, 0x74, 0x52, 0x02, 0x51, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x22, 0x45, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x44, 0x12, 0x0c, 0x0a, 0x01,
	0x52, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x52, 0x12, 0x0c, 0x0a, 0x01, 0x53, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x53, 0x22, 0x45, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x03, 0x53, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x03, 0x53, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x22,
	0x9c, 0x01, 0x0a, 0x0a, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x12, 0x2f,
	0x0a, 0x04, 0x53, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x2e,
	0x53, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x53, 0x69, 0x67, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x1a, 0x49, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x43,
	0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x56, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x56, 0x69, 0x65, 0x77,
	0x12, 0x21, 0x0a, 0x02, 0x51, 0x43, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x52,
	0x02, 0x51, 0x43, 0x32, 0xa7, 0x02, 0x0a, 0x08, 0x48, 0x6f, 0x74, 0x73, 0x74, 0x75, 0x66, 0x66,
	0x12, 0x35, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x04, 0x98, 0xb5, 0x18, 0x01, 0x12, 0x38, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x43,
	0x65, 0x72, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x90, 0xb5, 0x18,
	0x01, 0x12, 0x3a, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x56, 0x69, 0x65, 0x77, 0x12, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x56, 0x69, 0x65, 0x77, 0x4d, 0x73, 0x67, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x90, 0xb5, 0x18, 0x01, 0x12, 0x37, 0x0a,
	0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x04, 0x98, 0xb5, 0x18, 0x01, 0x12, 0x35, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0x90, 0xb5, 0x18, 0x01, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6c, 0x61,
	0x62, 0x2f, 0x68, 0x6f, 0x74, 0x73, 0x74, 0x75, 0x66, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_proto_hotstuff_proto_rawDescOnce sync.Once
	file_internal_proto_hotstuff_proto_rawDescData = file_internal_proto_hotstuff_proto_rawDesc
)

func file_internal_proto_hotstuff_proto_rawDescGZIP() []byte {
	file_internal_proto_hotstuff_proto_rawDescOnce.Do(func() {
		file_internal_proto_hotstuff_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_hotstuff_proto_rawDescData)
	})
	return file_internal_proto_hotstuff_proto_rawDescData
}

var file_internal_proto_hotstuff_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_proto_hotstuff_proto_goTypes = []interface{}{
	(*BlockHash)(nil),     // 0: proto.BlockHash
	(*Block)(nil),         // 1: proto.Block
	(*Signature)(nil),     // 2: proto.Signature
	(*PartialCert)(nil),   // 3: proto.PartialCert
	(*QuorumCert)(nil),    // 4: proto.QuorumCert
	(*NewViewMsg)(nil),    // 5: proto.NewViewMsg
	nil,                   // 6: proto.QuorumCert.SigsEntry
	(*emptypb.Empty)(nil), // 7: google.protobuf.Empty
}
var file_internal_proto_hotstuff_proto_depIdxs = []int32{
	4,  // 0: proto.Block.QC:type_name -> proto.QuorumCert
	2,  // 1: proto.PartialCert.Sig:type_name -> proto.Signature
	6,  // 2: proto.QuorumCert.Sigs:type_name -> proto.QuorumCert.SigsEntry
	4,  // 3: proto.NewViewMsg.QC:type_name -> proto.QuorumCert
	2,  // 4: proto.QuorumCert.SigsEntry.value:type_name -> proto.Signature
	1,  // 5: proto.Hotstuff.Propose:input_type -> proto.Block
	3,  // 6: proto.Hotstuff.Vote:input_type -> proto.PartialCert
	5,  // 7: proto.Hotstuff.NewView:input_type -> proto.NewViewMsg
	0,  // 8: proto.Hotstuff.Fetch:input_type -> proto.BlockHash
	1,  // 9: proto.Hotstuff.Deliver:input_type -> proto.Block
	7,  // 10: proto.Hotstuff.Propose:output_type -> google.protobuf.Empty
	7,  // 11: proto.Hotstuff.Vote:output_type -> google.protobuf.Empty
	7,  // 12: proto.Hotstuff.NewView:output_type -> google.protobuf.Empty
	7,  // 13: proto.Hotstuff.Fetch:output_type -> google.protobuf.Empty
	7,  // 14: proto.Hotstuff.Deliver:output_type -> google.protobuf.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_internal_proto_hotstuff_proto_init() }
func file_internal_proto_hotstuff_proto_init() {
	if File_internal_proto_hotstuff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_hotstuff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHash); i {
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
		file_internal_proto_hotstuff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_internal_proto_hotstuff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_internal_proto_hotstuff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartialCert); i {
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
		file_internal_proto_hotstuff_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuorumCert); i {
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
		file_internal_proto_hotstuff_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewViewMsg); i {
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
			RawDescriptor: file_internal_proto_hotstuff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_hotstuff_proto_goTypes,
		DependencyIndexes: file_internal_proto_hotstuff_proto_depIdxs,
		MessageInfos:      file_internal_proto_hotstuff_proto_msgTypes,
	}.Build()
	File_internal_proto_hotstuff_proto = out.File
	file_internal_proto_hotstuff_proto_rawDesc = nil
	file_internal_proto_hotstuff_proto_goTypes = nil
	file_internal_proto_hotstuff_proto_depIdxs = nil
}