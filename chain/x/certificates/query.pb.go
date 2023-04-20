// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: empowerchain/certificates/query.proto

package certificates

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83edfea194a14431, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83edfea194a14431, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryCertificateRequest struct {
	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *QueryCertificateRequest) Reset()         { *m = QueryCertificateRequest{} }
func (m *QueryCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCertificateRequest) ProtoMessage()    {}
func (*QueryCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83edfea194a14431, []int{2}
}
func (m *QueryCertificateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertificateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertificateRequest.Merge(m, src)
}
func (m *QueryCertificateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertificateRequest proto.InternalMessageInfo

func (m *QueryCertificateRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueryCertificateRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type QueryCertificateResponse struct {
	Certificate Certificate `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate"`
}

func (m *QueryCertificateResponse) Reset()         { *m = QueryCertificateResponse{} }
func (m *QueryCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCertificateResponse) ProtoMessage()    {}
func (*QueryCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83edfea194a14431, []int{3}
}
func (m *QueryCertificateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertificateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertificateResponse.Merge(m, src)
}
func (m *QueryCertificateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertificateResponse proto.InternalMessageInfo

func (m *QueryCertificateResponse) GetCertificate() Certificate {
	if m != nil {
		return m.Certificate
	}
	return Certificate{}
}

type QueryAllCertificatesByUserRequest struct {
	Owner      string            `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Pagination query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination"`
}

func (m *QueryAllCertificatesByUserRequest) Reset()         { *m = QueryAllCertificatesByUserRequest{} }
func (m *QueryAllCertificatesByUserRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCertificatesByUserRequest) ProtoMessage()    {}
func (*QueryAllCertificatesByUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83edfea194a14431, []int{4}
}
func (m *QueryAllCertificatesByUserRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCertificatesByUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCertificatesByUserRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCertificatesByUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCertificatesByUserRequest.Merge(m, src)
}
func (m *QueryAllCertificatesByUserRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCertificatesByUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCertificatesByUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCertificatesByUserRequest proto.InternalMessageInfo

func (m *QueryAllCertificatesByUserRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *QueryAllCertificatesByUserRequest) GetPagination() query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return query.PageRequest{}
}

type QueryAllCertificatesByUserResponse struct {
	Certificates []Certificate      `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates"`
	Pagination   query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination"`
}

func (m *QueryAllCertificatesByUserResponse) Reset()         { *m = QueryAllCertificatesByUserResponse{} }
func (m *QueryAllCertificatesByUserResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCertificatesByUserResponse) ProtoMessage()    {}
func (*QueryAllCertificatesByUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83edfea194a14431, []int{5}
}
func (m *QueryAllCertificatesByUserResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCertificatesByUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCertificatesByUserResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCertificatesByUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCertificatesByUserResponse.Merge(m, src)
}
func (m *QueryAllCertificatesByUserResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCertificatesByUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCertificatesByUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCertificatesByUserResponse proto.InternalMessageInfo

func (m *QueryAllCertificatesByUserResponse) GetCertificates() []Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *QueryAllCertificatesByUserResponse) GetPagination() query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return query.PageResponse{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "empowerchain.certificates.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "empowerchain.certificates.QueryParamsResponse")
	proto.RegisterType((*QueryCertificateRequest)(nil), "empowerchain.certificates.QueryCertificateRequest")
	proto.RegisterType((*QueryCertificateResponse)(nil), "empowerchain.certificates.QueryCertificateResponse")
	proto.RegisterType((*QueryAllCertificatesByUserRequest)(nil), "empowerchain.certificates.QueryAllCertificatesByUserRequest")
	proto.RegisterType((*QueryAllCertificatesByUserResponse)(nil), "empowerchain.certificates.QueryAllCertificatesByUserResponse")
}

func init() {
	proto.RegisterFile("empowerchain/certificates/query.proto", fileDescriptor_83edfea194a14431)
}

var fileDescriptor_83edfea194a14431 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xe3, 0xd0, 0x46, 0x62, 0x8a, 0x38, 0x98, 0x48, 0x6c, 0x23, 0xb4, 0xa4, 0x2b, 0x01,
	0x15, 0x52, 0x6d, 0x1a, 0x84, 0xc4, 0x0d, 0x35, 0x88, 0x1b, 0x7f, 0x42, 0x10, 0x48, 0x70, 0x41,
	0xce, 0xc6, 0x6c, 0x8d, 0x92, 0xf5, 0xd6, 0x76, 0x28, 0xb9, 0xf2, 0x04, 0x48, 0x3c, 0x41, 0xcf,
	0x5c, 0x79, 0x05, 0xa4, 0x1e, 0x2b, 0xb8, 0x70, 0x42, 0x28, 0xe1, 0x41, 0x50, 0x6c, 0x97, 0x3a,
	0x0a, 0x5d, 0xd2, 0x5b, 0x62, 0x7f, 0xf3, 0xcd, 0xf7, 0x9b, 0x1d, 0xc3, 0x35, 0x3e, 0x2c, 0xe4,
	0x3e, 0x57, 0xe9, 0x2e, 0x13, 0x39, 0x4d, 0xb9, 0x32, 0xe2, 0x8d, 0x48, 0x99, 0xe1, 0x9a, 0xee,
	0x8d, 0xb8, 0x1a, 0x93, 0x42, 0x49, 0x23, 0xf1, 0x7a, 0x28, 0x23, 0xa1, 0xac, 0x51, 0xcf, 0x64,
	0x26, 0xad, 0x8a, 0xce, 0x7e, 0xb9, 0x82, 0xc6, 0x95, 0x4c, 0xca, 0x6c, 0xc0, 0x29, 0x2b, 0x04,
	0x65, 0x79, 0x2e, 0x0d, 0x33, 0x42, 0xe6, 0xda, 0xdf, 0xde, 0x4c, 0xa5, 0x1e, 0x4a, 0x4d, 0x7b,
	0x4c, 0x73, 0xd7, 0x87, 0xbe, 0xdb, 0xee, 0x71, 0xc3, 0xb6, 0x69, 0xc1, 0x32, 0x91, 0x5b, 0xb1,
	0xd7, 0xae, 0x3b, 0xed, 0x6b, 0xd7, 0xc2, 0xfd, 0xf1, 0x57, 0x25, 0xe1, 0xcd, 0xb8, 0xe0, 0x5e,
	0x96, 0xd4, 0x01, 0x3f, 0x9d, 0xf5, 0xe8, 0x30, 0xc5, 0x86, 0xba, 0xcb, 0xf7, 0x46, 0x5c, 0x9b,
	0xe4, 0x05, 0x5c, 0x9a, 0x3b, 0xd5, 0x85, 0xcc, 0x35, 0xc7, 0xf7, 0xa0, 0x56, 0xd8, 0x93, 0x08,
	0x35, 0xd1, 0xe6, 0x5a, 0x6b, 0x83, 0x9c, 0x8a, 0x4e, 0x5c, 0x69, 0x7b, 0xe5, 0xf0, 0xe7, 0xd5,
	0x4a, 0xd7, 0x97, 0x25, 0x2f, 0xe1, 0xb2, 0xf5, 0xbd, 0x7f, 0xa2, 0xf4, 0x2d, 0xf1, 0x45, 0xa8,
	0x8a, 0xbe, 0xf5, 0x5d, 0xe9, 0x56, 0x45, 0x1f, 0x13, 0x58, 0x95, 0xfb, 0x39, 0x57, 0x51, 0xb5,
	0x89, 0x36, 0xcf, 0xb7, 0xa3, 0x6f, 0x5f, 0xb6, 0xea, 0x1e, 0x70, 0xa7, 0xdf, 0x57, 0x5c, 0xeb,
	0x67, 0x46, 0x89, 0x3c, 0xeb, 0x3a, 0x59, 0xf2, 0x16, 0xa2, 0x45, 0x6b, 0x9f, 0xfb, 0x31, 0xac,
	0x05, 0xd9, 0x7c, 0xf8, 0xeb, 0x25, 0xe1, 0x03, 0x13, 0x4f, 0x10, 0x1a, 0x24, 0x07, 0x08, 0x36,
	0x6c, 0xb3, 0x9d, 0xc1, 0x20, 0x90, 0xea, 0xf6, 0xf8, 0xb9, 0xe6, 0xea, 0x98, 0xe8, 0x2f, 0x01,
	0x5a, 0x8a, 0x00, 0x3f, 0x04, 0x38, 0xf9, 0xc0, 0x16, 0x7b, 0x16, 0xd2, 0x57, 0xcc, 0xb6, 0x81,
	0xb8, 0xad, 0xf3, 0xdb, 0x40, 0x3a, 0x2c, 0x3b, 0x9e, 0x9e, 0x0f, 0x19, 0xd4, 0x27, 0x5f, 0x11,
	0x24, 0x65, 0x19, 0xfd, 0x68, 0x3a, 0x70, 0x21, 0x24, 0x8f, 0x50, 0xf3, 0xdc, 0x99, 0x67, 0x33,
	0xe7, 0x80, 0x1f, 0xfd, 0x03, 0xe3, 0xc6, 0x7f, 0x31, 0x5c, 0x9c, 0x45, 0x8e, 0xd6, 0x67, 0x04,
	0xab, 0x96, 0x03, 0x1f, 0x20, 0xa8, 0xb9, 0xad, 0xc2, 0x5b, 0x25, 0xf9, 0x16, 0xd7, 0xb9, 0x41,
	0x96, 0x95, 0xbb, 0x14, 0xc9, 0xdd, 0x0f, 0xdf, 0x7f, 0x7f, 0xaa, 0xb6, 0xf0, 0x2d, 0xfa, 0xc0,
	0xd5, 0x75, 0x06, 0x4c, 0x1b, 0x91, 0xd2, 0xd3, 0xdf, 0x94, 0x5b, 0xf0, 0xf6, 0x93, 0xc3, 0x49,
	0x8c, 0x8e, 0x26, 0x31, 0xfa, 0x35, 0x89, 0xd1, 0xc7, 0x69, 0x5c, 0x39, 0x9a, 0xc6, 0x95, 0x1f,
	0xd3, 0xb8, 0xf2, 0xea, 0x4e, 0x26, 0xcc, 0xee, 0xa8, 0x47, 0x52, 0x39, 0x2c, 0x75, 0x7d, 0x3f,
	0xe7, 0xdb, 0xab, 0xd9, 0x67, 0x7a, 0xfb, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x42, 0x5e,
	0x6e, 0x8c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/empowerchain.certificates.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/empowerchain.certificates.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "empowerchain.certificates.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "empowerchain/certificates/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCertificateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertificateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertificateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCertificateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertificateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertificateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Certificate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllCertificatesByUserRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCertificatesByUserRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCertificatesByUserRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCertificatesByUserResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCertificatesByUserResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCertificatesByUserResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Certificates) > 0 {
		for iNdEx := len(m.Certificates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certificates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryCertificateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCertificateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Certificate.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCertificatesByUserRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = m.Pagination.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCertificatesByUserResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Certificates) > 0 {
		for _, e := range m.Certificates {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	l = m.Pagination.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCertificateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCertificateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertificateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCertificateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Certificate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCertificatesByUserRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllCertificatesByUserRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCertificatesByUserRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllCertificatesByUserResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllCertificatesByUserResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCertificatesByUserResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificates = append(m.Certificates, Certificate{})
			if err := m.Certificates[len(m.Certificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
