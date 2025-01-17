// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: compliance/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the compliance module's genesis state.
type GenesisState struct {
	ComplianceInfoList   []ComplianceInfo   `protobuf:"bytes,1,rep,name=complianceInfoList,proto3" json:"complianceInfoList"`
	CertifiedModelList   []CertifiedModel   `protobuf:"bytes,2,rep,name=certifiedModelList,proto3" json:"certifiedModelList"`
	RevokedModelList     []RevokedModel     `protobuf:"bytes,3,rep,name=revokedModelList,proto3" json:"revokedModelList"`
	ProvisionalModelList []ProvisionalModel `protobuf:"bytes,4,rep,name=provisionalModelList,proto3" json:"provisionalModelList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a491f9d8ec824f82, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetComplianceInfoList() []ComplianceInfo {
	if m != nil {
		return m.ComplianceInfoList
	}
	return nil
}

func (m *GenesisState) GetCertifiedModelList() []CertifiedModel {
	if m != nil {
		return m.CertifiedModelList
	}
	return nil
}

func (m *GenesisState) GetRevokedModelList() []RevokedModel {
	if m != nil {
		return m.RevokedModelList
	}
	return nil
}

func (m *GenesisState) GetProvisionalModelList() []ProvisionalModel {
	if m != nil {
		return m.ProvisionalModelList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zigbeealliance.distributedcomplianceledger.compliance.GenesisState")
}

func init() { proto.RegisterFile("compliance/genesis.proto", fileDescriptor_a491f9d8ec824f82) }

var fileDescriptor_a491f9d8ec824f82 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0x87, 0x5b, 0x21, 0x1e, 0xaa, 0x07, 0xd3, 0x70, 0x20, 0x1e, 0x56, 0xc2, 0xc9, 0x0b, 0x6d,
	0xa2, 0xf1, 0x05, 0x20, 0x86, 0x18, 0x35, 0x31, 0x78, 0xf3, 0x42, 0x5a, 0x3a, 0xac, 0x13, 0x4b,
	0xb7, 0xd9, 0x5d, 0x88, 0x7f, 0x4e, 0xc6, 0x17, 0xf0, 0xb1, 0x38, 0x72, 0xf4, 0x64, 0x0c, 0xbc,
	0x88, 0x61, 0xb7, 0xb8, 0x5b, 0xe0, 0x84, 0xb7, 0xc9, 0xce, 0xf4, 0xfb, 0x7e, 0x9d, 0x5d, 0xaf,
	0x3e, 0x60, 0xa3, 0x3c, 0xc5, 0x28, 0x1b, 0x40, 0x48, 0x21, 0x03, 0x81, 0x22, 0xc8, 0x39, 0x93,
	0xcc, 0xbf, 0x78, 0x45, 0x1a, 0x03, 0x44, 0xa9, 0xee, 0x06, 0x09, 0x0a, 0xc9, 0x31, 0x1e, 0x4b,
	0x48, 0xcc, 0x37, 0x29, 0x24, 0x14, 0x78, 0x60, 0x0e, 0x8e, 0x1b, 0x16, 0xd0, 0x94, 0x7d, 0xcc,
	0x86, 0x4c, 0x83, 0xcb, 0x13, 0xc0, 0x25, 0x0e, 0x11, 0x92, 0xfe, 0x88, 0x25, 0x90, 0x16, 0x13,
	0xc4, 0x9a, 0xe0, 0x30, 0x61, 0x4f, 0x6b, 0xfd, 0xa6, 0xd5, 0xcf, 0x39, 0x9b, 0xa0, 0x40, 0x96,
	0x45, 0x69, 0x69, 0xa6, 0x46, 0x19, 0x65, 0xaa, 0x0c, 0x97, 0x95, 0x3e, 0x6d, 0x7e, 0x54, 0xbd,
	0xc3, 0xae, 0xfe, 0xcd, 0x7b, 0x19, 0x49, 0xf0, 0xdf, 0x3c, 0xdf, 0xc0, 0xae, 0xb2, 0x21, 0xbb,
	0x41, 0x21, 0xeb, 0x6e, 0xa3, 0x72, 0x7a, 0x70, 0x76, 0x19, 0xec, 0xb4, 0x82, 0xa0, 0x53, 0x02,
	0xb6, 0xab, 0xd3, 0xef, 0x13, 0xa7, 0xb7, 0x45, 0xa3, 0xe4, 0xab, 0x05, 0xdc, 0x2e, 0xb3, 0x2b,
	0xf9, 0xde, 0xff, 0xe4, 0x25, 0xe0, 0x9f, 0x7c, 0x43, 0xe3, 0x8f, 0xbd, 0xa3, 0x62, 0xb7, 0x46,
	0x5d, 0x51, 0xea, 0xce, 0x8e, 0xea, 0x9e, 0x85, 0x2b, 0xc4, 0x1b, 0x0a, 0xff, 0xdd, 0xf5, 0x6a,
	0xd6, 0x9d, 0x19, 0x77, 0x55, 0xb9, 0xbb, 0x3b, 0xba, 0xef, 0xd6, 0x90, 0x85, 0x7f, 0xab, 0xaa,
	0x0d, 0xd3, 0x39, 0x71, 0x67, 0x73, 0xe2, 0xfe, 0xcc, 0x89, 0xfb, 0xb9, 0x20, 0xce, 0x6c, 0x41,
	0x9c, 0xaf, 0x05, 0x71, 0x1e, 0xae, 0x29, 0xca, 0xc7, 0x71, 0xbc, 0x24, 0x87, 0x3a, 0x48, 0x6b,
	0x95, 0x24, 0xb4, 0x92, 0xb4, 0x8c, 0xb9, 0xa5, 0xb3, 0x84, 0xcf, 0xd6, 0x6b, 0x0f, 0xe5, 0x4b,
	0x0e, 0x22, 0xde, 0x57, 0x6f, 0xee, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x70, 0xe0, 0xbc,
	0x64, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProvisionalModelList) > 0 {
		for iNdEx := len(m.ProvisionalModelList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProvisionalModelList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.RevokedModelList) > 0 {
		for iNdEx := len(m.RevokedModelList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RevokedModelList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.CertifiedModelList) > 0 {
		for iNdEx := len(m.CertifiedModelList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CertifiedModelList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ComplianceInfoList) > 0 {
		for iNdEx := len(m.ComplianceInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ComplianceInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ComplianceInfoList) > 0 {
		for _, e := range m.ComplianceInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CertifiedModelList) > 0 {
		for _, e := range m.CertifiedModelList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RevokedModelList) > 0 {
		for _, e := range m.RevokedModelList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProvisionalModelList) > 0 {
		for _, e := range m.ProvisionalModelList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComplianceInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ComplianceInfoList = append(m.ComplianceInfoList, ComplianceInfo{})
			if err := m.ComplianceInfoList[len(m.ComplianceInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertifiedModelList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertifiedModelList = append(m.CertifiedModelList, CertifiedModel{})
			if err := m.CertifiedModelList[len(m.CertifiedModelList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevokedModelList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevokedModelList = append(m.RevokedModelList, RevokedModel{})
			if err := m.RevokedModelList[len(m.RevokedModelList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvisionalModelList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProvisionalModelList = append(m.ProvisionalModelList, ProvisionalModel{})
			if err := m.ProvisionalModelList[len(m.ProvisionalModelList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
