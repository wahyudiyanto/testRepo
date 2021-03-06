// Code generated by protoc-gen-go. DO NOT EDIT.
// source: obat.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	obat.proto

It has these top-level messages:
	AddObatReq
	ReadObatByNamaReq
	ReadObatByNamaResp
	ReadObatResp
	UpdateObatReq
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type AddObatReq struct {
	Ido                int32  `protobuf:"varint,1,opt,name=Ido" json:"Ido,omitempty"`
	KodeObat           string `protobuf:"bytes,2,opt,name=KodeObat" json:"KodeObat,omitempty"`
	NamaObat           string `protobuf:"bytes,3,opt,name=NamaObat" json:"NamaObat,omitempty"`
	TanggalKadaluwarsa string `protobuf:"bytes,4,opt,name=TanggalKadaluwarsa" json:"TanggalKadaluwarsa,omitempty"`
	Harga              int64  `protobuf:"varint,5,opt,name=Harga" json:"Harga,omitempty"`
	CreatedBy          string `protobuf:"bytes,6,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn          string `protobuf:"bytes,7,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdateBy           string `protobuf:"bytes,8,opt,name=UpdateBy" json:"UpdateBy,omitempty"`
	UpdateOn           string `protobuf:"bytes,9,opt,name=UpdateOn" json:"UpdateOn,omitempty"`
	Status             int32  `protobuf:"varint,10,opt,name=Status" json:"Status,omitempty"`
}

func (m *AddObatReq) Reset()                    { *m = AddObatReq{} }
func (m *AddObatReq) String() string            { return proto.CompactTextString(m) }
func (*AddObatReq) ProtoMessage()               {}
func (*AddObatReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddObatReq) GetIdo() int32 {
	if m != nil {
		return m.Ido
	}
	return 0
}

func (m *AddObatReq) GetKodeObat() string {
	if m != nil {
		return m.KodeObat
	}
	return ""
}

func (m *AddObatReq) GetNamaObat() string {
	if m != nil {
		return m.NamaObat
	}
	return ""
}

func (m *AddObatReq) GetTanggalKadaluwarsa() string {
	if m != nil {
		return m.TanggalKadaluwarsa
	}
	return ""
}

func (m *AddObatReq) GetHarga() int64 {
	if m != nil {
		return m.Harga
	}
	return 0
}

func (m *AddObatReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *AddObatReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *AddObatReq) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *AddObatReq) GetUpdateOn() string {
	if m != nil {
		return m.UpdateOn
	}
	return ""
}

func (m *AddObatReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReadObatByNamaReq struct {
	NamaObat string `protobuf:"bytes,3,opt,name=NamaObat" json:"NamaObat,omitempty"`
}

func (m *ReadObatByNamaReq) Reset()                    { *m = ReadObatByNamaReq{} }
func (m *ReadObatByNamaReq) String() string            { return proto.CompactTextString(m) }
func (*ReadObatByNamaReq) ProtoMessage()               {}
func (*ReadObatByNamaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadObatByNamaReq) GetNamaObat() string {
	if m != nil {
		return m.NamaObat
	}
	return ""
}

type ReadObatByNamaResp struct {
	Ido                int32  `protobuf:"varint,1,opt,name=Ido" json:"Ido,omitempty"`
	KodeObat           string `protobuf:"bytes,2,opt,name=KodeObat" json:"KodeObat,omitempty"`
	NamaObat           string `protobuf:"bytes,3,opt,name=NamaObat" json:"NamaObat,omitempty"`
	TanggalKadaluwarsa string `protobuf:"bytes,4,opt,name=TanggalKadaluwarsa" json:"TanggalKadaluwarsa,omitempty"`
	Harga              int64  `protobuf:"varint,5,opt,name=Harga" json:"Harga,omitempty"`
	CreatedBy          string `protobuf:"bytes,6,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn          string `protobuf:"bytes,7,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdateBy           string `protobuf:"bytes,8,opt,name=UpdateBy" json:"UpdateBy,omitempty"`
	UpdateOn           string `protobuf:"bytes,9,opt,name=UpdateOn" json:"UpdateOn,omitempty"`
	Status             int32  `protobuf:"varint,10,opt,name=Status" json:"Status,omitempty"`
}

func (m *ReadObatByNamaResp) Reset()                    { *m = ReadObatByNamaResp{} }
func (m *ReadObatByNamaResp) String() string            { return proto.CompactTextString(m) }
func (*ReadObatByNamaResp) ProtoMessage()               {}
func (*ReadObatByNamaResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadObatByNamaResp) GetIdo() int32 {
	if m != nil {
		return m.Ido
	}
	return 0
}

func (m *ReadObatByNamaResp) GetKodeObat() string {
	if m != nil {
		return m.KodeObat
	}
	return ""
}

func (m *ReadObatByNamaResp) GetNamaObat() string {
	if m != nil {
		return m.NamaObat
	}
	return ""
}

func (m *ReadObatByNamaResp) GetTanggalKadaluwarsa() string {
	if m != nil {
		return m.TanggalKadaluwarsa
	}
	return ""
}

func (m *ReadObatByNamaResp) GetHarga() int64 {
	if m != nil {
		return m.Harga
	}
	return 0
}

func (m *ReadObatByNamaResp) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *ReadObatByNamaResp) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *ReadObatByNamaResp) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *ReadObatByNamaResp) GetUpdateOn() string {
	if m != nil {
		return m.UpdateOn
	}
	return ""
}

func (m *ReadObatByNamaResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReadObatResp struct {
	AllObat []*ReadObatByNamaResp `protobuf:"bytes,1,rep,name=allObat" json:"allObat,omitempty"`
}

func (m *ReadObatResp) Reset()                    { *m = ReadObatResp{} }
func (m *ReadObatResp) String() string            { return proto.CompactTextString(m) }
func (*ReadObatResp) ProtoMessage()               {}
func (*ReadObatResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadObatResp) GetAllObat() []*ReadObatByNamaResp {
	if m != nil {
		return m.AllObat
	}
	return nil
}

type UpdateObatReq struct {
	Ido                int32  `protobuf:"varint,1,opt,name=Ido" json:"Ido,omitempty"`
	KodeObat           string `protobuf:"bytes,2,opt,name=KodeObat" json:"KodeObat,omitempty"`
	NamaObat           string `protobuf:"bytes,3,opt,name=NamaObat" json:"NamaObat,omitempty"`
	TanggalKadaluwarsa string `protobuf:"bytes,4,opt,name=TanggalKadaluwarsa" json:"TanggalKadaluwarsa,omitempty"`
	Harga              int64  `protobuf:"varint,5,opt,name=Harga" json:"Harga,omitempty"`
	CreatedBy          string `protobuf:"bytes,6,opt,name=CreatedBy" json:"CreatedBy,omitempty"`
	CreatedOn          string `protobuf:"bytes,7,opt,name=CreatedOn" json:"CreatedOn,omitempty"`
	UpdateBy           string `protobuf:"bytes,8,opt,name=UpdateBy" json:"UpdateBy,omitempty"`
	UpdateOn           string `protobuf:"bytes,9,opt,name=UpdateOn" json:"UpdateOn,omitempty"`
	Status             int32  `protobuf:"varint,10,opt,name=Status" json:"Status,omitempty"`
}

func (m *UpdateObatReq) Reset()                    { *m = UpdateObatReq{} }
func (m *UpdateObatReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateObatReq) ProtoMessage()               {}
func (*UpdateObatReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateObatReq) GetIdo() int32 {
	if m != nil {
		return m.Ido
	}
	return 0
}

func (m *UpdateObatReq) GetKodeObat() string {
	if m != nil {
		return m.KodeObat
	}
	return ""
}

func (m *UpdateObatReq) GetNamaObat() string {
	if m != nil {
		return m.NamaObat
	}
	return ""
}

func (m *UpdateObatReq) GetTanggalKadaluwarsa() string {
	if m != nil {
		return m.TanggalKadaluwarsa
	}
	return ""
}

func (m *UpdateObatReq) GetHarga() int64 {
	if m != nil {
		return m.Harga
	}
	return 0
}

func (m *UpdateObatReq) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *UpdateObatReq) GetCreatedOn() string {
	if m != nil {
		return m.CreatedOn
	}
	return ""
}

func (m *UpdateObatReq) GetUpdateBy() string {
	if m != nil {
		return m.UpdateBy
	}
	return ""
}

func (m *UpdateObatReq) GetUpdateOn() string {
	if m != nil {
		return m.UpdateOn
	}
	return ""
}

func (m *UpdateObatReq) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*AddObatReq)(nil), "grpc.AddObatReq")
	proto.RegisterType((*ReadObatByNamaReq)(nil), "grpc.ReadObatByNamaReq")
	proto.RegisterType((*ReadObatByNamaResp)(nil), "grpc.ReadObatByNamaResp")
	proto.RegisterType((*ReadObatResp)(nil), "grpc.ReadObatResp")
	proto.RegisterType((*UpdateObatReq)(nil), "grpc.UpdateObatReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for ObatService service

type ObatServiceClient interface {
	AddObat(ctx context.Context, in *AddObatReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadObatByNama(ctx context.Context, in *ReadObatByNamaReq, opts ...grpc1.CallOption) (*ReadObatByNamaResp, error)
	ReadObat(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadObatResp, error)
	UpdateObat(ctx context.Context, in *UpdateObatReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
}

type obatServiceClient struct {
	cc *grpc1.ClientConn
}

func NewObatServiceClient(cc *grpc1.ClientConn) ObatServiceClient {
	return &obatServiceClient{cc}
}

func (c *obatServiceClient) AddObat(ctx context.Context, in *AddObatReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.ObatService/AddObat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *obatServiceClient) ReadObatByNama(ctx context.Context, in *ReadObatByNamaReq, opts ...grpc1.CallOption) (*ReadObatByNamaResp, error) {
	out := new(ReadObatByNamaResp)
	err := grpc1.Invoke(ctx, "/grpc.ObatService/ReadObatByNama", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *obatServiceClient) ReadObat(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadObatResp, error) {
	out := new(ReadObatResp)
	err := grpc1.Invoke(ctx, "/grpc.ObatService/ReadObat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *obatServiceClient) UpdateObat(ctx context.Context, in *UpdateObatReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.ObatService/UpdateObat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ObatService service

type ObatServiceServer interface {
	AddObat(context.Context, *AddObatReq) (*google_protobuf.Empty, error)
	ReadObatByNama(context.Context, *ReadObatByNamaReq) (*ReadObatByNamaResp, error)
	ReadObat(context.Context, *google_protobuf.Empty) (*ReadObatResp, error)
	UpdateObat(context.Context, *UpdateObatReq) (*google_protobuf.Empty, error)
}

func RegisterObatServiceServer(s *grpc1.Server, srv ObatServiceServer) {
	s.RegisterService(&_ObatService_serviceDesc, srv)
}

func _ObatService_AddObat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddObatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObatServiceServer).AddObat(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ObatService/AddObat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObatServiceServer).AddObat(ctx, req.(*AddObatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObatService_ReadObatByNama_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadObatByNamaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObatServiceServer).ReadObatByNama(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ObatService/ReadObatByNama",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObatServiceServer).ReadObatByNama(ctx, req.(*ReadObatByNamaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObatService_ReadObat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObatServiceServer).ReadObat(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ObatService/ReadObat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObatServiceServer).ReadObat(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObatService_UpdateObat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObatServiceServer).UpdateObat(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ObatService/UpdateObat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObatServiceServer).UpdateObat(ctx, req.(*UpdateObatReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ObatService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.ObatService",
	HandlerType: (*ObatServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddObat",
			Handler:    _ObatService_AddObat_Handler,
		},
		{
			MethodName: "ReadObatByNama",
			Handler:    _ObatService_ReadObatByNama_Handler,
		},
		{
			MethodName: "ReadObat",
			Handler:    _ObatService_ReadObat_Handler,
		},
		{
			MethodName: "UpdateObat",
			Handler:    _ObatService_UpdateObat_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "obat.proto",
}

func init() { proto.RegisterFile("obat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xcf, 0xea, 0xd3, 0x40,
	0x10, 0xc7, 0x9b, 0xa4, 0x7f, 0xa7, 0x2a, 0x75, 0x94, 0xba, 0x44, 0x0f, 0x21, 0xa7, 0x9c, 0x52,
	0xa8, 0x08, 0x82, 0x27, 0x23, 0x05, 0xa5, 0x60, 0x20, 0xd5, 0x07, 0x98, 0x34, 0x6b, 0x10, 0xd2,
	0x24, 0x4d, 0xb7, 0x4a, 0xae, 0x3e, 0x8d, 0xe0, 0xc5, 0x47, 0x94, 0xdd, 0x24, 0xa6, 0xf9, 0xfd,
	0xda, 0xfe, 0x5e, 0xa0, 0xb7, 0x9d, 0xf9, 0xce, 0xb0, 0x33, 0x9f, 0x19, 0x06, 0x20, 0x0b, 0x49,
	0xb8, 0x79, 0x91, 0x89, 0x0c, 0xfb, 0x71, 0x91, 0x6f, 0xcd, 0x97, 0x71, 0x96, 0xc5, 0x09, 0x5f,
	0x28, 0x5f, 0x78, 0xfc, 0xb6, 0xe0, 0xbb, 0x5c, 0x94, 0x55, 0x88, 0xfd, 0x5b, 0x07, 0x78, 0x1f,
	0x45, 0x7e, 0x48, 0x22, 0xe0, 0x7b, 0x9c, 0x81, 0xf1, 0x29, 0xca, 0x98, 0x66, 0x69, 0xce, 0x20,
	0x90, 0x4f, 0x34, 0x61, 0xbc, 0xce, 0x22, 0x2e, 0x03, 0x98, 0x6e, 0x69, 0xce, 0x24, 0xf8, 0x6f,
	0x4b, 0xed, 0x33, 0xed, 0x48, 0x69, 0x46, 0xa5, 0x35, 0x36, 0xba, 0x80, 0x5f, 0x28, 0x8d, 0x63,
	0x4a, 0xd6, 0x14, 0x51, 0x72, 0xfc, 0x49, 0xc5, 0x81, 0x58, 0x5f, 0x45, 0x9d, 0x51, 0xf0, 0x39,
	0x0c, 0x3e, 0x52, 0x11, 0x13, 0x1b, 0x58, 0x9a, 0x63, 0x04, 0x95, 0x81, 0xaf, 0x60, 0xf2, 0xa1,
	0xe0, 0x24, 0x78, 0xe4, 0x95, 0x6c, 0xa8, 0x92, 0x5b, 0xc7, 0x89, 0xea, 0xa7, 0x6c, 0xd4, 0x51,
	0xfd, 0x54, 0x56, 0xf7, 0x35, 0x8f, 0x48, 0x70, 0xaf, 0x64, 0xe3, 0xaa, 0xba, 0xc6, 0x6e, 0x35,
	0x3f, 0x65, 0x93, 0x53, 0xcd, 0x4f, 0x71, 0x0e, 0xc3, 0x8d, 0x20, 0x71, 0x3c, 0x30, 0x50, 0x18,
	0x6a, 0xcb, 0x5e, 0xc0, 0xd3, 0x80, 0x93, 0x42, 0xe5, 0x95, 0xb2, 0x4f, 0x09, 0xec, 0x0a, 0x02,
	0xfb, 0xaf, 0x0e, 0x78, 0x37, 0xe3, 0x90, 0xdf, 0x18, 0x5f, 0x61, 0xec, 0xc1, 0xa3, 0x86, 0x98,
	0x62, 0xb5, 0x84, 0x11, 0x25, 0x89, 0x6a, 0x5e, 0xb3, 0x0c, 0x67, 0xba, 0x64, 0xae, 0xdc, 0x69,
	0xf7, 0x3e, 0xd6, 0xa0, 0x09, 0xb4, 0xff, 0xe8, 0xf0, 0xb8, 0xfe, 0xe8, 0xb6, 0xd5, 0x0f, 0x11,
	0x5f, 0xfe, 0xd2, 0x61, 0x2a, 0x5b, 0xdb, 0xf0, 0xe2, 0xc7, 0xf7, 0x2d, 0xc7, 0x37, 0x30, 0xaa,
	0xef, 0x01, 0xce, 0x2a, 0xd6, 0xed, 0x79, 0x30, 0xe7, 0x6e, 0x75, 0x4b, 0xdc, 0xe6, 0x96, 0xb8,
	0x2b, 0x79, 0x4b, 0xec, 0x1e, 0xae, 0xe0, 0x49, 0x77, 0x26, 0xf8, 0xe2, 0xfc, 0xa4, 0xf6, 0xe6,
	0xc5, 0x11, 0xda, 0x3d, 0x7c, 0x0b, 0xe3, 0xc6, 0x8f, 0x17, 0x3e, 0x33, 0xb1, 0x9b, 0x5f, 0x67,
	0xbe, 0x03, 0x68, 0x87, 0x8e, 0xcf, 0xaa, 0x98, 0xce, 0x1a, 0x5c, 0xae, 0x3e, 0x1c, 0x2a, 0xcf,
	0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x63, 0xb9, 0x9b, 0x3d, 0x05, 0x00, 0x00,
}
