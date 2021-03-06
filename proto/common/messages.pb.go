// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common/messages.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Uint32List struct {
	List                 []uint32 `protobuf:"varint,1,rep,packed,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uint32List) Reset()         { *m = Uint32List{} }
func (m *Uint32List) String() string { return proto.CompactTextString(m) }
func (*Uint32List) ProtoMessage()    {}
func (*Uint32List) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_b5bf9291ad2f880e, []int{0}
}
func (m *Uint32List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uint32List.Unmarshal(m, b)
}
func (m *Uint32List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uint32List.Marshal(b, m, deterministic)
}
func (dst *Uint32List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uint32List.Merge(dst, src)
}
func (m *Uint32List) XXX_Size() int {
	return xxx_messageInfo_Uint32List.Size(m)
}
func (m *Uint32List) XXX_DiscardUnknown() {
	xxx_messageInfo_Uint32List.DiscardUnknown(m)
}

var xxx_messageInfo_Uint32List proto.InternalMessageInfo

func (m *Uint32List) GetList() []uint32 {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Uint32List)(nil), "ethereum.common.Uint32List")
}

func init() {
	proto.RegisterFile("proto/common/messages.proto", fileDescriptor_messages_b5bf9291ad2f880e)
}

var fileDescriptor_messages_b5bf9291ad2f880e = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x2d, 0xd6, 0x03, 0x8b, 0x0a, 0xf1, 0xa7, 0x96, 0x64, 0xa4, 0x16, 0xa5, 0x96, 0xe6, 0xea, 0x41,
	0xe4, 0x95, 0x14, 0xb8, 0xb8, 0x42, 0x33, 0xf3, 0x4a, 0x8c, 0x8d, 0x7c, 0x32, 0x8b, 0x4b, 0x84,
	0x84, 0xb8, 0x58, 0x72, 0x32, 0x8b, 0x4b, 0x24, 0x18, 0x15, 0x98, 0x35, 0x78, 0x83, 0xc0, 0xec,
	0x24, 0x36, 0xb0, 0x4e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x72, 0x32, 0x68, 0x58,
	0x00, 0x00, 0x00,
}
