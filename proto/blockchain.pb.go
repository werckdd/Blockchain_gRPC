// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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

type Block struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PrevBlockHash        string   `protobuf:"bytes,2,opt,name=prevBlockHash,proto3" json:"prevBlockHash,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{0}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Block) GetPrevBlockHash() string {
	if m != nil {
		return m.PrevBlockHash
	}
	return ""
}

func (m *Block) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AddBlockRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddBlockRequest) Reset()         { *m = AddBlockRequest{} }
func (m *AddBlockRequest) String() string { return proto.CompactTextString(m) }
func (*AddBlockRequest) ProtoMessage()    {}
func (*AddBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{1}
}

func (m *AddBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBlockRequest.Unmarshal(m, b)
}
func (m *AddBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBlockRequest.Marshal(b, m, deterministic)
}
func (m *AddBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBlockRequest.Merge(m, src)
}
func (m *AddBlockRequest) XXX_Size() int {
	return xxx_messageInfo_AddBlockRequest.Size(m)
}
func (m *AddBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddBlockRequest proto.InternalMessageInfo

func (m *AddBlockRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type AddBlockResponse struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddBlockResponse) Reset()         { *m = AddBlockResponse{} }
func (m *AddBlockResponse) String() string { return proto.CompactTextString(m) }
func (*AddBlockResponse) ProtoMessage()    {}
func (*AddBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{2}
}

func (m *AddBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBlockResponse.Unmarshal(m, b)
}
func (m *AddBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBlockResponse.Marshal(b, m, deterministic)
}
func (m *AddBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBlockResponse.Merge(m, src)
}
func (m *AddBlockResponse) XXX_Size() int {
	return xxx_messageInfo_AddBlockResponse.Size(m)
}
func (m *AddBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddBlockResponse proto.InternalMessageInfo

func (m *AddBlockResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetBlockchainRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockchainRequest) Reset()         { *m = GetBlockchainRequest{} }
func (m *GetBlockchainRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockchainRequest) ProtoMessage()    {}
func (*GetBlockchainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{3}
}

func (m *GetBlockchainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockchainRequest.Unmarshal(m, b)
}
func (m *GetBlockchainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockchainRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockchainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockchainRequest.Merge(m, src)
}
func (m *GetBlockchainRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockchainRequest.Size(m)
}
func (m *GetBlockchainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockchainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockchainRequest proto.InternalMessageInfo

type GetBlockchainResponse struct {
	Blocks               []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockchainResponse) Reset()         { *m = GetBlockchainResponse{} }
func (m *GetBlockchainResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockchainResponse) ProtoMessage()    {}
func (*GetBlockchainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ac6287ce250c9a, []int{4}
}

func (m *GetBlockchainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockchainResponse.Unmarshal(m, b)
}
func (m *GetBlockchainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockchainResponse.Marshal(b, m, deterministic)
}
func (m *GetBlockchainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockchainResponse.Merge(m, src)
}
func (m *GetBlockchainResponse) XXX_Size() int {
	return xxx_messageInfo_GetBlockchainResponse.Size(m)
}
func (m *GetBlockchainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockchainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockchainResponse proto.InternalMessageInfo

func (m *GetBlockchainResponse) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "proto.Block")
	proto.RegisterType((*AddBlockRequest)(nil), "proto.AddBlockRequest")
	proto.RegisterType((*AddBlockResponse)(nil), "proto.AddBlockResponse")
	proto.RegisterType((*GetBlockchainRequest)(nil), "proto.GetBlockchainRequest")
	proto.RegisterType((*GetBlockchainResponse)(nil), "proto.GetBlockchainResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockchainClient interface {
	AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error)
	GetBlockchain(ctx context.Context, in *GetBlockchainRequest, opts ...grpc.CallOption) (*GetBlockchainResponse, error)
}

type blockchainClient struct {
	cc *grpc.ClientConn
}

func NewBlockchainClient(cc *grpc.ClientConn) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error) {
	out := new(AddBlockResponse)
	err := c.cc.Invoke(ctx, "/proto.Blockchain/AddBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBlockchain(ctx context.Context, in *GetBlockchainRequest, opts ...grpc.CallOption) (*GetBlockchainResponse, error) {
	out := new(GetBlockchainResponse)
	err := c.cc.Invoke(ctx, "/proto.Blockchain/GetBlockchain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
type BlockchainServer interface {
	AddBlock(context.Context, *AddBlockRequest) (*AddBlockResponse, error)
	GetBlockchain(context.Context, *GetBlockchainRequest) (*GetBlockchainResponse, error)
}

func RegisterBlockchainServer(s *grpc.Server, srv BlockchainServer) {
	s.RegisterService(&_Blockchain_serviceDesc, srv)
}

func _Blockchain_AddBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).AddBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blockchain/AddBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).AddBlock(ctx, req.(*AddBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockchainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blockchain/GetBlockchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlockchain(ctx, req.(*GetBlockchainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blockchain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBlock",
			Handler:    _Blockchain_AddBlock_Handler,
		},
		{
			MethodName: "GetBlockchain",
			Handler:    _Blockchain_GetBlockchain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchain.proto",
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor_e9ac6287ce250c9a) }

var fileDescriptor_e9ac6287ce250c9a = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xca, 0xc9, 0x4f,
	0xce, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x4a, 0xa1, 0x5c, 0xac, 0x4e, 0x20, 0x29, 0x21, 0x21, 0x2e, 0x96, 0x8c, 0xc4, 0xe2, 0x0c, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x85, 0x8b, 0xb7, 0xa0, 0x28, 0xb5, 0x0c,
	0xac, 0xc0, 0x03, 0x24, 0xc9, 0x04, 0x96, 0x44, 0x15, 0x04, 0xe9, 0x4c, 0x49, 0x2c, 0x49, 0x94,
	0x60, 0x86, 0xe8, 0x04, 0xb1, 0x95, 0x54, 0xb9, 0xf8, 0x1d, 0x53, 0x52, 0xc0, 0x6a, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xe0, 0xca, 0x18, 0x91, 0x94, 0xa9, 0x71, 0x09, 0x20, 0x94, 0x15,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x62, 0x73, 0x88, 0x92, 0x18, 0x97, 0x88, 0x7b, 0x6a, 0x89, 0x13,
	0xdc, 0x0f, 0x50, 0x33, 0x95, 0x6c, 0xb9, 0x44, 0xd1, 0xc4, 0xa1, 0x86, 0xa8, 0x70, 0xb1, 0x81,
	0x7d, 0x5c, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0x03, 0xf1, 0xb5, 0x1e, 0xc4, 0x2a,
	0xa8, 0x9c, 0xd1, 0x4c, 0x46, 0x2e, 0x2e, 0x84, 0x66, 0x21, 0x5b, 0x2e, 0x0e, 0x98, 0x6b, 0x84,
	0xc4, 0xa0, 0x1a, 0xd0, 0x7c, 0x21, 0x25, 0x8e, 0x21, 0x0e, 0xb1, 0x51, 0x89, 0x41, 0xc8, 0x87,
	0x8b, 0x17, 0xc5, 0x31, 0x42, 0xd2, 0x50, 0xb5, 0xd8, 0x9c, 0x2e, 0x25, 0x83, 0x5d, 0x12, 0x66,
	0x5a, 0x12, 0x1b, 0x58, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x43, 0x9e, 0xd7, 0xba,
	0x01, 0x00, 0x00,
}