// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.2
// source: github.com/libp2p/go-libp2p-kad-dht/pb/dht.proto

package dht_pb

import (
	pb "github.com/libp2p/go-libp2p-record/pb"
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

type Message_MessageType int32

const (
	Message_PUT_VALUE     Message_MessageType = 0
	Message_GET_VALUE     Message_MessageType = 1
	Message_ADD_PROVIDER  Message_MessageType = 2
	Message_GET_PROVIDERS Message_MessageType = 3
	Message_FIND_NODE     Message_MessageType = 4
	Message_PING          Message_MessageType = 5
)

// Enum value maps for Message_MessageType.
var (
	Message_MessageType_name = map[int32]string{
		0: "PUT_VALUE",
		1: "GET_VALUE",
		2: "ADD_PROVIDER",
		3: "GET_PROVIDERS",
		4: "FIND_NODE",
		5: "PING",
	}
	Message_MessageType_value = map[string]int32{
		"PUT_VALUE":     0,
		"GET_VALUE":     1,
		"ADD_PROVIDER":  2,
		"GET_PROVIDERS": 3,
		"FIND_NODE":     4,
		"PING":          5,
	}
)

func (x Message_MessageType) Enum() *Message_MessageType {
	p := new(Message_MessageType)
	*p = x
	return p
}

func (x Message_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_enumTypes[0].Descriptor()
}

func (Message_MessageType) Type() protoreflect.EnumType {
	return &file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_enumTypes[0]
}

func (x Message_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_MessageType.Descriptor instead.
func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescGZIP(), []int{0, 0}
}

type Message_ConnectionType int32

const (
	// sender does not have a connection to peer, and no extra information
	// (default)
	Message_NOT_CONNECTED Message_ConnectionType = 0
	// sender has a live connection to peer
	Message_CONNECTED Message_ConnectionType = 1
	// sender recently connected to peer
	Message_CAN_CONNECT Message_ConnectionType = 2
	// sender recently tried to connect to peer repeatedly but failed to connect
	// ("try" here is loose, but this should signal "made strong effort,
	// failed")
	Message_CANNOT_CONNECT Message_ConnectionType = 3
)

// Enum value maps for Message_ConnectionType.
var (
	Message_ConnectionType_name = map[int32]string{
		0: "NOT_CONNECTED",
		1: "CONNECTED",
		2: "CAN_CONNECT",
		3: "CANNOT_CONNECT",
	}
	Message_ConnectionType_value = map[string]int32{
		"NOT_CONNECTED":  0,
		"CONNECTED":      1,
		"CAN_CONNECT":    2,
		"CANNOT_CONNECT": 3,
	}
)

func (x Message_ConnectionType) Enum() *Message_ConnectionType {
	p := new(Message_ConnectionType)
	*p = x
	return p
}

func (x Message_ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_enumTypes[1].Descriptor()
}

func (Message_ConnectionType) Type() protoreflect.EnumType {
	return &file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_enumTypes[1]
}

func (x Message_ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_ConnectionType.Descriptor instead.
func (Message_ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescGZIP(), []int{0, 1}
}

type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// defines what type of message it is.
	Type Message_MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=dht.pb.Message_MessageType" json:"type,omitempty"`
	// defines what coral cluster level this query/response belongs to.
	// in case we want to implement coral's cluster rings in the future.
	ClusterLevelRaw int32 `protobuf:"varint,10,opt,name=clusterLevelRaw,proto3" json:"clusterLevelRaw,omitempty"`
	// Used to specify the key associated with this message.
	// PUT_VALUE, GET_VALUE, ADD_PROVIDER, GET_PROVIDERS
	Key []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Used to return a value
	// PUT_VALUE, GET_VALUE
	Record *pb.Record `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`
	// Used to return peers closer to a key in a query
	// GET_VALUE, GET_PROVIDERS, FIND_NODE
	CloserPeers []*Message_Peer `protobuf:"bytes,8,rep,name=closerPeers,proto3" json:"closerPeers,omitempty"`
	// Used to return Providers
	// GET_VALUE, ADD_PROVIDER, GET_PROVIDERS
	ProviderPeers []*Message_Peer `protobuf:"bytes,9,rep,name=providerPeers,proto3" json:"providerPeers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() Message_MessageType {
	if x != nil {
		return x.Type
	}
	return Message_PUT_VALUE
}

func (x *Message) GetClusterLevelRaw() int32 {
	if x != nil {
		return x.ClusterLevelRaw
	}
	return 0
}

func (x *Message) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Message) GetRecord() *pb.Record {
	if x != nil {
		return x.Record
	}
	return nil
}

func (x *Message) GetCloserPeers() []*Message_Peer {
	if x != nil {
		return x.CloserPeers
	}
	return nil
}

func (x *Message) GetProviderPeers() []*Message_Peer {
	if x != nil {
		return x.ProviderPeers
	}
	return nil
}

type Message_Peer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of a given peer.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// multiaddrs for a given peer
	Addrs [][]byte `protobuf:"bytes,2,rep,name=addrs,proto3" json:"addrs,omitempty"`
	// used to signal the sender's connection capabilities to the peer
	Connection    Message_ConnectionType `protobuf:"varint,3,opt,name=connection,proto3,enum=dht.pb.Message_ConnectionType" json:"connection,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message_Peer) Reset() {
	*x = Message_Peer{}
	mi := &file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message_Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_Peer) ProtoMessage() {}

func (x *Message_Peer) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_Peer.ProtoReflect.Descriptor instead.
func (*Message_Peer) Descriptor() ([]byte, []int) {
	return file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Message_Peer) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Message_Peer) GetAddrs() [][]byte {
	if x != nil {
		return x.Addrs
	}
	return nil
}

func (x *Message_Peer) GetConnection() Message_ConnectionType {
	if x != nil {
		return x.Connection
	}
	return Message_NOT_CONNECTED
}

var File_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto protoreflect.FileDescriptor

var file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x62,
	0x70, 0x32, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x2d, 0x6b, 0x61,
	0x64, 0x2d, 0x64, 0x68, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x64, 0x68, 0x74, 0x2e, 0x70, 0x62, 0x1a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x2d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7,
	0x04, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x64, 0x68, 0x74, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x61, 0x77, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x61, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x68, 0x74, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x0b, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x64, 0x68, 0x74, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x50, 0x65, 0x65, 0x72, 0x73, 0x1a, 0x6c, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x61,
	0x64, 0x64, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x64, 0x68, 0x74, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x45, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x44, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45,
	0x52, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49,
	0x44, 0x45, 0x52, 0x53, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x49, 0x4e, 0x44, 0x5f, 0x4e,
	0x4f, 0x44, 0x45, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x22,
	0x57, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x4e, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x03, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x2f, 0x67, 0x6f,
	0x2d, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x2d, 0x6b, 0x61, 0x64, 0x2d, 0x64, 0x68, 0x74, 0x2f,
	0x70, 0x62, 0x3b, 0x64, 0x68, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescOnce sync.Once
	file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescData = file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDesc
)

func file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescGZIP() []byte {
	file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescOnce.Do(func() {
		file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescData)
	})
	return file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDescData
}

var file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_goTypes = []any{
	(Message_MessageType)(0),    // 0: dht.pb.Message.MessageType
	(Message_ConnectionType)(0), // 1: dht.pb.Message.ConnectionType
	(*Message)(nil),             // 2: dht.pb.Message
	(*Message_Peer)(nil),        // 3: dht.pb.Message.Peer
	(*pb.Record)(nil),           // 4: record.pb.Record
}
var file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_depIdxs = []int32{
	0, // 0: dht.pb.Message.type:type_name -> dht.pb.Message.MessageType
	4, // 1: dht.pb.Message.record:type_name -> record.pb.Record
	3, // 2: dht.pb.Message.closerPeers:type_name -> dht.pb.Message.Peer
	3, // 3: dht.pb.Message.providerPeers:type_name -> dht.pb.Message.Peer
	1, // 4: dht.pb.Message.Peer.connection:type_name -> dht.pb.Message.ConnectionType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_init() }
func file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_init() {
	if File_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_goTypes,
		DependencyIndexes: file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_depIdxs,
		EnumInfos:         file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_enumTypes,
		MessageInfos:      file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_msgTypes,
	}.Build()
	File_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto = out.File
	file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_rawDesc = nil
	file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_goTypes = nil
	file_github_com_libp2p_go_libp2p_kad_dht_pb_dht_proto_depIdxs = nil
}
