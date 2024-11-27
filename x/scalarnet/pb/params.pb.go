// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/scalarnet/v1beta1/params.proto

package pb

import (
	fmt "fmt"
	github_com_axelarnetwork_axelar_core_x_nexus_exported "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params represent the genesis parameters for the module
type Params struct {
	// IBC packet route timeout window
	RouteTimeoutWindow               uint64                           `protobuf:"varint,1,opt,name=route_timeout_window,json=routeTimeoutWindow,proto3" json:"route_timeout_window,omitempty"`
	TransferLimit                    uint64                           `protobuf:"varint,3,opt,name=transfer_limit,json=transferLimit,proto3" json:"transfer_limit,omitempty"`
	EndBlockerLimit                  uint64                           `protobuf:"varint,4,opt,name=end_blocker_limit,json=endBlockerLimit,proto3" json:"end_blocker_limit,omitempty"`
	CallContractsProposalMinDeposits []CallContractProposalMinDeposit `protobuf:"bytes,5,rep,name=call_contracts_proposal_min_deposits,json=callContractsProposalMinDeposits,proto3" json:"call_contracts_proposal_min_deposits"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_68386da21503c6a6, []int{0}
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

type CallContractProposalMinDeposit struct {
	Chain           github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	ContractAddress string                                                          `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	MinDeposits     []types.Coin                                                    `protobuf:"bytes,3,rep,name=min_deposits,json=minDeposits,proto3" json:"min_deposits"`
}

func (m *CallContractProposalMinDeposit) Reset()         { *m = CallContractProposalMinDeposit{} }
func (m *CallContractProposalMinDeposit) String() string { return proto.CompactTextString(m) }
func (*CallContractProposalMinDeposit) ProtoMessage()    {}
func (*CallContractProposalMinDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_68386da21503c6a6, []int{1}
}
func (m *CallContractProposalMinDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CallContractProposalMinDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CallContractProposalMinDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CallContractProposalMinDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallContractProposalMinDeposit.Merge(m, src)
}
func (m *CallContractProposalMinDeposit) XXX_Size() int {
	return m.Size()
}
func (m *CallContractProposalMinDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_CallContractProposalMinDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_CallContractProposalMinDeposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "scalar.scalarnet.v1beta1.Params")
	proto.RegisterType((*CallContractProposalMinDeposit)(nil), "scalar.scalarnet.v1beta1.CallContractProposalMinDeposit")
}

func init() {
	proto.RegisterFile("scalar/scalarnet/v1beta1/params.proto", fileDescriptor_68386da21503c6a6)
}

var fileDescriptor_68386da21503c6a6 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x93, 0xb6, 0x9b, 0xc0, 0x03, 0x36, 0xa2, 0x1d, 0xc2, 0x0e, 0x69, 0x35, 0x31, 0x51,
	0x90, 0x1a, 0xaf, 0x70, 0xe1, 0x86, 0x48, 0x39, 0x21, 0xfe, 0x4c, 0x11, 0x12, 0x82, 0x4b, 0xe4,
	0x38, 0xa6, 0xb3, 0x9a, 0xf8, 0x8d, 0x6c, 0x97, 0x86, 0x6f, 0xc1, 0x89, 0x0f, 0xc1, 0x8d, 0x4f,
	0x41, 0x8f, 0x3b, 0x72, 0x1a, 0xd0, 0x7e, 0x0b, 0x4e, 0x28, 0x76, 0xda, 0x15, 0x89, 0xed, 0x64,
	0xe7, 0x79, 0x7e, 0x79, 0x9c, 0x3c, 0x79, 0x83, 0x8e, 0x14, 0x25, 0x39, 0x91, 0xd8, 0x2e, 0x82,
	0x69, 0xfc, 0x71, 0x98, 0x32, 0x4d, 0x86, 0xb8, 0x24, 0x92, 0x14, 0x2a, 0x2c, 0x25, 0x68, 0xf0,
	0x7c, 0xeb, 0x87, 0x6b, 0x2c, 0x6c, 0xb0, 0x83, 0xfd, 0x31, 0x8c, 0xc1, 0x40, 0xb8, 0xde, 0x59,
	0xfe, 0x20, 0xa0, 0xa0, 0x0a, 0x50, 0x38, 0x25, 0x8a, 0xad, 0x13, 0x29, 0x70, 0x61, 0xfd, 0xc3,
	0xef, 0x2d, 0xb4, 0x7d, 0x62, 0x0e, 0xf0, 0x8e, 0xd1, 0xbe, 0x84, 0xa9, 0x66, 0x89, 0xe6, 0x05,
	0x83, 0xa9, 0x4e, 0x66, 0x5c, 0x64, 0x30, 0xf3, 0xdd, 0x9e, 0xdb, 0xef, 0xc4, 0x9e, 0xf1, 0xde,
	0x58, 0xeb, 0xad, 0x71, 0xbc, 0x23, 0x74, 0x4b, 0x4b, 0x22, 0xd4, 0x07, 0x26, 0x93, 0x9c, 0x17,
	0x5c, 0xfb, 0x6d, 0xc3, 0xde, 0x5c, 0xa9, 0x2f, 0x6a, 0xd1, 0x7b, 0x80, 0x6e, 0x33, 0x91, 0x25,
	0x69, 0x0e, 0x74, 0xb2, 0x26, 0x3b, 0x86, 0xdc, 0x65, 0x22, 0x8b, 0xac, 0x6e, 0xd9, 0x6f, 0x2e,
	0xba, 0x4b, 0x49, 0x9e, 0x27, 0x14, 0x84, 0x96, 0x84, 0x6a, 0x95, 0x94, 0x12, 0x4a, 0x50, 0x24,
	0x4f, 0x0a, 0x2e, 0x92, 0x8c, 0x95, 0xa0, 0xb8, 0x56, 0xfe, 0x56, 0xaf, 0xdd, 0xdf, 0x79, 0xf8,
	0x38, 0xbc, 0xac, 0x8f, 0x70, 0x44, 0xf2, 0x7c, 0xd4, 0x84, 0x9c, 0x34, 0x11, 0x2f, 0xb9, 0x78,
	0x66, 0x03, 0xa2, 0x7b, 0xf3, 0xf3, 0xae, 0xf3, 0xf5, 0x67, 0xb7, 0x7b, 0x35, 0xa7, 0xe2, 0x1e,
	0xdd, 0x00, 0xd4, 0x7f, 0x88, 0xe7, 0x9d, 0x6b, 0xad, 0xbd, 0xf6, 0xe1, 0x97, 0x16, 0x0a, 0xae,
	0xce, 0xf2, 0xde, 0xa1, 0x2d, 0x7a, 0x4a, 0xb8, 0x30, 0x95, 0x5e, 0x8f, 0x46, 0x7f, 0xce, 0xbb,
	0x4f, 0xc6, 0x5c, 0x9f, 0x4e, 0xd3, 0x90, 0x42, 0x81, 0x49, 0xc5, 0xec, 0x3b, 0xcc, 0x40, 0x4e,
	0x9a, 0xab, 0x01, 0x05, 0xc9, 0x70, 0x85, 0x05, 0xab, 0xa6, 0x0a, 0xb3, 0xaa, 0x04, 0xa9, 0x59,
	0x16, 0x8e, 0xea, 0x98, 0x57, 0xa4, 0x60, 0xb1, 0x4d, 0xf4, 0xee, 0xa3, 0xbd, 0x55, 0x63, 0x09,
	0xc9, 0x32, 0xc9, 0x94, 0xf2, 0x5b, 0xf5, 0x29, 0xf1, 0xee, 0x4a, 0x7f, 0x6a, 0x65, 0x4f, 0xa0,
	0x1b, 0xff, 0x34, 0xd9, 0x36, 0x4d, 0xde, 0x09, 0xed, 0xa4, 0x84, 0xf5, 0xa4, 0x5c, 0x94, 0x08,
	0x5c, 0x44, 0xc7, 0x4d, 0x55, 0xfd, 0x8d, 0x67, 0x6d, 0xc6, 0xca, 0x2e, 0x03, 0x95, 0x4d, 0xb0,
	0xfe, 0x54, 0x32, 0x65, 0x6e, 0x50, 0xf1, 0x4e, 0x71, 0x51, 0x4f, 0xf4, 0x7a, 0xfe, 0x3b, 0x70,
	0xe6, 0x8b, 0xc0, 0x3d, 0x5b, 0x04, 0xee, 0xaf, 0x45, 0xe0, 0x7e, 0x5e, 0x06, 0xce, 0xd9, 0x32,
	0x70, 0x7e, 0x2c, 0x03, 0xe7, 0xfd, 0x70, 0x23, 0xd4, 0x7e, 0x44, 0x90, 0xe3, 0x66, 0x37, 0x30,
	0x33, 0xaa, 0x70, 0xb5, 0xf1, 0x57, 0x94, 0x69, 0xba, 0x6d, 0xe4, 0x47, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x66, 0xeb, 0x95, 0xe3, 0x33, 0x03, 0x00, 0x00,
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
	if len(m.CallContractsProposalMinDeposits) > 0 {
		for iNdEx := len(m.CallContractsProposalMinDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CallContractsProposalMinDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.EndBlockerLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EndBlockerLimit))
		i--
		dAtA[i] = 0x20
	}
	if m.TransferLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TransferLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.RouteTimeoutWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RouteTimeoutWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CallContractProposalMinDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallContractProposalMinDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CallContractProposalMinDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinDeposits) > 0 {
		for iNdEx := len(m.MinDeposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinDeposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
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
	if m.RouteTimeoutWindow != 0 {
		n += 1 + sovParams(uint64(m.RouteTimeoutWindow))
	}
	if m.TransferLimit != 0 {
		n += 1 + sovParams(uint64(m.TransferLimit))
	}
	if m.EndBlockerLimit != 0 {
		n += 1 + sovParams(uint64(m.EndBlockerLimit))
	}
	if len(m.CallContractsProposalMinDeposits) > 0 {
		for _, e := range m.CallContractsProposalMinDeposits {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *CallContractProposalMinDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.MinDeposits) > 0 {
		for _, e := range m.MinDeposits {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field RouteTimeoutWindow", wireType)
			}
			m.RouteTimeoutWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RouteTimeoutWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferLimit", wireType)
			}
			m.TransferLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransferLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockerLimit", wireType)
			}
			m.EndBlockerLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockerLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallContractsProposalMinDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallContractsProposalMinDeposits = append(m.CallContractsProposalMinDeposits, CallContractProposalMinDeposit{})
			if err := m.CallContractsProposalMinDeposits[len(m.CallContractsProposalMinDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *CallContractProposalMinDeposit) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CallContractProposalMinDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallContractProposalMinDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinDeposits = append(m.MinDeposits, types.Coin{})
			if err := m.MinDeposits[len(m.MinDeposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
