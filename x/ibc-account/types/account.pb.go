// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/account/account.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// IBCAccount defines an account to which other chains have privileges
type IBCAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
	SourcePort         string `protobuf:"bytes,2,opt,name=sourcePort,proto3" json:"sourcePort,omitempty"`
	SourceChannel      string `protobuf:"bytes,3,opt,name=sourceChannel,proto3" json:"sourceChannel,omitempty"`
	DestinationPort    string `protobuf:"bytes,4,opt,name=destinationPort,proto3" json:"destinationPort,omitempty"`
	DestinationChannel string `protobuf:"bytes,5,opt,name=destinationChannel,proto3" json:"destinationChannel,omitempty"`
}

func (m *IBCAccount) Reset()      { *m = IBCAccount{} }
func (*IBCAccount) ProtoMessage() {}
func (*IBCAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_be5ed7ee65e0e021, []int{0}
}
func (m *IBCAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IBCAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IBCAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IBCAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IBCAccount.Merge(m, src)
}
func (m *IBCAccount) XXX_Size() int {
	return m.Size()
}
func (m *IBCAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_IBCAccount.DiscardUnknown(m)
}

var xxx_messageInfo_IBCAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IBCAccount)(nil), "ibc.account.IBCAccount")
}

func init() { proto.RegisterFile("ibc/account/account.proto", fileDescriptor_be5ed7ee65e0e021) }

var fileDescriptor_be5ed7ee65e0e021 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x93, 0xfe, 0x3f, 0x48, 0x38, 0x20, 0x24, 0x83, 0x44, 0x5b, 0x24, 0xa7, 0xaa, 0x18,
	0xba, 0x24, 0x51, 0x61, 0xeb, 0x46, 0x3a, 0xa0, 0x6e, 0xa8, 0x23, 0x4b, 0xb1, 0x5d, 0xab, 0xb1,
	0x68, 0xed, 0x2a, 0x76, 0x10, 0x7d, 0x03, 0x46, 0x46, 0xc6, 0x6e, 0xbc, 0x00, 0x0f, 0xc1, 0x58,
	0x31, 0x31, 0x55, 0xa8, 0x7d, 0x03, 0x9e, 0x00, 0x61, 0xbb, 0xb4, 0x20, 0xa6, 0xe4, 0x7c, 0xe7,
	0xdc, 0x73, 0x25, 0x5f, 0x50, 0xe1, 0x84, 0x26, 0x98, 0x52, 0x59, 0x08, 0xbd, 0xfa, 0xc6, 0xe3,
	0x5c, 0x6a, 0x09, 0x03, 0x4e, 0x68, 0xec, 0x50, 0xb5, 0x42, 0xa5, 0x1a, 0x49, 0xd5, 0x33, 0x56,
	0x62, 0x85, 0xcd, 0x55, 0x0f, 0x07, 0x72, 0x20, 0x2d, 0xff, 0xfa, 0x73, 0x14, 0xd9, 0x4c, 0x82,
	0x0b, 0x9d, 0x25, 0xb7, 0x4d, 0xc2, 0x34, 0x6e, 0x1a, 0x61, 0xfd, 0xfa, 0x53, 0x09, 0x80, 0x4e,
	0xda, 0x3e, 0xb7, 0xfd, 0xf0, 0x1a, 0xec, 0x12, 0xac, 0x58, 0xcf, 0xed, 0x2b, 0xfb, 0x35, 0xbf,
	0x11, 0x9c, 0xd6, 0x62, 0xb7, 0xc9, 0x0c, 0xba, 0x96, 0x38, 0xc5, 0x8a, 0xb9, 0xb9, 0xf4, 0x78,
	0x36, 0x0f, 0xfd, 0x8f, 0x79, 0x78, 0x30, 0xc1, 0xa3, 0x61, 0xab, 0xbe, 0xd9, 0x51, 0xef, 0x06,
	0x64, 0x9d, 0x84, 0x08, 0x00, 0x25, 0x8b, 0x9c, 0xb2, 0x4b, 0x99, 0xeb, 0x72, 0xa9, 0xe6, 0x37,
	0x76, 0xba, 0x1b, 0x04, 0x9e, 0x80, 0x3d, 0xab, 0xda, 0x19, 0x16, 0x82, 0x0d, 0xcb, 0xff, 0x4c,
	0xe4, 0x27, 0x84, 0x0d, 0xb0, 0xdf, 0x67, 0x4a, 0x73, 0x81, 0x35, 0x97, 0xc2, 0x54, 0xfd, 0x37,
	0xb9, 0xdf, 0x18, 0xc6, 0x00, 0x6e, 0xa0, 0x55, 0xe9, 0x96, 0x09, 0xff, 0xe1, 0xb4, 0x8e, 0xee,
	0xa7, 0xa1, 0xf7, 0x38, 0x0d, 0xbd, 0xd7, 0xe7, 0x28, 0x58, 0xbf, 0x4c, 0x27, 0xc5, 0x2f, 0x0b,
	0xe4, 0xcf, 0x16, 0xc8, 0x7f, 0x5f, 0x20, 0xff, 0x61, 0x89, 0xbc, 0xd9, 0x12, 0x79, 0x6f, 0x4b,
	0xe4, 0x5d, 0x5d, 0x0c, 0xb8, 0xce, 0x0a, 0x12, 0x53, 0x39, 0x4a, 0x68, 0x86, 0xb9, 0xc0, 0x63,
	0xc5, 0x95, 0xbb, 0x4e, 0xa4, 0xfa, 0x37, 0x11, 0x17, 0x9a, 0xe5, 0xc6, 0x8a, 0x56, 0x47, 0xbe,
	0x4b, 0x38, 0xa1, 0xdf, 0x4a, 0x4f, 0xc6, 0x4c, 0x91, 0x6d, 0x73, 0x93, 0xb3, 0xcf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x23, 0x9e, 0xdd, 0x3e, 0x0e, 0x02, 0x00, 0x00,
}

func (m *IBCAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IBCAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IBCAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DestinationChannel) > 0 {
		i -= len(m.DestinationChannel)
		copy(dAtA[i:], m.DestinationChannel)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.DestinationChannel)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DestinationPort) > 0 {
		i -= len(m.DestinationPort)
		copy(dAtA[i:], m.DestinationPort)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.DestinationPort)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0x12
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
func (m *IBCAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.DestinationPort)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.DestinationChannel)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IBCAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IBCAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IBCAccount: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
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
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
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
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationPort", wireType)
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
			m.DestinationPort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChannel", wireType)
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
			m.DestinationChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAccount
			}
			if (iNdEx + skippy) < 0 {
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
