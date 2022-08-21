// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: empowerchain/proofofexistence/proof.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Proof struct {
	// Hash is the SHA-256 hash of the data of which is being made a proof for.
	// The hash needs to be sent as a Base64 encoded string.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// Timestamp is the time the proof was added on-chain (block time)
	Timestamp time.Time `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp" yaml:"timestamp"`
	// Reporter is the account address that created the proof
	Reporter string `protobuf:"bytes,3,opt,name=reporter,proto3" json:"reporter,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_df2ebf73a2cd905c, []int{0}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Proof) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Proof) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

func init() {
	proto.RegisterType((*Proof)(nil), "empowerchain.empowerchain.proofofexistence.Proof")
}

func init() {
	proto.RegisterFile("empowerchain/proofofexistence/proof.proto", fileDescriptor_df2ebf73a2cd905c)
}

var fileDescriptor_df2ebf73a2cd905c = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcd, 0x2d, 0xc8,
	0x2f, 0x4f, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x28, 0xca, 0xcf, 0x4f, 0xcb, 0x4f,
	0x4b, 0xad, 0xc8, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x85, 0x08, 0xe8, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0x69, 0x21, 0x2b, 0xd5, 0x43, 0xe1, 0xa0, 0xeb, 0x93, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x01, 0xeb, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49,
	0xcc, 0x2d, 0x80, 0x18, 0x26, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x66, 0xea, 0x83, 0x58, 0x10,
	0x51, 0xa5, 0x7e, 0x46, 0x2e, 0xd6, 0x00, 0x90, 0x59, 0x42, 0x42, 0x5c, 0x2c, 0x19, 0x89, 0xc5,
	0x19, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x50, 0x18, 0x17, 0x27, 0xdc, 0x18,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x29, 0x3d, 0x88, 0x45, 0x7a, 0x30, 0x8b, 0xf4, 0x42,
	0x60, 0x2a, 0x9c, 0x64, 0x4e, 0xdc, 0x93, 0x67, 0xf8, 0x74, 0x4f, 0x5e, 0xa0, 0x32, 0x31, 0x37,
	0xc7, 0x4a, 0x09, 0xae, 0x55, 0x69, 0xc2, 0x7d, 0x79, 0xc6, 0x20, 0x84, 0x51, 0x42, 0x52, 0x5c,
	0x1c, 0x45, 0xa9, 0x05, 0xf9, 0x45, 0x25, 0xa9, 0x45, 0x12, 0xcc, 0x60, 0xfb, 0xe0, 0x7c, 0xa7,
	0xf0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0x09, 0x44, 0x14, 0x4e, 0x05, 0x66, 0x98, 0x96,
	0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x5d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x80,
	0x9f, 0x55, 0x9b, 0x81, 0x01, 0x00, 0x00,
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reporter) > 0 {
		i -= len(m.Reporter)
		copy(dAtA[i:], m.Reporter)
		i = encodeVarintProof(dAtA, i, uint64(len(m.Reporter)))
		i--
		dAtA[i] = 0x1a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintProof(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintProof(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProof(dAtA []byte, offset int, v uint64) int {
	offset -= sovProof(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovProof(uint64(l))
	l = len(m.Reporter)
	if l > 0 {
		n += 1 + l + sovProof(uint64(l))
	}
	return n
}

func sovProof(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProof(x uint64) (n int) {
	return sovProof(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProof
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProof
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
				return ErrInvalidLengthProof
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProof
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProof(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProof
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
func skipProof(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
					return 0, ErrIntOverflowProof
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
				return 0, ErrInvalidLengthProof
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProof
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProof
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProof        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProof          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProof = fmt.Errorf("proto: unexpected end of group")
)
