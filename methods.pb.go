// Code generated by protoc-gen-go. DO NOT EDIT.
// source: club/meeya/ginext/methods.proto

package ginext

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Plugins struct {
	FullName []string `protobuf:"bytes,1,rep,name=fullName" json:"fullName,omitempty" form:"fullName"`
}

func (m *Plugins) Reset()                    { *m = Plugins{} }
func (m *Plugins) String() string            { return proto.CompactTextString(m) }
func (*Plugins) ProtoMessage()               {}
func (*Plugins) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Plugins) GetFullName() []string {
	if m != nil {
		return m.FullName
	}
	return nil
}

func init() {
	proto.RegisterType((*Plugins)(nil), "club.meeya.ginext.Plugins")
}

func init() { proto.RegisterFile("club/meeya/ginext/methods.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xce, 0x29, 0x4d,
	0xd2, 0xcf, 0x4d, 0x4d, 0xad, 0x4c, 0xd4, 0x4f, 0xcf, 0xcc, 0x4b, 0xad, 0x28, 0xd1, 0xcf, 0x4d,
	0x2d, 0xc9, 0xc8, 0x4f, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x04, 0x29, 0xd0,
	0x03, 0x2b, 0xd0, 0x83, 0x28, 0x50, 0x52, 0xe5, 0x62, 0x0f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x2b,
	0x16, 0x92, 0xe2, 0xe2, 0x48, 0x2b, 0xcd, 0xc9, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x0c, 0x82, 0xf3, 0x9d, 0xa4, 0xa2, 0x24, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0xf3, 0xb3, 0x92, 0xb2, 0xd3, 0xf3, 0xa1, 0x76, 0x24, 0xb1, 0x81, 0x0d,
	0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x62, 0xd1, 0x42, 0x3e, 0x7f, 0x00, 0x00, 0x00,
}
