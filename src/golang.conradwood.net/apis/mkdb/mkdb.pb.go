// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/mkdb/mkdb.proto
// DO NOT EDIT!

/*
Package mkdb is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/mkdb/mkdb.proto

It has these top-level messages:
	ProtoDef
	ProtoField
	GetMessagesRequest
	AMessage
	GetMessageResponse
	CreateDBRequest
	CreateDBResponse
*/
package mkdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// preliminary definition...
type ProtoDef struct {
	Fields []*ProtoField `protobuf:"bytes,1,rep,name=Fields" json:"Fields,omitempty"`
	// e.g. "golang.conradwood.net/apis/common"
	ImportPath string `protobuf:"bytes,2,opt,name=ImportPath" json:"ImportPath,omitempty"`
	// e.g. Measurement
	Name string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
}

func (m *ProtoDef) Reset()                    { *m = ProtoDef{} }
func (m *ProtoDef) String() string            { return proto.CompactTextString(m) }
func (*ProtoDef) ProtoMessage()               {}
func (*ProtoDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoDef) GetFields() []*ProtoField {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *ProtoDef) GetImportPath() string {
	if m != nil {
		return m.ImportPath
	}
	return ""
}

func (m *ProtoDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ProtoField struct {
	Name       string            `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Type       int32             `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	PrimaryKey bool              `protobuf:"varint,3,opt,name=PrimaryKey" json:"PrimaryKey,omitempty"`
	Options    map[string]string `protobuf:"bytes,4,rep,name=Options" json:"Options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ProtoField) Reset()                    { *m = ProtoField{} }
func (m *ProtoField) String() string            { return proto.CompactTextString(m) }
func (*ProtoField) ProtoMessage()               {}
func (*ProtoField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProtoField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtoField) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ProtoField) GetPrimaryKey() bool {
	if m != nil {
		return m.PrimaryKey
	}
	return false
}

func (m *ProtoField) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetMessagesRequest struct {
	ProtoFile string `protobuf:"bytes,1,opt,name=ProtoFile" json:"ProtoFile,omitempty"`
}

func (m *GetMessagesRequest) Reset()                    { *m = GetMessagesRequest{} }
func (m *GetMessagesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMessagesRequest) ProtoMessage()               {}
func (*GetMessagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetMessagesRequest) GetProtoFile() string {
	if m != nil {
		return m.ProtoFile
	}
	return ""
}

type AMessage struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *AMessage) Reset()                    { *m = AMessage{} }
func (m *AMessage) String() string            { return proto.CompactTextString(m) }
func (*AMessage) ProtoMessage()               {}
func (*AMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetMessageResponse struct {
	Messages []*AMessage `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *GetMessageResponse) Reset()                    { *m = GetMessageResponse{} }
func (m *GetMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMessageResponse) ProtoMessage()               {}
func (*GetMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetMessageResponse) GetMessages() []*AMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type CreateDBRequest struct {
	ProtoFile     string `protobuf:"bytes,1,opt,name=ProtoFile" json:"ProtoFile,omitempty"`
	Message       string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Package       string `protobuf:"bytes,3,opt,name=Package" json:"Package,omitempty"`
	IDField       string `protobuf:"bytes,4,opt,name=IDField" json:"IDField,omitempty"`
	ImportPath    string `protobuf:"bytes,5,opt,name=ImportPath" json:"ImportPath,omitempty"`
	ProtoFileName string `protobuf:"bytes,6,opt,name=ProtoFileName" json:"ProtoFileName,omitempty"`
	TableName     string `protobuf:"bytes,7,opt,name=TableName" json:"TableName,omitempty"`
	TablePrefix   string `protobuf:"bytes,8,opt,name=TablePrefix" json:"TablePrefix,omitempty"`
}

func (m *CreateDBRequest) Reset()                    { *m = CreateDBRequest{} }
func (m *CreateDBRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDBRequest) ProtoMessage()               {}
func (*CreateDBRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateDBRequest) GetProtoFile() string {
	if m != nil {
		return m.ProtoFile
	}
	return ""
}

func (m *CreateDBRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateDBRequest) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *CreateDBRequest) GetIDField() string {
	if m != nil {
		return m.IDField
	}
	return ""
}

func (m *CreateDBRequest) GetImportPath() string {
	if m != nil {
		return m.ImportPath
	}
	return ""
}

func (m *CreateDBRequest) GetProtoFileName() string {
	if m != nil {
		return m.ProtoFileName
	}
	return ""
}

func (m *CreateDBRequest) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *CreateDBRequest) GetTablePrefix() string {
	if m != nil {
		return m.TablePrefix
	}
	return ""
}

type CreateDBResponse struct {
	GoFile string `protobuf:"bytes,1,opt,name=GoFile" json:"GoFile,omitempty"`
}

func (m *CreateDBResponse) Reset()                    { *m = CreateDBResponse{} }
func (m *CreateDBResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDBResponse) ProtoMessage()               {}
func (*CreateDBResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateDBResponse) GetGoFile() string {
	if m != nil {
		return m.GoFile
	}
	return ""
}

func init() {
	proto.RegisterType((*ProtoDef)(nil), "mkdb.ProtoDef")
	proto.RegisterType((*ProtoField)(nil), "mkdb.ProtoField")
	proto.RegisterType((*GetMessagesRequest)(nil), "mkdb.GetMessagesRequest")
	proto.RegisterType((*AMessage)(nil), "mkdb.AMessage")
	proto.RegisterType((*GetMessageResponse)(nil), "mkdb.GetMessageResponse")
	proto.RegisterType((*CreateDBRequest)(nil), "mkdb.CreateDBRequest")
	proto.RegisterType((*CreateDBResponse)(nil), "mkdb.CreateDBResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MKDB service

type MKDBClient interface {
	// given the contents of a .proto file will parse it and return the messages contained in the file
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessageResponse, error)
	// given a protofile and a message name will return Database Accessor helpers in a .go source file
	CreateDBFile(ctx context.Context, in *CreateDBRequest, opts ...grpc.CallOption) (*CreateDBResponse, error)
}

type mKDBClient struct {
	cc *grpc.ClientConn
}

func NewMKDBClient(cc *grpc.ClientConn) MKDBClient {
	return &mKDBClient{cc}
}

func (c *mKDBClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessageResponse, error) {
	out := new(GetMessageResponse)
	err := grpc.Invoke(ctx, "/mkdb.MKDB/GetMessages", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mKDBClient) CreateDBFile(ctx context.Context, in *CreateDBRequest, opts ...grpc.CallOption) (*CreateDBResponse, error) {
	out := new(CreateDBResponse)
	err := grpc.Invoke(ctx, "/mkdb.MKDB/CreateDBFile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MKDB service

type MKDBServer interface {
	// given the contents of a .proto file will parse it and return the messages contained in the file
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessageResponse, error)
	// given a protofile and a message name will return Database Accessor helpers in a .go source file
	CreateDBFile(context.Context, *CreateDBRequest) (*CreateDBResponse, error)
}

func RegisterMKDBServer(s *grpc.Server, srv MKDBServer) {
	s.RegisterService(&_MKDB_serviceDesc, srv)
}

func _MKDB_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MKDBServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mkdb.MKDB/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MKDBServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MKDB_CreateDBFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MKDBServer).CreateDBFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mkdb.MKDB/CreateDBFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MKDBServer).CreateDBFile(ctx, req.(*CreateDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MKDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mkdb.MKDB",
	HandlerType: (*MKDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessages",
			Handler:    _MKDB_GetMessages_Handler,
		},
		{
			MethodName: "CreateDBFile",
			Handler:    _MKDB_CreateDBFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.conradwood.net/apis/mkdb/mkdb.proto",
}

func init() { proto.RegisterFile("protos/golang.conradwood.net/apis/mkdb/mkdb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x13, 0x27, 0x71, 0x27, 0x05, 0xa2, 0x15, 0x54, 0x56, 0x04, 0x25, 0x58, 0x1c, 0xa2,
	0x1e, 0x1c, 0x11, 0x0e, 0xa0, 0x4a, 0x48, 0x34, 0x04, 0xaa, 0xaa, 0x2a, 0x58, 0x56, 0x5f, 0x60,
	0xd3, 0x4c, 0x53, 0x2b, 0x89, 0xd7, 0x78, 0xb7, 0x80, 0xcf, 0x5c, 0x78, 0x35, 0xde, 0x0a, 0xed,
	0x78, 0xed, 0x2c, 0x0e, 0x07, 0x2e, 0xd6, 0xce, 0xf7, 0x7d, 0xf3, 0xef, 0x81, 0x57, 0x59, 0x2e,
	0x94, 0x90, 0x93, 0x95, 0xd8, 0xf0, 0x74, 0x15, 0xde, 0x88, 0x34, 0xe7, 0xcb, 0xef, 0x42, 0x2c,
	0xc3, 0x14, 0xd5, 0x84, 0x67, 0x89, 0x9c, 0x6c, 0xd7, 0xcb, 0x05, 0x7d, 0x42, 0xd2, 0x32, 0x57,
	0xbf, 0x83, 0x3b, 0xf0, 0x22, 0x6d, 0xce, 0xf1, 0x96, 0x8d, 0xa1, 0xfb, 0x29, 0xc1, 0xcd, 0x52,
	0xfa, 0xce, 0xa8, 0x3d, 0xee, 0x4f, 0x07, 0x21, 0xc9, 0x89, 0x27, 0x22, 0x36, 0x3c, 0x3b, 0x06,
	0xb8, 0xd8, 0x66, 0x22, 0x57, 0x11, 0x57, 0x77, 0x7e, 0x6b, 0xe4, 0x8c, 0x0f, 0x62, 0x0b, 0x61,
	0x0c, 0xdc, 0xcf, 0x7c, 0x8b, 0x7e, 0x9b, 0x18, 0x7a, 0x07, 0xbf, 0x1d, 0x80, 0x5d, 0xa8, 0x5a,
	0xe2, 0xec, 0x24, 0x1a, 0xbb, 0x2e, 0x32, 0xa4, 0x80, 0x9d, 0x98, 0xde, 0x3a, 0x55, 0x94, 0x27,
	0x5b, 0x9e, 0x17, 0x97, 0x58, 0x50, 0x40, 0x2f, 0xb6, 0x10, 0xf6, 0x06, 0x7a, 0x5f, 0x32, 0x95,
	0x88, 0x54, 0xfa, 0x2e, 0x55, 0xfd, 0xac, 0x59, 0x75, 0x68, 0xf8, 0x8f, 0xa9, 0xca, 0x8b, 0xb8,
	0x52, 0x0f, 0x4f, 0xe1, 0xd0, 0x26, 0xd8, 0x00, 0xda, 0x6b, 0x2c, 0x4c, 0x3d, 0xfa, 0xc9, 0x1e,
	0x43, 0xe7, 0x1b, 0xdf, 0xdc, 0xa3, 0x69, 0xb0, 0x34, 0x4e, 0x5b, 0x6f, 0x9d, 0x60, 0x0a, 0xec,
	0x1c, 0xd5, 0x15, 0x4a, 0xc9, 0x57, 0x28, 0x63, 0xfc, 0x7a, 0x8f, 0x52, 0xb1, 0xa7, 0x70, 0x60,
	0xb2, 0x6e, 0xaa, 0xbe, 0x76, 0x40, 0x70, 0x0c, 0xde, 0x99, 0xf1, 0xf8, 0x57, 0xf3, 0xc1, 0x7b,
	0x3b, 0x66, 0x8c, 0x32, 0x13, 0xa9, 0x44, 0x76, 0x02, 0x5e, 0x95, 0xc6, 0x6c, 0xe5, 0x61, 0xd9,
	0x5f, 0x15, 0x2b, 0xae, 0xf9, 0xe0, 0x67, 0x0b, 0x1e, 0x7d, 0xc8, 0x91, 0x2b, 0x9c, 0xcf, 0xfe,
	0xab, 0x26, 0xe6, 0x43, 0xcf, 0x78, 0x9b, 0x1e, 0x2b, 0x53, 0x33, 0x11, 0xbf, 0x59, 0x6b, 0xa6,
	0x5c, 0x62, 0x65, 0x6a, 0xe6, 0x62, 0x4e, 0x83, 0xf5, 0xdd, 0x92, 0x31, 0x66, 0xe3, 0xaf, 0xe8,
	0xec, 0xfd, 0x15, 0x2f, 0xe1, 0x41, 0x9d, 0x9a, 0xda, 0xef, 0x92, 0xe4, 0x6f, 0x50, 0x57, 0x7c,
	0xcd, 0x17, 0x46, 0xd1, 0x2b, 0x2b, 0xae, 0x01, 0x36, 0x82, 0x3e, 0x19, 0x51, 0x8e, 0xb7, 0xc9,
	0x0f, 0xdf, 0x23, 0xde, 0x86, 0x82, 0x13, 0x18, 0xec, 0x86, 0x60, 0xa6, 0x78, 0x04, 0xdd, 0x73,
	0x7b, 0x04, 0xc6, 0x9a, 0xfe, 0x72, 0xc0, 0xbd, 0xba, 0x9c, 0xcf, 0xd8, 0x19, 0xf4, 0xad, 0x85,
	0x32, 0xbf, 0x9c, 0xf1, 0xfe, 0x8e, 0x87, 0x7b, 0x4c, 0x9d, 0xe3, 0x1d, 0x1c, 0x56, 0x79, 0x69,
	0xb6, 0x4f, 0x4a, 0x65, 0x63, 0x21, 0xc3, 0xa3, 0x26, 0x5c, 0xba, 0xcf, 0x5e, 0xc0, 0xf3, 0x14,
	0x95, 0x7d, 0xb9, 0xe6, 0x96, 0xf5, 0xf1, 0x92, 0xcf, 0xa2, 0x4b, 0x87, 0xfb, 0xfa, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x57, 0x29, 0xcd, 0x89, 0xed, 0x03, 0x00, 0x00,
}