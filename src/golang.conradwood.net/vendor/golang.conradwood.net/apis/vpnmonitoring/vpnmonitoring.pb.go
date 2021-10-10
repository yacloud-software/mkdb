// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/vpnmonitoring/vpnmonitoring.proto
// DO NOT EDIT!

/*
Package vpnmonitoring is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/vpnmonitoring/vpnmonitoring.proto

It has these top-level messages:
*/
package vpnmonitoring

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for VpnMonitoring service

type VpnMonitoringClient interface {
}

type vpnMonitoringClient struct {
	cc *grpc.ClientConn
}

func NewVpnMonitoringClient(cc *grpc.ClientConn) VpnMonitoringClient {
	return &vpnMonitoringClient{cc}
}

// Server API for VpnMonitoring service

type VpnMonitoringServer interface {
}

func RegisterVpnMonitoringServer(s *grpc.Server, srv VpnMonitoringServer) {
	s.RegisterService(&_VpnMonitoring_serviceDesc, srv)
}

var _VpnMonitoring_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vpnmonitoring.VpnMonitoring",
	HandlerType: (*VpnMonitoringServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "golang.conradwood.net/apis/vpnmonitoring/vpnmonitoring.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/vpnmonitoring/vpnmonitoring.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0xcf, 0x49,
	0xcc, 0x4b, 0xd7, 0x4b, 0xce, 0xcf, 0x2b, 0x4a, 0x4c, 0x29, 0xcf, 0xcf, 0x4f, 0xd1, 0xcb, 0x4b,
	0x2d, 0xd1, 0x4f, 0x2c, 0xc8, 0x2c, 0xd6, 0x2f, 0x2b, 0xc8, 0xcb, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9,
	0x2f, 0xca, 0xcc, 0x4b, 0x47, 0xe5, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa2, 0x08,
	0x1a, 0xf1, 0x73, 0xf1, 0x86, 0x15, 0xe4, 0xf9, 0xc2, 0x05, 0x9c, 0xb4, 0xb8, 0x34, 0xf2, 0x52,
	0x4b, 0x90, 0x0d, 0x87, 0x5a, 0x07, 0x32, 0x5f, 0x0f, 0x45, 0x73, 0x12, 0x1b, 0xd8, 0x48, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xbf, 0x8f, 0x8c, 0x92, 0x00, 0x00, 0x00,
}
