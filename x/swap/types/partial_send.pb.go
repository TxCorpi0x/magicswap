// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: magicswap/swap/partial_send.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type PartialSend struct {
	Id            uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Recipient     string     `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	BurntAmount   types.Coin `protobuf:"bytes,3,opt,name=burntAmount,proto3" json:"burntAmount"`
	SentAmount    types.Coin `protobuf:"bytes,4,opt,name=sentAmount,proto3" json:"sentAmount"`
	SwappedAmount types.Coin `protobuf:"bytes,5,opt,name=swappedAmount,proto3" json:"swappedAmount"`
	Creator       string     `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *PartialSend) Reset()         { *m = PartialSend{} }
func (m *PartialSend) String() string { return proto.CompactTextString(m) }
func (*PartialSend) ProtoMessage()    {}
func (*PartialSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ba424868b7efb2, []int{0}
}
func (m *PartialSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PartialSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PartialSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PartialSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialSend.Merge(m, src)
}
func (m *PartialSend) XXX_Size() int {
	return m.Size()
}
func (m *PartialSend) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialSend.DiscardUnknown(m)
}

var xxx_messageInfo_PartialSend proto.InternalMessageInfo

func (m *PartialSend) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PartialSend) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *PartialSend) GetBurntAmount() types.Coin {
	if m != nil {
		return m.BurntAmount
	}
	return types.Coin{}
}

func (m *PartialSend) GetSentAmount() types.Coin {
	if m != nil {
		return m.SentAmount
	}
	return types.Coin{}
}

func (m *PartialSend) GetSwappedAmount() types.Coin {
	if m != nil {
		return m.SwappedAmount
	}
	return types.Coin{}
}

func (m *PartialSend) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*PartialSend)(nil), "magicswap.swap.PartialSend")
}

func init() { proto.RegisterFile("magicswap/swap/partial_send.proto", fileDescriptor_62ba424868b7efb2) }

var fileDescriptor_62ba424868b7efb2 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0x32, 0x31,
	0x14, 0x85, 0xa7, 0xf3, 0xf3, 0x63, 0x28, 0x91, 0x45, 0xe3, 0x62, 0x24, 0xa6, 0xa2, 0x2b, 0x12,
	0x93, 0x56, 0xf4, 0x01, 0x0c, 0x10, 0xf6, 0x06, 0x5d, 0xb9, 0x31, 0x9d, 0x99, 0x66, 0x6c, 0xe2,
	0xf4, 0x36, 0x6d, 0x51, 0x7c, 0x0b, 0x1f, 0xc1, 0xc7, 0x61, 0xc9, 0xd2, 0x95, 0x31, 0xcc, 0x8b,
	0x98, 0x99, 0x41, 0xc0, 0x1d, 0x9b, 0x9b, 0xb6, 0xf7, 0x9e, 0xef, 0xe4, 0xf6, 0xe0, 0xb3, 0x5c,
	0x64, 0x2a, 0x71, 0xaf, 0xc2, 0xf0, 0xaa, 0x18, 0x61, 0xbd, 0x12, 0xcf, 0x8f, 0x4e, 0xea, 0x94,
	0x19, 0x0b, 0x1e, 0x48, 0x67, 0x33, 0xc2, 0xca, 0xd2, 0x3d, 0xca, 0x20, 0x83, 0xaa, 0xc5, 0xcb,
	0x53, 0x3d, 0xd5, 0xa5, 0x09, 0xb8, 0x1c, 0x1c, 0x8f, 0x85, 0x93, 0xfc, 0x65, 0x10, 0x4b, 0x2f,
	0x06, 0x3c, 0x01, 0xa5, 0xeb, 0xfe, 0xf9, 0x47, 0x88, 0xdb, 0xb7, 0x35, 0xfc, 0x4e, 0xea, 0x94,
	0x74, 0x70, 0xa8, 0xd2, 0x08, 0xf5, 0x50, 0xbf, 0x31, 0x0d, 0x55, 0x4a, 0x4e, 0x70, 0xcb, 0xca,
	0x44, 0x19, 0x25, 0xb5, 0x8f, 0xc2, 0x1e, 0xea, 0xb7, 0xa6, 0xdb, 0x07, 0x32, 0xc4, 0xed, 0x78,
	0x66, 0xb5, 0x1f, 0xe6, 0x30, 0xd3, 0x3e, 0xfa, 0xd7, 0x43, 0xfd, 0xf6, 0xd5, 0x31, 0xab, 0x3d,
	0x59, 0xe9, 0xc9, 0xd6, 0x9e, 0x6c, 0x0c, 0x4a, 0x8f, 0x1a, 0x8b, 0xaf, 0xd3, 0x60, 0xba, 0xab,
	0x21, 0x37, 0x18, 0x3b, 0xb9, 0x21, 0x34, 0xf6, 0x23, 0xec, 0x48, 0xc8, 0x04, 0x1f, 0x96, 0xfb,
	0x1b, 0x99, 0xae, 0x19, 0xff, 0xf7, 0x63, 0xfc, 0x55, 0x91, 0x08, 0x1f, 0x24, 0x56, 0x0a, 0x0f,
	0x36, 0x6a, 0x56, 0x6b, 0xfe, 0x5e, 0x47, 0x93, 0xc5, 0x8a, 0xa2, 0xe5, 0x8a, 0xa2, 0xef, 0x15,
	0x45, 0xef, 0x05, 0x0d, 0x96, 0x05, 0x0d, 0x3e, 0x0b, 0x1a, 0x3c, 0x5c, 0x64, 0xca, 0x3f, 0xcd,
	0x62, 0x96, 0x40, 0xce, 0xef, 0xe7, 0x63, 0xb0, 0x46, 0x5d, 0xce, 0xf9, 0x36, 0xba, 0x79, 0x1d,
	0x9e, 0x7f, 0x33, 0xd2, 0xc5, 0xcd, 0xea, 0xc3, 0xaf, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x30,
	0xdc, 0xa6, 0x7b, 0xdb, 0x01, 0x00, 0x00,
}

func (m *PartialSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartialSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PartialSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPartialSend(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.SwappedAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPartialSend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.SentAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPartialSend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.BurntAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPartialSend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintPartialSend(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPartialSend(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPartialSend(dAtA []byte, offset int, v uint64) int {
	offset -= sovPartialSend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PartialSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPartialSend(uint64(m.Id))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovPartialSend(uint64(l))
	}
	l = m.BurntAmount.Size()
	n += 1 + l + sovPartialSend(uint64(l))
	l = m.SentAmount.Size()
	n += 1 + l + sovPartialSend(uint64(l))
	l = m.SwappedAmount.Size()
	n += 1 + l + sovPartialSend(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPartialSend(uint64(l))
	}
	return n
}

func sovPartialSend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPartialSend(x uint64) (n int) {
	return sovPartialSend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PartialSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPartialSend
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
			return fmt.Errorf("proto: PartialSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartialSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartialSend
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
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartialSend
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
				return ErrInvalidLengthPartialSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPartialSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurntAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartialSend
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
				return ErrInvalidLengthPartialSend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPartialSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurntAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartialSend
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
				return ErrInvalidLengthPartialSend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPartialSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SentAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwappedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartialSend
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
				return ErrInvalidLengthPartialSend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPartialSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwappedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartialSend
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
				return ErrInvalidLengthPartialSend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPartialSend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPartialSend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPartialSend
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
func skipPartialSend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPartialSend
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
					return 0, ErrIntOverflowPartialSend
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
					return 0, ErrIntOverflowPartialSend
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
				return 0, ErrInvalidLengthPartialSend
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPartialSend
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPartialSend
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPartialSend        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPartialSend          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPartialSend = fmt.Errorf("proto: unexpected end of group")
)
