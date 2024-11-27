// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/scalarnet/v1beta1/proposal.proto

package pb

import (
	fmt "fmt"
	_ "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	github_com_axelarnetwork_axelar_core_x_nexus_exported "github.com/axelarnetwork/axelar-core/x/nexus/exported"
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

// CallContractsProposal is a gov Content type for calling contracts on other
// chains
type CallContractsProposal struct {
	Title         string         `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ContractCalls []ContractCall `protobuf:"bytes,3,rep,name=contract_calls,json=contractCalls,proto3" json:"contract_calls"`
}

func (m *CallContractsProposal) Reset()         { *m = CallContractsProposal{} }
func (m *CallContractsProposal) String() string { return proto.CompactTextString(m) }
func (*CallContractsProposal) ProtoMessage()    {}
func (*CallContractsProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_199eedd75915a81d, []int{0}
}
func (m *CallContractsProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CallContractsProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CallContractsProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CallContractsProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallContractsProposal.Merge(m, src)
}
func (m *CallContractsProposal) XXX_Size() int {
	return m.Size()
}
func (m *CallContractsProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CallContractsProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CallContractsProposal proto.InternalMessageInfo

type ContractCall struct {
	Chain           github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	ContractAddress string                                                          `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Payload         []byte                                                          `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *ContractCall) Reset()         { *m = ContractCall{} }
func (m *ContractCall) String() string { return proto.CompactTextString(m) }
func (*ContractCall) ProtoMessage()    {}
func (*ContractCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_199eedd75915a81d, []int{1}
}
func (m *ContractCall) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractCall.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCall.Merge(m, src)
}
func (m *ContractCall) XXX_Size() int {
	return m.Size()
}
func (m *ContractCall) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCall.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCall proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CallContractsProposal)(nil), "scalar.scalarnet.v1beta1.CallContractsProposal")
	proto.RegisterType((*ContractCall)(nil), "scalar.scalarnet.v1beta1.ContractCall")
}

func init() {
	proto.RegisterFile("scalar/scalarnet/v1beta1/proposal.proto", fileDescriptor_199eedd75915a81d)
}

var fileDescriptor_199eedd75915a81d = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0x6e, 0xdb, 0x30,
	0x10, 0xc6, 0xc5, 0xda, 0x6e, 0x51, 0xda, 0xfd, 0x03, 0xc1, 0x05, 0x04, 0x0f, 0xb2, 0xe0, 0xa1,
	0xb5, 0x07, 0x8b, 0x70, 0xbb, 0x75, 0x29, 0x6a, 0xed, 0x49, 0xa0, 0x4c, 0xc9, 0x12, 0x50, 0x14,
	0x21, 0x0b, 0xa1, 0x45, 0x82, 0xa4, 0x13, 0xf9, 0x2d, 0x32, 0x66, 0xcc, 0x9c, 0x39, 0x0f, 0xe1,
	0xd1, 0x63, 0x26, 0x23, 0xb1, 0xdf, 0x22, 0x53, 0x20, 0x51, 0x36, 0x84, 0x00, 0x99, 0x74, 0xf7,
	0xe9, 0x87, 0xef, 0xbe, 0x3b, 0x09, 0xfe, 0x52, 0x04, 0x33, 0x2c, 0x91, 0x79, 0x64, 0x54, 0xa3,
	0xab, 0x49, 0x44, 0x35, 0x9e, 0x20, 0x21, 0xb9, 0xe0, 0x0a, 0x33, 0x5f, 0x48, 0xae, 0xb9, 0xed,
	0x18, 0xc2, 0x3f, 0x80, 0x7e, 0x05, 0xf6, 0xba, 0x09, 0x4f, 0x78, 0x09, 0xa1, 0xa2, 0x32, 0x7c,
	0x6f, 0x84, 0x73, 0x5a, 0x18, 0x67, 0x34, 0x5f, 0x28, 0x44, 0x73, 0xc1, 0xa5, 0xa6, 0xf1, 0xc1,
	0x5d, 0x2f, 0x05, 0x55, 0x06, 0x1d, 0xdc, 0x03, 0xf8, 0x23, 0xc0, 0x8c, 0x05, 0x3c, 0xd3, 0x12,
	0x13, 0xad, 0x4e, 0xaa, 0xd1, 0x76, 0x17, 0xb6, 0x74, 0xaa, 0x19, 0x75, 0x80, 0x07, 0x86, 0x9f,
	0x43, 0xd3, 0xd8, 0x1e, 0x6c, 0xc7, 0x54, 0x11, 0x99, 0x0a, 0x9d, 0xf2, 0xcc, 0xf9, 0x50, 0xbe,
	0xab, 0x4b, 0xf6, 0x29, 0xfc, 0x4a, 0x2a, 0xb3, 0x0b, 0x82, 0x19, 0x53, 0x4e, 0xc3, 0x6b, 0x0c,
	0xdb, 0xbf, 0x7f, 0xfa, 0xef, 0x6d, 0xe1, 0xef, 0x87, 0x17, 0x41, 0xa6, 0xcd, 0xd5, 0xa6, 0x6f,
	0x85, 0x5f, 0x48, 0x4d, 0x53, 0x7f, 0x9b, 0xb7, 0x77, 0x7d, 0x30, 0x78, 0x00, 0xb0, 0x53, 0x67,
	0xed, 0x33, 0xd8, 0x22, 0x33, 0x9c, 0x66, 0x26, 0xe3, 0x34, 0x78, 0xd9, 0xf4, 0xff, 0x25, 0xa9,
	0x9e, 0x2d, 0x22, 0x9f, 0xf0, 0x39, 0x32, 0x67, 0xc8, 0xa8, 0xbe, 0xe6, 0xf2, 0xb2, 0xea, 0xc6,
	0x84, 0x4b, 0x8a, 0xf2, 0x37, 0xb7, 0xf1, 0x83, 0xc2, 0xe6, 0x08, 0xcf, 0x69, 0x68, 0x1c, 0xed,
	0x11, 0xfc, 0x7e, 0x58, 0x03, 0xc7, 0xb1, 0xa4, 0x4a, 0x55, 0xdb, 0x7e, 0xdb, 0xeb, 0xff, 0x8d,
	0x6c, 0x3b, 0xf0, 0x93, 0xc0, 0x4b, 0xc6, 0x71, 0xec, 0x34, 0x3c, 0x30, 0xec, 0x84, 0xfb, 0xd6,
	0xc4, 0x9e, 0x1e, 0xaf, 0x9e, 0x5d, 0x6b, 0xb5, 0x75, 0xc1, 0x7a, 0xeb, 0x82, 0xa7, 0xad, 0x0b,
	0x6e, 0x76, 0xae, 0xb5, 0xde, 0xb9, 0xd6, 0xe3, 0xce, 0xb5, 0xce, 0x27, 0xb5, 0xc0, 0xe6, 0x34,
	0x5c, 0x26, 0x55, 0x35, 0x2e, 0x3f, 0x92, 0x42, 0x79, 0xed, 0x1f, 0x11, 0x51, 0xf4, 0xb1, 0x94,
	0xff, 0xbc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x48, 0xca, 0xc7, 0x41, 0x02, 0x00, 0x00,
}

func (m *CallContractsProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CallContractsProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CallContractsProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractCalls) > 0 {
		for iNdEx := len(m.ContractCalls) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractCalls[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContractCall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractCall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractCall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CallContractsProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.ContractCalls) > 0 {
		for _, e := range m.ContractCalls {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *ContractCall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CallContractsProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: CallContractsProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CallContractsProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractCalls", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractCalls = append(m.ContractCalls, ContractCall{})
			if err := m.ContractCalls[len(m.ContractCalls)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *ContractCall) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: ContractCall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractCall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
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
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)