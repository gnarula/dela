// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package cosipbft

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// Player is part of the change set of a forward link. It provides the
// information necessary to add a player.
type Player struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PublicKey            *any.Any `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Player) GetPublicKey() *any.Any {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

// ChangeSet contains the information about changes related to the authority for
// the next proposal. Leader does not affect the authority but can be used by
// the viewchange to decide on the next leader.
type ChangeSet struct {
	Remove               []uint32  `protobuf:"varint,1,rep,packed,name=remove,proto3" json:"remove,omitempty"`
	Add                  []*Player `protobuf:"bytes,2,rep,name=add,proto3" json:"add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ChangeSet) Reset()         { *m = ChangeSet{} }
func (m *ChangeSet) String() string { return proto.CompactTextString(m) }
func (*ChangeSet) ProtoMessage()    {}
func (*ChangeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *ChangeSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeSet.Unmarshal(m, b)
}
func (m *ChangeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeSet.Marshal(b, m, deterministic)
}
func (m *ChangeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeSet.Merge(m, src)
}
func (m *ChangeSet) XXX_Size() int {
	return xxx_messageInfo_ChangeSet.Size(m)
}
func (m *ChangeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeSet.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeSet proto.InternalMessageInfo

func (m *ChangeSet) GetRemove() []uint32 {
	if m != nil {
		return m.Remove
	}
	return nil
}

func (m *ChangeSet) GetAdd() []*Player {
	if m != nil {
		return m.Add
	}
	return nil
}

// ForwardLinkProto is the message representing a forward link between two
// proposals. It contains both hash and the prepare and commit signatures.
type ForwardLinkProto struct {
	From                 []byte     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   []byte     `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Prepare              *any.Any   `protobuf:"bytes,3,opt,name=prepare,proto3" json:"prepare,omitempty"`
	Commit               *any.Any   `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
	ChangeSet            *ChangeSet `protobuf:"bytes,5,opt,name=changeSet,proto3" json:"changeSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ForwardLinkProto) Reset()         { *m = ForwardLinkProto{} }
func (m *ForwardLinkProto) String() string { return proto.CompactTextString(m) }
func (*ForwardLinkProto) ProtoMessage()    {}
func (*ForwardLinkProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *ForwardLinkProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardLinkProto.Unmarshal(m, b)
}
func (m *ForwardLinkProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardLinkProto.Marshal(b, m, deterministic)
}
func (m *ForwardLinkProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardLinkProto.Merge(m, src)
}
func (m *ForwardLinkProto) XXX_Size() int {
	return xxx_messageInfo_ForwardLinkProto.Size(m)
}
func (m *ForwardLinkProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardLinkProto.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardLinkProto proto.InternalMessageInfo

func (m *ForwardLinkProto) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ForwardLinkProto) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *ForwardLinkProto) GetPrepare() *any.Any {
	if m != nil {
		return m.Prepare
	}
	return nil
}

func (m *ForwardLinkProto) GetCommit() *any.Any {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *ForwardLinkProto) GetChangeSet() *ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

// ChainProto is the message representing a list of forward links that creates
// a verifiable chain.
type ChainProto struct {
	Links                []*ForwardLinkProto `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ChainProto) Reset()         { *m = ChainProto{} }
func (m *ChainProto) String() string { return proto.CompactTextString(m) }
func (*ChainProto) ProtoMessage()    {}
func (*ChainProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *ChainProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainProto.Unmarshal(m, b)
}
func (m *ChainProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainProto.Marshal(b, m, deterministic)
}
func (m *ChainProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainProto.Merge(m, src)
}
func (m *ChainProto) XXX_Size() int {
	return xxx_messageInfo_ChainProto.Size(m)
}
func (m *ChainProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainProto.DiscardUnknown(m)
}

var xxx_messageInfo_ChainProto proto.InternalMessageInfo

func (m *ChainProto) GetLinks() []*ForwardLinkProto {
	if m != nil {
		return m.Links
	}
	return nil
}

// PrepareRequest is the message sent to start a consensus for a proposal.
type PrepareRequest struct {
	Proposal *any.Any `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
	// Signature proves the proposal comes from an expected identity.
	Signature            *any.Any `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareRequest) Reset()         { *m = PrepareRequest{} }
func (m *PrepareRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareRequest) ProtoMessage()    {}
func (*PrepareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *PrepareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareRequest.Unmarshal(m, b)
}
func (m *PrepareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareRequest.Marshal(b, m, deterministic)
}
func (m *PrepareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareRequest.Merge(m, src)
}
func (m *PrepareRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareRequest.Size(m)
}
func (m *PrepareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareRequest proto.InternalMessageInfo

func (m *PrepareRequest) GetProposal() *any.Any {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func (m *PrepareRequest) GetSignature() *any.Any {
	if m != nil {
		return m.Signature
	}
	return nil
}

// CommitRequest is the message sent to commit to a proposal.
type CommitRequest struct {
	To                   []byte   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Prepare              *any.Any `protobuf:"bytes,2,opt,name=prepare,proto3" json:"prepare,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitRequest) Reset()         { *m = CommitRequest{} }
func (m *CommitRequest) String() string { return proto.CompactTextString(m) }
func (*CommitRequest) ProtoMessage()    {}
func (*CommitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{5}
}

func (m *CommitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitRequest.Unmarshal(m, b)
}
func (m *CommitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitRequest.Marshal(b, m, deterministic)
}
func (m *CommitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitRequest.Merge(m, src)
}
func (m *CommitRequest) XXX_Size() int {
	return xxx_messageInfo_CommitRequest.Size(m)
}
func (m *CommitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommitRequest proto.InternalMessageInfo

func (m *CommitRequest) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *CommitRequest) GetPrepare() *any.Any {
	if m != nil {
		return m.Prepare
	}
	return nil
}

// PropagateRequest is the last message of a consensus process to send the valid
// forward link to participants.
type PropagateRequest struct {
	To                   []byte   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Commit               *any.Any `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PropagateRequest) Reset()         { *m = PropagateRequest{} }
func (m *PropagateRequest) String() string { return proto.CompactTextString(m) }
func (*PropagateRequest) ProtoMessage()    {}
func (*PropagateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{6}
}

func (m *PropagateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PropagateRequest.Unmarshal(m, b)
}
func (m *PropagateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PropagateRequest.Marshal(b, m, deterministic)
}
func (m *PropagateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PropagateRequest.Merge(m, src)
}
func (m *PropagateRequest) XXX_Size() int {
	return xxx_messageInfo_PropagateRequest.Size(m)
}
func (m *PropagateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PropagateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PropagateRequest proto.InternalMessageInfo

func (m *PropagateRequest) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PropagateRequest) GetCommit() *any.Any {
	if m != nil {
		return m.Commit
	}
	return nil
}

func init() {
	proto.RegisterType((*Player)(nil), "cosipbft.Player")
	proto.RegisterType((*ChangeSet)(nil), "cosipbft.ChangeSet")
	proto.RegisterType((*ForwardLinkProto)(nil), "cosipbft.ForwardLinkProto")
	proto.RegisterType((*ChainProto)(nil), "cosipbft.ChainProto")
	proto.RegisterType((*PrepareRequest)(nil), "cosipbft.PrepareRequest")
	proto.RegisterType((*CommitRequest)(nil), "cosipbft.CommitRequest")
	proto.RegisterType((*PropagateRequest)(nil), "cosipbft.PropagateRequest")
}

func init() {
	proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5)
}

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x74, 0xeb, 0xd6, 0xd7, 0xae, 0xaa, 0x0c, 0x42, 0x66, 0xa7, 0xc8, 0xa7, 0x1c,
	0x90, 0x37, 0xc2, 0x11, 0x09, 0x09, 0x2a, 0xc1, 0x01, 0x24, 0x22, 0x23, 0x71, 0xe0, 0xe6, 0x24,
	0xaf, 0x59, 0xb4, 0x24, 0x36, 0xb6, 0x33, 0x94, 0xbf, 0x90, 0x7f, 0x0b, 0xcd, 0x69, 0x1a, 0x09,
	0xd1, 0xb2, 0x5b, 0x5e, 0xf2, 0x79, 0x2f, 0xdf, 0x1f, 0xb0, 0x6e, 0xd0, 0x5a, 0x59, 0xa2, 0xe5,
	0xda, 0x28, 0xa7, 0xc8, 0x65, 0xae, 0x6c, 0xa5, 0xb3, 0x9d, 0xbb, 0x7e, 0x59, 0x2a, 0x55, 0xd6,
	0x78, 0xe3, 0xdf, 0x67, 0xdd, 0xee, 0x46, 0xb6, 0xfd, 0x00, 0xb1, 0xef, 0x30, 0x4f, 0x6b, 0xd9,
	0xa3, 0x21, 0x14, 0x2e, 0x64, 0x51, 0x18, 0xb4, 0x96, 0x06, 0x51, 0x10, 0xaf, 0xc4, 0x38, 0x92,
	0x04, 0x16, 0xba, 0xcb, 0xea, 0x2a, 0xff, 0x8c, 0x3d, 0x0d, 0xa3, 0x20, 0x5e, 0x26, 0xcf, 0xf9,
	0x70, 0x92, 0x8f, 0x27, 0xf9, 0xfb, 0xb6, 0x17, 0x13, 0xc6, 0x3e, 0xc1, 0x62, 0x7b, 0x27, 0xdb,
	0x12, 0xbf, 0xa1, 0x23, 0x2f, 0x60, 0x6e, 0xb0, 0x51, 0x0f, 0x48, 0x83, 0x68, 0x16, 0x5f, 0x89,
	0xfd, 0x44, 0x18, 0xcc, 0x64, 0x51, 0xd0, 0x30, 0x9a, 0xc5, 0xcb, 0x64, 0xc3, 0x47, 0xbd, 0x7c,
	0x50, 0x24, 0x1e, 0x3f, 0xb2, 0xdf, 0x01, 0x6c, 0x3e, 0x2a, 0xf3, 0x4b, 0x9a, 0xe2, 0x4b, 0xd5,
	0xde, 0xa7, 0xde, 0x1a, 0x81, 0xb3, 0x9d, 0x51, 0xcd, 0x5e, 0xa8, 0x7f, 0x26, 0x6b, 0x08, 0x9d,
	0xf2, 0xf2, 0x56, 0x22, 0x74, 0x8a, 0x70, 0xb8, 0xd0, 0x06, 0xb5, 0x34, 0x48, 0x67, 0x27, 0x34,
	0x8f, 0x10, 0x79, 0x05, 0xf3, 0x5c, 0x35, 0x4d, 0xe5, 0xe8, 0xd9, 0x09, 0x7c, 0xcf, 0x90, 0xd7,
	0xb0, 0xc8, 0x47, 0x7f, 0xf4, 0xdc, 0x2f, 0x3c, 0x9b, 0x0c, 0x1c, 0xac, 0x8b, 0x89, 0x62, 0xef,
	0x00, 0xb6, 0x77, 0xb2, 0x6a, 0x07, 0x0b, 0xb7, 0x70, 0x5e, 0x57, 0xed, 0xbd, 0xf5, 0x91, 0x2c,
	0x93, 0xeb, 0x69, 0xf9, 0x6f, 0xb7, 0x62, 0x00, 0xd9, 0x03, 0xac, 0xd3, 0x41, 0xab, 0xc0, 0x9f,
	0x1d, 0x5a, 0x47, 0x6e, 0xe1, 0x52, 0x1b, 0xa5, 0x95, 0x95, 0xb5, 0x8f, 0xe2, 0x98, 0xe8, 0x03,
	0xf5, 0x58, 0xa5, 0xad, 0xca, 0x56, 0xba, 0xce, 0xe0, 0xe9, 0x2a, 0x0f, 0x18, 0xfb, 0x0a, 0x57,
	0x5b, 0x6f, 0x7a, 0xfc, 0xed, 0x90, 0x74, 0xf0, 0xaf, 0xa4, 0xc3, 0x27, 0x24, 0xcd, 0x52, 0xd8,
	0xa4, 0x46, 0x69, 0x59, 0x4a, 0x87, 0xc7, 0x6e, 0x4e, 0x6d, 0x84, 0xff, 0x6f, 0xe3, 0xc3, 0xea,
	0x07, 0xf0, 0xb7, 0x63, 0x80, 0xd9, 0xdc, 0x33, 0x6f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2c,
	0x65, 0xae, 0xad, 0x11, 0x03, 0x00, 0x00,
}
