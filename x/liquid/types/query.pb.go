// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/liquid/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryDelegatedBalanceRequest defines the request type for Query/DelegatedBalance method.
type QueryDelegatedBalanceRequest struct {
	// delegator is the address of the account to query
	Delegator string `protobuf:"bytes,1,opt,name=delegator,proto3" json:"delegator,omitempty"`
}

func (m *QueryDelegatedBalanceRequest) Reset()         { *m = QueryDelegatedBalanceRequest{} }
func (m *QueryDelegatedBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatedBalanceRequest) ProtoMessage()    {}
func (*QueryDelegatedBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d745428489be444, []int{0}
}
func (m *QueryDelegatedBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatedBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatedBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatedBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatedBalanceRequest.Merge(m, src)
}
func (m *QueryDelegatedBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatedBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatedBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatedBalanceRequest proto.InternalMessageInfo

// DelegatedBalanceResponse defines the response type for the Query/DelegatedBalance method.
type QueryDelegatedBalanceResponse struct {
	// vested is the amount of all delegated coins that have vested (ie not locked)
	Vested types.Coin `protobuf:"bytes,1,opt,name=vested,proto3" json:"vested"`
	// vesting is the amount of all delegated coins that are still vesting (ie locked)
	Vesting types.Coin `protobuf:"bytes,2,opt,name=vesting,proto3" json:"vesting"`
}

func (m *QueryDelegatedBalanceResponse) Reset()         { *m = QueryDelegatedBalanceResponse{} }
func (m *QueryDelegatedBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDelegatedBalanceResponse) ProtoMessage()    {}
func (*QueryDelegatedBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d745428489be444, []int{1}
}
func (m *QueryDelegatedBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDelegatedBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDelegatedBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDelegatedBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDelegatedBalanceResponse.Merge(m, src)
}
func (m *QueryDelegatedBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDelegatedBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDelegatedBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDelegatedBalanceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryDelegatedBalanceRequest)(nil), "kava.liquid.v1beta1.QueryDelegatedBalanceRequest")
	proto.RegisterType((*QueryDelegatedBalanceResponse)(nil), "kava.liquid.v1beta1.QueryDelegatedBalanceResponse")
}

func init() { proto.RegisterFile("kava/liquid/v1beta1/query.proto", fileDescriptor_0d745428489be444) }

var fileDescriptor_0d745428489be444 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3d, 0x6f, 0xd3, 0x40,
	0x1c, 0xc6, 0x7d, 0x11, 0x04, 0xe5, 0x58, 0x90, 0xc9, 0x90, 0x44, 0xe1, 0x82, 0x32, 0x65, 0xc9,
	0x9d, 0x62, 0x10, 0x08, 0x36, 0x0c, 0x62, 0x27, 0x48, 0x0c, 0x2c, 0xd1, 0xd9, 0x3e, 0x1d, 0x27,
	0x9c, 0x3b, 0xc7, 0x77, 0x8e, 0x88, 0x10, 0x0b, 0x9f, 0xa0, 0x52, 0xbe, 0x4a, 0x3b, 0x77, 0xcd,
	0x18, 0xb5, 0x4b, 0xa7, 0xaa, 0x4d, 0xfa, 0x41, 0x2a, 0xfb, 0xec, 0x54, 0xaa, 0xd2, 0xaa, 0xdd,
	0xce, 0x7a, 0x7e, 0xcf, 0xf3, 0x7f, 0x33, 0xec, 0xfd, 0xa6, 0x73, 0x4a, 0x62, 0x31, 0xcb, 0x44,
	0x44, 0xe6, 0xa3, 0x80, 0x19, 0x3a, 0x22, 0xb3, 0x8c, 0xa5, 0x0b, 0x9c, 0xa4, 0xca, 0x28, 0xf7,
	0x65, 0x0e, 0x60, 0x0b, 0xe0, 0x12, 0xe8, 0x34, 0xb9, 0xe2, 0xaa, 0xd0, 0x49, 0xfe, 0xb2, 0x68,
	0xa7, 0xcb, 0x95, 0xe2, 0x31, 0x23, 0x34, 0x11, 0x84, 0x4a, 0xa9, 0x0c, 0x35, 0x42, 0x49, 0x5d,
	0xaa, 0x28, 0x54, 0x7a, 0xaa, 0x34, 0x09, 0xa8, 0x66, 0xbb, 0x4a, 0xa1, 0x12, 0xb2, 0xd4, 0xdb,
	0x56, 0x9f, 0xd8, 0x58, 0xfb, 0x61, 0xa5, 0xfe, 0x0f, 0xd8, 0xfd, 0x96, 0xb7, 0xf4, 0x85, 0xc5,
	0x8c, 0x53, 0xc3, 0x22, 0x9f, 0xc6, 0x54, 0x86, 0x6c, 0xcc, 0x66, 0x19, 0xd3, 0xc6, 0x7d, 0x07,
	0x1b, 0x91, 0x95, 0x54, 0xda, 0x02, 0xaf, 0xc1, 0xa0, 0xe1, 0xb7, 0x4e, 0x0e, 0x87, 0xcd, 0x32,
	0xe4, 0x53, 0x14, 0xa5, 0x4c, 0xeb, 0xef, 0x26, 0x15, 0x92, 0x8f, 0x6f, 0xd0, 0xfe, 0x12, 0xc0,
	0x57, 0x77, 0x04, 0xeb, 0x44, 0x49, 0xcd, 0xdc, 0xf7, 0xb0, 0x3e, 0x67, 0xda, 0xb0, 0xa8, 0x88,
	0x7d, 0xee, 0xb5, 0x71, 0x99, 0x99, 0x4f, 0x51, 0xad, 0x03, 0x7f, 0x56, 0x42, 0xfa, 0x4f, 0x56,
	0xe7, 0x3d, 0x67, 0x5c, 0xe2, 0xee, 0x07, 0xf8, 0x2c, 0x7f, 0x09, 0xc9, 0x5b, 0xb5, 0x87, 0x39,
	0x2b, 0xde, 0x3b, 0x06, 0xf0, 0x69, 0xd1, 0x95, 0x7b, 0x04, 0xe0, 0x8b, 0xdb, 0xad, 0xb9, 0x23,
	0xbc, 0xe7, 0x22, 0xf8, 0xbe, 0xfd, 0x74, 0xbc, 0xc7, 0x58, 0xec, 0xe4, 0xfd, 0x8f, 0xff, 0x4f,
	0xaf, 0x96, 0xb5, 0xb7, 0xae, 0x47, 0xf6, 0xfd, 0x21, 0x51, 0x65, 0x9b, 0x04, 0xd6, 0x47, 0xfe,
	0xee, 0xd6, 0xfa, 0xcf, 0xff, 0xba, 0xba, 0x44, 0xce, 0x6a, 0x83, 0xc0, 0x7a, 0x83, 0xc0, 0xc5,
	0x06, 0x81, 0x83, 0x2d, 0x72, 0xd6, 0x5b, 0xe4, 0x9c, 0x6d, 0x91, 0xf3, 0x73, 0xc0, 0x85, 0xf9,
	0x95, 0x05, 0x38, 0x54, 0xd3, 0x22, 0x7b, 0x18, 0xd3, 0x40, 0xdb, 0x2a, 0x7f, 0xaa, 0x3a, 0x66,
	0x91, 0x30, 0x1d, 0xd4, 0x8b, 0xf3, 0xbf, 0xb9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x32, 0x91, 0xae,
	0x51, 0xa5, 0x02, 0x00, 0x00,
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
	// DelegatedBalance returns an account's vesting and vested coins currently delegated to validators.
	// It ignores coins in unbonding delegations.
	DelegatedBalance(ctx context.Context, in *QueryDelegatedBalanceRequest, opts ...grpc.CallOption) (*QueryDelegatedBalanceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DelegatedBalance(ctx context.Context, in *QueryDelegatedBalanceRequest, opts ...grpc.CallOption) (*QueryDelegatedBalanceResponse, error) {
	out := new(QueryDelegatedBalanceResponse)
	err := c.cc.Invoke(ctx, "/kava.liquid.v1beta1.Query/DelegatedBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DelegatedBalance returns an account's vesting and vested coins currently delegated to validators.
	// It ignores coins in unbonding delegations.
	DelegatedBalance(context.Context, *QueryDelegatedBalanceRequest) (*QueryDelegatedBalanceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DelegatedBalance(ctx context.Context, req *QueryDelegatedBalanceRequest) (*QueryDelegatedBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegatedBalance not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DelegatedBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatedBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatedBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kava.liquid.v1beta1.Query/DelegatedBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatedBalance(ctx, req.(*QueryDelegatedBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kava.liquid.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DelegatedBalance",
			Handler:    _Query_DelegatedBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kava/liquid/v1beta1/query.proto",
}

func (m *QueryDelegatedBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatedBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatedBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDelegatedBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDelegatedBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDelegatedBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Vesting.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Vested.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryDelegatedBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDelegatedBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Vested.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.Vesting.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDelegatedBalanceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDelegatedBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatedBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
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
			m.Delegator = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDelegatedBalanceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDelegatedBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDelegatedBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vested", wireType)
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
			if err := m.Vested.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vesting", wireType)
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
			if err := m.Vesting.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
