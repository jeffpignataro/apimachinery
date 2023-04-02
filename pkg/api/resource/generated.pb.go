/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/jeffpignataro/kubernetes/vendor/github.com/jeffpignataro/apimachinery/pkg/api/resource/generated.proto

package resource

import (
	fmt "fmt"

	math "math"

	proto "github.com/gogo/protobuf/proto"
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

func (m *Quantity) Reset()      { *m = Quantity{} }
func (*Quantity) ProtoMessage() {}
func (*Quantity) Descriptor() ([]byte, []int) {
	return fileDescriptor_612bba87bd70906c, []int{0}
}
func (m *Quantity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quantity.Unmarshal(m, b)
}
func (m *Quantity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quantity.Marshal(b, m, deterministic)
}
func (m *Quantity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quantity.Merge(m, src)
}
func (m *Quantity) XXX_Size() int {
	return xxx_messageInfo_Quantity.Size(m)
}
func (m *Quantity) XXX_DiscardUnknown() {
	xxx_messageInfo_Quantity.DiscardUnknown(m)
}

var xxx_messageInfo_Quantity proto.InternalMessageInfo

func (m *QuantityValue) Reset()      { *m = QuantityValue{} }
func (*QuantityValue) ProtoMessage() {}
func (*QuantityValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_612bba87bd70906c, []int{1}
}
func (m *QuantityValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuantityValue.Unmarshal(m, b)
}
func (m *QuantityValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuantityValue.Marshal(b, m, deterministic)
}
func (m *QuantityValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuantityValue.Merge(m, src)
}
func (m *QuantityValue) XXX_Size() int {
	return xxx_messageInfo_QuantityValue.Size(m)
}
func (m *QuantityValue) XXX_DiscardUnknown() {
	xxx_messageInfo_QuantityValue.DiscardUnknown(m)
}

var xxx_messageInfo_QuantityValue proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Quantity)(nil), "k8s.io.apimachinery.pkg.api.resource.Quantity")
	proto.RegisterType((*QuantityValue)(nil), "k8s.io.apimachinery.pkg.api.resource.QuantityValue")
}

func init() {
	proto.RegisterFile("github.com/jeffpignataro/kubernetes/vendor/github.com/jeffpignataro/apimachinery/pkg/api/resource/generated.proto", fileDescriptor_612bba87bd70906c)
}

var fileDescriptor_612bba87bd70906c = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0xcd, 0xb6, 0x28, 0xd6,
	0xcb, 0xcc, 0xd7, 0xcf, 0x2e, 0x4d, 0x4a, 0x2d, 0xca, 0x4b, 0x2d, 0x49, 0x2d, 0xd6, 0x2f, 0x4b,
	0xcd, 0x4b, 0xc9, 0x2f, 0xd2, 0x87, 0x4a, 0x24, 0x16, 0x64, 0xe6, 0x26, 0x26, 0x67, 0x64, 0xe6,
	0xa5, 0x16, 0x55, 0xea, 0x17, 0x64, 0xa7, 0x83, 0x04, 0xf4, 0x8b, 0x52, 0x8b, 0xf3, 0x4b, 0x8b,
	0x92, 0x53, 0xf5, 0xd3, 0x53, 0xf3, 0x52, 0x8b, 0x12, 0x4b, 0x52, 0x53, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x54, 0x20, 0xba, 0xf4, 0x90, 0x75, 0xe9, 0x15, 0x64, 0xa7, 0x83, 0x04, 0xf4,
	0x60, 0xba, 0xa4, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3,
	0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x9a, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x18,
	0xaa, 0x64, 0xc1, 0xc5, 0x11, 0x58, 0x9a, 0x98, 0x57, 0x92, 0x59, 0x52, 0x29, 0x24, 0xc6, 0xc5,
	0x56, 0x5c, 0x52, 0x94, 0x99, 0x97, 0x2e, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x59,
	0x89, 0xcc, 0x58, 0x20, 0xcf, 0xd0, 0xb1, 0x50, 0x9e, 0x61, 0xc2, 0x42, 0x79, 0x86, 0x05, 0x0b,
	0xe5, 0x19, 0x1a, 0xee, 0x28, 0x30, 0x28, 0xd9, 0x72, 0xf1, 0xc2, 0x74, 0x86, 0x25, 0xe6, 0x94,
	0xa6, 0x92, 0xa6, 0xdd, 0xc9, 0xeb, 0xc4, 0x43, 0x39, 0x86, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c,
	0x94, 0x63, 0x68, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x37,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x43, 0x94, 0x0a, 0x31, 0x21,
	0x05, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x70, 0x98, 0xa3, 0x69, 0x01, 0x00, 0x00,
}
