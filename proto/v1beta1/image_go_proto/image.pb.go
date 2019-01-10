// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1beta1/image.proto

package image_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Instructions from Dockerfile.
type Layer_Directive int32

const (
	// Default value for unsupported/missing directive.
	Layer_DIRECTIVE_UNSPECIFIED Layer_Directive = 0
	// https://docs.docker.com/reference/builder/#maintainer
	Layer_MAINTAINER Layer_Directive = 1
	// https://docs.docker.com/reference/builder/#run
	Layer_RUN Layer_Directive = 2
	// https://docs.docker.com/reference/builder/#cmd
	Layer_CMD Layer_Directive = 3
	// https://docs.docker.com/reference/builder/#label
	Layer_LABEL Layer_Directive = 4
	// https://docs.docker.com/reference/builder/#expose
	Layer_EXPOSE Layer_Directive = 5
	// https://docs.docker.com/reference/builder/#env
	Layer_ENV Layer_Directive = 6
	// https://docs.docker.com/reference/builder/#add
	Layer_ADD Layer_Directive = 7
	// https://docs.docker.com/reference/builder/#copy
	Layer_COPY Layer_Directive = 8
	// https://docs.docker.com/reference/builder/#entrypoint
	Layer_ENTRYPOINT Layer_Directive = 9
	// https://docs.docker.com/reference/builder/#volume
	Layer_VOLUME Layer_Directive = 10
	// https://docs.docker.com/reference/builder/#user
	Layer_USER Layer_Directive = 11
	// https://docs.docker.com/reference/builder/#workdir
	Layer_WORKDIR Layer_Directive = 12
	// https://docs.docker.com/reference/builder/#arg
	Layer_ARG Layer_Directive = 13
	// https://docs.docker.com/reference/builder/#onbuild
	Layer_ONBUILD Layer_Directive = 14
	// https://docs.docker.com/reference/builder/#stopsignal
	Layer_STOPSIGNAL Layer_Directive = 15
	// https://docs.docker.com/reference/builder/#healthcheck
	Layer_HEALTHCHECK Layer_Directive = 16
	// https://docs.docker.com/reference/builder/#shell
	Layer_SHELL Layer_Directive = 17
)

var Layer_Directive_name = map[int32]string{
	0:  "DIRECTIVE_UNSPECIFIED",
	1:  "MAINTAINER",
	2:  "RUN",
	3:  "CMD",
	4:  "LABEL",
	5:  "EXPOSE",
	6:  "ENV",
	7:  "ADD",
	8:  "COPY",
	9:  "ENTRYPOINT",
	10: "VOLUME",
	11: "USER",
	12: "WORKDIR",
	13: "ARG",
	14: "ONBUILD",
	15: "STOPSIGNAL",
	16: "HEALTHCHECK",
	17: "SHELL",
}

var Layer_Directive_value = map[string]int32{
	"DIRECTIVE_UNSPECIFIED": 0,
	"MAINTAINER":            1,
	"RUN":                   2,
	"CMD":                   3,
	"LABEL":                 4,
	"EXPOSE":                5,
	"ENV":                   6,
	"ADD":                   7,
	"COPY":                  8,
	"ENTRYPOINT":            9,
	"VOLUME":                10,
	"USER":                  11,
	"WORKDIR":               12,
	"ARG":                   13,
	"ONBUILD":               14,
	"STOPSIGNAL":            15,
	"HEALTHCHECK":           16,
	"SHELL":                 17,
}

func (x Layer_Directive) String() string {
	return proto.EnumName(Layer_Directive_name, int32(x))
}

func (Layer_Directive) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab0265a3e321e6f0, []int{0, 0}
}

// Layer holds metadata specific to a layer of a Docker image.
type Layer struct {
	// Required. The recovered Dockerfile directive used to construct this layer.
	Directive Layer_Directive `protobuf:"varint,1,opt,name=directive,proto3,enum=grafeas.v1beta1.image.Layer_Directive" json:"directive,omitempty"`
	// The recovered arguments to the Dockerfile directive.
	Arguments            string   `protobuf:"bytes,2,opt,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Layer) Reset()         { *m = Layer{} }
func (m *Layer) String() string { return proto.CompactTextString(m) }
func (*Layer) ProtoMessage()    {}
func (*Layer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0265a3e321e6f0, []int{0}
}

func (m *Layer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Layer.Unmarshal(m, b)
}
func (m *Layer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Layer.Marshal(b, m, deterministic)
}
func (m *Layer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Layer.Merge(m, src)
}
func (m *Layer) XXX_Size() int {
	return xxx_messageInfo_Layer.Size(m)
}
func (m *Layer) XXX_DiscardUnknown() {
	xxx_messageInfo_Layer.DiscardUnknown(m)
}

var xxx_messageInfo_Layer proto.InternalMessageInfo

func (m *Layer) GetDirective() Layer_Directive {
	if m != nil {
		return m.Directive
	}
	return Layer_DIRECTIVE_UNSPECIFIED
}

func (m *Layer) GetArguments() string {
	if m != nil {
		return m.Arguments
	}
	return ""
}

// A set of properties that uniquely identify a given Docker image.
type Fingerprint struct {
	// Required. The layer ID of the final layer in the Docker image's v1
	// representation.
	V1Name string `protobuf:"bytes,1,opt,name=v1_name,json=v1Name,proto3" json:"v1_name,omitempty"`
	// Required. The ordered list of v2 blobs that represent a given image.
	V2Blob []string `protobuf:"bytes,2,rep,name=v2_blob,json=v2Blob,proto3" json:"v2_blob,omitempty"`
	// Output only. The name of the image's v2 blobs computed via:
	//   [bottom] := v2_blob[bottom]
	//   [N] := sha256(v2_blob[N] + " " + v2_name[N+1])
	// Only the name of the final blob is kept.
	V2Name               string   `protobuf:"bytes,3,opt,name=v2_name,json=v2Name,proto3" json:"v2_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fingerprint) Reset()         { *m = Fingerprint{} }
func (m *Fingerprint) String() string { return proto.CompactTextString(m) }
func (*Fingerprint) ProtoMessage()    {}
func (*Fingerprint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0265a3e321e6f0, []int{1}
}

func (m *Fingerprint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fingerprint.Unmarshal(m, b)
}
func (m *Fingerprint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fingerprint.Marshal(b, m, deterministic)
}
func (m *Fingerprint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fingerprint.Merge(m, src)
}
func (m *Fingerprint) XXX_Size() int {
	return xxx_messageInfo_Fingerprint.Size(m)
}
func (m *Fingerprint) XXX_DiscardUnknown() {
	xxx_messageInfo_Fingerprint.DiscardUnknown(m)
}

var xxx_messageInfo_Fingerprint proto.InternalMessageInfo

func (m *Fingerprint) GetV1Name() string {
	if m != nil {
		return m.V1Name
	}
	return ""
}

func (m *Fingerprint) GetV2Blob() []string {
	if m != nil {
		return m.V2Blob
	}
	return nil
}

func (m *Fingerprint) GetV2Name() string {
	if m != nil {
		return m.V2Name
	}
	return ""
}

// Basis describes the base image portion (Note) of the DockerImage
// relationship. Linked occurrences are derived from this or an
// equivalent image via:
//   FROM <Basis.resource_url>
// Or an equivalent reference, e.g. a tag of the resource_url.
type Basis struct {
	// Required. Immutable. The resource_url for the resource representing the
	// basis of associated occurrence images.
	ResourceUrl string `protobuf:"bytes,1,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// Required. Immutable. The fingerprint of the base image.
	Fingerprint          *Fingerprint `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Basis) Reset()         { *m = Basis{} }
func (m *Basis) String() string { return proto.CompactTextString(m) }
func (*Basis) ProtoMessage()    {}
func (*Basis) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0265a3e321e6f0, []int{2}
}

func (m *Basis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Basis.Unmarshal(m, b)
}
func (m *Basis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Basis.Marshal(b, m, deterministic)
}
func (m *Basis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Basis.Merge(m, src)
}
func (m *Basis) XXX_Size() int {
	return xxx_messageInfo_Basis.Size(m)
}
func (m *Basis) XXX_DiscardUnknown() {
	xxx_messageInfo_Basis.DiscardUnknown(m)
}

var xxx_messageInfo_Basis proto.InternalMessageInfo

func (m *Basis) GetResourceUrl() string {
	if m != nil {
		return m.ResourceUrl
	}
	return ""
}

func (m *Basis) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

// Details of an image occurrence.
type Details struct {
	// Required. Immutable. The child image derived from the base image.
	DerivedImage         *Derived `protobuf:"bytes,1,opt,name=derived_image,json=derivedImage,proto3" json:"derived_image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Details) Reset()         { *m = Details{} }
func (m *Details) String() string { return proto.CompactTextString(m) }
func (*Details) ProtoMessage()    {}
func (*Details) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0265a3e321e6f0, []int{3}
}

func (m *Details) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Details.Unmarshal(m, b)
}
func (m *Details) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Details.Marshal(b, m, deterministic)
}
func (m *Details) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Details.Merge(m, src)
}
func (m *Details) XXX_Size() int {
	return xxx_messageInfo_Details.Size(m)
}
func (m *Details) XXX_DiscardUnknown() {
	xxx_messageInfo_Details.DiscardUnknown(m)
}

var xxx_messageInfo_Details proto.InternalMessageInfo

func (m *Details) GetDerivedImage() *Derived {
	if m != nil {
		return m.DerivedImage
	}
	return nil
}

// Derived describes the derived image portion (Occurrence) of the DockerImage
// relationship. This image would be produced from a Dockerfile with FROM
// <DockerImage.Basis in attached Note>.
type Derived struct {
	// Required. The fingerprint of the derived image.
	Fingerprint *Fingerprint `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Output only. The number of layers by which this image differs from the
	// associated image basis.
	Distance int32 `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
	// This contains layer-specific metadata, if populated it has length
	// "distance" and is ordered with [distance] being the layer immediately
	// following the base image and [1] being the final layer.
	LayerInfo []*Layer `protobuf:"bytes,3,rep,name=layer_info,json=layerInfo,proto3" json:"layer_info,omitempty"`
	// Output only. This contains the base image URL for the derived image
	// occurrence.
	BaseResourceUrl      string   `protobuf:"bytes,4,opt,name=base_resource_url,json=baseResourceUrl,proto3" json:"base_resource_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Derived) Reset()         { *m = Derived{} }
func (m *Derived) String() string { return proto.CompactTextString(m) }
func (*Derived) ProtoMessage()    {}
func (*Derived) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab0265a3e321e6f0, []int{4}
}

func (m *Derived) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Derived.Unmarshal(m, b)
}
func (m *Derived) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Derived.Marshal(b, m, deterministic)
}
func (m *Derived) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Derived.Merge(m, src)
}
func (m *Derived) XXX_Size() int {
	return xxx_messageInfo_Derived.Size(m)
}
func (m *Derived) XXX_DiscardUnknown() {
	xxx_messageInfo_Derived.DiscardUnknown(m)
}

var xxx_messageInfo_Derived proto.InternalMessageInfo

func (m *Derived) GetFingerprint() *Fingerprint {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

func (m *Derived) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *Derived) GetLayerInfo() []*Layer {
	if m != nil {
		return m.LayerInfo
	}
	return nil
}

func (m *Derived) GetBaseResourceUrl() string {
	if m != nil {
		return m.BaseResourceUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("grafeas.v1beta1.image.Layer_Directive", Layer_Directive_name, Layer_Directive_value)
	proto.RegisterType((*Layer)(nil), "grafeas.v1beta1.image.Layer")
	proto.RegisterType((*Fingerprint)(nil), "grafeas.v1beta1.image.Fingerprint")
	proto.RegisterType((*Basis)(nil), "grafeas.v1beta1.image.Basis")
	proto.RegisterType((*Details)(nil), "grafeas.v1beta1.image.Details")
	proto.RegisterType((*Derived)(nil), "grafeas.v1beta1.image.Derived")
}

func init() { proto.RegisterFile("proto/v1beta1/image.proto", fileDescriptor_ab0265a3e321e6f0) }

var fileDescriptor_ab0265a3e321e6f0 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0x9b, 0x30,
	0x18, 0xc6, 0x47, 0x68, 0x92, 0xf2, 0xd2, 0x3f, 0xae, 0xa5, 0x6a, 0xe9, 0x54, 0x4d, 0x59, 0x0e,
	0x53, 0xb5, 0x03, 0x55, 0xd8, 0x61, 0x87, 0x9d, 0x08, 0xb8, 0x0d, 0x2a, 0x85, 0xc8, 0x49, 0xba,
	0x76, 0x3b, 0x20, 0x48, 0x9c, 0xcc, 0x52, 0x02, 0x95, 0x21, 0x91, 0xf6, 0x75, 0xf6, 0x99, 0x76,
	0xde, 0x67, 0xd8, 0x47, 0x98, 0xec, 0xd2, 0xa6, 0x9b, 0xda, 0x1d, 0x76, 0xc2, 0x7e, 0x9e, 0xf7,
	0xf7, 0xf0, 0xf2, 0xda, 0xc0, 0xd1, 0xad, 0xc8, 0xcb, 0xfc, 0x74, 0xdd, 0x4d, 0x59, 0x99, 0x74,
	0x4f, 0xf9, 0x32, 0x99, 0x33, 0x4b, 0x69, 0xf8, 0x70, 0x2e, 0x92, 0x19, 0x4b, 0x0a, 0xab, 0x32,
	0x2d, 0x65, 0x76, 0x7e, 0xd6, 0xa0, 0x1e, 0x24, 0xdf, 0x98, 0xc0, 0x1e, 0x18, 0x53, 0x2e, 0xd8,
	0xa4, 0xe4, 0x6b, 0xd6, 0xd2, 0xda, 0xda, 0xc9, 0x9e, 0xfd, 0xd6, 0x7a, 0x12, 0xb2, 0x14, 0x60,
	0x79, 0xf7, 0xd5, 0x74, 0x03, 0xe2, 0x63, 0x30, 0x12, 0x31, 0x5f, 0x2d, 0x59, 0x56, 0x16, 0xad,
	0x5a, 0x5b, 0x3b, 0x31, 0xe8, 0x46, 0xe8, 0xfc, 0xd2, 0xc0, 0x78, 0xc0, 0xf0, 0x11, 0x1c, 0x7a,
	0x3e, 0x25, 0xee, 0xc8, 0xbf, 0x22, 0xf1, 0x38, 0x1c, 0x0e, 0x88, 0xeb, 0x9f, 0xf9, 0xc4, 0x43,
	0x2f, 0xf0, 0x1e, 0xc0, 0xa5, 0xe3, 0x87, 0x23, 0xc7, 0x0f, 0x09, 0x45, 0x1a, 0x6e, 0x82, 0x4e,
	0xc7, 0x21, 0xaa, 0xc9, 0x85, 0x7b, 0xe9, 0x21, 0x1d, 0x1b, 0x50, 0x0f, 0x9c, 0x1e, 0x09, 0xd0,
	0x16, 0x06, 0x68, 0x90, 0xeb, 0x41, 0x34, 0x24, 0xa8, 0x2e, 0x7d, 0x12, 0x5e, 0xa1, 0x86, 0x5c,
	0x38, 0x9e, 0x87, 0x9a, 0x78, 0x1b, 0xb6, 0xdc, 0x68, 0x70, 0x83, 0xb6, 0x65, 0x28, 0x09, 0x47,
	0xf4, 0x66, 0x10, 0xf9, 0xe1, 0x08, 0x19, 0x92, 0xbb, 0x8a, 0x82, 0xf1, 0x25, 0x41, 0x20, 0xab,
	0xc6, 0x43, 0x42, 0x91, 0x89, 0x4d, 0x68, 0x7e, 0x8a, 0xe8, 0x85, 0xe7, 0x53, 0xb4, 0xa3, 0x52,
	0xe8, 0x39, 0xda, 0x95, 0x6a, 0x14, 0xf6, 0xc6, 0x7e, 0xe0, 0xa1, 0x3d, 0x19, 0x34, 0x1c, 0x45,
	0x83, 0xa1, 0x7f, 0x1e, 0x3a, 0x01, 0xda, 0xc7, 0xfb, 0x60, 0xf6, 0x89, 0x13, 0x8c, 0xfa, 0x6e,
	0x9f, 0xb8, 0x17, 0x08, 0xc9, 0xe6, 0x86, 0x7d, 0x12, 0x04, 0xe8, 0xa0, 0x73, 0x0d, 0xe6, 0x19,
	0xcf, 0xe6, 0x4c, 0xdc, 0x0a, 0x9e, 0x95, 0xf8, 0x25, 0x34, 0xd7, 0xdd, 0x38, 0x4b, 0x96, 0x77,
	0x33, 0x36, 0x68, 0x63, 0xdd, 0x0d, 0x93, 0x25, 0x53, 0x86, 0x1d, 0xa7, 0x8b, 0x3c, 0x6d, 0xd5,
	0xda, 0xba, 0x32, 0xec, 0xde, 0x22, 0x4f, 0x2b, 0x43, 0x11, 0x7a, 0x45, 0xd8, 0x92, 0xe8, 0xdc,
	0x42, 0xbd, 0x97, 0x14, 0xbc, 0xc0, 0x6f, 0x60, 0x47, 0xb0, 0x22, 0x5f, 0x89, 0x09, 0x8b, 0x57,
	0x62, 0x51, 0x05, 0x9b, 0xf7, 0xda, 0x58, 0x2c, 0xb0, 0x07, 0xe6, 0x6c, 0xd3, 0x85, 0x3a, 0x18,
	0xd3, 0xee, 0x3c, 0x73, 0xbc, 0x8f, 0xfa, 0xa5, 0x8f, 0xb1, 0x4e, 0x08, 0x4d, 0x8f, 0x95, 0x09,
	0x5f, 0x14, 0xd8, 0x85, 0xdd, 0x29, 0x13, 0x7c, 0xcd, 0xa6, 0xb1, 0x82, 0xd4, 0x4b, 0x4d, 0xfb,
	0xf5, 0x33, 0x91, 0xde, 0x5d, 0x2d, 0xdd, 0xa9, 0x20, 0x5f, 0x5d, 0xbe, 0x1f, 0x9a, 0x0c, 0x54,
	0xc2, 0xdf, 0x1d, 0x6a, 0xff, 0xd5, 0x21, 0x7e, 0x05, 0xdb, 0x53, 0x5e, 0x94, 0x49, 0x36, 0x61,
	0xea, 0x23, 0xeb, 0xf4, 0x61, 0x8f, 0x3f, 0x02, 0x2c, 0xe4, 0xc5, 0x8d, 0x79, 0x36, 0xcb, 0x5b,
	0x7a, 0x5b, 0x3f, 0x31, 0xed, 0xe3, 0x7f, 0xdd, 0x70, 0x6a, 0xa8, 0x7a, 0x3f, 0x9b, 0xe5, 0xf8,
	0x1d, 0x1c, 0xa4, 0x49, 0xc1, 0xe2, 0x3f, 0x06, 0xbd, 0xa5, 0x06, 0xbd, 0x2f, 0x0d, 0xba, 0x19,
	0x76, 0xef, 0x0b, 0xb4, 0x78, 0xfe, 0x74, 0xf0, 0x40, 0xfb, 0xfc, 0x61, 0xce, 0xcb, 0xaf, 0xab,
	0xd4, 0x9a, 0xe4, 0xcb, 0xd3, 0xaa, 0xe6, 0xe1, 0xf9, 0xc4, 0xef, 0x1b, 0xcf, 0xf3, 0x58, 0xc9,
	0xdf, 0x6b, 0xfa, 0x39, 0x75, 0xd2, 0x86, 0xda, 0xbc, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x99, 0x11, 0xd9, 0xeb, 0x03, 0x00, 0x00,
}
