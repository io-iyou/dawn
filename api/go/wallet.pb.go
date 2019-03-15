// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wallet.proto

package backend

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string            `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Fee         float32           `protobuf:"fixed32,3,opt,name=fee,proto3" json:"fee,omitempty"`
	OrderId     string            `protobuf:"bytes,4,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Created     *types.Timestamp  `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Annotations map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Account) GetFee() float32 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Account) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Account) GetCreated() *types.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Account) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type AccountList struct {
	Items []*Account `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *AccountList) Reset()         { *m = AccountList{} }
func (m *AccountList) String() string { return proto.CompactTextString(m) }
func (*AccountList) ProtoMessage()    {}
func (*AccountList) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{1}
}
func (m *AccountList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountList.Merge(m, src)
}
func (m *AccountList) XXX_Size() int {
	return m.Size()
}
func (m *AccountList) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountList.DiscardUnknown(m)
}

var xxx_messageInfo_AccountList proto.InternalMessageInfo

func (m *AccountList) GetItems() []*Account {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "backend.Account")
	proto.RegisterMapType((map[string]string)(nil), "backend.Account.AnnotationsEntry")
	proto.RegisterType((*AccountList)(nil), "backend.AccountList")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor_b88fd140af4deb6f) }

var fileDescriptor_b88fd140af4deb6f = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x33, 0x49, 0xd3, 0x70, 0x4f, 0xee, 0xbd, 0x94, 0xa1, 0xc8, 0x90, 0x45, 0x8c, 0x5d,
	0x48, 0x40, 0x4c, 0xa1, 0x2a, 0x88, 0x0b, 0xa1, 0x8a, 0x0b, 0xc1, 0x55, 0xa8, 0xb8, 0x9e, 0x66,
	0xa6, 0x25, 0x34, 0xcd, 0x94, 0x64, 0xa2, 0xf4, 0x2d, 0x5c, 0xb8, 0xf0, 0x91, 0x5c, 0x76, 0xe9,
	0x52, 0xda, 0x17, 0x91, 0x4c, 0x12, 0x95, 0x76, 0xe3, 0x6e, 0xfe, 0x33, 0xdf, 0xcc, 0xf9, 0x0e,
	0x07, 0xfe, 0x3e, 0xd1, 0x24, 0xe1, 0x32, 0x58, 0x64, 0x42, 0x0a, 0x6c, 0x8d, 0x69, 0x34, 0xe3,
	0x29, 0x73, 0xa0, 0xc8, 0x79, 0x56, 0x15, 0x9d, 0xfd, 0xa9, 0x10, 0xd3, 0x84, 0xf7, 0x55, 0x1a,
	0x17, 0x93, 0xbe, 0x8c, 0xe7, 0x3c, 0x97, 0x74, 0xbe, 0xa8, 0x80, 0xde, 0xab, 0x0e, 0xd6, 0x30,
	0x8a, 0x44, 0x91, 0x4a, 0xfc, 0x1f, 0xf4, 0x98, 0x11, 0xe4, 0x21, 0xff, 0x4f, 0xa8, 0xc7, 0x0c,
	0xef, 0x41, 0xbb, 0xfc, 0xea, 0x96, 0x11, 0x5d, 0xd5, 0xea, 0x84, 0x3b, 0x60, 0x4c, 0x38, 0x27,
	0x86, 0x87, 0x7c, 0x3d, 0x2c, 0x8f, 0x98, 0x80, 0x25, 0x32, 0xa6, 0xd0, 0x96, 0x42, 0x9b, 0x88,
	0x4f, 0xc1, 0x8a, 0x32, 0x4e, 0x25, 0x67, 0xc4, 0xf4, 0x90, 0x6f, 0x0f, 0x9c, 0xa0, 0x52, 0x0a,
	0x1a, 0xa5, 0x60, 0xd4, 0x28, 0x85, 0x0d, 0x8a, 0xaf, 0xc1, 0xa6, 0x69, 0x2a, 0x24, 0x95, 0xb1,
	0x48, 0x73, 0xd2, 0xf6, 0x0c, 0xdf, 0x1e, 0x1c, 0x04, 0xf5, 0x84, 0x41, 0x2d, 0x1c, 0x0c, 0xbf,
	0x99, 0x9b, 0x54, 0x66, 0xcb, 0xf0, 0xe7, 0x2b, 0xe7, 0x12, 0x3a, 0xdb, 0x40, 0xa9, 0x3e, 0xe3,
	0xcb, 0x7a, 0xc6, 0xf2, 0x88, 0xbb, 0x60, 0x3e, 0xd2, 0xa4, 0xe0, 0xf5, 0x8c, 0x55, 0xb8, 0xd0,
	0xcf, 0x51, 0xef, 0x0c, 0xec, 0xba, 0xd1, 0x5d, 0x9c, 0x4b, 0x7c, 0x08, 0x66, 0x2c, 0xf9, 0x3c,
	0x27, 0x48, 0xd9, 0x74, 0xb6, 0x6d, 0xc2, 0xea, 0x7a, 0xf0, 0x82, 0xc0, 0x7a, 0x50, 0x8b, 0xc9,
	0xf1, 0x11, 0x18, 0x43, 0xc6, 0xf0, 0x0e, 0xeb, 0xec, 0x54, 0x7a, 0x1a, 0x3e, 0x86, 0x96, 0x6a,
	0xf4, 0xef, 0xeb, 0xee, 0x3e, 0xe7, 0x99, 0xd3, 0xdd, 0x46, 0x4b, 0x48, 0xe1, 0xe6, 0x48, 0x48,
	0x9a, 0xfc, 0xee, 0xf7, 0x2b, 0xf2, 0xb6, 0x76, 0xd1, 0x6a, 0xed, 0xa2, 0x8f, 0xb5, 0x8b, 0x9e,
	0x37, 0xae, 0xb6, 0xda, 0xb8, 0xda, 0xfb, 0xc6, 0xd5, 0xc6, 0x6d, 0xb5, 0x89, 0x93, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa5, 0x85, 0xb3, 0xa5, 0x4f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WalletsClient is the client API for Wallets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletsClient interface {
	Add(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error)
	List(ctx context.Context, in *User, opts ...grpc.CallOption) (*AccountList, error)
	Total(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error)
}

type walletsClient struct {
	cc *grpc.ClientConn
}

func NewWalletsClient(cc *grpc.ClientConn) WalletsClient {
	return &walletsClient{cc}
}

func (c *walletsClient) Add(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/backend.Wallets/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletsClient) List(ctx context.Context, in *User, opts ...grpc.CallOption) (*AccountList, error) {
	out := new(AccountList)
	err := c.cc.Invoke(ctx, "/backend.Wallets/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletsClient) Total(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/backend.Wallets/Total", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletsServer is the server API for Wallets service.
type WalletsServer interface {
	Add(context.Context, *Account) (*Account, error)
	List(context.Context, *User) (*AccountList, error)
	Total(context.Context, *Account) (*Account, error)
}

func RegisterWalletsServer(s *grpc.Server, srv WalletsServer) {
	s.RegisterService(&_Wallets_serviceDesc, srv)
}

func _Wallets_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Wallets/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletsServer).Add(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallets_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Wallets/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletsServer).List(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallets_Total_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletsServer).Total(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Wallets/Total",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletsServer).Total(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wallets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backend.Wallets",
	HandlerType: (*WalletsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Wallets_Add_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Wallets_List_Handler,
		},
		{
			MethodName: "Total",
			Handler:    _Wallets_Total_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.UserId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if m.Fee != 0 {
		dAtA[i] = 0x1d
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Fee))))
		i += 4
	}
	if len(m.OrderId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintWallet(dAtA, i, uint64(len(m.OrderId)))
		i += copy(dAtA[i:], m.OrderId)
	}
	if m.Created != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintWallet(dAtA, i, uint64(m.Created.Size()))
		n1, err := m.Created.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Annotations) > 0 {
		for k, _ := range m.Annotations {
			dAtA[i] = 0x32
			i++
			v := m.Annotations[k]
			mapSize := 1 + len(k) + sovWallet(uint64(len(k))) + 1 + len(v) + sovWallet(uint64(len(v)))
			i = encodeVarintWallet(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintWallet(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintWallet(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *AccountList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintWallet(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintWallet(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	if m.Fee != 0 {
		n += 5
	}
	l = len(m.OrderId)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	if m.Created != nil {
		l = m.Created.Size()
		n += 1 + l + sovWallet(uint64(l))
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWallet(uint64(len(k))) + 1 + len(v) + sovWallet(uint64(len(v)))
			n += mapEntrySize + 1 + sovWallet(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *AccountList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovWallet(uint64(l))
		}
	}
	return n
}

func sovWallet(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWallet(x uint64) (n int) {
	return sovWallet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Fee = float32(math.Float32frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Created == nil {
				m.Created = &types.Timestamp{}
			}
			if err := m.Created.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Annotations == nil {
				m.Annotations = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWallet
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWallet
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWallet
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWallet
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWallet
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWallet
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthWallet
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWallet(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthWallet
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Annotations[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWallet
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
func (m *AccountList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
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
			return fmt.Errorf("proto: AccountList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Account{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWallet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWallet
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
func skipWallet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWallet
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
					return 0, ErrIntOverflowWallet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWallet
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
				return 0, ErrInvalidLengthWallet
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthWallet
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWallet
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWallet(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthWallet
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWallet = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWallet   = fmt.Errorf("proto: integer overflow")
)