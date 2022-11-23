// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/genesis.proto

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

// GenesisState defines the dex module's genesis state.
type GenesisState struct {
	Params        Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ContractState []ContractState `protobuf:"bytes,2,rep,name=contractState,proto3" json:"contractState"`
	LastEpoch     uint64          `protobuf:"varint,3,opt,name=lastEpoch,proto3" json:"lastEpoch,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a803aaabd08db59d, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetContractState() []ContractState {
	if m != nil {
		return m.ContractState
	}
	return nil
}

func (m *GenesisState) GetLastEpoch() uint64 {
	if m != nil {
		return m.LastEpoch
	}
	return 0
}

type ContractState struct {
	ContractInfo        ContractInfoV2       `protobuf:"bytes,1,opt,name=contractInfo,proto3" json:"contractInfo"`
	LongBookList        []LongBook           `protobuf:"bytes,2,rep,name=longBookList,proto3" json:"longBookList"`
	ShortBookList       []ShortBook          `protobuf:"bytes,3,rep,name=shortBookList,proto3" json:"shortBookList"`
	TriggeredOrdersList []Order              `protobuf:"bytes,4,rep,name=triggeredOrdersList,proto3" json:"triggeredOrdersList"`
	PairList            []Pair               `protobuf:"bytes,5,rep,name=pairList,proto3" json:"pairList"`
	PriceList           []ContractPairPrices `protobuf:"bytes,6,rep,name=priceList,proto3" json:"priceList"`
}

func (m *ContractState) Reset()         { *m = ContractState{} }
func (m *ContractState) String() string { return proto.CompactTextString(m) }
func (*ContractState) ProtoMessage()    {}
func (*ContractState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a803aaabd08db59d, []int{1}
}
func (m *ContractState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractState.Merge(m, src)
}
func (m *ContractState) XXX_Size() int {
	return m.Size()
}
func (m *ContractState) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractState.DiscardUnknown(m)
}

var xxx_messageInfo_ContractState proto.InternalMessageInfo

func (m *ContractState) GetContractInfo() ContractInfoV2 {
	if m != nil {
		return m.ContractInfo
	}
	return ContractInfoV2{}
}

func (m *ContractState) GetLongBookList() []LongBook {
	if m != nil {
		return m.LongBookList
	}
	return nil
}

func (m *ContractState) GetShortBookList() []ShortBook {
	if m != nil {
		return m.ShortBookList
	}
	return nil
}

func (m *ContractState) GetTriggeredOrdersList() []Order {
	if m != nil {
		return m.TriggeredOrdersList
	}
	return nil
}

func (m *ContractState) GetPairList() []Pair {
	if m != nil {
		return m.PairList
	}
	return nil
}

func (m *ContractState) GetPriceList() []ContractPairPrices {
	if m != nil {
		return m.PriceList
	}
	return nil
}

type ContractPairPrices struct {
	PricePair Pair     `protobuf:"bytes,1,opt,name=pricePair,proto3" json:"pricePair"`
	Prices    []*Price `protobuf:"bytes,6,rep,name=prices,proto3" json:"prices,omitempty"`
}

func (m *ContractPairPrices) Reset()         { *m = ContractPairPrices{} }
func (m *ContractPairPrices) String() string { return proto.CompactTextString(m) }
func (*ContractPairPrices) ProtoMessage()    {}
func (*ContractPairPrices) Descriptor() ([]byte, []int) {
	return fileDescriptor_a803aaabd08db59d, []int{2}
}
func (m *ContractPairPrices) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractPairPrices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractPairPrices.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractPairPrices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractPairPrices.Merge(m, src)
}
func (m *ContractPairPrices) XXX_Size() int {
	return m.Size()
}
func (m *ContractPairPrices) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractPairPrices.DiscardUnknown(m)
}

var xxx_messageInfo_ContractPairPrices proto.InternalMessageInfo

func (m *ContractPairPrices) GetPricePair() Pair {
	if m != nil {
		return m.PricePair
	}
	return Pair{}
}

func (m *ContractPairPrices) GetPrices() []*Price {
	if m != nil {
		return m.Prices
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "seiprotocol.seichain.dex.GenesisState")
	proto.RegisterType((*ContractState)(nil), "seiprotocol.seichain.dex.ContractState")
	proto.RegisterType((*ContractPairPrices)(nil), "seiprotocol.seichain.dex.ContractPairPrices")
}

func init() { proto.RegisterFile("dex/genesis.proto", fileDescriptor_a803aaabd08db59d) }

var fileDescriptor_a803aaabd08db59d = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0x1b, 0x5b, 0x8b, 0x9b, 0x6d, 0xfd, 0x93, 0x5d, 0xa4, 0x14, 0x99, 0x2d, 0xf5, 0x60,
	0x0f, 0xee, 0x0c, 0xd4, 0x83, 0x37, 0x91, 0x8a, 0x2c, 0x42, 0x61, 0x4b, 0x0b, 0x0a, 0x5e, 0xca,
	0x34, 0x8d, 0x33, 0xa1, 0xed, 0x64, 0x48, 0xb2, 0x50, 0xfd, 0x14, 0xfa, 0x91, 0xbc, 0xed, 0x71,
	0x8f, 0x9e, 0x44, 0xda, 0x0f, 0xe0, 0x57, 0x90, 0xbc, 0x7d, 0xb3, 0x6d, 0xd1, 0xd9, 0xf5, 0x36,
	0x79, 0xf2, 0xbc, 0xbf, 0x3c, 0xc9, 0xfb, 0x0e, 0x7d, 0x34, 0x15, 0xcb, 0x28, 0x11, 0x99, 0x30,
	0xd2, 0x84, 0xb9, 0x56, 0x56, 0xb1, 0x86, 0x11, 0x12, 0xbe, 0xb8, 0x9a, 0x87, 0x46, 0x48, 0x9e,
	0xc6, 0x32, 0x0b, 0xa7, 0x62, 0xd9, 0x3c, 0x4e, 0x54, 0xa2, 0x60, 0x2b, 0x72, 0x5f, 0x1b, 0x7f,
	0xf3, 0xa1, 0x43, 0xe4, 0xb1, 0x8e, 0x17, 0x48, 0x68, 0x1e, 0x39, 0x65, 0xae, 0xb2, 0x64, 0x3c,
	0x51, 0x6a, 0x86, 0xe2, 0xb1, 0x13, 0x4d, 0xaa, 0xb4, 0xdd, 0x55, 0xc1, 0x6a, 0x25, 0x9f, 0x8d,
	0x8d, 0xfc, 0x22, 0x50, 0x7c, 0xe0, 0x44, 0xa5, 0xa7, 0x42, 0xa3, 0xc0, 0x9c, 0xc0, 0x55, 0x66,
	0x75, 0xcc, 0x2d, 0x6a, 0xf7, 0x37, 0xc7, 0x4a, 0xef, 0x79, 0xec, 0xd6, 0x8b, 0xd8, 0xf2, 0x74,
	0xac, 0x85, 0xb9, 0x98, 0xdb, 0x5d, 0x58, 0xae, 0x25, 0x47, 0x7a, 0xfb, 0x3b, 0xa1, 0xb5, 0xb3,
	0xcd, 0x8d, 0x47, 0x36, 0xb6, 0x82, 0xbd, 0xa2, 0xd5, 0x4d, 0xfc, 0x06, 0x69, 0x91, 0xce, 0x61,
	0xb7, 0x15, 0x16, 0xbd, 0x40, 0x38, 0x00, 0x5f, 0xaf, 0x72, 0xf9, 0xf3, 0xa4, 0x34, 0xc4, 0x2a,
	0x36, 0xa2, 0x75, 0x9f, 0x0d, 0x80, 0x8d, 0x3b, 0xad, 0x72, 0xe7, 0xb0, 0xfb, 0xac, 0x18, 0xf3,
	0x66, 0xd7, 0x8e, 0xb4, 0x7d, 0x06, 0x7b, 0x42, 0x0f, 0xe6, 0xb1, 0xb1, 0x6f, 0x73, 0xc5, 0xd3,
	0x46, 0xb9, 0x45, 0x3a, 0x95, 0xe1, 0x56, 0x68, 0xff, 0x2e, 0xd3, 0xfa, 0x1e, 0x84, 0x0d, 0x69,
	0xcd, 0x03, 0xde, 0x65, 0x9f, 0x14, 0x5e, 0xa5, 0x73, 0x7b, 0x06, 0xe7, 0x7e, 0xdf, 0xc5, 0x10,
	0x7b, 0x0c, 0xd6, 0xa7, 0x35, 0xd7, 0xc5, 0x9e, 0x52, 0xb3, 0xbe, 0x34, 0x16, 0xef, 0xd5, 0x2e,
	0x66, 0xf6, 0xd1, 0xed, 0x69, 0xbb, 0xd5, 0xec, 0x9c, 0xd6, 0xa1, 0xfd, 0xd7, 0xb8, 0x32, 0xe0,
	0x9e, 0x16, 0xe3, 0x46, 0xde, 0xee, 0x9f, 0x68, 0xaf, 0x9e, 0x7d, 0xa0, 0x47, 0x56, 0xcb, 0x24,
	0x11, 0x5a, 0x4c, 0xcf, 0xdd, 0xb4, 0x18, 0xc0, 0x56, 0x00, 0x7b, 0x52, 0x8c, 0x05, 0x2f, 0x22,
	0xff, 0x45, 0x60, 0xaf, 0xe9, 0x3d, 0x37, 0x58, 0x40, 0xbb, 0x0b, 0xb4, 0xe0, 0xa6, 0x91, 0x90,
	0x1e, 0x76, 0x5d, 0xc5, 0x06, 0xf4, 0x00, 0x46, 0x0e, 0x10, 0x55, 0x40, 0x3c, 0xbf, 0xbd, 0x15,
	0x0e, 0x35, 0x70, 0x65, 0x7e, 0xc2, 0xb6, 0x90, 0xf6, 0x37, 0x42, 0xd9, 0xdf, 0x3e, 0xd6, 0xc3,
	0x83, 0x9c, 0x84, 0x3d, 0xff, 0xbf, 0xac, 0xdb, 0x32, 0xf6, 0x92, 0x56, 0x61, 0x61, 0x30, 0xe9,
	0x0d, 0x4f, 0x07, 0xa7, 0x0e, 0xd1, 0xde, 0x3b, 0xbb, 0x5c, 0x05, 0xe4, 0x6a, 0x15, 0x90, 0x5f,
	0xab, 0x80, 0x7c, 0x5d, 0x07, 0xa5, 0xab, 0x75, 0x50, 0xfa, 0xb1, 0x0e, 0x4a, 0x1f, 0x4f, 0x13,
	0x69, 0xd3, 0x8b, 0x49, 0xc8, 0xd5, 0x22, 0x32, 0x42, 0x9e, 0x7a, 0x1a, 0x2c, 0x00, 0x17, 0x2d,
	0x23, 0xf8, 0xf5, 0x3f, 0xe7, 0xc2, 0x4c, 0xaa, 0xb0, 0xff, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9c, 0x61, 0x6b, 0x1e, 0x8e, 0x04, 0x00, 0x00,
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
	if m.LastEpoch != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastEpoch))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContractState) > 0 {
		for iNdEx := len(m.ContractState) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractState[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ContractState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PriceList) > 0 {
		for iNdEx := len(m.PriceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PriceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PairList) > 0 {
		for iNdEx := len(m.PairList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PairList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TriggeredOrdersList) > 0 {
		for iNdEx := len(m.TriggeredOrdersList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TriggeredOrdersList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ShortBookList) > 0 {
		for iNdEx := len(m.ShortBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ShortBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.LongBookList) > 0 {
		for iNdEx := len(m.LongBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LongBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.ContractInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ContractPairPrices) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractPairPrices) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractPairPrices) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size, err := m.PricePair.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ContractState) > 0 {
		for _, e := range m.ContractState {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastEpoch != 0 {
		n += 1 + sovGenesis(uint64(m.LastEpoch))
	}
	return n
}

func (m *ContractState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ContractInfo.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LongBookList) > 0 {
		for _, e := range m.LongBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ShortBookList) > 0 {
		for _, e := range m.ShortBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TriggeredOrdersList) > 0 {
		for _, e := range m.TriggeredOrdersList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PairList) > 0 {
		for _, e := range m.PairList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PriceList) > 0 {
		for _, e := range m.PriceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ContractPairPrices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PricePair.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractState", wireType)
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
			m.ContractState = append(m.ContractState, ContractState{})
			if err := m.ContractState[len(m.ContractState)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEpoch", wireType)
			}
			m.LastEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *ContractState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractInfo", wireType)
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
			if err := m.ContractInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LongBookList", wireType)
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
			m.LongBookList = append(m.LongBookList, LongBook{})
			if err := m.LongBookList[len(m.LongBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortBookList", wireType)
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
			m.ShortBookList = append(m.ShortBookList, ShortBook{})
			if err := m.ShortBookList[len(m.ShortBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TriggeredOrdersList", wireType)
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
			m.TriggeredOrdersList = append(m.TriggeredOrdersList, Order{})
			if err := m.TriggeredOrdersList[len(m.TriggeredOrdersList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairList", wireType)
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
			m.PairList = append(m.PairList, Pair{})
			if err := m.PairList[len(m.PairList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceList", wireType)
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
			m.PriceList = append(m.PriceList, ContractPairPrices{})
			if err := m.PriceList[len(m.PriceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ContractPairPrices) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractPairPrices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractPairPrices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PricePair", wireType)
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
			if err := m.PricePair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
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
			m.Prices = append(m.Prices, &Price{})
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
