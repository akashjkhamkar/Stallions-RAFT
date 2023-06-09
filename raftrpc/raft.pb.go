// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.3
// source: raftrpc/raft.proto

package raftrpc

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

// Vote req
type RequestVoteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Server       int32 `protobuf:"varint,2,opt,name=server,proto3" json:"server,omitempty"`
	Lastlogindex int32 `protobuf:"varint,3,opt,name=lastlogindex,proto3" json:"lastlogindex,omitempty"`
	Lastlogterm  int32 `protobuf:"varint,4,opt,name=lastlogterm,proto3" json:"lastlogterm,omitempty"`
}

func (x *RequestVoteArgs) Reset() {
	*x = RequestVoteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftrpc_raft_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteArgs) ProtoMessage() {}

func (x *RequestVoteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_raftrpc_raft_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteArgs.ProtoReflect.Descriptor instead.
func (*RequestVoteArgs) Descriptor() ([]byte, []int) {
	return file_raftrpc_raft_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVoteArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteArgs) GetServer() int32 {
	if x != nil {
		return x.Server
	}
	return 0
}

func (x *RequestVoteArgs) GetLastlogindex() int32 {
	if x != nil {
		return x.Lastlogindex
	}
	return 0
}

func (x *RequestVoteArgs) GetLastlogterm() int32 {
	if x != nil {
		return x.Lastlogterm
	}
	return 0
}

// Vote res
type RequestVoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Vote bool  `protobuf:"varint,2,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (x *RequestVoteReply) Reset() {
	*x = RequestVoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftrpc_raft_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteReply) ProtoMessage() {}

func (x *RequestVoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_raftrpc_raft_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteReply.ProtoReflect.Descriptor instead.
func (*RequestVoteReply) Descriptor() ([]byte, []int) {
	return file_raftrpc_raft_proto_rawDescGZIP(), []int{1}
}

func (x *RequestVoteReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteReply) GetVote() bool {
	if x != nil {
		return x.Vote
	}
	return false
}

type HeartBeatArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int32       `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Id           int32       `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Prevlogindex int32       `protobuf:"varint,3,opt,name=prevlogindex,proto3" json:"prevlogindex,omitempty"`
	Prevlogterm  int32       `protobuf:"varint,4,opt,name=prevlogterm,proto3" json:"prevlogterm,omitempty"`
	Entries      []*LogEntry `protobuf:"bytes,5,rep,name=entries,proto3" json:"entries,omitempty"`
	Leadercommit int32       `protobuf:"varint,6,opt,name=leadercommit,proto3" json:"leadercommit,omitempty"`
}

func (x *HeartBeatArgs) Reset() {
	*x = HeartBeatArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftrpc_raft_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatArgs) ProtoMessage() {}

func (x *HeartBeatArgs) ProtoReflect() protoreflect.Message {
	mi := &file_raftrpc_raft_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatArgs.ProtoReflect.Descriptor instead.
func (*HeartBeatArgs) Descriptor() ([]byte, []int) {
	return file_raftrpc_raft_proto_rawDescGZIP(), []int{2}
}

func (x *HeartBeatArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *HeartBeatArgs) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HeartBeatArgs) GetPrevlogindex() int32 {
	if x != nil {
		return x.Prevlogindex
	}
	return 0
}

func (x *HeartBeatArgs) GetPrevlogterm() int32 {
	if x != nil {
		return x.Prevlogterm
	}
	return 0
}

func (x *HeartBeatArgs) GetEntries() []*LogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *HeartBeatArgs) GetLeadercommit() int32 {
	if x != nil {
		return x.Leadercommit
	}
	return 0
}

type HeartBeatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term      int32 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success   bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Jumpindex int32 `protobuf:"varint,3,opt,name=jumpindex,proto3" json:"jumpindex,omitempty"`
}

func (x *HeartBeatReply) Reset() {
	*x = HeartBeatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftrpc_raft_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatReply) ProtoMessage() {}

func (x *HeartBeatReply) ProtoReflect() protoreflect.Message {
	mi := &file_raftrpc_raft_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatReply.ProtoReflect.Descriptor instead.
func (*HeartBeatReply) Descriptor() ([]byte, []int) {
	return file_raftrpc_raft_proto_rawDescGZIP(), []int{3}
}

func (x *HeartBeatReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *HeartBeatReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *HeartBeatReply) GetJumpindex() int32 {
	if x != nil {
		return x.Jumpindex
	}
	return 0
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Term    int32  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftrpc_raft_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_raftrpc_raft_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_raftrpc_raft_proto_rawDescGZIP(), []int{4}
}

func (x *LogEntry) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *LogEntry) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

type ApplyMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commandvalid bool   `protobuf:"varint,1,opt,name=commandvalid,proto3" json:"commandvalid,omitempty"`
	Command      string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Commandindex int32  `protobuf:"varint,3,opt,name=commandindex,proto3" json:"commandindex,omitempty"`
}

func (x *ApplyMsg) Reset() {
	*x = ApplyMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raftrpc_raft_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyMsg) ProtoMessage() {}

func (x *ApplyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_raftrpc_raft_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyMsg.ProtoReflect.Descriptor instead.
func (*ApplyMsg) Descriptor() ([]byte, []int) {
	return file_raftrpc_raft_proto_rawDescGZIP(), []int{5}
}

func (x *ApplyMsg) GetCommandvalid() bool {
	if x != nil {
		return x.Commandvalid
	}
	return false
}

func (x *ApplyMsg) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ApplyMsg) GetCommandindex() int32 {
	if x != nil {
		return x.Commandindex
	}
	return 0
}

var File_raftrpc_raft_proto protoreflect.FileDescriptor

var file_raftrpc_raft_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x61, 0x66, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x61, 0x66, 0x74, 0x72, 0x70, 0x63, 0x22, 0x83, 0x01,
	0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x67,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x6c, 0x6f, 0x67, 0x74, 0x65, 0x72, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x6c, 0x6f, 0x67, 0x74,
	0x65, 0x72, 0x6d, 0x22, 0x3a, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x76,
	0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x22,
	0xca, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x41, 0x72, 0x67,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x65,
	0x76, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x65,
	0x76, 0x6c, 0x6f, 0x67, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x70, 0x72, 0x65, 0x76, 0x6c, 0x6f, 0x67, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x2b, 0x0a, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72,
	0x61, 0x66, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x5c, 0x0a, 0x0e,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x6a, 0x75, 0x6d, 0x70, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6a, 0x75, 0x6d, 0x70, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x38, 0x0a, 0x08, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x65, 0x72, 0x6d, 0x22, 0x6c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4d, 0x73, 0x67,
	0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x32, 0x9d, 0x01, 0x0a, 0x07, 0x52, 0x61, 0x66, 0x74, 0x52, 0x70, 0x63, 0x12, 0x4b,
	0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x19,
	0x2e, 0x72, 0x61, 0x66, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x10, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x17, 0x2e, 0x72, 0x61, 0x66, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x69, 0x6f, 0x6e, 0x72, 0x61,
	0x66, 0x74, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_raftrpc_raft_proto_rawDescOnce sync.Once
	file_raftrpc_raft_proto_rawDescData = file_raftrpc_raft_proto_rawDesc
)

func file_raftrpc_raft_proto_rawDescGZIP() []byte {
	file_raftrpc_raft_proto_rawDescOnce.Do(func() {
		file_raftrpc_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_raftrpc_raft_proto_rawDescData)
	})
	return file_raftrpc_raft_proto_rawDescData
}

var file_raftrpc_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_raftrpc_raft_proto_goTypes = []interface{}{
	(*RequestVoteArgs)(nil),  // 0: raftrpc.RequestVoteArgs
	(*RequestVoteReply)(nil), // 1: raftrpc.RequestVoteReply
	(*HeartBeatArgs)(nil),    // 2: raftrpc.HeartBeatArgs
	(*HeartBeatReply)(nil),   // 3: raftrpc.HeartBeatReply
	(*LogEntry)(nil),         // 4: raftrpc.LogEntry
	(*ApplyMsg)(nil),         // 5: raftrpc.ApplyMsg
}
var file_raftrpc_raft_proto_depIdxs = []int32{
	4, // 0: raftrpc.HeartBeatArgs.entries:type_name -> raftrpc.LogEntry
	0, // 1: raftrpc.RaftRpc.RequestVoteHandler:input_type -> raftrpc.RequestVoteArgs
	2, // 2: raftrpc.RaftRpc.HeartbeatHandler:input_type -> raftrpc.HeartBeatArgs
	1, // 3: raftrpc.RaftRpc.RequestVoteHandler:output_type -> raftrpc.RequestVoteReply
	3, // 4: raftrpc.RaftRpc.HeartbeatHandler:output_type -> raftrpc.HeartBeatReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_raftrpc_raft_proto_init() }
func file_raftrpc_raft_proto_init() {
	if File_raftrpc_raft_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raftrpc_raft_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteArgs); i {
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
		file_raftrpc_raft_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteReply); i {
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
		file_raftrpc_raft_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatArgs); i {
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
		file_raftrpc_raft_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatReply); i {
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
		file_raftrpc_raft_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
		file_raftrpc_raft_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyMsg); i {
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
			RawDescriptor: file_raftrpc_raft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raftrpc_raft_proto_goTypes,
		DependencyIndexes: file_raftrpc_raft_proto_depIdxs,
		MessageInfos:      file_raftrpc_raft_proto_msgTypes,
	}.Build()
	File_raftrpc_raft_proto = out.File
	file_raftrpc_raft_proto_rawDesc = nil
	file_raftrpc_raft_proto_goTypes = nil
	file_raftrpc_raft_proto_depIdxs = nil
}
