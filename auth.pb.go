// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TokenType int32

const (
	TokenType_ACCESS_TOKEN  TokenType = 0
	TokenType_REFRESH_TOKEN TokenType = 1
)

var TokenType_name = map[int32]string{
	0: "ACCESS_TOKEN",
	1: "REFRESH_TOKEN",
}

var TokenType_value = map[string]int32{
	"ACCESS_TOKEN":  0,
	"REFRESH_TOKEN": 1,
}

func (x TokenType) String() string {
	return proto.EnumName(TokenType_name, int32(x))
}

func (TokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

type AuthRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type AuthResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresIn            int64    `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	UserId               int64    `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *AuthResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *AuthResponse) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *AuthResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type VerifyRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type VerifyResponse struct {
	IsValid              bool     `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *VerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResponse.Unmarshal(m, b)
}
func (m *VerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResponse.Marshal(b, m, deterministic)
}
func (m *VerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResponse.Merge(m, src)
}
func (m *VerifyResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyResponse.Size(m)
}
func (m *VerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResponse proto.InternalMessageInfo

func (m *VerifyResponse) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

type RefreshRequest struct {
	RefreshToken         string   `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshRequest) Reset()         { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()    {}
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *RefreshRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshRequest.Unmarshal(m, b)
}
func (m *RefreshRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshRequest.Marshal(b, m, deterministic)
}
func (m *RefreshRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshRequest.Merge(m, src)
}
func (m *RefreshRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshRequest.Size(m)
}
func (m *RefreshRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshRequest proto.InternalMessageInfo

func (m *RefreshRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresIn            int64    `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	UserId               int64    `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshResponse) Reset()         { *m = RefreshResponse{} }
func (m *RefreshResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshResponse) ProtoMessage()    {}
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *RefreshResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshResponse.Unmarshal(m, b)
}
func (m *RefreshResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshResponse.Marshal(b, m, deterministic)
}
func (m *RefreshResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshResponse.Merge(m, src)
}
func (m *RefreshResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshResponse.Size(m)
}
func (m *RefreshResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshResponse proto.InternalMessageInfo

func (m *RefreshResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RefreshResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *RefreshResponse) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *RefreshResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ParseRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseRequest) Reset()         { *m = ParseRequest{} }
func (m *ParseRequest) String() string { return proto.CompactTextString(m) }
func (*ParseRequest) ProtoMessage()    {}
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *ParseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseRequest.Unmarshal(m, b)
}
func (m *ParseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseRequest.Marshal(b, m, deterministic)
}
func (m *ParseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseRequest.Merge(m, src)
}
func (m *ParseRequest) XXX_Size() int {
	return xxx_messageInfo_ParseRequest.Size(m)
}
func (m *ParseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParseRequest proto.InternalMessageInfo

func (m *ParseRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type ParseResponse struct {
	UserId               int64     `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TokenType            TokenType `protobuf:"varint,2,opt,name=token_type,json=tokenType,proto3,enum=auth.TokenType" json:"token_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ParseResponse) Reset()         { *m = ParseResponse{} }
func (m *ParseResponse) String() string { return proto.CompactTextString(m) }
func (*ParseResponse) ProtoMessage()    {}
func (*ParseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *ParseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseResponse.Unmarshal(m, b)
}
func (m *ParseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseResponse.Marshal(b, m, deterministic)
}
func (m *ParseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseResponse.Merge(m, src)
}
func (m *ParseResponse) XXX_Size() int {
	return xxx_messageInfo_ParseResponse.Size(m)
}
func (m *ParseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParseResponse proto.InternalMessageInfo

func (m *ParseResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ParseResponse) GetTokenType() TokenType {
	if m != nil {
		return m.TokenType
	}
	return TokenType_ACCESS_TOKEN
}

type JWK struct {
	Kty                  string   `protobuf:"bytes,1,opt,name=kty,proto3" json:"kty,omitempty"`
	E                    string   `protobuf:"bytes,2,opt,name=e,proto3" json:"e,omitempty"`
	N                    string   `protobuf:"bytes,3,opt,name=n,proto3" json:"n,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWK) Reset()         { *m = JWK{} }
func (m *JWK) String() string { return proto.CompactTextString(m) }
func (*JWK) ProtoMessage()    {}
func (*JWK) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{8}
}

func (m *JWK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWK.Unmarshal(m, b)
}
func (m *JWK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWK.Marshal(b, m, deterministic)
}
func (m *JWK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWK.Merge(m, src)
}
func (m *JWK) XXX_Size() int {
	return xxx_messageInfo_JWK.Size(m)
}
func (m *JWK) XXX_DiscardUnknown() {
	xxx_messageInfo_JWK.DiscardUnknown(m)
}

var xxx_messageInfo_JWK proto.InternalMessageInfo

func (m *JWK) GetKty() string {
	if m != nil {
		return m.Kty
	}
	return ""
}

func (m *JWK) GetE() string {
	if m != nil {
		return m.E
	}
	return ""
}

func (m *JWK) GetN() string {
	if m != nil {
		return m.N
	}
	return ""
}

type JwksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwksRequest) Reset()         { *m = JwksRequest{} }
func (m *JwksRequest) String() string { return proto.CompactTextString(m) }
func (*JwksRequest) ProtoMessage()    {}
func (*JwksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{9}
}

func (m *JwksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwksRequest.Unmarshal(m, b)
}
func (m *JwksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwksRequest.Marshal(b, m, deterministic)
}
func (m *JwksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwksRequest.Merge(m, src)
}
func (m *JwksRequest) XXX_Size() int {
	return xxx_messageInfo_JwksRequest.Size(m)
}
func (m *JwksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JwksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JwksRequest proto.InternalMessageInfo

type JwksResponse struct {
	Keys                 []*JWK   `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwksResponse) Reset()         { *m = JwksResponse{} }
func (m *JwksResponse) String() string { return proto.CompactTextString(m) }
func (*JwksResponse) ProtoMessage()    {}
func (*JwksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{10}
}

func (m *JwksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwksResponse.Unmarshal(m, b)
}
func (m *JwksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwksResponse.Marshal(b, m, deterministic)
}
func (m *JwksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwksResponse.Merge(m, src)
}
func (m *JwksResponse) XXX_Size() int {
	return xxx_messageInfo_JwksResponse.Size(m)
}
func (m *JwksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JwksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JwksResponse proto.InternalMessageInfo

func (m *JwksResponse) GetKeys() []*JWK {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterEnum("auth.TokenType", TokenType_name, TokenType_value)
	proto.RegisterType((*AuthRequest)(nil), "auth.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "auth.AuthResponse")
	proto.RegisterType((*VerifyRequest)(nil), "auth.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "auth.VerifyResponse")
	proto.RegisterType((*RefreshRequest)(nil), "auth.RefreshRequest")
	proto.RegisterType((*RefreshResponse)(nil), "auth.RefreshResponse")
	proto.RegisterType((*ParseRequest)(nil), "auth.ParseRequest")
	proto.RegisterType((*ParseResponse)(nil), "auth.ParseResponse")
	proto.RegisterType((*JWK)(nil), "auth.JWK")
	proto.RegisterType((*JwksRequest)(nil), "auth.JwksRequest")
	proto.RegisterType((*JwksResponse)(nil), "auth.JwksResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0xad, 0x63, 0x27, 0xa9, 0x6f, 0xec, 0x34, 0x9d, 0xbe, 0xa7, 0x06, 0x8b, 0xa2, 0xe0, 0x6e,
	0x2a, 0x20, 0x31, 0x4a, 0x85, 0x84, 0x44, 0x91, 0x68, 0xaa, 0x20, 0x9a, 0x48, 0x80, 0x9c, 0xa8,
	0x45, 0x6c, 0x2c, 0x37, 0x99, 0x24, 0x26, 0xc1, 0x36, 0x9e, 0x09, 0xc5, 0x5b, 0x3e, 0x00, 0x09,
	0xb1, 0xe3, 0x53, 0x58, 0xf1, 0x0f, 0xfc, 0x02, 0x12, 0xe2, 0x2f, 0x90, 0x67, 0xc6, 0xc1, 0x4e,
	0x59, 0xb1, 0x62, 0xe5, 0x99, 0x33, 0xf7, 0xcc, 0x3d, 0xe7, 0xf8, 0xda, 0x00, 0xee, 0x92, 0xce,
	0x5a, 0x61, 0x14, 0xd0, 0x00, 0x29, 0xc9, 0xda, 0xb8, 0x3e, 0x0d, 0x82, 0xe9, 0x02, 0x5b, 0x6e,
	0xe8, 0x59, 0xae, 0xef, 0x07, 0xd4, 0xa5, 0x5e, 0xe0, 0x13, 0x5e, 0x63, 0x3c, 0x98, 0x7a, 0x74,
	0xb6, 0xbc, 0x68, 0x8d, 0x82, 0xd7, 0xd6, 0x22, 0x9e, 0x50, 0x8b, 0xc1, 0xa3, 0xe6, 0x14, 0xfb,
	0xcd, 0xb7, 0xee, 0xc2, 0x1b, 0xbb, 0x14, 0x5b, 0x57, 0x16, 0x9c, 0x6c, 0xb6, 0xa1, 0x72, 0xbc,
	0xa4, 0x33, 0x1b, 0xbf, 0x59, 0x62, 0x42, 0xd1, 0x3e, 0x94, 0x97, 0x04, 0x47, 0x8e, 0x37, 0xae,
	0x4b, 0x0d, 0xe9, 0x40, 0xee, 0xc0, 0x97, 0x9f, 0x5f, 0xe5, 0xa2, 0x29, 0x37, 0x7e, 0x94, 0xed,
	0x52, 0x72, 0x74, 0x3a, 0x36, 0x3f, 0x48, 0xa0, 0x71, 0x12, 0x09, 0x03, 0x9f, 0x60, 0x74, 0x13,
	0x34, 0x77, 0x34, 0xc2, 0x84, 0x38, 0x34, 0x98, 0x63, 0x9f, 0x51, 0x55, 0xbb, 0xc2, 0xb1, 0x61,
	0x02, 0xa1, 0x7d, 0xd0, 0x23, 0x3c, 0x89, 0x30, 0x99, 0x89, 0x9a, 0x02, 0xab, 0xd1, 0x04, 0xc8,
	0x8b, 0xf6, 0x00, 0xf0, 0xbb, 0xd0, 0x8b, 0x30, 0x71, 0x3c, 0xbf, 0x2e, 0x27, 0x02, 0x6c, 0x55,
	0x20, 0xa7, 0x3e, 0xda, 0xfd, 0x2d, 0x4e, 0x61, 0x67, 0xa9, 0xa0, 0x87, 0xa0, 0x9f, 0xe1, 0xc8,
	0x9b, 0xc4, 0xa9, 0x8d, 0x3b, 0x7f, 0x12, 0xd4, 0x51, 0x13, 0x2f, 0x4a, 0x54, 0xa8, 0xdd, 0xc8,
	0x69, 0x33, 0x6f, 0x43, 0x35, 0xa5, 0x0b, 0x43, 0xd7, 0x60, 0xd3, 0x23, 0x0e, 0x8b, 0x8a, 0x71,
	0x37, 0xed, 0xb2, 0x47, 0xce, 0x92, 0xad, 0xf9, 0x08, 0xaa, 0x36, 0xd7, 0x9c, 0x36, 0x6b, 0xad,
	0x5b, 0xbb, 0xd2, 0x2d, 0xe7, 0xd2, 0xfc, 0x28, 0xc1, 0xd6, 0xea, 0x8a, 0x7f, 0x24, 0xc1, 0x23,
	0xd0, 0x9e, 0xbb, 0x11, 0xc1, 0x7f, 0x17, 0xe0, 0x0b, 0xd0, 0x05, 0x5b, 0xd8, 0xd9, 0x5d, 0x1b,
	0xa3, 0xb4, 0x0f, 0x6a, 0x01, 0xb0, 0x0b, 0x1d, 0x1a, 0x87, 0x98, 0x39, 0xa8, 0xb6, 0xb7, 0x5a,
	0x6c, 0xe0, 0xd9, 0x55, 0xc3, 0x38, 0xc4, 0xb6, 0x4a, 0xd3, 0xa5, 0x79, 0x08, 0x72, 0xef, 0xbc,
	0x8f, 0x6a, 0x20, 0xcf, 0x69, 0x2c, 0x52, 0x49, 0x96, 0x48, 0x03, 0x09, 0x8b, 0x04, 0x24, 0x9c,
	0xec, 0xb8, 0x5b, 0xd5, 0x96, 0x7c, 0x53, 0x87, 0x4a, 0xef, 0x72, 0x4e, 0x84, 0x17, 0xb3, 0x09,
	0x1a, 0xdf, 0x0a, 0x71, 0x7b, 0xa0, 0xcc, 0x71, 0x4c, 0xea, 0x52, 0x43, 0x3e, 0xa8, 0xb4, 0x55,
	0xde, 0xbd, 0x77, 0xde, 0xb7, 0x19, 0x7c, 0xeb, 0x2e, 0xa8, 0x2b, 0x29, 0xa8, 0x06, 0xda, 0xf1,
	0xc9, 0x49, 0x77, 0x30, 0x70, 0x86, 0xcf, 0xfa, 0xdd, 0xa7, 0xb5, 0x0d, 0xb4, 0x0d, 0xba, 0xdd,
	0x7d, 0x6c, 0x77, 0x07, 0x4f, 0x04, 0x24, 0xb5, 0x3f, 0x17, 0x40, 0x49, 0xbe, 0x07, 0x64, 0x89,
	0xe7, 0x36, 0xbf, 0x33, 0xf3, 0x61, 0x19, 0x28, 0x0b, 0x71, 0x21, 0xe6, 0x06, 0xba, 0x07, 0x25,
	0x3e, 0x79, 0x68, 0x87, 0x9f, 0xe7, 0xc6, 0xd8, 0xf8, 0x2f, 0x0f, 0xae, 0x68, 0xf7, 0xa1, 0x2c,
	0x06, 0x08, 0x89, 0x92, 0xfc, 0x48, 0x1a, 0xff, 0xaf, 0xa1, 0x2b, 0x66, 0x1b, 0x8a, 0xec, 0x4d,
	0x21, 0xa1, 0x27, 0xfb, 0xd2, 0x8d, 0x9d, 0x1c, 0xb6, 0xe2, 0x1c, 0x81, 0x92, 0xe4, 0x97, 0xba,
	0xca, 0x44, 0x9b, 0xba, 0xca, 0xc6, 0x6b, 0xea, 0xef, 0xbf, 0x7d, 0xff, 0x54, 0x28, 0xa3, 0xa2,
	0xf5, 0xea, 0x72, 0x4e, 0x3a, 0xc6, 0xcb, 0x7a, 0xe6, 0xff, 0x44, 0x5d, 0x3c, 0x0b, 0x9a, 0x5e,
	0x60, 0x25, 0xbc, 0x8b, 0x12, 0xfb, 0x07, 0x1d, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xa0,
	0xfb, 0xed, 0xf2, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
	Jwks(ctx context.Context, in *JwksRequest, opts ...grpc.CallOption) (*JwksResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Jwks(ctx context.Context, in *JwksRequest, opts ...grpc.CallOption) (*JwksResponse, error) {
	out := new(JwksResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/Jwks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
	Jwks(context.Context, *JwksRequest) (*JwksResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Jwks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JwksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Jwks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Jwks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Jwks(ctx, req.(*JwksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Auth_Auth_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Auth_Verify_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Auth_Refresh_Handler,
		},
		{
			MethodName: "Parse",
			Handler:    _Auth_Parse_Handler,
		},
		{
			MethodName: "Jwks",
			Handler:    _Auth_Jwks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
