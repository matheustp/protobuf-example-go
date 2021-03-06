// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adressbook/adressbook.proto

/*
Package tutorial is a generated protocol buffer package.

It is generated from these files:
	adressbook/adressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package tutorial

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// [START messages]
type Person struct {
	Name        string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id          int32                      `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email       string                     `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones      []*Person_PhoneNumber      `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
	LastUpdated *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("adressbook/adressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x51, 0x4b, 0xeb, 0x30,
	0x18, 0x86, 0x4f, 0xbb, 0xae, 0x6c, 0x5f, 0x0f, 0xa3, 0x27, 0x8c, 0x43, 0xa9, 0x82, 0x65, 0x78,
	0x51, 0x10, 0x32, 0x98, 0x82, 0x57, 0x5e, 0x58, 0x18, 0x2a, 0x3a, 0x57, 0xca, 0x86, 0x97, 0x92,
	0xda, 0x38, 0xcb, 0xda, 0x26, 0x34, 0x29, 0xb8, 0xbf, 0xe4, 0x5f, 0xf0, 0xcf, 0x49, 0xd3, 0x76,
	0x0e, 0xf1, 0xee, 0x4d, 0xbe, 0xa7, 0x5f, 0xdf, 0x3c, 0x70, 0x44, 0x92, 0x92, 0x0a, 0x11, 0x33,
	0xb6, 0x9d, 0x7e, 0x47, 0xcc, 0x4b, 0x26, 0x19, 0x1a, 0xc8, 0x4a, 0xb2, 0x32, 0x25, 0x99, 0x7b,
	0xb2, 0x61, 0x6c, 0x93, 0xd1, 0xa9, 0xba, 0x8f, 0xab, 0xd7, 0xa9, 0x4c, 0x73, 0x2a, 0x24, 0xc9,
	0x79, 0x83, 0x4e, 0x3e, 0x75, 0x30, 0x43, 0x5a, 0x0a, 0x56, 0x20, 0x04, 0x46, 0x41, 0x72, 0xea,
	0x68, 0x9e, 0xe6, 0x0f, 0x23, 0x95, 0xd1, 0x08, 0xf4, 0x34, 0x71, 0x74, 0x4f, 0xf3, 0xfb, 0x91,
	0x9e, 0x26, 0x68, 0x0c, 0x7d, 0x9a, 0x93, 0x34, 0x73, 0x7a, 0x0a, 0x6a, 0x0e, 0xe8, 0x02, 0x4c,
	0xfe, 0xc6, 0x0a, 0x2a, 0x1c, 0xc3, 0xeb, 0xf9, 0xd6, 0xec, 0x18, 0x77, 0x05, 0x70, 0xb3, 0x1b,
	0x87, 0xf5, 0xf8, 0xb1, 0xca, 0x63, 0x5a, 0x46, 0x2d, 0x8b, 0xae, 0xe0, 0x6f, 0x46, 0x84, 0x7c,
	0xae, 0x78, 0x42, 0x24, 0x4d, 0x9c, 0xbe, 0xa7, 0xf9, 0xd6, 0xcc, 0xc5, 0x4d, 0x65, 0xdc, 0x55,
	0xc6, 0xab, 0xae, 0x72, 0x64, 0xd5, 0xfc, 0xba, 0xc1, 0xdd, 0x35, 0x58, 0x07, 0x5b, 0xd1, 0x7f,
	0x30, 0x0b, 0x95, 0xda, 0xfe, 0xed, 0x09, 0x61, 0x30, 0xe4, 0x8e, 0x53, 0xf5, 0x86, 0xd1, 0xcc,
	0xfd, 0xbd, 0xd9, 0x6a, 0xc7, 0x69, 0xa4, 0xb8, 0xc9, 0x19, 0x0c, 0xf7, 0x57, 0x08, 0xc0, 0x5c,
	0x2c, 0x83, 0xbb, 0x87, 0xb9, 0xfd, 0x07, 0x0d, 0xc0, 0xb8, 0x5d, 0x2e, 0xe6, 0xb6, 0x56, 0xa7,
	0xa7, 0x65, 0x74, 0x6f, 0xeb, 0x93, 0x4b, 0xb0, 0xae, 0x13, 0x65, 0x3f, 0x60, 0x6c, 0x8b, 0x7c,
	0x30, 0x39, 0x65, 0x3c, 0xab, 0x1d, 0xd6, 0x1e, 0xec, 0x9f, 0x7f, 0x8b, 0xda, 0x79, 0x10, 0xc2,
	0xf8, 0x85, 0xe5, 0x98, 0xbe, 0x93, 0x9c, 0x67, 0x74, 0x8f, 0x05, 0xff, 0x0e, 0xd6, 0x85, 0xb5,
	0x00, 0xf1, 0xa1, 0x9f, 0xde, 0x34, 0x42, 0xc2, 0x4e, 0xc8, 0xbc, 0xf9, 0x4a, 0xe0, 0x03, 0x38,
	0x36, 0x95, 0xaf, 0xf3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x23, 0x36, 0x2a, 0x84, 0x19, 0x02,
	0x00, 0x00,
}
