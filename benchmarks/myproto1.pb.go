// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myproto1.proto

/*
Package benchmarks is a generated protocol buffer package.

It is generated from these files:
	myproto1.proto

It has these top-level messages:
	Myproto1
*/
package benchmarks

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

type Myproto1 struct {
	F1  bool    `protobuf:"varint,1,opt,name=F1" json:"F1,omitempty"`
	F2  float32 `protobuf:"fixed32,2,opt,name=F2" json:"F2,omitempty"`
	F3  float64 `protobuf:"fixed64,3,opt,name=F3" json:"F3,omitempty"`
	F4  string  `protobuf:"bytes,4,opt,name=F4" json:"F4,omitempty"`
	F5  []byte  `protobuf:"bytes,5,opt,name=F5,proto3" json:"F5,omitempty"`
	F6  int32   `protobuf:"varint,6,opt,name=F6" json:"F6,omitempty"`
	F7  int32   `protobuf:"varint,7,opt,name=F7" json:"F7,omitempty"`
	F8  int32   `protobuf:"varint,8,opt,name=F8" json:"F8,omitempty"`
	F9  int32   `protobuf:"varint,9,opt,name=F9" json:"F9,omitempty"`
	F10 int64   `protobuf:"varint,10,opt,name=F10" json:"F10,omitempty"`
	F11 int64   `protobuf:"varint,11,opt,name=F11" json:"F11,omitempty"`
	F12 uint32  `protobuf:"varint,12,opt,name=F12" json:"F12,omitempty"`
	F13 uint32  `protobuf:"varint,13,opt,name=F13" json:"F13,omitempty"`
	F14 uint32  `protobuf:"varint,14,opt,name=F14" json:"F14,omitempty"`
	F15 uint32  `protobuf:"varint,15,opt,name=F15" json:"F15,omitempty"`
	F16 uint64  `protobuf:"varint,16,opt,name=F16" json:"F16,omitempty"`
	F17 uint64  `protobuf:"varint,17,opt,name=F17" json:"F17,omitempty"`
}

func (m *Myproto1) Reset()                    { *m = Myproto1{} }
func (m *Myproto1) String() string            { return proto.CompactTextString(m) }
func (*Myproto1) ProtoMessage()               {}
func (*Myproto1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Myproto1) GetF1() bool {
	if m != nil {
		return m.F1
	}
	return false
}

func (m *Myproto1) GetF2() float32 {
	if m != nil {
		return m.F2
	}
	return 0
}

func (m *Myproto1) GetF3() float64 {
	if m != nil {
		return m.F3
	}
	return 0
}

func (m *Myproto1) GetF4() string {
	if m != nil {
		return m.F4
	}
	return ""
}

func (m *Myproto1) GetF5() []byte {
	if m != nil {
		return m.F5
	}
	return nil
}

func (m *Myproto1) GetF6() int32 {
	if m != nil {
		return m.F6
	}
	return 0
}

func (m *Myproto1) GetF7() int32 {
	if m != nil {
		return m.F7
	}
	return 0
}

func (m *Myproto1) GetF8() int32 {
	if m != nil {
		return m.F8
	}
	return 0
}

func (m *Myproto1) GetF9() int32 {
	if m != nil {
		return m.F9
	}
	return 0
}

func (m *Myproto1) GetF10() int64 {
	if m != nil {
		return m.F10
	}
	return 0
}

func (m *Myproto1) GetF11() int64 {
	if m != nil {
		return m.F11
	}
	return 0
}

func (m *Myproto1) GetF12() uint32 {
	if m != nil {
		return m.F12
	}
	return 0
}

func (m *Myproto1) GetF13() uint32 {
	if m != nil {
		return m.F13
	}
	return 0
}

func (m *Myproto1) GetF14() uint32 {
	if m != nil {
		return m.F14
	}
	return 0
}

func (m *Myproto1) GetF15() uint32 {
	if m != nil {
		return m.F15
	}
	return 0
}

func (m *Myproto1) GetF16() uint64 {
	if m != nil {
		return m.F16
	}
	return 0
}

func (m *Myproto1) GetF17() uint64 {
	if m != nil {
		return m.F17
	}
	return 0
}

func init() {
	proto.RegisterType((*Myproto1)(nil), "benchmarks.Myproto1")
}

func init() { proto.RegisterFile("myproto1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xd0, 0x3b, 0x4e, 0x03, 0x41,
	0x0c, 0xc6, 0x71, 0xd9, 0x79, 0xb0, 0x19, 0x92, 0x25, 0x6c, 0xf5, 0x95, 0x16, 0x95, 0x2b, 0xc4,
	0xec, 0xec, 0x23, 0xb9, 0xc0, 0x74, 0x34, 0x73, 0x03, 0x82, 0x90, 0x90, 0x50, 0x08, 0x0a, 0x34,
	0x5c, 0x89, 0x53, 0xa2, 0xf5, 0xec, 0x54, 0xf6, 0xff, 0xe7, 0xce, 0xae, 0x3e, 0xff, 0x7e, 0x5d,
	0x2f, 0x3f, 0x17, 0xff, 0x68, 0xa3, 0x71, 0xa7, 0xb7, 0xcf, 0xd7, 0xf7, 0xf3, 0xcb, 0xf5, 0xe3,
	0xfb, 0xe1, 0x8f, 0x5d, 0xf5, 0x3c, 0x9f, 0x9b, 0xda, 0x71, 0xf4, 0x20, 0x21, 0xad, 0x12, 0xc7,
	0xdc, 0x2d, 0x58, 0x48, 0x39, 0x71, 0x6c, 0xad, 0x03, 0x16, 0x42, 0x4a, 0x89, 0x63, 0xb0, 0xee,
	0xb0, 0x14, 0xd2, 0x4d, 0xe2, 0xd8, 0x59, 0xf7, 0x58, 0x09, 0xe9, 0x36, 0x71, 0xec, 0xad, 0x07,
	0xac, 0x85, 0x74, 0x95, 0x38, 0x0e, 0xd6, 0x23, 0x6e, 0xe6, 0x1e, 0xad, 0x0f, 0xa8, 0xe6, 0x3e,
	0x58, 0x1f, 0xb1, 0x99, 0xfb, 0xd8, 0xec, 0xdd, 0x22, 0xfa, 0x27, 0x38, 0x21, 0x5d, 0xa4, 0x69,
	0xcd, 0xe2, 0x71, 0x5b, 0xc4, 0x67, 0x69, 0xb1, 0x15, 0xd2, 0xdd, 0x24, 0x6d, 0x96, 0x80, 0x5d,
	0x91, 0x90, 0xa5, 0x43, 0x5d, 0xa4, 0xcb, 0xd2, 0xe3, 0xae, 0x48, 0x9f, 0x65, 0xc0, 0x5e, 0x48,
	0x97, 0x93, 0x0c, 0x59, 0x46, 0xdc, 0x17, 0x19, 0x4f, 0x6b, 0xfb, 0x53, 0xf8, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0xd4, 0x6d, 0x2f, 0x74, 0x51, 0x01, 0x00, 0x00,
}