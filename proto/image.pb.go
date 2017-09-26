// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/proto/image.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	gomatcha.io/matcha/proto/image.proto

It has these top-level messages:
	Image
	ImageProperties
	ImageOrResource
	Color
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Image struct {
	Width  int64  `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height int64  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Stride int64  `protobuf:"varint,4,opt,name=stride" json:"stride,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto1.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Image) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetStride() int64 {
	if m != nil {
		return m.Stride
	}
	return 0
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ImageProperties struct {
	Width  int64   `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height int64   `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Scale  float64 `protobuf:"fixed64,3,opt,name=scale" json:"scale,omitempty"`
}

func (m *ImageProperties) Reset()                    { *m = ImageProperties{} }
func (m *ImageProperties) String() string            { return proto1.CompactTextString(m) }
func (*ImageProperties) ProtoMessage()               {}
func (*ImageProperties) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ImageProperties) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ImageProperties) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ImageProperties) GetScale() float64 {
	if m != nil {
		return m.Scale
	}
	return 0
}

type ImageOrResource struct {
	Image *Image `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Path  string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *ImageOrResource) Reset()                    { *m = ImageOrResource{} }
func (m *ImageOrResource) String() string            { return proto1.CompactTextString(m) }
func (*ImageOrResource) ProtoMessage()               {}
func (*ImageOrResource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ImageOrResource) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImageOrResource) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Color struct {
	Red   uint32 `protobuf:"varint,1,opt,name=red" json:"red,omitempty"`
	Blue  uint32 `protobuf:"varint,2,opt,name=blue" json:"blue,omitempty"`
	Green uint32 `protobuf:"varint,3,opt,name=green" json:"green,omitempty"`
	Alpha uint32 `protobuf:"varint,4,opt,name=alpha" json:"alpha,omitempty"`
}

func (m *Color) Reset()                    { *m = Color{} }
func (m *Color) String() string            { return proto1.CompactTextString(m) }
func (*Color) ProtoMessage()               {}
func (*Color) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Color) GetRed() uint32 {
	if m != nil {
		return m.Red
	}
	return 0
}

func (m *Color) GetBlue() uint32 {
	if m != nil {
		return m.Blue
	}
	return 0
}

func (m *Color) GetGreen() uint32 {
	if m != nil {
		return m.Green
	}
	return 0
}

func (m *Color) GetAlpha() uint32 {
	if m != nil {
		return m.Alpha
	}
	return 0
}

func init() {
	proto1.RegisterType((*Image)(nil), "matcha.Image")
	proto1.RegisterType((*ImageProperties)(nil), "matcha.ImageProperties")
	proto1.RegisterType((*ImageOrResource)(nil), "matcha.ImageOrResource")
	proto1.RegisterType((*Color)(nil), "matcha.Color")
}

func init() { proto1.RegisterFile("gomatcha.io/matcha/proto/image.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0xee, 0xa6, 0xe8, 0x68, 0x51, 0xc2, 0x22, 0x39, 0x2e, 0xd5, 0xc3, 0x9e, 0x5a,
	0x70, 0xdf, 0xa0, 0x9e, 0x14, 0xc4, 0x12, 0xf0, 0xa0, 0xb7, 0xd9, 0x76, 0x68, 0x02, 0x5d, 0x53,
	0xd2, 0x2c, 0xbe, 0x8f, 0x4f, 0x2a, 0x99, 0x14, 0xef, 0x9e, 0xf2, 0x7f, 0xff, 0x4c, 0x67, 0xa6,
	0x33, 0xf0, 0x30, 0xb8, 0x13, 0x86, 0xce, 0x60, 0x65, 0x5d, 0x9d, 0x54, 0x3d, 0x79, 0x17, 0x5c,
	0x6d, 0x4f, 0x38, 0x50, 0xc5, 0x5a, 0xe6, 0x29, 0x52, 0x22, 0x88, 0xe7, 0x68, 0xcb, 0x2d, 0x88,
	0x6f, 0xdb, 0x07, 0xa3, 0xb2, 0x5d, 0xb6, 0x5f, 0xeb, 0x04, 0xf2, 0x0e, 0x72, 0x43, 0x76, 0x30,
	0x41, 0xad, 0xd8, 0x5e, 0x28, 0xfa, 0x73, 0xf0, 0xb6, 0x27, 0xb5, 0x49, 0x7e, 0x22, 0x29, 0x61,
	0xd3, 0x63, 0x40, 0xb5, 0xde, 0x65, 0xfb, 0x6b, 0xcd, 0xba, 0x7c, 0x87, 0x1b, 0x6e, 0xd1, 0x7a,
	0x37, 0x91, 0x0f, 0x96, 0xe6, 0x7f, 0x36, 0xdb, 0x82, 0x98, 0x3b, 0x1c, 0x89, 0xab, 0x66, 0x3a,
	0x41, 0xf9, 0xb2, 0x94, 0x7d, 0xf3, 0x9a, 0x66, 0x77, 0xf6, 0x1d, 0xc9, 0x7b, 0x10, 0xfc, 0x8f,
	0x5c, 0xf6, 0xea, 0xb1, 0xa8, 0x96, 0x45, 0x70, 0x9e, 0x4e, 0xb1, 0x38, 0xe2, 0x84, 0xc1, 0x70,
	0x8f, 0x4b, 0xcd, 0xba, 0xfc, 0x00, 0xf1, 0xe4, 0x46, 0xe7, 0xe5, 0x2d, 0xac, 0x3d, 0xf5, 0xfc,
	0x7d, 0xa1, 0xa3, 0x8c, 0xe9, 0xc7, 0xf1, 0x4c, 0x9c, 0x5e, 0x68, 0xd6, 0x71, 0xa0, 0xc1, 0x13,
	0x7d, 0xf1, 0x40, 0x85, 0x4e, 0x10, 0x5d, 0x1c, 0x27, 0x83, 0xbc, 0x92, 0x42, 0x27, 0x68, 0x0e,
	0xa0, 0xac, 0xab, 0xfe, 0x6e, 0xb2, 0x3c, 0x7c, 0x84, 0x46, 0xb4, 0xf1, 0xf9, 0x14, 0x4c, 0x3f,
	0xab, 0x8b, 0x57, 0x0e, 0xb6, 0xcd, 0x31, 0x67, 0xe7, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd9,
	0x5c, 0x4e, 0x3e, 0xcc, 0x01, 0x00, 0x00,
}