// Code generated by protoc-gen-go.
// source: elwin.proto
// DO NOT EDIT!

/*
Package elwin is a generated protocol buffer package.

It is generated from these files:
	elwin.proto

It has these top-level messages:
	GetRequest
	GetReply
	Requirement
	ExperimentList
	Experiment
	Param
*/
package elwin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Operation int32

const (
	Operation_EXISTS    Operation = 0
	Operation_EQUAL     Operation = 1
	Operation_NOT_EQUAL Operation = 2
	Operation_IN        Operation = 3
	Operation_NOT_IN    Operation = 4
)

var Operation_name = map[int32]string{
	0: "EXISTS",
	1: "EQUAL",
	2: "NOT_EQUAL",
	3: "IN",
	4: "NOT_IN",
}
var Operation_value = map[string]int32{
	"EXISTS":    0,
	"EQUAL":     1,
	"NOT_EQUAL": 2,
	"IN":        3,
	"NOT_IN":    4,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetRequest struct {
	UserID       string         `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Requirements []*Requirement `protobuf:"bytes,2,rep,name=requirements" json:"requirements,omitempty"`
	By           string         `protobuf:"bytes,3,opt,name=by" json:"by,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetRequest) GetRequirements() []*Requirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *GetRequest) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

type GetReply struct {
	Experiments []*Experiment              `protobuf:"bytes,1,rep,name=experiments" json:"experiments,omitempty"`
	Group       map[string]*ExperimentList `protobuf:"bytes,2,rep,name=group" json:"group,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetReply) GetExperiments() []*Experiment {
	if m != nil {
		return m.Experiments
	}
	return nil
}

func (m *GetReply) GetGroup() map[string]*ExperimentList {
	if m != nil {
		return m.Group
	}
	return nil
}

type Requirement struct {
	Key    string    `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Op     Operation `protobuf:"varint,2,opt,name=op,enum=elwin.api.Operation" json:"op,omitempty"`
	Values []string  `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *Requirement) Reset()                    { *m = Requirement{} }
func (m *Requirement) String() string            { return proto.CompactTextString(m) }
func (*Requirement) ProtoMessage()               {}
func (*Requirement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Requirement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Requirement) GetOp() Operation {
	if m != nil {
		return m.Op
	}
	return Operation_EXISTS
}

func (m *Requirement) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type ExperimentList struct {
	Experiments []*Experiment `protobuf:"bytes,1,rep,name=experiments" json:"experiments,omitempty"`
}

func (m *ExperimentList) Reset()                    { *m = ExperimentList{} }
func (m *ExperimentList) String() string            { return proto.CompactTextString(m) }
func (*ExperimentList) ProtoMessage()               {}
func (*ExperimentList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExperimentList) GetExperiments() []*Experiment {
	if m != nil {
		return m.Experiments
	}
	return nil
}

type Experiment struct {
	Name      string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace string            `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Params    []*Param          `protobuf:"bytes,3,rep,name=params" json:"params,omitempty"`
	Labels    map[string]string `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Experiment) Reset()                    { *m = Experiment{} }
func (m *Experiment) String() string            { return proto.CompactTextString(m) }
func (*Experiment) ProtoMessage()               {}
func (*Experiment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Experiment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Experiment) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Experiment) GetParams() []*Param {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Experiment) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type Param struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Param) Reset()                    { *m = Param{} }
func (m *Param) String() string            { return proto.CompactTextString(m) }
func (*Param) ProtoMessage()               {}
func (*Param) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Param) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Param) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "elwin.api.GetRequest")
	proto.RegisterType((*GetReply)(nil), "elwin.api.GetReply")
	proto.RegisterType((*Requirement)(nil), "elwin.api.Requirement")
	proto.RegisterType((*ExperimentList)(nil), "elwin.api.ExperimentList")
	proto.RegisterType((*Experiment)(nil), "elwin.api.Experiment")
	proto.RegisterType((*Param)(nil), "elwin.api.Param")
	proto.RegisterEnum("elwin.api.Operation", Operation_name, Operation_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Elwin service

type ElwinClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
}

type elwinClient struct {
	cc *grpc.ClientConn
}

func NewElwinClient(cc *grpc.ClientConn) ElwinClient {
	return &elwinClient{cc}
}

func (c *elwinClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/elwin.api.Elwin/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Elwin service

type ElwinServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
}

func RegisterElwinServer(s *grpc.Server, srv ElwinServer) {
	s.RegisterService(&_Elwin_serviceDesc, srv)
}

func _Elwin_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElwinServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elwin.api.Elwin/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElwinServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Elwin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elwin.api.Elwin",
	HandlerType: (*ElwinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Elwin_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "elwin.proto",
}

func init() { proto.RegisterFile("elwin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x36, 0x49, 0x1b, 0xcc, 0x8b, 0x96, 0x30, 0x6e, 0x97, 0x58, 0x16, 0xa9, 0x41, 0xa4, 0xec,
	0x21, 0x61, 0xab, 0xa0, 0xdb, 0x9b, 0x62, 0x28, 0x85, 0xd2, 0xd5, 0xe9, 0x0a, 0x22, 0x8a, 0x4c,
	0x97, 0xb1, 0x06, 0xd3, 0xcc, 0x6c, 0x66, 0xb2, 0xda, 0xab, 0x7f, 0xc1, 0xff, 0xe5, 0xc5, 0xbb,
	0x27, 0x7f, 0x88, 0x64, 0x26, 0xb6, 0xb3, 0x92, 0x93, 0xa7, 0xcc, 0x7b, 0xef, 0x9b, 0xef, 0x7d,
	0xef, 0x7d, 0x19, 0xf0, 0x69, 0xfe, 0x25, 0x2b, 0x62, 0x5e, 0x32, 0xc9, 0x90, 0xa7, 0x03, 0xc2,
	0xb3, 0xc1, 0xd1, 0x9a, 0xb1, 0x75, 0x4e, 0x13, 0xc2, 0xb3, 0x84, 0x14, 0x05, 0x93, 0x44, 0x66,
	0xac, 0x10, 0x1a, 0x18, 0x71, 0x80, 0x29, 0x95, 0x98, 0x5e, 0x56, 0x54, 0x48, 0x74, 0x08, 0x6e,
	0x25, 0x68, 0x39, 0x7b, 0x11, 0x5a, 0x43, 0x6b, 0xe4, 0xe1, 0x26, 0x42, 0x13, 0xb8, 0x55, 0xd2,
	0xcb, 0x2a, 0x2b, 0xe9, 0x86, 0x16, 0x52, 0x84, 0xf6, 0xd0, 0x19, 0xf9, 0xe3, 0xc3, 0x78, 0xd7,
	0x25, 0xc6, 0xfb, 0x32, 0xbe, 0x86, 0x45, 0x3d, 0xb0, 0x57, 0xdb, 0xd0, 0x51, 0x7c, 0xf6, 0x6a,
	0x1b, 0xfd, 0xb0, 0xe0, 0xa6, 0x6a, 0xc9, 0xf3, 0x2d, 0x7a, 0x02, 0x3e, 0xfd, 0xca, 0x69, 0x99,
	0x69, 0x5e, 0x4b, 0xf1, 0xf6, 0x0d, 0xde, 0x74, 0x57, 0xc5, 0x26, 0x12, 0x3d, 0x86, 0xee, 0xba,
	0x64, 0x15, 0x6f, 0xa4, 0xdc, 0x33, 0xae, 0xfc, 0x25, 0x8f, 0xa7, 0x35, 0x20, 0x2d, 0x64, 0xb9,
	0xc5, 0x1a, 0x3c, 0x58, 0x02, 0xec, 0x93, 0x28, 0x00, 0xe7, 0x33, 0xdd, 0x36, 0xa3, 0xd6, 0x47,
	0x94, 0x40, 0xf7, 0x8a, 0xe4, 0x15, 0x0d, 0xed, 0xa1, 0x35, 0xf2, 0xc7, 0x77, 0x5b, 0x85, 0xcc,
	0x33, 0x21, 0xb1, 0xc6, 0x4d, 0xec, 0xa7, 0x56, 0xf4, 0x1e, 0x7c, 0x63, 0xfa, 0x16, 0xd6, 0x07,
	0x60, 0x33, 0xae, 0x28, 0x7b, 0xe3, 0x03, 0x83, 0xf2, 0x8c, 0xd3, 0x52, 0x99, 0x81, 0x6d, 0xc6,
	0xeb, 0xdd, 0x2b, 0x4e, 0x11, 0x3a, 0x43, 0xa7, 0xde, 0xbd, 0x8e, 0xa2, 0x19, 0xf4, 0xae, 0xf7,
	0xfe, 0xef, 0xa5, 0x45, 0xbf, 0x2c, 0x80, 0x7d, 0x0d, 0x21, 0xe8, 0x14, 0x64, 0x43, 0x1b, 0xa9,
	0xea, 0x8c, 0x8e, 0xc0, 0xab, 0xbf, 0x82, 0x93, 0x0b, 0xbd, 0x05, 0x0f, 0xef, 0x13, 0x68, 0x04,
	0x2e, 0x27, 0x25, 0xd9, 0x68, 0x8d, 0xfe, 0x38, 0x30, 0x9a, 0xbe, 0xac, 0x0b, 0xb8, 0xa9, 0xa3,
	0x53, 0x70, 0x73, 0xb2, 0xa2, 0xb9, 0x08, 0x3b, 0x0a, 0x79, 0xbf, 0x55, 0x5e, 0x3c, 0x57, 0x18,
	0xed, 0x51, 0x73, 0x61, 0x70, 0x0a, 0xbe, 0x91, 0x6e, 0xd9, 0xe7, 0x81, 0xe9, 0x92, 0x67, 0x5a,
	0x71, 0x02, 0x5d, 0x25, 0xa3, 0x75, 0xb4, 0xd6, 0x6b, 0xc7, 0x29, 0x78, 0x3b, 0x1f, 0x10, 0x80,
	0x9b, 0xbe, 0x99, 0x2d, 0xcf, 0x97, 0xc1, 0x0d, 0xe4, 0x41, 0x37, 0x7d, 0xf5, 0xfa, 0xd9, 0x3c,
	0xb0, 0xd0, 0x6d, 0xf0, 0x16, 0x67, 0xe7, 0x1f, 0x74, 0x68, 0x23, 0x17, 0xec, 0xd9, 0x22, 0x70,
	0x6a, 0x74, 0x9d, 0x9e, 0x2d, 0x82, 0xce, 0xf8, 0x1d, 0x74, 0xd3, 0x7a, 0x40, 0xb4, 0x04, 0x67,
	0x4a, 0x25, 0xea, 0xff, 0xfb, 0x43, 0xaa, 0x07, 0x36, 0xb8, 0xd3, 0xf2, 0x9f, 0x46, 0xc3, 0x6f,
	0x3f, 0x7f, 0x7f, 0xb7, 0x07, 0x51, 0x3f, 0x51, 0xc5, 0xe4, 0xea, 0x24, 0x31, 0x5c, 0x9b, 0x58,
	0xc7, 0xcf, 0x47, 0x6f, 0x1f, 0xae, 0x33, 0xf9, 0xa9, 0x5a, 0xc5, 0x17, 0x6c, 0x93, 0x7c, 0x64,
	0x2c, 0xaf, 0x44, 0xc6, 0x0a, 0x8d, 0x57, 0x0f, 0x59, 0xe8, 0xf3, 0xca, 0x55, 0xd1, 0xa3, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x98, 0x78, 0xd2, 0x0c, 0x0e, 0x04, 0x00, 0x00,
}
