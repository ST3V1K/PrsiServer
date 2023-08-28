// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: grpc/server.proto

package server

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

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{1}
}

func (x *Game) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GameJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game   *Game   `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	Player *Player `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *GameJoinRequest) Reset() {
	*x = GameJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameJoinRequest) ProtoMessage() {}

func (x *GameJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameJoinRequest.ProtoReflect.Descriptor instead.
func (*GameJoinRequest) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{2}
}

func (x *GameJoinRequest) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

func (x *GameJoinRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type GameJoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Joined bool `protobuf:"varint,1,opt,name=joined,proto3" json:"joined,omitempty"`
}

func (x *GameJoinResponse) Reset() {
	*x = GameJoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameJoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameJoinResponse) ProtoMessage() {}

func (x *GameJoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameJoinResponse.ProtoReflect.Descriptor instead.
func (*GameJoinResponse) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{3}
}

func (x *GameJoinResponse) GetJoined() bool {
	if x != nil {
		return x.Joined
	}
	return false
}

type GameCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *GameCreateRequest) Reset() {
	*x = GameCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameCreateRequest) ProtoMessage() {}

func (x *GameCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameCreateRequest.ProtoReflect.Descriptor instead.
func (*GameCreateRequest) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{4}
}

func (x *GameCreateRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type GameCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GameCreateResponse) Reset() {
	*x = GameCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameCreateResponse) ProtoMessage() {}

func (x *GameCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameCreateResponse.ProtoReflect.Descriptor instead.
func (*GameCreateResponse) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{5}
}

func (x *GameCreateResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ListGamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *ListGamesRequest) Reset() {
	*x = ListGamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGamesRequest) ProtoMessage() {}

func (x *ListGamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGamesRequest.ProtoReflect.Descriptor instead.
func (*ListGamesRequest) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{6}
}

func (x *ListGamesRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type ListGamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games []*Game `protobuf:"bytes,1,rep,name=games,proto3" json:"games,omitempty"`
}

func (x *ListGamesResponse) Reset() {
	*x = ListGamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGamesResponse) ProtoMessage() {}

func (x *ListGamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGamesResponse.ProtoReflect.Descriptor instead.
func (*ListGamesResponse) Descriptor() ([]byte, []int) {
	return file_grpc_server_proto_rawDescGZIP(), []int{7}
}

func (x *ListGamesResponse) GetGames() []*Game {
	if x != nil {
		return x.Games
	}
	return nil
}

var File_grpc_server_proto protoreflect.FileDescriptor

var file_grpc_server_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x06, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1a, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0x5b, 0x0a, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x2a,
	0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x61,
	0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x61, 0x6d, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0x3a, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x37, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x05, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x32, 0xd2, 0x01, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4a, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_server_proto_rawDescOnce sync.Once
	file_grpc_server_proto_rawDescData = file_grpc_server_proto_rawDesc
)

func file_grpc_server_proto_rawDescGZIP() []byte {
	file_grpc_server_proto_rawDescOnce.Do(func() {
		file_grpc_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_server_proto_rawDescData)
	})
	return file_grpc_server_proto_rawDescData
}

var file_grpc_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_server_proto_goTypes = []interface{}{
	(*Player)(nil),             // 0: server.Player
	(*Game)(nil),               // 1: server.Game
	(*GameJoinRequest)(nil),    // 2: server.GameJoinRequest
	(*GameJoinResponse)(nil),   // 3: server.GameJoinResponse
	(*GameCreateRequest)(nil),  // 4: server.GameCreateRequest
	(*GameCreateResponse)(nil), // 5: server.GameCreateResponse
	(*ListGamesRequest)(nil),   // 6: server.ListGamesRequest
	(*ListGamesResponse)(nil),  // 7: server.ListGamesResponse
}
var file_grpc_server_proto_depIdxs = []int32{
	1, // 0: server.GameJoinRequest.game:type_name -> server.Game
	0, // 1: server.GameJoinRequest.player:type_name -> server.Player
	0, // 2: server.GameCreateRequest.player:type_name -> server.Player
	0, // 3: server.ListGamesRequest.player:type_name -> server.Player
	1, // 4: server.ListGamesResponse.games:type_name -> server.Game
	2, // 5: server.GameService.Join:input_type -> server.GameJoinRequest
	4, // 6: server.GameService.NewGame:input_type -> server.GameCreateRequest
	6, // 7: server.GameService.ListGames:input_type -> server.ListGamesRequest
	3, // 8: server.GameService.Join:output_type -> server.GameJoinResponse
	5, // 9: server.GameService.NewGame:output_type -> server.GameCreateResponse
	7, // 10: server.GameService.ListGames:output_type -> server.ListGamesResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_server_proto_init() }
func file_grpc_server_proto_init() {
	if File_grpc_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_grpc_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_grpc_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameJoinRequest); i {
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
		file_grpc_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameJoinResponse); i {
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
		file_grpc_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameCreateRequest); i {
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
		file_grpc_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameCreateResponse); i {
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
		file_grpc_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGamesRequest); i {
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
		file_grpc_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGamesResponse); i {
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
			RawDescriptor: file_grpc_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_server_proto_goTypes,
		DependencyIndexes: file_grpc_server_proto_depIdxs,
		MessageInfos:      file_grpc_server_proto_msgTypes,
	}.Build()
	File_grpc_server_proto = out.File
	file_grpc_server_proto_rawDesc = nil
	file_grpc_server_proto_goTypes = nil
	file_grpc_server_proto_depIdxs = nil
}
