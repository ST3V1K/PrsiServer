// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: game.proto

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

type Card_Color int32

const (
	Card_Hearts   Card_Color = 0
	Card_Spades   Card_Color = 1
	Card_Diamonds Card_Color = 2
	Card_Clubs    Card_Color = 3
)

// Enum value maps for Card_Color.
var (
	Card_Color_name = map[int32]string{
		0: "Hearts",
		1: "Spades",
		2: "Diamonds",
		3: "Clubs",
	}
	Card_Color_value = map[string]int32{
		"Hearts":   0,
		"Spades":   1,
		"Diamonds": 2,
		"Clubs":    3,
	}
)

func (x Card_Color) Enum() *Card_Color {
	p := new(Card_Color)
	*p = x
	return p
}

func (x Card_Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Card_Color) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[0].Descriptor()
}

func (Card_Color) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x Card_Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Card_Color.Descriptor instead.
func (Card_Color) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{7, 0}
}

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
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
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
	return file_game_proto_rawDescGZIP(), []int{0}
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
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
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
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *Game) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Game   *Game   `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	Player *Player `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *GameRequest) Reset() {
	*x = GameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRequest) ProtoMessage() {}

func (x *GameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRequest.ProtoReflect.Descriptor instead.
func (*GameRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *GameRequest) GetGame() *Game {
	if x != nil {
		return x.Game
	}
	return nil
}

func (x *GameRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type FilteredGamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Filter string  `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *FilteredGamesRequest) Reset() {
	*x = FilteredGamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilteredGamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilteredGamesRequest) ProtoMessage() {}

func (x *FilteredGamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilteredGamesRequest.ProtoReflect.Descriptor instead.
func (*FilteredGamesRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *FilteredGamesRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

func (x *FilteredGamesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type GameStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string  `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Seed         int32   `protobuf:"varint,3,opt,name=seed,proto3" json:"seed,omitempty"`
	OpponentName *string `protobuf:"bytes,2,opt,name=opponentName,proto3,oneof" json:"opponentName,omitempty"`
	Card         *Card   `protobuf:"bytes,4,opt,name=card,proto3,oneof" json:"card,omitempty"`
	Draw         *int32  `protobuf:"varint,5,opt,name=draw,proto3,oneof" json:"draw,omitempty"`
	Tie          *bool   `protobuf:"varint,6,opt,name=tie,proto3,oneof" json:"tie,omitempty"`
}

func (x *GameStream) Reset() {
	*x = GameStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStream) ProtoMessage() {}

func (x *GameStream) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStream.ProtoReflect.Descriptor instead.
func (*GameStream) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *GameStream) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *GameStream) GetSeed() int32 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *GameStream) GetOpponentName() string {
	if x != nil && x.OpponentName != nil {
		return *x.OpponentName
	}
	return ""
}

func (x *GameStream) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

func (x *GameStream) GetDraw() int32 {
	if x != nil && x.Draw != nil {
		return *x.Draw
	}
	return 0
}

func (x *GameStream) GetTie() bool {
	if x != nil && x.Tie != nil {
		return *x.Tie
	}
	return false
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *SuccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListGamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameUuids   []string `protobuf:"bytes,1,rep,name=gameUuids,proto3" json:"gameUuids,omitempty"`
	PlayerNames []string `protobuf:"bytes,2,rep,name=playerNames,proto3" json:"playerNames,omitempty"`
}

func (x *ListGamesResponse) Reset() {
	*x = ListGamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGamesResponse) ProtoMessage() {}

func (x *ListGamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[6]
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
	return file_game_proto_rawDescGZIP(), []int{6}
}

func (x *ListGamesResponse) GetGameUuids() []string {
	if x != nil {
		return x.GameUuids
	}
	return nil
}

func (x *ListGamesResponse) GetPlayerNames() []string {
	if x != nil {
		return x.PlayerNames
	}
	return nil
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color    Card_Color  `protobuf:"varint,1,opt,name=color,proto3,enum=server.Card_Color" json:"color,omitempty"`
	Value    int32       `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	ChangeTo *Card_Color `protobuf:"varint,3,opt,name=changeTo,proto3,enum=server.Card_Color,oneof" json:"changeTo,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{7}
}

func (x *Card) GetColor() Card_Color {
	if x != nil {
		return x.Color
	}
	return Card_Hearts
}

func (x *Card) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Card) GetChangeTo() Card_Color {
	if x != nil && x.ChangeTo != nil {
		return *x.ChangeTo
	}
	return Card_Hearts
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1a,
	0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x0b, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x67, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xdf, 0x01, 0x0a, 0x0a,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x65,
	0x65, 0x64, 0x12, 0x27, 0x0a, 0x0c, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x70, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x04, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x48, 0x01, 0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x72, 0x61, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x02, 0x52, 0x04, 0x64, 0x72, 0x61, 0x77, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x74,
	0x69, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x03, 0x74, 0x69, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x64, 0x72, 0x61, 0x77, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x69, 0x65, 0x22, 0x2b, 0x0a,
	0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x53, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x55, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x55, 0x75, 0x69, 0x64, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0xc2, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x88, 0x01, 0x01, 0x22, 0x38, 0x0a,
	0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x72, 0x74, 0x73,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x70, 0x61, 0x64, 0x65, 0x73, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x73, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x43, 0x6c, 0x75, 0x62, 0x73, 0x10, 0x03, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x6f, 0x32, 0xa6, 0x03, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x1a,
	0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47,
	0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x65, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a,
	0x08, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0xaa, 0x02, 0x0b, 0x50, 0x72, 0x73, 0x69,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_game_proto_goTypes = []interface{}{
	(Card_Color)(0),              // 0: server.Card.Color
	(*Player)(nil),               // 1: server.Player
	(*Game)(nil),                 // 2: server.Game
	(*GameRequest)(nil),          // 3: server.GameRequest
	(*FilteredGamesRequest)(nil), // 4: server.FilteredGamesRequest
	(*GameStream)(nil),           // 5: server.GameStream
	(*SuccessResponse)(nil),      // 6: server.SuccessResponse
	(*ListGamesResponse)(nil),    // 7: server.ListGamesResponse
	(*Card)(nil),                 // 8: server.Card
}
var file_game_proto_depIdxs = []int32{
	2,  // 0: server.GameRequest.game:type_name -> server.Game
	1,  // 1: server.GameRequest.player:type_name -> server.Player
	1,  // 2: server.FilteredGamesRequest.player:type_name -> server.Player
	8,  // 3: server.GameStream.card:type_name -> server.Card
	0,  // 4: server.Card.color:type_name -> server.Card.Color
	0,  // 5: server.Card.changeTo:type_name -> server.Card.Color
	1,  // 6: server.GameService.NewGame:input_type -> server.Player
	3,  // 7: server.GameService.RemoveGame:input_type -> server.GameRequest
	3,  // 8: server.GameService.Join:input_type -> server.GameRequest
	3,  // 9: server.GameService.Leave:input_type -> server.GameRequest
	3,  // 10: server.GameService.RequestTie:input_type -> server.GameRequest
	1,  // 11: server.GameService.ListGames:input_type -> server.Player
	4,  // 12: server.GameService.ListGamesFiltered:input_type -> server.FilteredGamesRequest
	5,  // 13: server.GameService.NewGame:output_type -> server.GameStream
	6,  // 14: server.GameService.RemoveGame:output_type -> server.SuccessResponse
	5,  // 15: server.GameService.Join:output_type -> server.GameStream
	6,  // 16: server.GameService.Leave:output_type -> server.SuccessResponse
	6,  // 17: server.GameService.RequestTie:output_type -> server.SuccessResponse
	7,  // 18: server.GameService.ListGames:output_type -> server.ListGamesResponse
	7,  // 19: server.GameService.ListGamesFiltered:output_type -> server.ListGamesResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRequest); i {
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
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilteredGamesRequest); i {
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
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStream); i {
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
		file_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
	file_game_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_game_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		EnumInfos:         file_game_proto_enumTypes,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
