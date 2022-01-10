// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid/superfluid.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type SuperfluidAssetType int32

const (
	SuperfluidAssetTypeNative  SuperfluidAssetType = 0
	SuperfluidAssetTypeLPShare SuperfluidAssetType = 1
)

var SuperfluidAssetType_name = map[int32]string{
	0: "SuperfluidAssetTypeNative",
	1: "SuperfluidAssetTypeLPShare",
}

var SuperfluidAssetType_value = map[string]int32{
	"SuperfluidAssetTypeNative":  0,
	"SuperfluidAssetTypeLPShare": 1,
}

func (x SuperfluidAssetType) String() string {
	return proto.EnumName(SuperfluidAssetType_name, int32(x))
}

func (SuperfluidAssetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{0}
}

// SuperfluidAsset stores the pair of superfluid asset type and denom pair
type SuperfluidAsset struct {
	Denom     string              `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	AssetType SuperfluidAssetType `protobuf:"varint,2,opt,name=asset_type,json=assetType,proto3,enum=osmosis.superfluid.SuperfluidAssetType" json:"asset_type,omitempty"`
}

func (m *SuperfluidAsset) Reset()         { *m = SuperfluidAsset{} }
func (m *SuperfluidAsset) String() string { return proto.CompactTextString(m) }
func (*SuperfluidAsset) ProtoMessage()    {}
func (*SuperfluidAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{0}
}
func (m *SuperfluidAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuperfluidAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuperfluidAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuperfluidAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperfluidAsset.Merge(m, src)
}
func (m *SuperfluidAsset) XXX_Size() int {
	return m.Size()
}
func (m *SuperfluidAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperfluidAsset.DiscardUnknown(m)
}

var xxx_messageInfo_SuperfluidAsset proto.InternalMessageInfo

// SuperfluidIntermediaryAccount takes the role of intermediary between LP token and OSMO tokens for superfluid staking
type SuperfluidIntermediaryAccount struct {
	Denom   string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	ValAddr string `protobuf:"bytes,2,opt,name=val_addr,json=valAddr,proto3" json:"val_addr,omitempty"`
	// id of perpetual gauge to send rewards to for distribution
	GaugeId uint64 `protobuf:"varint,3,opt,name=gauge_id,json=gaugeId,proto3" json:"gauge_id,omitempty"`
}

func (m *SuperfluidIntermediaryAccount) Reset()         { *m = SuperfluidIntermediaryAccount{} }
func (m *SuperfluidIntermediaryAccount) String() string { return proto.CompactTextString(m) }
func (*SuperfluidIntermediaryAccount) ProtoMessage()    {}
func (*SuperfluidIntermediaryAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{1}
}
func (m *SuperfluidIntermediaryAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuperfluidIntermediaryAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuperfluidIntermediaryAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuperfluidIntermediaryAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperfluidIntermediaryAccount.Merge(m, src)
}
func (m *SuperfluidIntermediaryAccount) XXX_Size() int {
	return m.Size()
}
func (m *SuperfluidIntermediaryAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperfluidIntermediaryAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SuperfluidIntermediaryAccount proto.InternalMessageInfo

func (m *SuperfluidIntermediaryAccount) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *SuperfluidIntermediaryAccount) GetValAddr() string {
	if m != nil {
		return m.ValAddr
	}
	return ""
}

func (m *SuperfluidIntermediaryAccount) GetGaugeId() uint64 {
	if m != nil {
		return m.GaugeId
	}
	return 0
}

func init() {
	proto.RegisterEnum("osmosis.superfluid.SuperfluidAssetType", SuperfluidAssetType_name, SuperfluidAssetType_value)
	proto.RegisterType((*SuperfluidAsset)(nil), "osmosis.superfluid.SuperfluidAsset")
	proto.RegisterType((*SuperfluidIntermediaryAccount)(nil), "osmosis.superfluid.SuperfluidIntermediaryAccount")
}

func init() {
	proto.RegisterFile("osmosis/superfluid/superfluid.proto", fileDescriptor_79d3c29d82dbb734)
}

var fileDescriptor_79d3c29d82dbb734 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x8e, 0xda, 0x30,
	0x1c, 0xc6, 0xe3, 0x96, 0x16, 0xf0, 0xd0, 0xa2, 0x94, 0x01, 0x22, 0x61, 0x10, 0x1d, 0x8a, 0x2a,
	0x35, 0x16, 0x74, 0xeb, 0x46, 0x87, 0x4a, 0x48, 0xa8, 0xaa, 0x42, 0x27, 0x16, 0xe4, 0xc4, 0x26,
	0x58, 0x4d, 0xe2, 0xc8, 0x76, 0xa2, 0xe6, 0x0d, 0x18, 0xef, 0x11, 0x90, 0xee, 0x65, 0x6e, 0x64,
	0xbc, 0xf1, 0x04, 0xcb, 0x3d, 0xc6, 0x29, 0x21, 0x1c, 0x88, 0xe3, 0xb6, 0xff, 0x97, 0xdf, 0xa7,
	0x2f, 0x9f, 0xff, 0x36, 0xfc, 0x2c, 0x54, 0x28, 0x14, 0x57, 0x58, 0x25, 0x31, 0x93, 0xcb, 0x20,
	0xe1, 0xf4, 0x6c, 0xb4, 0x63, 0x29, 0xb4, 0x30, 0xcd, 0xd2, 0x64, 0x9f, 0x88, 0xd5, 0xf4, 0x85,
	0x2f, 0x0a, 0x8c, 0xf3, 0xe9, 0xe0, 0xb4, 0x90, 0x2f, 0x84, 0x1f, 0x30, 0x5c, 0x28, 0x37, 0x59,
	0x62, 0x9a, 0x48, 0xa2, 0xb9, 0x88, 0x4a, 0xde, 0xbd, 0xe4, 0x9a, 0x87, 0x4c, 0x69, 0x12, 0xc6,
	0xc7, 0x00, 0xaf, 0xf8, 0x17, 0x76, 0x89, 0x62, 0x38, 0x1d, 0xba, 0x4c, 0x93, 0x21, 0xf6, 0x04,
	0x2f, 0x03, 0xfa, 0x19, 0xfc, 0x38, 0x7b, 0x2e, 0x31, 0x56, 0x8a, 0x69, 0xb3, 0x09, 0xdf, 0x51,
	0x16, 0x89, 0xb0, 0x05, 0x7a, 0x60, 0x50, 0x77, 0x0e, 0xc2, 0xfc, 0x05, 0x21, 0xc9, 0xf1, 0x42,
	0x67, 0x31, 0x6b, 0xbd, 0xe9, 0x81, 0xc1, 0x87, 0xd1, 0x17, 0xfb, 0xe5, 0x41, 0xec, 0x8b, 0xb8,
	0xbf, 0x59, 0xcc, 0x9c, 0x3a, 0x39, 0x8e, 0x3f, 0x6a, 0xeb, 0x4d, 0xd7, 0x78, 0xdc, 0x74, 0x41,
	0xff, 0x1f, 0xec, 0x9c, 0xbc, 0x93, 0x48, 0x33, 0x19, 0x32, 0xca, 0x89, 0xcc, 0xc6, 0x9e, 0x27,
	0x92, 0xe8, 0xb5, 0x22, 0x6d, 0x58, 0x4b, 0x49, 0xb0, 0x20, 0x94, 0xca, 0xa2, 0x46, 0xdd, 0xa9,
	0xa6, 0x24, 0x18, 0x53, 0x2a, 0x73, 0xe4, 0x93, 0xc4, 0x67, 0x0b, 0x4e, 0x5b, 0x6f, 0x7b, 0x60,
	0x50, 0x71, 0xaa, 0x85, 0x9e, 0xd0, 0xaf, 0x73, 0xf8, 0xe9, 0x4a, 0x31, 0xb3, 0x03, 0xdb, 0x57,
	0x3e, 0xff, 0x26, 0x9a, 0xa7, 0xac, 0x61, 0x98, 0x08, 0x5a, 0x57, 0xf0, 0xf4, 0xcf, 0x6c, 0x45,
	0x24, 0x6b, 0x00, 0xab, 0xb2, 0xbe, 0x45, 0xc6, 0xcf, 0xe9, 0xdd, 0x0e, 0x81, 0xed, 0x0e, 0x81,
	0x87, 0x1d, 0x02, 0x37, 0x7b, 0x64, 0x6c, 0xf7, 0xc8, 0xb8, 0xdf, 0x23, 0x63, 0x3e, 0xf2, 0xb9,
	0x5e, 0x25, 0xae, 0xed, 0x89, 0x10, 0x97, 0xab, 0xfa, 0x16, 0x10, 0x57, 0x1d, 0x05, 0xfe, 0x7f,
	0xfe, 0x4e, 0xf2, 0xcd, 0x2a, 0xf7, 0x7d, 0x71, 0x31, 0xdf, 0x9f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x91, 0x33, 0x0e, 0x96, 0x4a, 0x02, 0x00, 0x00,
}

func (this *SuperfluidAsset) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SuperfluidAsset)
	if !ok {
		that2, ok := that.(SuperfluidAsset)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if this.AssetType != that1.AssetType {
		return false
	}
	return true
}
func (m *SuperfluidAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuperfluidAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuperfluidAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AssetType != 0 {
		i = encodeVarintSuperfluid(dAtA, i, uint64(m.AssetType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SuperfluidIntermediaryAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuperfluidIntermediaryAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuperfluidIntermediaryAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GaugeId != 0 {
		i = encodeVarintSuperfluid(dAtA, i, uint64(m.GaugeId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ValAddr) > 0 {
		i -= len(m.ValAddr)
		copy(dAtA[i:], m.ValAddr)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.ValAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSuperfluid(dAtA []byte, offset int, v uint64) int {
	offset -= sovSuperfluid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SuperfluidAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	if m.AssetType != 0 {
		n += 1 + sovSuperfluid(uint64(m.AssetType))
	}
	return n
}

func (m *SuperfluidIntermediaryAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	l = len(m.ValAddr)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	if m.GaugeId != 0 {
		n += 1 + sovSuperfluid(uint64(m.GaugeId))
	}
	return n
}

func sovSuperfluid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSuperfluid(x uint64) (n int) {
	return sovSuperfluid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SuperfluidAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperfluid
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
			return fmt.Errorf("proto: SuperfluidAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuperfluidAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetType", wireType)
			}
			m.AssetType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetType |= SuperfluidAssetType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSuperfluid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSuperfluid
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
func (m *SuperfluidIntermediaryAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperfluid
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
			return fmt.Errorf("proto: SuperfluidIntermediaryAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuperfluidIntermediaryAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeId", wireType)
			}
			m.GaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSuperfluid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSuperfluid
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
func skipSuperfluid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSuperfluid
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
					return 0, ErrIntOverflowSuperfluid
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
					return 0, ErrIntOverflowSuperfluid
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
				return 0, ErrInvalidLengthSuperfluid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSuperfluid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSuperfluid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSuperfluid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSuperfluid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSuperfluid = fmt.Errorf("proto: unexpected end of group")
)
