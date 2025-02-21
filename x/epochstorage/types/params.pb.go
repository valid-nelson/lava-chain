// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/epochstorage/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	UnstakeHoldBlocks       uint64 `protobuf:"varint,1,opt,name=unstakeHoldBlocks,proto3" json:"unstakeHoldBlocks,omitempty" yaml:"unstake_hold_blocks"`
	EpochBlocks             uint64 `protobuf:"varint,2,opt,name=epochBlocks,proto3" json:"epochBlocks,omitempty" yaml:"epoch_blocks"`
	EpochsToSave            uint64 `protobuf:"varint,3,opt,name=epochsToSave,proto3" json:"epochsToSave,omitempty" yaml:"epochs_to_save"`
	LatestParamChange       uint64 `protobuf:"varint,4,opt,name=latestParamChange,proto3" json:"latestParamChange,omitempty" yaml:"latest_param_change"`
	UnstakeHoldBlocksStatic uint64 `protobuf:"varint,5,opt,name=unstakeHoldBlocksStatic,proto3" json:"unstakeHoldBlocksStatic,omitempty" yaml:"unstake_hold_blocks_static"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_09513f4cf6e403e6, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetUnstakeHoldBlocks() uint64 {
	if m != nil {
		return m.UnstakeHoldBlocks
	}
	return 0
}

func (m *Params) GetEpochBlocks() uint64 {
	if m != nil {
		return m.EpochBlocks
	}
	return 0
}

func (m *Params) GetEpochsToSave() uint64 {
	if m != nil {
		return m.EpochsToSave
	}
	return 0
}

func (m *Params) GetLatestParamChange() uint64 {
	if m != nil {
		return m.LatestParamChange
	}
	return 0
}

func (m *Params) GetUnstakeHoldBlocksStatic() uint64 {
	if m != nil {
		return m.UnstakeHoldBlocksStatic
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "lavanet.lava.epochstorage.Params")
}

func init() {
	proto.RegisterFile("lavanet/lava/epochstorage/params.proto", fileDescriptor_09513f4cf6e403e6)
}

var fileDescriptor_09513f4cf6e403e6 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x4f, 0xf2, 0x40,
	0x18, 0xc7, 0xdb, 0x17, 0x5e, 0x86, 0xea, 0x62, 0xd5, 0x50, 0x18, 0xae, 0xda, 0x44, 0xe3, 0xd4,
	0x4b, 0x74, 0x92, 0xc4, 0x05, 0x17, 0x4d, 0x1c, 0x4c, 0x71, 0x72, 0xb9, 0x1c, 0xe5, 0xd2, 0x12,
	0x0e, 0x9e, 0x86, 0x3b, 0x1a, 0xf9, 0x00, 0xee, 0x8e, 0x8e, 0x7e, 0x1c, 0x47, 0x46, 0xa7, 0xc6,
	0xc0, 0x37, 0xe8, 0x27, 0x30, 0x7d, 0x4a, 0x0c, 0xd8, 0x30, 0x5d, 0x93, 0xfe, 0xfe, 0xbf, 0xf6,
	0xff, 0xdc, 0x63, 0x9d, 0x4b, 0x9e, 0xf2, 0x89, 0xd0, 0xb4, 0x38, 0xa9, 0x48, 0x20, 0x8c, 0x95,
	0x86, 0x29, 0x8f, 0x04, 0x4d, 0xf8, 0x94, 0x8f, 0x95, 0x9f, 0x4c, 0x41, 0x83, 0xdd, 0x5a, 0x73,
	0x7e, 0x71, 0xfa, 0x9b, 0x5c, 0xfb, 0x28, 0x82, 0x08, 0x90, 0xa2, 0xc5, 0x53, 0x19, 0xf0, 0x5e,
	0x6b, 0x56, 0xe3, 0x11, 0x0d, 0xf6, 0x83, 0x75, 0x30, 0x9b, 0x28, 0xcd, 0x47, 0xe2, 0x0e, 0xe4,
	0xa0, 0x2b, 0x21, 0x1c, 0x29, 0xc7, 0x3c, 0x31, 0x2f, 0xea, 0x5d, 0x92, 0x67, 0x6e, 0x7b, 0xce,
	0xc7, 0xb2, 0xe3, 0xad, 0x11, 0x16, 0x83, 0x1c, 0xb0, 0x3e, 0x42, 0x5e, 0x50, 0x0d, 0xda, 0xd7,
	0xd6, 0x1e, 0x7e, 0x7e, 0xed, 0xf9, 0x87, 0x9e, 0x66, 0x9e, 0xb9, 0x87, 0xa5, 0x07, 0x5f, 0xfe,
	0x0a, 0x36, 0x59, 0xfb, 0xc6, 0xda, 0x2f, 0xff, 0xfc, 0x09, 0x7a, 0x3c, 0x15, 0x4e, 0x0d, 0xb3,
	0xad, 0x3c, 0x73, 0x8f, 0x37, 0xb2, 0x8a, 0x69, 0x60, 0x8a, 0xa7, 0xc2, 0x0b, 0xb6, 0xf0, 0xa2,
	0x87, 0xe4, 0x5a, 0x28, 0x8d, 0xbd, 0x6e, 0x63, 0x3e, 0x89, 0x84, 0x53, 0xff, 0xdb, 0xa3, 0x44,
	0x18, 0x4e, 0x8f, 0x85, 0x08, 0x79, 0x41, 0x35, 0x68, 0x33, 0xab, 0x59, 0x29, 0xd7, 0xd3, 0x5c,
	0x0f, 0x43, 0xe7, 0x3f, 0x3a, 0xcf, 0xf2, 0xcc, 0x3d, 0xdd, 0x39, 0x1b, 0xa6, 0x90, 0xf5, 0x82,
	0x5d, 0x96, 0x4e, 0xfd, 0xfd, 0xc3, 0x35, 0xba, 0xf7, 0x9f, 0x4b, 0x62, 0x2e, 0x96, 0xc4, 0xfc,
	0x5e, 0x12, 0xf3, 0x6d, 0x45, 0x8c, 0xc5, 0x8a, 0x18, 0x5f, 0x2b, 0x62, 0x3c, 0xd3, 0x68, 0xa8,
	0xe3, 0x59, 0xdf, 0x0f, 0x61, 0x4c, 0xb7, 0xb6, 0x20, 0xbd, 0xa4, 0x2f, 0xdb, 0xab, 0xa0, 0xe7,
	0x89, 0x50, 0xfd, 0x06, 0xde, 0xec, 0xd5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x4f, 0x74,
	0x80, 0x34, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnstakeHoldBlocksStatic != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnstakeHoldBlocksStatic))
		i--
		dAtA[i] = 0x28
	}
	if m.LatestParamChange != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LatestParamChange))
		i--
		dAtA[i] = 0x20
	}
	if m.EpochsToSave != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EpochsToSave))
		i--
		dAtA[i] = 0x18
	}
	if m.EpochBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EpochBlocks))
		i--
		dAtA[i] = 0x10
	}
	if m.UnstakeHoldBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnstakeHoldBlocks))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UnstakeHoldBlocks != 0 {
		n += 1 + sovParams(uint64(m.UnstakeHoldBlocks))
	}
	if m.EpochBlocks != 0 {
		n += 1 + sovParams(uint64(m.EpochBlocks))
	}
	if m.EpochsToSave != 0 {
		n += 1 + sovParams(uint64(m.EpochsToSave))
	}
	if m.LatestParamChange != 0 {
		n += 1 + sovParams(uint64(m.LatestParamChange))
	}
	if m.UnstakeHoldBlocksStatic != 0 {
		n += 1 + sovParams(uint64(m.UnstakeHoldBlocksStatic))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakeHoldBlocks", wireType)
			}
			m.UnstakeHoldBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnstakeHoldBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlocks", wireType)
			}
			m.EpochBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochsToSave", wireType)
			}
			m.EpochsToSave = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochsToSave |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestParamChange", wireType)
			}
			m.LatestParamChange = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestParamChange |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakeHoldBlocksStatic", wireType)
			}
			m.UnstakeHoldBlocksStatic = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnstakeHoldBlocksStatic |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
