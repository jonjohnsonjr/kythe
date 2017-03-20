// Code generated by protoc-gen-gogo.
// source: kythe/proto/entryset.proto
// DO NOT EDIT!

/*
	Package entryset_proto is a generated protocol buffer package.

	It is generated from these files:
		kythe/proto/entryset.proto

	It has these top-level messages:
		EntrySet
*/
package entryset_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An EntrySet represents a compact collection of entries.
// Compaction is achieved by encoding node names and string-valued data as
// table offsets, and by sorting and prefix-encoding all the strings.
//
// Even without more sophisticated compression this provides a fairly large
// savings over fully-separate entries, even when encoded with protobuf
// wire-format overhead.
//
// The format is somewhat expensive to construct, but not asymptotically bad,
// and decoding is both simpler and less memory-intensive than encoding.
type EntrySet struct {
	// One entry for each unique node named in the entry set, in canonical order.
	// The index of a node in this field is its id.
	Nodes []*EntrySet_Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	// One entry for each node in the entry set. The index of a group in this
	// field matches the id of its corresponding node.
	FactGroups []*EntrySet_FactGroup `protobuf:"bytes,2,rep,name=fact_groups,json=factGroups" json:"fact_groups,omitempty"`
	// One entry for each node in the entry set. The index of a group in this
	// field matches the id of its corresponding source node.
	EdgeGroups []*EntrySet_EdgeGroup `protobuf:"bytes,3,rep,name=edge_groups,json=edgeGroups" json:"edge_groups,omitempty"`
	// A prefix-coded table of all the symbols referenced by the messages above.
	// The entries in this field are lexicographically ordered.  The string table
	// always implicitly contains the empty string as its first entry, but it is
	// not represented explicitly in the message.
	Symbols []*EntrySet_String `protobuf:"bytes,4,rep,name=symbols" json:"symbols,omitempty"`
}

func (m *EntrySet) Reset()                    { *m = EntrySet{} }
func (m *EntrySet) String() string            { return proto.CompactTextString(m) }
func (*EntrySet) ProtoMessage()               {}
func (*EntrySet) Descriptor() ([]byte, []int) { return fileDescriptorEntryset, []int{0} }

func (m *EntrySet) GetNodes() []*EntrySet_Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *EntrySet) GetFactGroups() []*EntrySet_FactGroup {
	if m != nil {
		return m.FactGroups
	}
	return nil
}

func (m *EntrySet) GetEdgeGroups() []*EntrySet_EdgeGroup {
	if m != nil {
		return m.EdgeGroups
	}
	return nil
}

func (m *EntrySet) GetSymbols() []*EntrySet_String {
	if m != nil {
		return m.Symbols
	}
	return nil
}

// The order of these fields reflects the standard ordering for vnames.
// Each of the fields is an index into the symbol table.
type EntrySet_Node struct {
	Corpus    int32 `protobuf:"varint,1,opt,name=corpus,proto3" json:"corpus,omitempty"`
	Language  int32 `protobuf:"varint,2,opt,name=language,proto3" json:"language,omitempty"`
	Path      int32 `protobuf:"varint,3,opt,name=path,proto3" json:"path,omitempty"`
	Root      int32 `protobuf:"varint,4,opt,name=root,proto3" json:"root,omitempty"`
	Signature int32 `protobuf:"varint,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *EntrySet_Node) Reset()                    { *m = EntrySet_Node{} }
func (m *EntrySet_Node) String() string            { return proto.CompactTextString(m) }
func (*EntrySet_Node) ProtoMessage()               {}
func (*EntrySet_Node) Descriptor() ([]byte, []int) { return fileDescriptorEntryset, []int{0, 0} }

func (m *EntrySet_Node) GetCorpus() int32 {
	if m != nil {
		return m.Corpus
	}
	return 0
}

func (m *EntrySet_Node) GetLanguage() int32 {
	if m != nil {
		return m.Language
	}
	return 0
}

func (m *EntrySet_Node) GetPath() int32 {
	if m != nil {
		return m.Path
	}
	return 0
}

func (m *EntrySet_Node) GetRoot() int32 {
	if m != nil {
		return m.Root
	}
	return 0
}

func (m *EntrySet_Node) GetSignature() int32 {
	if m != nil {
		return m.Signature
	}
	return 0
}

type EntrySet_Fact struct {
	Name  int32 `protobuf:"varint,1,opt,name=name,proto3" json:"name,omitempty"`
	Value int32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *EntrySet_Fact) Reset()                    { *m = EntrySet_Fact{} }
func (m *EntrySet_Fact) String() string            { return proto.CompactTextString(m) }
func (*EntrySet_Fact) ProtoMessage()               {}
func (*EntrySet_Fact) Descriptor() ([]byte, []int) { return fileDescriptorEntryset, []int{0, 1} }

func (m *EntrySet_Fact) GetName() int32 {
	if m != nil {
		return m.Name
	}
	return 0
}

func (m *EntrySet_Fact) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type EntrySet_FactGroup struct {
	Facts []*EntrySet_Fact `protobuf:"bytes,1,rep,name=facts" json:"facts,omitempty"`
}

func (m *EntrySet_FactGroup) Reset()                    { *m = EntrySet_FactGroup{} }
func (m *EntrySet_FactGroup) String() string            { return proto.CompactTextString(m) }
func (*EntrySet_FactGroup) ProtoMessage()               {}
func (*EntrySet_FactGroup) Descriptor() ([]byte, []int) { return fileDescriptorEntryset, []int{0, 2} }

func (m *EntrySet_FactGroup) GetFacts() []*EntrySet_Fact {
	if m != nil {
		return m.Facts
	}
	return nil
}

type EntrySet_Edge struct {
	Kind   int32 `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Target int32 `protobuf:"varint,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (m *EntrySet_Edge) Reset()                    { *m = EntrySet_Edge{} }
func (m *EntrySet_Edge) String() string            { return proto.CompactTextString(m) }
func (*EntrySet_Edge) ProtoMessage()               {}
func (*EntrySet_Edge) Descriptor() ([]byte, []int) { return fileDescriptorEntryset, []int{0, 3} }

func (m *EntrySet_Edge) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *EntrySet_Edge) GetTarget() int32 {
	if m != nil {
		return m.Target
	}
	return 0
}

type EntrySet_EdgeGroup struct {
	Edges []*EntrySet_Edge `protobuf:"bytes,1,rep,name=edges" json:"edges,omitempty"`
}

func (m *EntrySet_EdgeGroup) Reset()                    { *m = EntrySet_EdgeGroup{} }
func (m *EntrySet_EdgeGroup) String() string            { return proto.CompactTextString(m) }
func (*EntrySet_EdgeGroup) ProtoMessage()               {}
func (*EntrySet_EdgeGroup) Descriptor() ([]byte, []int) { return fileDescriptorEntryset, []int{0, 4} }

func (m *EntrySet_EdgeGroup) GetEdges() []*EntrySet_Edge {
	if m != nil {
		return m.Edges
	}
	return nil
}

type EntrySet_String struct {
	Prefix int32  `protobuf:"varint,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix []byte `protobuf:"bytes,2,opt,name=suffix,proto3" json:"suffix,omitempty"`
}

func (m *EntrySet_String) Reset()                    { *m = EntrySet_String{} }
func (m *EntrySet_String) String() string            { return proto.CompactTextString(m) }
func (*EntrySet_String) ProtoMessage()               {}
func (*EntrySet_String) Descriptor() ([]byte, []int) { return fileDescriptorEntryset, []int{0, 5} }

func (m *EntrySet_String) GetPrefix() int32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *EntrySet_String) GetSuffix() []byte {
	if m != nil {
		return m.Suffix
	}
	return nil
}

func init() {
	proto.RegisterType((*EntrySet)(nil), "kythe.storage.EntrySet")
	proto.RegisterType((*EntrySet_Node)(nil), "kythe.storage.EntrySet.Node")
	proto.RegisterType((*EntrySet_Fact)(nil), "kythe.storage.EntrySet.Fact")
	proto.RegisterType((*EntrySet_FactGroup)(nil), "kythe.storage.EntrySet.FactGroup")
	proto.RegisterType((*EntrySet_Edge)(nil), "kythe.storage.EntrySet.Edge")
	proto.RegisterType((*EntrySet_EdgeGroup)(nil), "kythe.storage.EntrySet.EdgeGroup")
	proto.RegisterType((*EntrySet_String)(nil), "kythe.storage.EntrySet.String")
}
func (m *EntrySet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntrySet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, msg := range m.Nodes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintEntryset(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.FactGroups) > 0 {
		for _, msg := range m.FactGroups {
			dAtA[i] = 0x12
			i++
			i = encodeVarintEntryset(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.EdgeGroups) > 0 {
		for _, msg := range m.EdgeGroups {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintEntryset(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Symbols) > 0 {
		for _, msg := range m.Symbols {
			dAtA[i] = 0x22
			i++
			i = encodeVarintEntryset(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *EntrySet_Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntrySet_Node) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Corpus != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Corpus))
	}
	if m.Language != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Language))
	}
	if m.Path != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Path))
	}
	if m.Root != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Root))
	}
	if m.Signature != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Signature))
	}
	return i, nil
}

func (m *EntrySet_Fact) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntrySet_Fact) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Name))
	}
	if m.Value != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Value))
	}
	return i, nil
}

func (m *EntrySet_FactGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntrySet_FactGroup) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Facts) > 0 {
		for _, msg := range m.Facts {
			dAtA[i] = 0xa
			i++
			i = encodeVarintEntryset(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *EntrySet_Edge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntrySet_Edge) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Kind))
	}
	if m.Target != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Target))
	}
	return i, nil
}

func (m *EntrySet_EdgeGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntrySet_EdgeGroup) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Edges) > 0 {
		for _, msg := range m.Edges {
			dAtA[i] = 0xa
			i++
			i = encodeVarintEntryset(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *EntrySet_String) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntrySet_String) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Prefix != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(m.Prefix))
	}
	if len(m.Suffix) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEntryset(dAtA, i, uint64(len(m.Suffix)))
		i += copy(dAtA[i:], m.Suffix)
	}
	return i, nil
}

func encodeFixed64Entryset(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Entryset(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEntryset(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EntrySet) Size() (n int) {
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovEntryset(uint64(l))
		}
	}
	if len(m.FactGroups) > 0 {
		for _, e := range m.FactGroups {
			l = e.Size()
			n += 1 + l + sovEntryset(uint64(l))
		}
	}
	if len(m.EdgeGroups) > 0 {
		for _, e := range m.EdgeGroups {
			l = e.Size()
			n += 1 + l + sovEntryset(uint64(l))
		}
	}
	if len(m.Symbols) > 0 {
		for _, e := range m.Symbols {
			l = e.Size()
			n += 1 + l + sovEntryset(uint64(l))
		}
	}
	return n
}

func (m *EntrySet_Node) Size() (n int) {
	var l int
	_ = l
	if m.Corpus != 0 {
		n += 1 + sovEntryset(uint64(m.Corpus))
	}
	if m.Language != 0 {
		n += 1 + sovEntryset(uint64(m.Language))
	}
	if m.Path != 0 {
		n += 1 + sovEntryset(uint64(m.Path))
	}
	if m.Root != 0 {
		n += 1 + sovEntryset(uint64(m.Root))
	}
	if m.Signature != 0 {
		n += 1 + sovEntryset(uint64(m.Signature))
	}
	return n
}

func (m *EntrySet_Fact) Size() (n int) {
	var l int
	_ = l
	if m.Name != 0 {
		n += 1 + sovEntryset(uint64(m.Name))
	}
	if m.Value != 0 {
		n += 1 + sovEntryset(uint64(m.Value))
	}
	return n
}

func (m *EntrySet_FactGroup) Size() (n int) {
	var l int
	_ = l
	if len(m.Facts) > 0 {
		for _, e := range m.Facts {
			l = e.Size()
			n += 1 + l + sovEntryset(uint64(l))
		}
	}
	return n
}

func (m *EntrySet_Edge) Size() (n int) {
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovEntryset(uint64(m.Kind))
	}
	if m.Target != 0 {
		n += 1 + sovEntryset(uint64(m.Target))
	}
	return n
}

func (m *EntrySet_EdgeGroup) Size() (n int) {
	var l int
	_ = l
	if len(m.Edges) > 0 {
		for _, e := range m.Edges {
			l = e.Size()
			n += 1 + l + sovEntryset(uint64(l))
		}
	}
	return n
}

func (m *EntrySet_String) Size() (n int) {
	var l int
	_ = l
	if m.Prefix != 0 {
		n += 1 + sovEntryset(uint64(m.Prefix))
	}
	l = len(m.Suffix)
	if l > 0 {
		n += 1 + l + sovEntryset(uint64(l))
	}
	return n
}

func sovEntryset(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEntryset(x uint64) (n int) {
	return sovEntryset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EntrySet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntryset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EntrySet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntrySet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntryset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &EntrySet_Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FactGroups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntryset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FactGroups = append(m.FactGroups, &EntrySet_FactGroup{})
			if err := m.FactGroups[len(m.FactGroups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdgeGroups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntryset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EdgeGroups = append(m.EdgeGroups, &EntrySet_EdgeGroup{})
			if err := m.EdgeGroups[len(m.EdgeGroups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntryset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbols = append(m.Symbols, &EntrySet_String{})
			if err := m.Symbols[len(m.Symbols)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntryset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntryset
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
func (m *EntrySet_Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntryset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Corpus", wireType)
			}
			m.Corpus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Corpus |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Language", wireType)
			}
			m.Language = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Language |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			m.Path = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Path |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			m.Root = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Root |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			m.Signature = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Signature |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEntryset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntryset
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
func (m *EntrySet_Fact) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntryset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Fact: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fact: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			m.Name = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Name |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEntryset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntryset
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
func (m *EntrySet_FactGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntryset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FactGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FactGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Facts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntryset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Facts = append(m.Facts, &EntrySet_Fact{})
			if err := m.Facts[len(m.Facts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntryset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntryset
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
func (m *EntrySet_Edge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntryset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Edge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Edge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			m.Target = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Target |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEntryset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntryset
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
func (m *EntrySet_EdgeGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntryset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EdgeGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EdgeGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Edges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntryset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Edges = append(m.Edges, &EntrySet_Edge{})
			if err := m.Edges[len(m.Edges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntryset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntryset
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
func (m *EntrySet_String) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntryset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: String: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: String: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			m.Prefix = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Prefix |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suffix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEntryset
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Suffix = append(m.Suffix[:0], dAtA[iNdEx:postIndex]...)
			if m.Suffix == nil {
				m.Suffix = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEntryset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntryset
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
func skipEntryset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntryset
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
					return 0, ErrIntOverflowEntryset
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEntryset
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEntryset
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEntryset
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEntryset(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEntryset = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntryset   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("kythe/proto/entryset.proto", fileDescriptorEntryset) }

var fileDescriptorEntryset = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0x41, 0x6e, 0xda, 0x40,
	0x14, 0xad, 0xc1, 0xa6, 0xf0, 0x69, 0xa5, 0x6a, 0x54, 0x55, 0x96, 0x85, 0xac, 0xb6, 0x2b, 0x56,
	0xa6, 0xa2, 0x1b, 0x76, 0x95, 0x90, 0x68, 0x77, 0x5d, 0x98, 0x03, 0x44, 0x03, 0xfe, 0x1e, 0x2c,
	0xc0, 0x63, 0xcd, 0x8c, 0xa3, 0xb0, 0xcb, 0x31, 0x72, 0x81, 0xdc, 0x25, 0xcb, 0x1c, 0x21, 0x22,
	0x17, 0x89, 0xfe, 0xd8, 0x43, 0x94, 0x05, 0x61, 0xf7, 0xdf, 0xfb, 0xef, 0xfd, 0x27, 0xbf, 0x31,
	0x44, 0xdb, 0x83, 0xd9, 0xe0, 0xa4, 0x52, 0xd2, 0xc8, 0x09, 0x96, 0x46, 0x1d, 0x34, 0x9a, 0xc4,
	0x42, 0xf6, 0xd9, 0xee, 0x12, 0x6d, 0xa4, 0xe2, 0x02, 0x7f, 0xde, 0x07, 0xd0, 0x5f, 0x90, 0x62,
	0x89, 0x86, 0x4d, 0x21, 0x28, 0x65, 0x86, 0x3a, 0xf4, 0xbe, 0x77, 0xc7, 0xc3, 0xe9, 0x28, 0x79,
	0xa3, 0x4d, 0x9c, 0x2e, 0xf9, 0x2f, 0x33, 0x4c, 0x1b, 0x29, 0x9b, 0xc3, 0x30, 0xe7, 0x6b, 0x73,
	0x25, 0x94, 0xac, 0x2b, 0x1d, 0x76, 0xac, 0xf3, 0xc7, 0x39, 0xe7, 0x5f, 0xbe, 0x36, 0xff, 0x48,
	0x99, 0x42, 0xee, 0x46, 0x7b, 0x03, 0x33, 0x81, 0xee, 0x46, 0xf7, 0xfd, 0x1b, 0x8b, 0x4c, 0x60,
	0x7b, 0x03, 0xdd, 0xa8, 0xd9, 0x0c, 0x3e, 0xea, 0xc3, 0x7e, 0x25, 0x77, 0x3a, 0xf4, 0xad, 0x3f,
	0x3e, 0xe7, 0x5f, 0x1a, 0x55, 0x94, 0x22, 0x75, 0xf2, 0xe8, 0xd6, 0x03, 0x9f, 0xbe, 0x88, 0x7d,
	0x83, 0xde, 0x5a, 0xaa, 0xaa, 0xa6, 0xef, 0xf7, 0xc6, 0x41, 0xda, 0x22, 0x16, 0x41, 0x7f, 0xc7,
	0x4b, 0x51, 0x73, 0x81, 0x61, 0xc7, 0x6e, 0x4e, 0x98, 0x31, 0xf0, 0x2b, 0x6e, 0x36, 0x61, 0xd7,
	0xf2, 0x76, 0x26, 0x4e, 0x49, 0x69, 0x42, 0xbf, 0xe1, 0x68, 0x66, 0x23, 0x18, 0xe8, 0x42, 0x94,
	0xdc, 0xd4, 0x0a, 0xc3, 0xc0, 0x2e, 0x5e, 0x89, 0xe8, 0x17, 0xf8, 0xd4, 0x0c, 0x39, 0x4b, 0xbe,
	0xc7, 0x36, 0xdf, 0xce, 0xec, 0x2b, 0x04, 0xd7, 0x7c, 0x57, 0xbb, 0xe8, 0x06, 0x44, 0x7f, 0x60,
	0x70, 0xea, 0x92, 0xde, 0x8d, 0xda, 0xbc, 0xf8, 0x6e, 0xe4, 0x48, 0x1b, 0x69, 0x34, 0x05, 0x9f,
	0x8a, 0xa4, 0xc8, 0x6d, 0x51, 0x66, 0x2e, 0x92, 0x66, 0x2a, 0xc2, 0x70, 0x25, 0xd0, 0xb4, 0x99,
	0x2d, 0xa2, 0xd0, 0x53, 0xf9, 0x14, 0x4a, 0xf5, 0x5f, 0x0c, 0x25, 0x47, 0xda, 0x48, 0xa3, 0x19,
	0xf4, 0x9a, 0xf6, 0x29, 0xa2, 0x52, 0x98, 0x17, 0x37, 0xae, 0xeb, 0x06, 0x11, 0xaf, 0xeb, 0x9c,
	0x78, 0x8a, 0xfe, 0x94, 0xb6, 0x68, 0xfe, 0xe5, 0xe1, 0x18, 0x7b, 0x8f, 0xc7, 0xd8, 0x7b, 0x3a,
	0xc6, 0xde, 0xdd, 0x73, 0xfc, 0x61, 0xd5, 0xb3, 0xff, 0xf3, 0xef, 0x97, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0x1c, 0xf5, 0xdf, 0xed, 0x02, 0x00, 0x00,
}