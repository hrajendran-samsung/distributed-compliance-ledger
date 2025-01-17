// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/proposed_certificate_revocation.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type ProposedCertificateRevocation struct {
	Subject      string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyId string   `protobuf:"bytes,2,opt,name=subject_key_id,json=subjectKeyId,proto3" json:"subject_key_id,omitempty"`
	Approvals    []string `protobuf:"bytes,3,rep,name=approvals,proto3" json:"approvals,omitempty"`
}

func (m *ProposedCertificateRevocation) Reset()         { *m = ProposedCertificateRevocation{} }
func (m *ProposedCertificateRevocation) String() string { return proto.CompactTextString(m) }
func (*ProposedCertificateRevocation) ProtoMessage()    {}
func (*ProposedCertificateRevocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_24b0dc6e71a9ad57, []int{0}
}
func (m *ProposedCertificateRevocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposedCertificateRevocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposedCertificateRevocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposedCertificateRevocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposedCertificateRevocation.Merge(m, src)
}
func (m *ProposedCertificateRevocation) XXX_Size() int {
	return m.Size()
}
func (m *ProposedCertificateRevocation) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposedCertificateRevocation.DiscardUnknown(m)
}

var xxx_messageInfo_ProposedCertificateRevocation proto.InternalMessageInfo

func (m *ProposedCertificateRevocation) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ProposedCertificateRevocation) GetSubjectKeyId() string {
	if m != nil {
		return m.SubjectKeyId
	}
	return ""
}

func (m *ProposedCertificateRevocation) GetApprovals() []string {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposedCertificateRevocation)(nil), "zigbeealliance.distributedcomplianceledger.pki.ProposedCertificateRevocation")
}

func init() {
	proto.RegisterFile("pki/proposed_certificate_revocation.proto", fileDescriptor_24b0dc6e71a9ad57)
}

var fileDescriptor_24b0dc6e71a9ad57 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4a, 0xf4, 0x40,
	0x14, 0x85, 0x77, 0xfe, 0x85, 0x5f, 0x36, 0x88, 0x45, 0xb0, 0x88, 0x82, 0xc3, 0x22, 0x16, 0x6b,
	0x91, 0x4c, 0x21, 0xd8, 0xbb, 0xda, 0x88, 0x8d, 0xac, 0x9d, 0x85, 0x61, 0x32, 0x73, 0x8d, 0xd7,
	0x4d, 0x32, 0xc3, 0xcc, 0x64, 0x31, 0x3e, 0x85, 0x9d, 0x2f, 0xe2, 0x43, 0x58, 0x2e, 0x56, 0x96,
	0x92, 0xbc, 0x88, 0x98, 0x64, 0x8d, 0xdd, 0x3d, 0x87, 0x73, 0xef, 0xb9, 0x7c, 0xde, 0xb1, 0x5e,
	0x22, 0xd3, 0x46, 0x69, 0x65, 0x41, 0xc6, 0x02, 0x8c, 0xc3, 0x7b, 0x14, 0xdc, 0x41, 0x6c, 0x60,
	0xa5, 0x04, 0x77, 0xa8, 0x8a, 0x48, 0x1b, 0xe5, 0x94, 0x1f, 0x3d, 0x63, 0x9a, 0x00, 0xf0, 0x2c,
	0x43, 0x5e, 0x08, 0x88, 0x24, 0x5a, 0x67, 0x30, 0x29, 0x1d, 0x48, 0xa1, 0x72, 0xdd, 0xb9, 0x19,
	0xc8, 0x14, 0x4c, 0xa4, 0x97, 0xb8, 0xbf, 0x27, 0x94, 0xcd, 0x95, 0x8d, 0xdb, 0x6d, 0xd6, 0x89,
	0xee, 0xd4, 0xe1, 0x2b, 0xf1, 0x0e, 0xae, 0xfb, 0xd2, 0xf3, 0xa1, 0x73, 0xf1, 0x5b, 0xe9, 0x07,
	0xde, 0x96, 0x2d, 0x93, 0x47, 0x10, 0x2e, 0x20, 0x53, 0x32, 0x9b, 0x2c, 0x36, 0xd2, 0x3f, 0xf2,
	0x76, 0xfa, 0x31, 0x5e, 0x42, 0x15, 0xa3, 0x0c, 0xfe, 0xb5, 0x81, 0xed, 0xde, 0xbd, 0x82, 0xea,
	0x52, 0xfa, 0xa7, 0xde, 0x84, 0x6b, 0x6d, 0xd4, 0x8a, 0x67, 0x36, 0x18, 0x4f, 0xc7, 0xb3, 0xc9,
	0x3c, 0xf8, 0x78, 0x0b, 0x77, 0xfb, 0x37, 0xce, 0xa4, 0x34, 0x60, 0xed, 0x8d, 0x33, 0x58, 0xa4,
	0x8b, 0x21, 0x3a, 0xbf, 0x7b, 0xaf, 0x29, 0x59, 0xd7, 0x94, 0x7c, 0xd5, 0x94, 0xbc, 0x34, 0x74,
	0xb4, 0x6e, 0xe8, 0xe8, 0xb3, 0xa1, 0xa3, 0xdb, 0x8b, 0x14, 0xdd, 0x43, 0x99, 0x44, 0x42, 0xe5,
	0xac, 0x23, 0x11, 0x6e, 0x50, 0xb0, 0x3f, 0x28, 0xc2, 0x81, 0x45, 0xd8, 0xc1, 0x60, 0x4f, 0xec,
	0x07, 0xb2, 0xab, 0x34, 0xd8, 0xe4, 0x7f, 0x0b, 0xe0, 0xe4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x6c,
	0xc4, 0xaa, 0x7d, 0x78, 0x01, 0x00, 0x00,
}

func (m *ProposedCertificateRevocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposedCertificateRevocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposedCertificateRevocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Approvals) > 0 {
		for iNdEx := len(m.Approvals) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Approvals[iNdEx])
			copy(dAtA[i:], m.Approvals[iNdEx])
			i = encodeVarintProposedCertificateRevocation(dAtA, i, uint64(len(m.Approvals[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SubjectKeyId) > 0 {
		i -= len(m.SubjectKeyId)
		copy(dAtA[i:], m.SubjectKeyId)
		i = encodeVarintProposedCertificateRevocation(dAtA, i, uint64(len(m.SubjectKeyId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintProposedCertificateRevocation(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposedCertificateRevocation(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposedCertificateRevocation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposedCertificateRevocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovProposedCertificateRevocation(uint64(l))
	}
	l = len(m.SubjectKeyId)
	if l > 0 {
		n += 1 + l + sovProposedCertificateRevocation(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, s := range m.Approvals {
			l = len(s)
			n += 1 + l + sovProposedCertificateRevocation(uint64(l))
		}
	}
	return n
}

func sovProposedCertificateRevocation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposedCertificateRevocation(x uint64) (n int) {
	return sovProposedCertificateRevocation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposedCertificateRevocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposedCertificateRevocation
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
			return fmt.Errorf("proto: ProposedCertificateRevocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposedCertificateRevocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificateRevocation
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
				return ErrInvalidLengthProposedCertificateRevocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificateRevocation
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
				return ErrInvalidLengthProposedCertificateRevocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificateRevocation
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
				return ErrInvalidLengthProposedCertificateRevocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposedCertificateRevocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
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
func skipProposedCertificateRevocation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposedCertificateRevocation
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
					return 0, ErrIntOverflowProposedCertificateRevocation
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
					return 0, ErrIntOverflowProposedCertificateRevocation
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
				return 0, ErrInvalidLengthProposedCertificateRevocation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposedCertificateRevocation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposedCertificateRevocation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposedCertificateRevocation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposedCertificateRevocation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposedCertificateRevocation = fmt.Errorf("proto: unexpected end of group")
)
