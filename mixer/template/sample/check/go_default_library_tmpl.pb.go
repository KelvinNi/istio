// Code generated by protoc-gen-gogo.
// source: bazel-out/local-fastbuild/genfiles/mixer/template/sample/check/go_default_library_tmpl.proto
// DO NOT EDIT!

/*
	Package istio_mixer_adapter_sample_check is a generated protocol buffer package.

	It is generated from these files:
		bazel-out/local-fastbuild/genfiles/mixer/template/sample/check/go_default_library_tmpl.proto

	It has these top-level messages:
		Type
		InstanceParam
*/
package istio_mixer_adapter_sample_check

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "istio.io/api/mixer/v1/template"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Type struct {
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorGoDefaultLibraryTmpl, []int{0} }

type InstanceParam struct {
	CheckExpression string            `protobuf:"bytes,1,opt,name=check_expression,json=checkExpression,proto3" json:"check_expression,omitempty"`
	StringMap       map[string]string `protobuf:"bytes,2,rep,name=stringMap" json:"stringMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptorGoDefaultLibraryTmpl, []int{1}
}

func (m *InstanceParam) GetCheckExpression() string {
	if m != nil {
		return m.CheckExpression
	}
	return ""
}

func (m *InstanceParam) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "istio.mixer.adapter.sample.check.Type")
	proto.RegisterType((*InstanceParam)(nil), "istio.mixer.adapter.sample.check.InstanceParam")
}
func (this *Type) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Type)
	if !ok {
		that2, ok := that.(Type)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *InstanceParam) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*InstanceParam)
	if !ok {
		that2, ok := that.(InstanceParam)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.CheckExpression != that1.CheckExpression {
		return false
	}
	if len(this.StringMap) != len(that1.StringMap) {
		return false
	}
	for i := range this.StringMap {
		if this.StringMap[i] != that1.StringMap[i] {
			return false
		}
	}
	return true
}
func (this *Type) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&istio_mixer_adapter_sample_check.Type{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InstanceParam) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&istio_mixer_adapter_sample_check.InstanceParam{")
	s = append(s, "CheckExpression: "+fmt.Sprintf("%#v", this.CheckExpression)+",\n")
	keysForStringMap := make([]string, 0, len(this.StringMap))
	for k, _ := range this.StringMap {
		keysForStringMap = append(keysForStringMap, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForStringMap)
	mapStringForStringMap := "map[string]string{"
	for _, k := range keysForStringMap {
		mapStringForStringMap += fmt.Sprintf("%#v: %#v,", k, this.StringMap[k])
	}
	mapStringForStringMap += "}"
	if this.StringMap != nil {
		s = append(s, "StringMap: "+mapStringForStringMap+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGoDefaultLibraryTmpl(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CheckExpression) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(m.CheckExpression)))
		i += copy(dAtA[i:], m.CheckExpression)
	}
	if len(m.StringMap) > 0 {
		for k, _ := range m.StringMap {
			dAtA[i] = 0x12
			i++
			v := m.StringMap[k]
			mapSize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + len(v) + sovGoDefaultLibraryTmpl(uint64(len(v)))
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeFixed64GoDefaultLibraryTmpl(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32GoDefaultLibraryTmpl(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGoDefaultLibraryTmpl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Type) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	l = len(m.CheckExpression)
	if l > 0 {
		n += 1 + l + sovGoDefaultLibraryTmpl(uint64(l))
	}
	if len(m.StringMap) > 0 {
		for k, v := range m.StringMap {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + len(v) + sovGoDefaultLibraryTmpl(uint64(len(v)))
			n += mapEntrySize + 1 + sovGoDefaultLibraryTmpl(uint64(mapEntrySize))
		}
	}
	return n
}

func sovGoDefaultLibraryTmpl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGoDefaultLibraryTmpl(x uint64) (n int) {
	return sovGoDefaultLibraryTmpl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Type{`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	keysForStringMap := make([]string, 0, len(this.StringMap))
	for k, _ := range this.StringMap {
		keysForStringMap = append(keysForStringMap, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForStringMap)
	mapStringForStringMap := "map[string]string{"
	for _, k := range keysForStringMap {
		mapStringForStringMap += fmt.Sprintf("%v: %v,", k, this.StringMap[k])
	}
	mapStringForStringMap += "}"
	s := strings.Join([]string{`&InstanceParam{`,
		`CheckExpression:` + fmt.Sprintf("%v", this.CheckExpression) + `,`,
		`StringMap:` + mapStringForStringMap + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGoDefaultLibraryTmpl(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoDefaultLibraryTmpl
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
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGoDefaultLibraryTmpl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
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
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGoDefaultLibraryTmpl
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
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckExpression", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CheckExpression = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
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
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGoDefaultLibraryTmpl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.StringMap == nil {
				m.StringMap = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthGoDefaultLibraryTmpl
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.StringMap[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.StringMap[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGoDefaultLibraryTmpl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGoDefaultLibraryTmpl
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
func skipGoDefaultLibraryTmpl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGoDefaultLibraryTmpl
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
					return 0, ErrIntOverflowGoDefaultLibraryTmpl
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
					return 0, ErrIntOverflowGoDefaultLibraryTmpl
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
				return 0, ErrInvalidLengthGoDefaultLibraryTmpl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGoDefaultLibraryTmpl
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
				next, err := skipGoDefaultLibraryTmpl(dAtA[start:])
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
	ErrInvalidLengthGoDefaultLibraryTmpl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGoDefaultLibraryTmpl   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("bazel-out/local-fastbuild/genfiles/mixer/template/sample/check/go_default_library_tmpl.proto", fileDescriptorGoDefaultLibraryTmpl)
}

var fileDescriptorGoDefaultLibraryTmpl = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0x07, 0xf0, 0x5c, 0xaa, 0x85, 0x9e, 0xa8, 0x25, 0x38, 0x94, 0x0e, 0x47, 0xe9, 0x54, 0x87,
	0x5e, 0xb0, 0x2e, 0x22, 0xe2, 0x20, 0x74, 0x70, 0x10, 0xa4, 0x3a, 0x16, 0xc2, 0x25, 0xfd, 0x5a,
	0x8f, 0x5e, 0x92, 0xe3, 0xee, 0x52, 0x1a, 0x27, 0x1f, 0x41, 0xf0, 0x25, 0x7c, 0x14, 0xc7, 0xe2,
	0x20, 0x8e, 0x36, 0x3a, 0x38, 0x76, 0x74, 0x94, 0x26, 0x6a, 0xe9, 0xe4, 0x76, 0xf7, 0x3f, 0xfe,
	0xbf, 0xfb, 0xf8, 0x70, 0xdf, 0x67, 0xb7, 0x20, 0xda, 0x71, 0x62, 0x5c, 0x11, 0x07, 0x4c, 0xb4,
	0x87, 0x4c, 0x1b, 0x3f, 0xe1, 0x62, 0xe0, 0x8e, 0x20, 0x1a, 0x72, 0x01, 0xda, 0x0d, 0xf9, 0x14,
	0x94, 0x6b, 0x20, 0x94, 0x82, 0x19, 0x70, 0x35, 0x0b, 0xa5, 0x00, 0x37, 0xb8, 0x81, 0x60, 0xec,
	0x8e, 0x62, 0x6f, 0x00, 0x43, 0x96, 0x08, 0xe3, 0x09, 0xee, 0x2b, 0xa6, 0x52, 0xcf, 0x84, 0x52,
	0x50, 0xa9, 0x62, 0x13, 0x3b, 0x0d, 0xae, 0x0d, 0x8f, 0x69, 0x0e, 0x50, 0x36, 0x60, 0xd2, 0x80,
	0xa2, 0x45, 0x9f, 0xe6, 0xfd, 0x7a, 0xb3, 0xc0, 0x27, 0x07, 0x2b, 0x1f, 0xa6, 0x06, 0x22, 0xcd,
	0xe3, 0x48, 0x17, 0x4a, 0xb3, 0x8c, 0x37, 0xae, 0x53, 0x09, 0xcd, 0x17, 0x84, 0xb7, 0xcf, 0x23,
	0x6d, 0x58, 0x14, 0xc0, 0x25, 0x53, 0x2c, 0x74, 0xf6, 0x71, 0x35, 0x67, 0x3c, 0x98, 0x4a, 0x05,
	0x7a, 0x59, 0xaa, 0xa1, 0x06, 0x6a, 0x55, 0x7a, 0xbb, 0x79, 0xde, 0xfd, 0x8b, 0x9d, 0x3e, 0xae,
	0x68, 0xa3, 0x78, 0x34, 0xba, 0x60, 0xb2, 0x66, 0x37, 0x4a, 0xad, 0xad, 0xce, 0x29, 0xfd, 0x6f,
	0x3c, 0xba, 0xf6, 0x1d, 0xbd, 0xfa, 0x05, 0xba, 0x91, 0x51, 0x69, 0x6f, 0x05, 0xd6, 0x4f, 0xf0,
	0xce, 0xfa, 0xa3, 0x53, 0xc5, 0xa5, 0x31, 0xa4, 0x3f, 0xd3, 0x2c, 0x8f, 0xce, 0x1e, 0xde, 0x9c,
	0x30, 0x91, 0x40, 0xcd, 0xce, 0xb3, 0xe2, 0x72, 0x6c, 0x1f, 0xa1, 0xb3, 0xce, 0x6c, 0x4e, 0xac,
	0xd7, 0x39, 0xb1, 0x16, 0x73, 0x82, 0xee, 0x32, 0x82, 0x1e, 0x33, 0x82, 0x9e, 0x32, 0x82, 0x66,
	0x19, 0x41, 0x6f, 0x19, 0x41, 0x9f, 0x19, 0xb1, 0x16, 0x19, 0x41, 0xf7, 0xef, 0xc4, 0xfa, 0x7a,
	0xfe, 0x78, 0xb0, 0x2d, 0xbf, 0x9c, 0xef, 0xe6, 0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x54, 0x1a,
	0xfb, 0x68, 0xc1, 0x01, 0x00, 0x00,
}
