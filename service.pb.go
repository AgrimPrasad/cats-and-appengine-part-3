// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package cats is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	PostAddFormatRequest
	GetListFormatRequest
	Cat
	CatsResponse
	ErrorResponse
*/
package cats

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

type PostAddFormatRequest struct {
	Cat *Cat `protobuf:"bytes,1,opt,name=cat" json:"cat,omitempty"`
	// "json" or "proto" are acceptable.
	Format string `protobuf:"bytes,2,opt,name=format" json:"format,omitempty"`
}

func (m *PostAddFormatRequest) Reset()                    { *m = PostAddFormatRequest{} }
func (m *PostAddFormatRequest) String() string            { return proto.CompactTextString(m) }
func (*PostAddFormatRequest) ProtoMessage()               {}
func (*PostAddFormatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PostAddFormatRequest) GetCat() *Cat {
	if m != nil {
		return m.Cat
	}
	return nil
}

func (m *PostAddFormatRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

type GetListFormatRequest struct {
	// "json" or "proto" are acceptable.
	Format string `protobuf:"bytes,1,opt,name=format" json:"format,omitempty"`
}

func (m *GetListFormatRequest) Reset()                    { *m = GetListFormatRequest{} }
func (m *GetListFormatRequest) String() string            { return proto.CompactTextString(m) }
func (*GetListFormatRequest) ProtoMessage()               {}
func (*GetListFormatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetListFormatRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

type Cat struct {
	Breed  string  `protobuf:"bytes,3,opt,name=breed" json:"breed,omitempty"`
	Key    int64   `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Weight float64 `protobuf:"fixed64,4,opt,name=weight" json:"weight,omitempty"`
}

func (m *Cat) Reset()                    { *m = Cat{} }
func (m *Cat) String() string            { return proto.CompactTextString(m) }
func (*Cat) ProtoMessage()               {}
func (*Cat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Cat) GetBreed() string {
	if m != nil {
		return m.Breed
	}
	return ""
}

func (m *Cat) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Cat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cat) GetWeight() float64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type CatsResponse struct {
	Cats  []*Cat `protobuf:"bytes,2,rep,name=cats" json:"cats,omitempty"`
	Total int32  `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
}

func (m *CatsResponse) Reset()                    { *m = CatsResponse{} }
func (m *CatsResponse) String() string            { return proto.CompactTextString(m) }
func (*CatsResponse) ProtoMessage()               {}
func (*CatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CatsResponse) GetCats() []*Cat {
	if m != nil {
		return m.Cats
	}
	return nil
}

func (m *CatsResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ErrorResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *ErrorResponse) Reset()                    { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()               {}
func (*ErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ErrorResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*PostAddFormatRequest)(nil), "cats.PostAddFormatRequest")
	proto.RegisterType((*GetListFormatRequest)(nil), "cats.GetListFormatRequest")
	proto.RegisterType((*Cat)(nil), "cats.Cat")
	proto.RegisterType((*CatsResponse)(nil), "cats.CatsResponse")
	proto.RegisterType((*ErrorResponse)(nil), "cats.ErrorResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xed, 0x76, 0xdb, 0x42, 0xa7, 0x5f, 0xe0, 0x63, 0x08, 0x12, 0x2a, 0x42, 0x58, 0x10, 0x7a,
	0xca, 0xa1, 0x5e, 0xbc, 0x96, 0xa0, 0x1e, 0xf4, 0x20, 0xeb, 0xc9, 0xe3, 0xb6, 0x19, 0x35, 0x68,
	0xbb, 0x75, 0x77, 0x54, 0xfc, 0x0b, 0xfe, 0x6a, 0xd9, 0x4d, 0x4c, 0x09, 0xf4, 0x36, 0x6f, 0xf6,
	0xcd, 0xbc, 0x79, 0x6f, 0x21, 0xf1, 0xe4, 0x3e, 0xeb, 0x0d, 0x15, 0x7b, 0x67, 0xd9, 0xe2, 0x68,
	0x63, 0xd8, 0xab, 0x5b, 0x48, 0xef, 0xad, 0xe7, 0x55, 0x55, 0x5d, 0x5b, 0xb7, 0x35, 0xac, 0xe9,
	0xfd, 0x83, 0x3c, 0xe3, 0x29, 0xc8, 0x8d, 0xe1, 0x4c, 0xe4, 0x62, 0x31, 0x5b, 0x4e, 0x8b, 0xc0,
	0x2d, 0x4a, 0xc3, 0x3a, 0x74, 0xf1, 0x04, 0x26, 0x4f, 0x91, 0x9d, 0x0d, 0x73, 0xb1, 0x98, 0xea,
	0x16, 0xa9, 0x02, 0xd2, 0x1b, 0xe2, 0xbb, 0xda, 0x73, 0x7f, 0xd9, 0x81, 0x2f, 0x7a, 0xfc, 0x47,
	0x90, 0xa5, 0x61, 0x4c, 0x61, 0xbc, 0x76, 0x44, 0x55, 0x26, 0xe3, 0x6b, 0x03, 0xf0, 0x3f, 0xc8,
	0x57, 0xfa, 0x8e, 0x13, 0x52, 0x87, 0x12, 0x11, 0x46, 0x3b, 0xb3, 0xa5, 0x56, 0x34, 0xd6, 0x61,
	0xf5, 0x17, 0xd5, 0xcf, 0x2f, 0x9c, 0x8d, 0x72, 0xb1, 0x10, 0xba, 0x45, 0xaa, 0x84, 0x7f, 0xa5,
	0x61, 0xaf, 0xc9, 0xef, 0xed, 0xce, 0x13, 0x9e, 0x41, 0xf4, 0x9b, 0x0d, 0x73, 0xd9, 0x37, 0x14,
	0xdb, 0xe1, 0x04, 0xb6, 0x6c, 0xde, 0xa2, 0xdc, 0x58, 0x37, 0x40, 0x9d, 0x43, 0x72, 0xe5, 0x9c,
	0x75, 0xdd, 0x96, 0x14, 0xc6, 0x14, 0x1a, 0xad, 0x8f, 0x06, 0x2c, 0x7f, 0x04, 0xcc, 0x82, 0xd8,
	0x43, 0x93, 0x2f, 0x5e, 0x42, 0xd2, 0xcb, 0x14, 0xe7, 0x8d, 0xdc, 0xb1, 0xa0, 0xe7, 0x87, 0x53,
	0xd4, 0x00, 0x57, 0x90, 0xf4, 0x02, 0xfc, 0x9b, 0x3c, 0x96, 0xea, 0x1c, 0xbb, 0xc9, 0xce, 0xa6,
	0x1a, 0xac, 0x27, 0xf1, 0x77, 0x2f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x52, 0x8a, 0x80,
	0xee, 0x01, 0x00, 0x00,
}