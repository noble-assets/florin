// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: florin/blacklist/v1/events.proto

package blacklist

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Emitted when a validator makes a decision.
type Decision struct {
	// from is the sender address.
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// to is the recipient address.
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// amount is the number of tokens.
	Amount cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
	// valid is true if transfer approved, false if rejected.
	Valid bool `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (m *Decision) Reset()         { *m = Decision{} }
func (m *Decision) String() string { return proto.CompactTextString(m) }
func (*Decision) ProtoMessage()    {}
func (*Decision) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d94ec933c59aa91, []int{0}
}
func (m *Decision) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Decision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Decision.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Decision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decision.Merge(m, src)
}
func (m *Decision) XXX_Size() int {
	return m.Size()
}
func (m *Decision) XXX_DiscardUnknown() {
	xxx_messageInfo_Decision.DiscardUnknown(m)
}

var xxx_messageInfo_Decision proto.InternalMessageInfo

func (m *Decision) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Decision) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Decision) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

// Emitted when an address is added to the blacklist.
type Ban struct {
	// adversary is the address that was added.
	Adversary string `protobuf:"bytes,1,opt,name=adversary,proto3" json:"adversary,omitempty"`
}

func (m *Ban) Reset()         { *m = Ban{} }
func (m *Ban) String() string { return proto.CompactTextString(m) }
func (*Ban) ProtoMessage()    {}
func (*Ban) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d94ec933c59aa91, []int{1}
}
func (m *Ban) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ban) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ban.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ban) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ban.Merge(m, src)
}
func (m *Ban) XXX_Size() int {
	return m.Size()
}
func (m *Ban) XXX_DiscardUnknown() {
	xxx_messageInfo_Ban.DiscardUnknown(m)
}

var xxx_messageInfo_Ban proto.InternalMessageInfo

func (m *Ban) GetAdversary() string {
	if m != nil {
		return m.Adversary
	}
	return ""
}

// Emitted when an address is removed from the blacklist.
type Unban struct {
	// friend is the address that was removed.
	Friend string `protobuf:"bytes,1,opt,name=friend,proto3" json:"friend,omitempty"`
}

func (m *Unban) Reset()         { *m = Unban{} }
func (m *Unban) String() string { return proto.CompactTextString(m) }
func (*Unban) ProtoMessage()    {}
func (*Unban) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d94ec933c59aa91, []int{2}
}
func (m *Unban) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Unban) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Unban.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Unban) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unban.Merge(m, src)
}
func (m *Unban) XXX_Size() int {
	return m.Size()
}
func (m *Unban) XXX_DiscardUnknown() {
	xxx_messageInfo_Unban.DiscardUnknown(m)
}

var xxx_messageInfo_Unban proto.InternalMessageInfo

func (m *Unban) GetFriend() string {
	if m != nil {
		return m.Friend
	}
	return ""
}

// Emitted when admin account is added.
type AdminAccountAdded struct {
	// account is the address that was added.
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *AdminAccountAdded) Reset()         { *m = AdminAccountAdded{} }
func (m *AdminAccountAdded) String() string { return proto.CompactTextString(m) }
func (*AdminAccountAdded) ProtoMessage()    {}
func (*AdminAccountAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d94ec933c59aa91, []int{3}
}
func (m *AdminAccountAdded) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdminAccountAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdminAccountAdded.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AdminAccountAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminAccountAdded.Merge(m, src)
}
func (m *AdminAccountAdded) XXX_Size() int {
	return m.Size()
}
func (m *AdminAccountAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminAccountAdded.DiscardUnknown(m)
}

var xxx_messageInfo_AdminAccountAdded proto.InternalMessageInfo

func (m *AdminAccountAdded) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

// Emitted when admin account is removed.
type AdminAccountRemoved struct {
	// account is the address that was removed.
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *AdminAccountRemoved) Reset()         { *m = AdminAccountRemoved{} }
func (m *AdminAccountRemoved) String() string { return proto.CompactTextString(m) }
func (*AdminAccountRemoved) ProtoMessage()    {}
func (*AdminAccountRemoved) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d94ec933c59aa91, []int{4}
}
func (m *AdminAccountRemoved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdminAccountRemoved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdminAccountRemoved.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AdminAccountRemoved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminAccountRemoved.Merge(m, src)
}
func (m *AdminAccountRemoved) XXX_Size() int {
	return m.Size()
}
func (m *AdminAccountRemoved) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminAccountRemoved.DiscardUnknown(m)
}

var xxx_messageInfo_AdminAccountRemoved proto.InternalMessageInfo

func (m *AdminAccountRemoved) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

// Emitted when an ownership transfer is started.
type OwnershipTransferStarted struct {
	// previous_owner is the address of the previous owner.
	PreviousOwner string `protobuf:"bytes,1,opt,name=previous_owner,json=previousOwner,proto3" json:"previous_owner,omitempty"`
	// new_owner is the address of the new owner.
	NewOwner string `protobuf:"bytes,2,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
}

func (m *OwnershipTransferStarted) Reset()         { *m = OwnershipTransferStarted{} }
func (m *OwnershipTransferStarted) String() string { return proto.CompactTextString(m) }
func (*OwnershipTransferStarted) ProtoMessage()    {}
func (*OwnershipTransferStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d94ec933c59aa91, []int{5}
}
func (m *OwnershipTransferStarted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnershipTransferStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnershipTransferStarted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnershipTransferStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnershipTransferStarted.Merge(m, src)
}
func (m *OwnershipTransferStarted) XXX_Size() int {
	return m.Size()
}
func (m *OwnershipTransferStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnershipTransferStarted.DiscardUnknown(m)
}

var xxx_messageInfo_OwnershipTransferStarted proto.InternalMessageInfo

func (m *OwnershipTransferStarted) GetPreviousOwner() string {
	if m != nil {
		return m.PreviousOwner
	}
	return ""
}

func (m *OwnershipTransferStarted) GetNewOwner() string {
	if m != nil {
		return m.NewOwner
	}
	return ""
}

// Emitted when an ownership transfer is finalized.
type OwnershipTransferred struct {
	// previous_owner is the address of the previous owner.
	PreviousOwner string `protobuf:"bytes,1,opt,name=previous_owner,json=previousOwner,proto3" json:"previous_owner,omitempty"`
	// new_owner is the address of the new owner.
	NewOwner string `protobuf:"bytes,2,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
}

func (m *OwnershipTransferred) Reset()         { *m = OwnershipTransferred{} }
func (m *OwnershipTransferred) String() string { return proto.CompactTextString(m) }
func (*OwnershipTransferred) ProtoMessage()    {}
func (*OwnershipTransferred) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d94ec933c59aa91, []int{6}
}
func (m *OwnershipTransferred) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnershipTransferred) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnershipTransferred.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnershipTransferred) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnershipTransferred.Merge(m, src)
}
func (m *OwnershipTransferred) XXX_Size() int {
	return m.Size()
}
func (m *OwnershipTransferred) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnershipTransferred.DiscardUnknown(m)
}

var xxx_messageInfo_OwnershipTransferred proto.InternalMessageInfo

func (m *OwnershipTransferred) GetPreviousOwner() string {
	if m != nil {
		return m.PreviousOwner
	}
	return ""
}

func (m *OwnershipTransferred) GetNewOwner() string {
	if m != nil {
		return m.NewOwner
	}
	return ""
}

func init() {
	proto.RegisterType((*Decision)(nil), "florin.blacklist.v1.Decision")
	proto.RegisterType((*Ban)(nil), "florin.blacklist.v1.Ban")
	proto.RegisterType((*Unban)(nil), "florin.blacklist.v1.Unban")
	proto.RegisterType((*AdminAccountAdded)(nil), "florin.blacklist.v1.AdminAccountAdded")
	proto.RegisterType((*AdminAccountRemoved)(nil), "florin.blacklist.v1.AdminAccountRemoved")
	proto.RegisterType((*OwnershipTransferStarted)(nil), "florin.blacklist.v1.OwnershipTransferStarted")
	proto.RegisterType((*OwnershipTransferred)(nil), "florin.blacklist.v1.OwnershipTransferred")
}

func init() { proto.RegisterFile("florin/blacklist/v1/events.proto", fileDescriptor_0d94ec933c59aa91) }

var fileDescriptor_0d94ec933c59aa91 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x6a, 0x13, 0x51,
	0x14, 0xc6, 0x33, 0x69, 0x1b, 0x93, 0x03, 0x16, 0x7a, 0x1b, 0x65, 0xac, 0x32, 0x09, 0x23, 0x42,
	0x11, 0x32, 0xd7, 0xda, 0x27, 0x48, 0x70, 0x61, 0x17, 0x22, 0x44, 0xdd, 0x74, 0x61, 0xb9, 0x33,
	0xf7, 0x26, 0xb9, 0x74, 0xe6, 0x9c, 0x70, 0xef, 0x9d, 0x09, 0x7d, 0x09, 0xf1, 0x31, 0x5c, 0xba,
	0xf0, 0x21, 0xba, 0x2c, 0xae, 0xc4, 0x45, 0x91, 0x64, 0xe1, 0x6b, 0xc8, 0xfc, 0xa9, 0x15, 0x04,
	0x57, 0x6e, 0x86, 0xf3, 0x7d, 0xf7, 0xf7, 0x9d, 0x73, 0x18, 0x0e, 0x0c, 0x67, 0x29, 0x19, 0x8d,
	0x3c, 0x4e, 0x45, 0x72, 0x9e, 0x6a, 0xeb, 0x78, 0x71, 0xc4, 0x55, 0xa1, 0xd0, 0xd9, 0x68, 0x69,
	0xc8, 0x11, 0xdb, 0xaf, 0x89, 0xe8, 0x37, 0x11, 0x15, 0x47, 0x07, 0x7b, 0x22, 0xd3, 0x48, 0xbc,
	0xfa, 0xd6, 0xdc, 0xc1, 0x83, 0x84, 0x6c, 0x46, 0xf6, 0xac, 0x52, 0xbc, 0x16, 0xcd, 0x53, 0x7f,
	0x4e, 0x73, 0xaa, 0xfd, 0xb2, 0xaa, 0xdd, 0xf0, 0x83, 0x07, 0xdd, 0x17, 0x2a, 0xd1, 0x56, 0x13,
	0x32, 0x06, 0xdb, 0x33, 0x43, 0x99, 0xef, 0x0d, 0xbd, 0xc3, 0xde, 0xb4, 0xaa, 0xd9, 0x2e, 0xb4,
	0x1d, 0xf9, 0xed, 0xca, 0x69, 0x3b, 0x62, 0x2f, 0xa1, 0x23, 0x32, 0xca, 0xd1, 0xf9, 0x5b, 0xa5,
	0x37, 0x79, 0x76, 0x79, 0x3d, 0x68, 0x7d, 0xbf, 0x1e, 0xdc, 0xab, 0x87, 0x59, 0x79, 0x1e, 0x69,
	0xe2, 0x99, 0x70, 0x8b, 0xe8, 0x04, 0xdd, 0xd7, 0x2f, 0x23, 0x68, 0xb6, 0x38, 0x41, 0xf7, 0xe9,
	0xe7, 0xe7, 0xa7, 0xde, 0xb4, 0xc9, 0xb3, 0x3e, 0xec, 0x14, 0x22, 0xd5, 0xd2, 0xdf, 0x1e, 0x7a,
	0x87, 0xdd, 0x69, 0x2d, 0xc2, 0xc7, 0xb0, 0x35, 0x11, 0xc8, 0x1e, 0x41, 0x4f, 0xc8, 0x42, 0x19,
	0x2b, 0xcc, 0x45, 0xb3, 0xcf, 0xad, 0x11, 0x0e, 0x60, 0xe7, 0x1d, 0xc6, 0x02, 0xd9, 0x7d, 0xe8,
	0xcc, 0x8c, 0x56, 0x28, 0x1b, 0xa6, 0x51, 0xe1, 0x08, 0xf6, 0xc6, 0x32, 0xd3, 0x38, 0x4e, 0x92,
	0x72, 0xd6, 0x58, 0x4a, 0x25, 0x99, 0x0f, 0x77, 0x44, 0xad, 0x1b, 0xfa, 0x46, 0x86, 0x1c, 0xf6,
	0xff, 0xc4, 0xa7, 0x2a, 0xa3, 0xe2, 0x9f, 0x81, 0xf7, 0xe0, 0xbf, 0x5e, 0xa1, 0x32, 0x76, 0xa1,
	0x97, 0x6f, 0x8d, 0x40, 0x3b, 0x53, 0xe6, 0x8d, 0x13, 0xc6, 0x29, 0xc9, 0x9e, 0xc0, 0xee, 0xd2,
	0xa8, 0x42, 0x53, 0x6e, 0xcf, 0xa8, 0x84, 0x9a, 0xf0, 0xdd, 0x1b, 0xb7, 0x4a, 0xb2, 0x87, 0xd0,
	0x43, 0xb5, 0x6a, 0x88, 0xfa, 0xff, 0x76, 0x51, 0xad, 0xaa, 0xc7, 0xf0, 0x14, 0xfa, 0x7f, 0xf5,
	0x37, 0xff, 0xa7, 0xf7, 0xe4, 0xd5, 0xe5, 0x3a, 0xf0, 0xae, 0xd6, 0x81, 0xf7, 0x63, 0x1d, 0x78,
	0x1f, 0x37, 0x41, 0xeb, 0x6a, 0x13, 0xb4, 0xbe, 0x6d, 0x82, 0xd6, 0xe9, 0xf1, 0x5c, 0xbb, 0x45,
	0x1e, 0x47, 0x09, 0x65, 0x3c, 0x23, 0x54, 0x46, 0xe7, 0x65, 0x21, 0xf3, 0x54, 0x8d, 0x90, 0xe2,
	0x54, 0xf1, 0xe2, 0x39, 0x77, 0x17, 0x4b, 0x65, 0x6f, 0x4f, 0x35, 0xee, 0x54, 0x87, 0x74, 0xfc,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x2a, 0x03, 0xab, 0xc5, 0x02, 0x00, 0x00,
}

func (m *Decision) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Decision) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Decision) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Valid {
		i--
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Ban) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ban) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ban) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Adversary) > 0 {
		i -= len(m.Adversary)
		copy(dAtA[i:], m.Adversary)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Adversary)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Unban) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unban) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Unban) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Friend) > 0 {
		i -= len(m.Friend)
		copy(dAtA[i:], m.Friend)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Friend)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AdminAccountAdded) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdminAccountAdded) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AdminAccountAdded) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AdminAccountRemoved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdminAccountRemoved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AdminAccountRemoved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OwnershipTransferStarted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnershipTransferStarted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnershipTransferStarted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewOwner) > 0 {
		i -= len(m.NewOwner)
		copy(dAtA[i:], m.NewOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PreviousOwner) > 0 {
		i -= len(m.PreviousOwner)
		copy(dAtA[i:], m.PreviousOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PreviousOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OwnershipTransferred) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnershipTransferred) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnershipTransferred) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewOwner) > 0 {
		i -= len(m.NewOwner)
		copy(dAtA[i:], m.NewOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PreviousOwner) > 0 {
		i -= len(m.PreviousOwner)
		copy(dAtA[i:], m.PreviousOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PreviousOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Decision) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovEvents(uint64(l))
	if m.Valid {
		n += 2
	}
	return n
}

func (m *Ban) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Adversary)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *Unban) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Friend)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *AdminAccountAdded) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *AdminAccountRemoved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *OwnershipTransferStarted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PreviousOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *OwnershipTransferred) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PreviousOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Decision) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Decision: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Decision: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Ban) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ban: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ban: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Adversary", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Adversary = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Unban) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Unban: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unban: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Friend", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Friend = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdminAccountAdded) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdminAccountAdded: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdminAccountAdded: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdminAccountRemoved) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdminAccountRemoved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdminAccountRemoved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OwnershipTransferStarted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OwnershipTransferStarted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnershipTransferStarted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OwnershipTransferred) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OwnershipTransferred: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnershipTransferred: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)