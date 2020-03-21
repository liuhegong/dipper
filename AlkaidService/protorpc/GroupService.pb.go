// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GroupService.proto

package com_dipper_proto_rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type GroupPro struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GroupNum             int32    `protobuf:"varint,2,opt,name=group_num,json=groupNum,proto3" json:"group_num,omitempty"`
	GroupName            string   `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsDel                int32    `protobuf:"varint,5,opt,name=is_del,json=isDel,proto3" json:"is_del,omitempty"`
	IsRead               int32    `protobuf:"varint,6,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	SendTime             string   `protobuf:"bytes,7,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	IsBack               int32    `protobuf:"varint,8,opt,name=is_back,json=isBack,proto3" json:"is_back,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupPro) Reset()         { *m = GroupPro{} }
func (m *GroupPro) String() string { return proto.CompactTextString(m) }
func (*GroupPro) ProtoMessage()    {}
func (*GroupPro) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b14facae47a5964, []int{0}
}

func (m *GroupPro) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupPro.Unmarshal(m, b)
}
func (m *GroupPro) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupPro.Marshal(b, m, deterministic)
}
func (m *GroupPro) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupPro.Merge(m, src)
}
func (m *GroupPro) XXX_Size() int {
	return xxx_messageInfo_GroupPro.Size(m)
}
func (m *GroupPro) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupPro.DiscardUnknown(m)
}

var xxx_messageInfo_GroupPro proto.InternalMessageInfo

func (m *GroupPro) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GroupPro) GetGroupNum() int32 {
	if m != nil {
		return m.GroupNum
	}
	return 0
}

func (m *GroupPro) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *GroupPro) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *GroupPro) GetIsDel() int32 {
	if m != nil {
		return m.IsDel
	}
	return 0
}

func (m *GroupPro) GetIsRead() int32 {
	if m != nil {
		return m.IsRead
	}
	return 0
}

func (m *GroupPro) GetSendTime() string {
	if m != nil {
		return m.SendTime
	}
	return ""
}

func (m *GroupPro) GetIsBack() int32 {
	if m != nil {
		return m.IsBack
	}
	return 0
}

func init() {
	proto.RegisterType((*GroupPro)(nil), "com.dipper.proto.rpc.GroupPro")
}

func init() {
	proto.RegisterFile("GroupService.proto", fileDescriptor_4b14facae47a5964)
}

var fileDescriptor_4b14facae47a5964 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xd9, 0xd5, 0x4d, 0xb7, 0x83, 0xf4, 0x10, 0xfc, 0x13, 0x10, 0xa1, 0x88, 0x87, 0x9e,
	0xf6, 0xa0, 0x6f, 0xb0, 0x08, 0xde, 0x4a, 0x59, 0xbd, 0x2f, 0xd3, 0x64, 0x90, 0xa1, 0x4d, 0x13,
	0x92, 0xdd, 0x3e, 0xa9, 0x0f, 0x24, 0x9d, 0x6e, 0xc1, 0x83, 0xc7, 0xfc, 0x7e, 0x33, 0x1f, 0xf9,
	0x06, 0xf4, 0x47, 0x0a, 0x63, 0xfc, 0xa4, 0x74, 0x64, 0x4b, 0x4d, 0x4c, 0x61, 0x08, 0xfa, 0xd6,
	0x06, 0xdf, 0x38, 0x8e, 0x91, 0xd2, 0x99, 0x34, 0x29, 0xda, 0xe7, 0x9f, 0x02, 0x6a, 0x19, 0xde,
	0xa4, 0xa0, 0x17, 0x50, 0xb2, 0x33, 0xc5, 0xb2, 0x58, 0x55, 0x5d, 0xc9, 0x4e, 0x3f, 0xc2, 0xfc,
	0xfb, 0xe4, 0xfa, 0xc3, 0xe8, 0x4d, 0x29, 0xb8, 0x16, 0xb0, 0x1e, 0xbd, 0x7e, 0x02, 0x98, 0x24,
	0x7a, 0x32, 0x57, 0xcb, 0x62, 0x35, 0xef, 0xce, 0xe3, 0x6b, 0xf4, 0xa4, 0xef, 0x41, 0xe1, 0x11,
	0x07, 0x4c, 0xe6, 0x5a, 0xd4, 0xf4, 0xd2, 0x77, 0xa0, 0x38, 0xf7, 0x8e, 0xf6, 0xa6, 0x92, 0xc0,
	0x8a, 0xf3, 0x3b, 0xed, 0xf5, 0x03, 0xcc, 0x38, 0xf7, 0x89, 0xd0, 0x19, 0x25, 0x5c, 0x71, 0xee,
	0x08, 0xe5, 0x0f, 0x99, 0x0e, 0xae, 0x1f, 0xd8, 0x93, 0x99, 0x49, 0x54, 0x7d, 0x02, 0x5f, 0xec,
	0x69, 0xda, 0xda, 0xa2, 0xdd, 0x99, 0xfa, 0xb2, 0xd5, 0xa2, 0xdd, 0xbd, 0x2e, 0xe0, 0xe6, 0xef,
	0x09, 0xda, 0x17, 0xf8, 0xb7, 0x7e, 0x0b, 0x97, 0xee, 0x43, 0xd8, 0x14, 0x5b, 0x25, 0xe2, 0xed,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x86, 0xae, 0xf7, 0x11, 0x3f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GroupServiceClient is the client API for GroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupServiceClient interface {
}

type groupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupServiceClient(cc grpc.ClientConnInterface) GroupServiceClient {
	return &groupServiceClient{cc}
}

// GroupServiceServer is the server API for GroupService service.
type GroupServiceServer interface {
}

// UnimplementedGroupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGroupServiceServer struct {
}

func RegisterGroupServiceServer(s *grpc.Server, srv GroupServiceServer) {
	s.RegisterService(&_GroupService_serviceDesc, srv)
}

var _GroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.dipper.proto.rpc.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "GroupService.proto",
}