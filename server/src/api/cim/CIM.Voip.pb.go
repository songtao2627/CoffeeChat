// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CIM.Voip.proto

package cim

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 音视频呼叫邀请
type CIMVoipInviteReq struct {
	// cmd id:		0x0401
	CreatorUserId        uint64            `protobuf:"varint,1,opt,name=creator_user_id,json=creatorUserId,proto3" json:"creator_user_id,omitempty"`
	InviteUserList       []uint64          `protobuf:"varint,2,rep,packed,name=invite_user_list,json=inviteUserList,proto3" json:"invite_user_list,omitempty"`
	InviteType           CIMVoipInviteType `protobuf:"varint,3,opt,name=invite_type,json=inviteType,proto3,enum=CIM.Def.CIMVoipInviteType" json:"invite_type,omitempty"`
	ChannelInfo          *CIMChannelInfo   `protobuf:"bytes,4,opt,name=channel_info,json=channelInfo,proto3" json:"channel_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CIMVoipInviteReq) Reset()         { *m = CIMVoipInviteReq{} }
func (m *CIMVoipInviteReq) String() string { return proto.CompactTextString(m) }
func (*CIMVoipInviteReq) ProtoMessage()    {}
func (*CIMVoipInviteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{0}
}

func (m *CIMVoipInviteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipInviteReq.Unmarshal(m, b)
}
func (m *CIMVoipInviteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipInviteReq.Marshal(b, m, deterministic)
}
func (m *CIMVoipInviteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipInviteReq.Merge(m, src)
}
func (m *CIMVoipInviteReq) XXX_Size() int {
	return xxx_messageInfo_CIMVoipInviteReq.Size(m)
}
func (m *CIMVoipInviteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipInviteReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipInviteReq proto.InternalMessageInfo

func (m *CIMVoipInviteReq) GetCreatorUserId() uint64 {
	if m != nil {
		return m.CreatorUserId
	}
	return 0
}

func (m *CIMVoipInviteReq) GetInviteUserList() []uint64 {
	if m != nil {
		return m.InviteUserList
	}
	return nil
}

func (m *CIMVoipInviteReq) GetInviteType() CIMVoipInviteType {
	if m != nil {
		return m.InviteType
	}
	return CIMVoipInviteType_kCIM_VOIP_INVITE_TYPE_UNKNOWN
}

func (m *CIMVoipInviteReq) GetChannelInfo() *CIMChannelInfo {
	if m != nil {
		return m.ChannelInfo
	}
	return nil
}

// 音视频呼叫应答状态
type CIMVoipInviteReply struct {
	// cmd id:		0x0402
	UserId               uint64           `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RspCode              CIMInviteRspCode `protobuf:"varint,2,opt,name=rsp_code,json=rspCode,proto3,enum=CIM.Def.CIMInviteRspCode" json:"rsp_code,omitempty"`
	ChannelInfo          *CIMChannelInfo  `protobuf:"bytes,3,opt,name=channel_info,json=channelInfo,proto3" json:"channel_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CIMVoipInviteReply) Reset()         { *m = CIMVoipInviteReply{} }
func (m *CIMVoipInviteReply) String() string { return proto.CompactTextString(m) }
func (*CIMVoipInviteReply) ProtoMessage()    {}
func (*CIMVoipInviteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{1}
}

func (m *CIMVoipInviteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipInviteReply.Unmarshal(m, b)
}
func (m *CIMVoipInviteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipInviteReply.Marshal(b, m, deterministic)
}
func (m *CIMVoipInviteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipInviteReply.Merge(m, src)
}
func (m *CIMVoipInviteReply) XXX_Size() int {
	return xxx_messageInfo_CIMVoipInviteReply.Size(m)
}
func (m *CIMVoipInviteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipInviteReply.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipInviteReply proto.InternalMessageInfo

func (m *CIMVoipInviteReply) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMVoipInviteReply) GetRspCode() CIMInviteRspCode {
	if m != nil {
		return m.RspCode
	}
	return CIMInviteRspCode_kCIM_VOIP_INVITE_CODE_UNKNOWN
}

func (m *CIMVoipInviteReply) GetChannelInfo() *CIMChannelInfo {
	if m != nil {
		return m.ChannelInfo
	}
	return nil
}

// 心跳
type CIMVoipHeartbeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMVoipHeartbeat) Reset()         { *m = CIMVoipHeartbeat{} }
func (m *CIMVoipHeartbeat) String() string { return proto.CompactTextString(m) }
func (*CIMVoipHeartbeat) ProtoMessage()    {}
func (*CIMVoipHeartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{2}
}

func (m *CIMVoipHeartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipHeartbeat.Unmarshal(m, b)
}
func (m *CIMVoipHeartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipHeartbeat.Marshal(b, m, deterministic)
}
func (m *CIMVoipHeartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipHeartbeat.Merge(m, src)
}
func (m *CIMVoipHeartbeat) XXX_Size() int {
	return xxx_messageInfo_CIMVoipHeartbeat.Size(m)
}
func (m *CIMVoipHeartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipHeartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipHeartbeat proto.InternalMessageInfo

// 挂断请求
type CIMVoipByeReq struct {
	LocalCallTimeLen     uint64   `protobuf:"varint,1,opt,name=local_call_time_len,json=localCallTimeLen,proto3" json:"local_call_time_len,omitempty"`
	UserId               uint64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMVoipByeReq) Reset()         { *m = CIMVoipByeReq{} }
func (m *CIMVoipByeReq) String() string { return proto.CompactTextString(m) }
func (*CIMVoipByeReq) ProtoMessage()    {}
func (*CIMVoipByeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{3}
}

func (m *CIMVoipByeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipByeReq.Unmarshal(m, b)
}
func (m *CIMVoipByeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipByeReq.Marshal(b, m, deterministic)
}
func (m *CIMVoipByeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipByeReq.Merge(m, src)
}
func (m *CIMVoipByeReq) XXX_Size() int {
	return xxx_messageInfo_CIMVoipByeReq.Size(m)
}
func (m *CIMVoipByeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipByeReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipByeReq proto.InternalMessageInfo

func (m *CIMVoipByeReq) GetLocalCallTimeLen() uint64 {
	if m != nil {
		return m.LocalCallTimeLen
	}
	return 0
}

func (m *CIMVoipByeReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CIMVoipByeRsp struct {
	ByeReason            CIMVoipByeReason `protobuf:"varint,1,opt,name=byeReason,proto3,enum=CIM.Def.CIMVoipByeReason" json:"byeReason,omitempty"`
	UserId               uint64           `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CIMVoipByeRsp) Reset()         { *m = CIMVoipByeRsp{} }
func (m *CIMVoipByeRsp) String() string { return proto.CompactTextString(m) }
func (*CIMVoipByeRsp) ProtoMessage()    {}
func (*CIMVoipByeRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{4}
}

func (m *CIMVoipByeRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipByeRsp.Unmarshal(m, b)
}
func (m *CIMVoipByeRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipByeRsp.Marshal(b, m, deterministic)
}
func (m *CIMVoipByeRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipByeRsp.Merge(m, src)
}
func (m *CIMVoipByeRsp) XXX_Size() int {
	return xxx_messageInfo_CIMVoipByeRsp.Size(m)
}
func (m *CIMVoipByeRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipByeRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipByeRsp proto.InternalMessageInfo

func (m *CIMVoipByeRsp) GetByeReason() CIMVoipByeReason {
	if m != nil {
		return m.ByeReason
	}
	return CIMVoipByeReason_kCIM_VOIP_BYE_REASON_UNKNOWN
}

func (m *CIMVoipByeRsp) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*CIMVoipInviteReq)(nil), "CIM.Voip.CIMVoipInviteReq")
	proto.RegisterType((*CIMVoipInviteReply)(nil), "CIM.Voip.CIMVoipInviteReply")
	proto.RegisterType((*CIMVoipHeartbeat)(nil), "CIM.Voip.CIMVoipHeartbeat")
	proto.RegisterType((*CIMVoipByeReq)(nil), "CIM.Voip.CIMVoipByeReq")
	proto.RegisterType((*CIMVoipByeRsp)(nil), "CIM.Voip.CIMVoipByeRsp")
}

func init() { proto.RegisterFile("CIM.Voip.proto", fileDescriptor_d1f178c14f4443d3) }

var fileDescriptor_d1f178c14f4443d3 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x31, 0xcf, 0xd3, 0x30,
	0x10, 0x55, 0x9a, 0xea, 0x6b, 0x71, 0x69, 0xa9, 0xcc, 0xd0, 0xd0, 0x29, 0xca, 0x80, 0xb2, 0x90,
	0xa1, 0x20, 0x21, 0xc1, 0xd6, 0x30, 0x34, 0x52, 0xbb, 0x44, 0x05, 0x24, 0x96, 0xc8, 0x71, 0x2e,
	0xaa, 0x25, 0x27, 0x36, 0xb6, 0x8b, 0x94, 0x5f, 0xc3, 0x2f, 0xe3, 0xbf, 0xa0, 0xc4, 0x69, 0x9b,
	0x16, 0x31, 0x7c, 0xdb, 0xf9, 0xbd, 0x77, 0x77, 0xef, 0xe9, 0x8c, 0x16, 0x71, 0x72, 0x88, 0xbe,
	0x09, 0x26, 0x23, 0xa9, 0x84, 0x11, 0x78, 0x7a, 0x79, 0xaf, 0xe7, 0x6d, 0xf5, 0x05, 0x4a, 0x4b,
	0x04, 0x7f, 0x1c, 0xb4, 0x8c, 0x93, 0x43, 0x4b, 0x25, 0xf5, 0x2f, 0x66, 0x20, 0x85, 0x9f, 0xf8,
	0x2d, 0x7a, 0x45, 0x15, 0x10, 0x23, 0x54, 0x76, 0xd6, 0xa0, 0x32, 0x56, 0x78, 0x8e, 0xef, 0x84,
	0xe3, 0x74, 0xde, 0xc3, 0x5f, 0x35, 0xa8, 0xa4, 0xc0, 0x21, 0x5a, 0xb2, 0xae, 0xc9, 0xca, 0x38,
	0xd3, 0xc6, 0x1b, 0xf9, 0x6e, 0x38, 0x4e, 0x17, 0x16, 0x6f, 0x75, 0x7b, 0xa6, 0x0d, 0xfe, 0x8c,
	0x66, 0xbd, 0xd2, 0x34, 0x12, 0x3c, 0xd7, 0x77, 0xc2, 0xc5, 0x66, 0x1d, 0x5d, 0xbc, 0xdc, 0x39,
	0x38, 0x36, 0x12, 0x52, 0xc4, 0xae, 0x35, 0xfe, 0x84, 0x5e, 0xd2, 0x13, 0xa9, 0x6b, 0xe0, 0x19,
	0xab, 0x4b, 0xe1, 0x8d, 0x7d, 0x27, 0x9c, 0x6d, 0x56, 0xc3, 0xee, 0xd8, 0xf2, 0x49, 0x5d, 0x8a,
	0x74, 0x46, 0x6f, 0x8f, 0xe0, 0xb7, 0x83, 0xf0, 0x43, 0x3e, 0xc9, 0x1b, 0xbc, 0x42, 0x93, 0xfb,
	0x64, 0x4f, 0x67, 0x1b, 0xe9, 0x03, 0x9a, 0x2a, 0x2d, 0x33, 0x2a, 0x0a, 0xf0, 0x46, 0x9d, 0xcb,
	0x37, 0xc3, 0x3d, 0xfd, 0x0c, 0x2d, 0x63, 0x51, 0x40, 0x3a, 0x51, 0xb6, 0xf8, 0xc7, 0xa1, 0xfb,
	0x0c, 0x87, 0xf8, 0x7a, 0x80, 0x1d, 0x10, 0x65, 0x72, 0x20, 0x26, 0xf8, 0x8e, 0xe6, 0x3d, 0xb6,
	0x6d, 0xba, 0x8b, 0xbc, 0x43, 0xaf, 0xb9, 0xa0, 0x84, 0x67, 0x94, 0x70, 0x9e, 0x19, 0x56, 0x41,
	0xc6, 0xa1, 0xee, 0xbd, 0x2f, 0x3b, 0x2a, 0x26, 0x9c, 0x1f, 0x59, 0x05, 0x7b, 0xa8, 0x87, 0xf1,
	0x46, 0xc3, 0x78, 0x01, 0xb9, 0x1b, 0xac, 0x25, 0xfe, 0x88, 0x5e, 0xe4, 0xed, 0x0a, 0xa2, 0x85,
	0x1d, 0xf7, 0x10, 0xf8, 0xea, 0xa1, 0x15, 0xa4, 0x37, 0xed, 0x7f, 0x57, 0x6c, 0x7d, 0xb4, 0xa2,
	0xa2, 0x8a, 0xa8, 0x28, 0x4b, 0x00, 0x7a, 0x22, 0xc6, 0xfe, 0xb4, 0xfc, 0x5c, 0xee, 0xdc, 0x1f,
	0x2e, 0x65, 0x55, 0xfe, 0xd4, 0x01, 0xef, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x19, 0x1c, 0x31,
	0x9b, 0xa5, 0x02, 0x00, 0x00,
}