// Code generated by protoc-gen-gogo.
// source: bazel-out/local-fastbuild/genfiles/mixer/template/sample/quota/go_default_library_tmpl.proto
// DO NOT EDIT!

/*
	Package istio_mixer_adapter_sample_quota is a generated protocol buffer package.

	It is generated from these files:
		bazel-out/local-fastbuild/genfiles/mixer/template/sample/quota/go_default_library_tmpl.proto

	It has these top-level messages:
		Type
		InstanceParam
*/
package istio_mixer_adapter_sample_quota

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "istio.io/api/mixer/v1/template"
import istio_mixer_v1_config_descriptor "istio.io/api/mixer/v1/config/descriptor"

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
	Dimensions map[string]istio_mixer_v1_config_descriptor.ValueType `protobuf:"bytes,1,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=istio.mixer.v1.config.descriptor.ValueType"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorGoDefaultLibraryTmpl, []int{0} }

func (m *Type) GetDimensions() map[string]istio_mixer_v1_config_descriptor.ValueType {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

type InstanceParam struct {
	Dimensions map[string]string `protobuf:"bytes,1,rep,name=dimensions" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BoolMap    map[string]string `protobuf:"bytes,2,rep,name=boolMap" json:"boolMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptorGoDefaultLibraryTmpl, []int{1}
}

func (m *InstanceParam) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *InstanceParam) GetBoolMap() map[string]string {
	if m != nil {
		return m.BoolMap
	}
	return nil
}

func init() {
	proto.RegisterType((*Type)(nil), "istio.mixer.adapter.sample.quota.Type")
	proto.RegisterType((*InstanceParam)(nil), "istio.mixer.adapter.sample.quota.InstanceParam")
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
	if len(this.Dimensions) != len(that1.Dimensions) {
		return false
	}
	for i := range this.Dimensions {
		if this.Dimensions[i] != that1.Dimensions[i] {
			return false
		}
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
	if len(this.Dimensions) != len(that1.Dimensions) {
		return false
	}
	for i := range this.Dimensions {
		if this.Dimensions[i] != that1.Dimensions[i] {
			return false
		}
	}
	if len(this.BoolMap) != len(that1.BoolMap) {
		return false
	}
	for i := range this.BoolMap {
		if this.BoolMap[i] != that1.BoolMap[i] {
			return false
		}
	}
	return true
}
func (this *Type) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&istio_mixer_adapter_sample_quota.Type{")
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]istio_mixer_v1_config_descriptor.ValueType{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%#v: %#v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	if this.Dimensions != nil {
		s = append(s, "Dimensions: "+mapStringForDimensions+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *InstanceParam) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&istio_mixer_adapter_sample_quota.InstanceParam{")
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%#v: %#v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	if this.Dimensions != nil {
		s = append(s, "Dimensions: "+mapStringForDimensions+",\n")
	}
	keysForBoolMap := make([]string, 0, len(this.BoolMap))
	for k, _ := range this.BoolMap {
		keysForBoolMap = append(keysForBoolMap, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForBoolMap)
	mapStringForBoolMap := "map[string]string{"
	for _, k := range keysForBoolMap {
		mapStringForBoolMap += fmt.Sprintf("%#v: %#v,", k, this.BoolMap[k])
	}
	mapStringForBoolMap += "}"
	if this.BoolMap != nil {
		s = append(s, "BoolMap: "+mapStringForBoolMap+",\n")
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
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0xa
			i++
			v := m.Dimensions[k]
			mapSize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + sovGoDefaultLibraryTmpl(uint64(v))
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintGoDefaultLibraryTmpl(dAtA, i, uint64(v))
		}
	}
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
	if len(m.Dimensions) > 0 {
		for k, _ := range m.Dimensions {
			dAtA[i] = 0xa
			i++
			v := m.Dimensions[k]
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
	if len(m.BoolMap) > 0 {
		for k, _ := range m.BoolMap {
			dAtA[i] = 0x12
			i++
			v := m.BoolMap[k]
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
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + sovGoDefaultLibraryTmpl(uint64(v))
			n += mapEntrySize + 1 + sovGoDefaultLibraryTmpl(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	if len(m.Dimensions) > 0 {
		for k, v := range m.Dimensions {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGoDefaultLibraryTmpl(uint64(len(k))) + 1 + len(v) + sovGoDefaultLibraryTmpl(uint64(len(v)))
			n += mapEntrySize + 1 + sovGoDefaultLibraryTmpl(uint64(mapEntrySize))
		}
	}
	if len(m.BoolMap) > 0 {
		for k, v := range m.BoolMap {
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
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]istio_mixer_v1_config_descriptor.ValueType{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	s := strings.Join([]string{`&Type{`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	keysForDimensions := make([]string, 0, len(this.Dimensions))
	for k, _ := range this.Dimensions {
		keysForDimensions = append(keysForDimensions, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForDimensions)
	mapStringForDimensions := "map[string]string{"
	for _, k := range keysForDimensions {
		mapStringForDimensions += fmt.Sprintf("%v: %v,", k, this.Dimensions[k])
	}
	mapStringForDimensions += "}"
	keysForBoolMap := make([]string, 0, len(this.BoolMap))
	for k, _ := range this.BoolMap {
		keysForBoolMap = append(keysForBoolMap, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForBoolMap)
	mapStringForBoolMap := "map[string]string{"
	for _, k := range keysForBoolMap {
		mapStringForBoolMap += fmt.Sprintf("%v: %v,", k, this.BoolMap[k])
	}
	mapStringForBoolMap += "}"
	s := strings.Join([]string{`&InstanceParam{`,
		`Dimensions:` + mapStringForDimensions + `,`,
		`BoolMap:` + mapStringForBoolMap + `,`,
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
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
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]istio_mixer_v1_config_descriptor.ValueType)
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
				var mapvalue istio_mixer_v1_config_descriptor.ValueType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGoDefaultLibraryTmpl
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapvalue |= (istio_mixer_v1_config_descriptor.ValueType(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Dimensions[mapkey] = mapvalue
			} else {
				var mapvalue istio_mixer_v1_config_descriptor.ValueType
				m.Dimensions[mapkey] = mapvalue
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
				return fmt.Errorf("proto: wrong wireType = %d for field Dimensions", wireType)
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
			if m.Dimensions == nil {
				m.Dimensions = make(map[string]string)
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
				m.Dimensions[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Dimensions[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolMap", wireType)
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
			if m.BoolMap == nil {
				m.BoolMap = make(map[string]string)
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
				m.BoolMap[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.BoolMap[mapkey] = mapvalue
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
	proto.RegisterFile("bazel-out/local-fastbuild/genfiles/mixer/template/sample/quota/go_default_library_tmpl.proto", fileDescriptorGoDefaultLibraryTmpl)
}

var fileDescriptorGoDefaultLibraryTmpl = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3d, 0x8b, 0xd4, 0x40,
	0x18, 0xc7, 0x77, 0x72, 0xbe, 0x70, 0xe3, 0x2b, 0xc1, 0x62, 0x49, 0x31, 0x84, 0xad, 0x16, 0x8e,
	0x9b, 0xe1, 0x56, 0x10, 0x39, 0x14, 0xf1, 0xd0, 0xc2, 0x42, 0x90, 0x45, 0xb6, 0x12, 0xc2, 0x24,
	0x79, 0xb2, 0x8c, 0x4e, 0x32, 0xe3, 0xcc, 0x64, 0xb9, 0x58, 0xf9, 0x11, 0x04, 0xbf, 0x84, 0xdf,
	0xc3, 0x46, 0xac, 0x0e, 0x2b, 0x4b, 0x37, 0x5a, 0x58, 0x5e, 0x69, 0x29, 0xc9, 0xe8, 0x5d, 0x4e,
	0xc4, 0xf3, 0xba, 0x40, 0x9e, 0xdf, 0xef, 0x79, 0xfe, 0x7f, 0x06, 0x3f, 0x4b, 0xf9, 0x2b, 0x90,
	0xdb, 0xaa, 0x76, 0x4c, 0xaa, 0x8c, 0xcb, 0xed, 0x82, 0x5b, 0x97, 0xd6, 0x42, 0xe6, 0x6c, 0x09,
	0x55, 0x21, 0x24, 0x58, 0x56, 0x8a, 0x7d, 0x30, 0xcc, 0x41, 0xa9, 0x25, 0x77, 0xc0, 0x2c, 0x2f,
	0xb5, 0x04, 0xf6, 0xb2, 0x56, 0x8e, 0xb3, 0xa5, 0x4a, 0x72, 0x28, 0x78, 0x2d, 0x5d, 0x22, 0x45,
	0x6a, 0xb8, 0x69, 0x12, 0x57, 0x6a, 0x49, 0xb5, 0x51, 0x4e, 0x85, 0xb1, 0xb0, 0x4e, 0x28, 0xda,
	0x0b, 0x28, 0xcf, 0xb9, 0x76, 0x60, 0xa8, 0xe7, 0x69, 0xcf, 0x47, 0x13, 0x2f, 0x5f, 0xed, 0x1c,
	0xfb, 0x61, 0xdf, 0x41, 0x65, 0x85, 0xaa, 0xac, 0xb7, 0x44, 0x5b, 0x47, 0x33, 0x99, 0xaa, 0x0a,
	0xb1, 0x64, 0x39, 0xd8, 0xcc, 0x08, 0xed, 0x94, 0x61, 0x2b, 0x2e, 0x6b, 0x48, 0x5c, 0xa3, 0xc1,
	0x0f, 0x4f, 0x3e, 0x22, 0x7c, 0xee, 0x69, 0xa3, 0x21, 0x5c, 0x60, 0x9c, 0x8b, 0xf2, 0x97, 0x69,
	0x8c, 0xe2, 0x8d, 0xe9, 0xa5, 0xd9, 0x2d, 0x7a, 0xda, 0x41, 0xb4, 0x63, 0xe9, 0x83, 0x23, 0xf0,
	0x61, 0xe5, 0x4c, 0x33, 0x1f, 0x98, 0xa2, 0xe7, 0xf8, 0xda, 0x1f, 0xbf, 0xc3, 0xeb, 0x78, 0xe3,
	0x05, 0x34, 0x63, 0x14, 0xa3, 0xe9, 0xe6, 0xbc, 0xfb, 0x0c, 0xef, 0xe3, 0xf3, 0xfd, 0x65, 0xe3,
	0x20, 0x46, 0xd3, 0xab, 0xb3, 0xad, 0x13, 0x7b, 0x57, 0x3b, 0xd4, 0x07, 0xa1, 0xc7, 0x41, 0xe8,
	0xa2, 0x1b, 0xef, 0x96, 0xcf, 0x3d, 0xb9, 0x1b, 0xdc, 0x46, 0x93, 0xf7, 0x01, 0xbe, 0xf2, 0xa8,
	0xb2, 0x8e, 0x57, 0x19, 0x3c, 0xe1, 0x86, 0x97, 0x61, 0xf2, 0x97, 0x54, 0xf7, 0x4e, 0x4f, 0x75,
	0x42, 0xf2, 0xaf, 0x78, 0xe1, 0x02, 0x5f, 0x4c, 0x95, 0x92, 0x8f, 0xb9, 0x1e, 0x07, 0xbd, 0xfd,
	0xce, 0x59, 0xed, 0x7b, 0x1e, 0xf7, 0xea, 0xdf, 0xb2, 0xe8, 0xee, 0xff, 0xd4, 0x76, 0x63, 0x58,
	0xdb, 0xe6, 0xa0, 0x89, 0x68, 0x17, 0x5f, 0x1e, 0x7a, 0xcf, 0xc2, 0xee, 0xcd, 0x0e, 0xd6, 0x64,
	0xf4, 0x79, 0x4d, 0x46, 0x87, 0x6b, 0x82, 0x5e, 0xb7, 0x04, 0xbd, 0x6b, 0x09, 0xfa, 0xd0, 0x12,
	0x74, 0xd0, 0x12, 0xf4, 0xa5, 0x25, 0xe8, 0x7b, 0x4b, 0x46, 0x87, 0x2d, 0x41, 0x6f, 0xbe, 0x92,
	0xd1, 0x8f, 0x4f, 0xdf, 0xde, 0x06, 0x41, 0x7a, 0xa1, 0x7f, 0x4d, 0x37, 0x7f, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xbb, 0x9f, 0x89, 0xb1, 0x20, 0x03, 0x00, 0x00,
}
