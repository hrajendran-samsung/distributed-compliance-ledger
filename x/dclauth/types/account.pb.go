// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dclauth/account.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Account struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	// NOTE. we do not user AccountRoles casting here to preserve repeated form
	//       so protobuf takes care about repeated items in generated code,
	//       (but that might be not the final solution)
	Roles    []AccountRole `protobuf:"bytes,2,rep,name=roles,proto3,casttype=AccountRole" json:"roles,omitempty"`
	VendorID int32         `protobuf:"varint,3,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
}

func (m *Account) Reset()      { *m = Account{} }
func (*Account) ProtoMessage() {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a2d3e1e8208016c, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
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

func init() {
	proto.RegisterType((*Account)(nil), "zigbeealliance.distributedcomplianceledger.dclauth.Account")
}

func init() { proto.RegisterFile("dclauth/account.proto", fileDescriptor_3a2d3e1e8208016c) }

var fileDescriptor_3a2d3e1e8208016c = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xbf, 0x7f, 0xa1, 0xa4, 0x48, 0x48, 0x11, 0x48, 0xa5, 0x83, 0x13, 0x21, 0x21,
	0x75, 0x69, 0xac, 0x96, 0x8d, 0x8d, 0x88, 0x81, 0xae, 0x19, 0x59, 0x2a, 0xdb, 0xb9, 0x4a, 0x23,
	0xb9, 0xbd, 0x55, 0xec, 0x54, 0xc0, 0x13, 0x30, 0x32, 0x32, 0x96, 0xb7, 0x61, 0xec, 0xc8, 0x54,
	0xa1, 0xf6, 0x2d, 0x98, 0x10, 0xb1, 0x4b, 0xd9, 0x7c, 0xee, 0xd1, 0x39, 0xbe, 0xfa, 0xae, 0x7f,
	0x96, 0x49, 0xc5, 0x2b, 0x33, 0x61, 0x5c, 0x4a, 0xac, 0x66, 0x26, 0x9e, 0x97, 0x68, 0x30, 0x18,
	0x3e, 0x15, 0xb9, 0x00, 0xe0, 0x4a, 0x15, 0x7c, 0x26, 0x21, 0xce, 0x0a, 0x6d, 0xca, 0x42, 0x54,
	0x06, 0x32, 0x89, 0xd3, 0xb9, 0x9d, 0x2a, 0xc8, 0x72, 0x28, 0x63, 0xd7, 0xd0, 0x3d, 0xcd, 0x31,
	0xc7, 0x3a, 0xce, 0x7e, 0x5e, 0xb6, 0xa9, 0x4b, 0x25, 0xea, 0x29, 0x6a, 0x56, 0x7f, 0xb2, 0x18,
	0x08, 0x30, 0x7c, 0x50, 0x0b, 0xe7, 0x9f, 0x5b, 0x7f, 0x6c, 0x83, 0x56, 0x58, 0xeb, 0xe2, 0x8d,
	0xf8, 0x87, 0x37, 0x76, 0xad, 0x60, 0xe4, 0x1f, 0x0b, 0xae, 0x61, 0xec, 0xd6, 0xec, 0x90, 0x88,
	0xf4, 0xda, 0xc3, 0x28, 0x76, 0x81, 0xba, 0xd0, 0xb5, 0xc7, 0x09, 0xd7, 0xe0, 0x72, 0xc9, 0xff,
	0xd5, 0x3a, 0x24, 0x69, 0x5b, 0xec, 0x47, 0xc1, 0xa5, 0xdf, 0x2c, 0x51, 0x81, 0xee, 0xfc, 0x8b,
	0x1a, 0xbd, 0xa3, 0xe4, 0xe4, 0x6b, 0x1d, 0xb6, 0x9d, 0x97, 0xa2, 0x82, 0xd4, 0xba, 0x41, 0xd7,
	0x6f, 0x2d, 0x60, 0x96, 0x61, 0x39, 0xba, 0xed, 0x34, 0x22, 0xd2, 0x6b, 0xa6, 0xbf, 0xfa, 0xba,
	0xf5, 0xbc, 0x0c, 0xbd, 0xd7, 0x65, 0xe8, 0x25, 0xe2, 0x7d, 0x43, 0xc9, 0x6a, 0x43, 0xc9, 0xe7,
	0x86, 0x92, 0x97, 0x2d, 0xf5, 0x56, 0x5b, 0xea, 0x7d, 0x6c, 0xa9, 0x77, 0x7f, 0x97, 0x17, 0x66,
	0x52, 0x89, 0x58, 0xe2, 0x94, 0x59, 0x9a, 0xfd, 0x1d, 0x4e, 0xf6, 0x07, 0x67, 0x7f, 0xcf, 0xb3,
	0x6f, 0x81, 0xb2, 0x07, 0xb6, 0x3b, 0x8a, 0x79, 0x9c, 0x83, 0x16, 0x07, 0x35, 0x8e, 0xab, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x76, 0x92, 0x8c, 0xac, 0x01, 0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VendorID != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintAccount(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAccount(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovAccount(uint64(l))
		}
	}
	if m.VendorID != 0 {
		n += 1 + sovAccount(uint64(m.VendorID))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, AccountRole(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
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
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)
