// Code generated by protoc-gen-go.
// source: examples/servers/rpc/service/nyt-proxy.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	examples/servers/rpc/service/nyt-proxy.proto

It has these top-level messages:
	MostPopularRequest
	MostPopularResponse
	CatsRequest
	CatsResponse
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import nyt "github.com/nytimes/gizmo/examples/nyt"
import nyt1 "github.com/nytimes/gizmo/examples/nyt"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MostPopularRequest struct {
	ResourceType   string `protobuf:"bytes,1,opt,name=resourceType" json:"resourceType,omitempty"`
	Section        string `protobuf:"bytes,2,opt,name=section" json:"section,omitempty"`
	TimePeriodDays uint32 `protobuf:"varint,3,opt,name=timePeriodDays" json:"timePeriodDays,omitempty"`
}

func (m *MostPopularRequest) Reset()                    { *m = MostPopularRequest{} }
func (m *MostPopularRequest) String() string            { return proto.CompactTextString(m) }
func (*MostPopularRequest) ProtoMessage()               {}
func (*MostPopularRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MostPopularResponse struct {
	Results []*nyt.MostPopularResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *MostPopularResponse) Reset()                    { *m = MostPopularResponse{} }
func (m *MostPopularResponse) String() string            { return proto.CompactTextString(m) }
func (*MostPopularResponse) ProtoMessage()               {}
func (*MostPopularResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MostPopularResponse) GetResults() []*nyt.MostPopularResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type CatsRequest struct {
}

func (m *CatsRequest) Reset()                    { *m = CatsRequest{} }
func (m *CatsRequest) String() string            { return proto.CompactTextString(m) }
func (*CatsRequest) ProtoMessage()               {}
func (*CatsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CatsResponse struct {
	Results []*nyt1.SemanticConceptArticle `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *CatsResponse) Reset()                    { *m = CatsResponse{} }
func (m *CatsResponse) String() string            { return proto.CompactTextString(m) }
func (*CatsResponse) ProtoMessage()               {}
func (*CatsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CatsResponse) GetResults() []*nyt1.SemanticConceptArticle {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*MostPopularRequest)(nil), "service.MostPopularRequest")
	proto.RegisterType((*MostPopularResponse)(nil), "service.MostPopularResponse")
	proto.RegisterType((*CatsRequest)(nil), "service.CatsRequest")
	proto.RegisterType((*CatsResponse)(nil), "service.CatsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for NYTProxyService service

type NYTProxyServiceClient interface {
	GetMostPopular(ctx context.Context, in *MostPopularRequest, opts ...grpc.CallOption) (*MostPopularResponse, error)
	GetCats(ctx context.Context, in *CatsRequest, opts ...grpc.CallOption) (*CatsResponse, error)
}

type nYTProxyServiceClient struct {
	cc *grpc.ClientConn
}

func NewNYTProxyServiceClient(cc *grpc.ClientConn) NYTProxyServiceClient {
	return &nYTProxyServiceClient{cc}
}

func (c *nYTProxyServiceClient) GetMostPopular(ctx context.Context, in *MostPopularRequest, opts ...grpc.CallOption) (*MostPopularResponse, error) {
	out := new(MostPopularResponse)
	err := grpc.Invoke(ctx, "/service.NYTProxyService/GetMostPopular", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nYTProxyServiceClient) GetCats(ctx context.Context, in *CatsRequest, opts ...grpc.CallOption) (*CatsResponse, error) {
	out := new(CatsResponse)
	err := grpc.Invoke(ctx, "/service.NYTProxyService/GetCats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NYTProxyService service

type NYTProxyServiceServer interface {
	GetMostPopular(context.Context, *MostPopularRequest) (*MostPopularResponse, error)
	GetCats(context.Context, *CatsRequest) (*CatsResponse, error)
}

func RegisterNYTProxyServiceServer(s *grpc.Server, srv NYTProxyServiceServer) {
	s.RegisterService(&_NYTProxyService_serviceDesc, srv)
}

func _NYTProxyService_GetMostPopular_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MostPopularRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NYTProxyServiceServer).GetMostPopular(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _NYTProxyService_GetCats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(NYTProxyServiceServer).GetCats(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _NYTProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.NYTProxyService",
	HandlerType: (*NYTProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMostPopular",
			Handler:    _NYTProxyService_GetMostPopular_Handler,
		},
		{
			MethodName: "GetCats",
			Handler:    _NYTProxyService_GetCats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x66, 0x0e, 0x1c, 0xbe, 0xfd, 0x82, 0x38, 0x47, 0xe9, 0x44, 0xa4, 0x17, 0x77, 0x98, 0x2d,
	0x4c, 0xf0, 0x24, 0x82, 0x4c, 0xf0, 0x20, 0x4a, 0xb1, 0xbb, 0x78, 0xac, 0xf1, 0x1d, 0x0a, 0x6d,
	0x13, 0xf3, 0x52, 0x59, 0xff, 0x11, 0xff, 0x5e, 0xb3, 0x26, 0x8c, 0x75, 0x7a, 0x29, 0xe9, 0xfb,
	0x7e, 0xe4, 0x7b, 0x5f, 0x60, 0x81, 0x9b, 0xb4, 0x90, 0x39, 0x52, 0x44, 0xa8, 0xbe, 0x51, 0x51,
	0xa4, 0x24, 0x6f, 0xce, 0x19, 0xc7, 0xa8, 0xac, 0xf5, 0xb5, 0x54, 0x62, 0x53, 0x87, 0xe6, 0xab,
	0x05, 0xeb, 0x39, 0xc0, 0xbf, 0xd8, 0xc9, 0x0c, 0x25, 0x2a, 0x04, 0x69, 0x29, 0x64, 0x95, 0xa7,
	0xca, 0x12, 0xfd, 0xa0, 0x85, 0x13, 0x16, 0x69, 0xa9, 0x33, 0xce, 0x45, 0xc9, 0x51, 0x6a, 0xcb,
	0x09, 0x12, 0x60, 0x2f, 0x46, 0x18, 0x5b, 0xe1, 0x1b, 0x7e, 0x55, 0x48, 0x9a, 0x4d, 0x60, 0xa0,
	0x90, 0x44, 0xa5, 0x38, 0xae, 0x6b, 0x89, 0x5e, 0xe7, 0xb2, 0x33, 0x3f, 0x61, 0x63, 0x30, 0x57,
	0x73, 0x9d, 0x89, 0xd2, 0x3b, 0x6a, 0x06, 0x53, 0x18, 0xe9, 0xac, 0xc0, 0x18, 0x55, 0x26, 0x3e,
	0x1f, 0xd3, 0x9a, 0xbc, 0xae, 0x99, 0x0f, 0x83, 0x7b, 0x38, 0x6d, 0x99, 0x92, 0x14, 0x25, 0x21,
	0xbb, 0x82, 0x9e, 0x71, 0xad, 0x72, 0x4d, 0xc6, 0xb0, 0x3b, 0xef, 0x2f, 0xa7, 0xa1, 0x09, 0x16,
	0xb6, 0xa9, 0x06, 0x0e, 0x86, 0xd0, 0x5f, 0xa5, 0x9a, 0x5c, 0x9a, 0xe0, 0x0e, 0x06, 0xf6, 0xd7,
	0xf9, 0x2c, 0x0e, 0x7d, 0x66, 0x8d, 0x4f, 0xe2, 0x16, 0x5c, 0xd9, 0x05, 0x1f, 0x94, 0x39, 0xe7,
	0xb8, 0xfc, 0xe9, 0xc0, 0xf8, 0xf5, 0x7d, 0x1d, 0x6f, 0x1b, 0x4c, 0x6c, 0x73, 0xec, 0x19, 0x46,
	0x4f, 0xa8, 0xf7, 0x2e, 0x66, 0xb3, 0xd0, 0xb5, 0x1a, 0xfe, 0xad, 0xc3, 0x3f, 0xff, 0x1f, 0x74,
	0x71, 0x6e, 0xa1, 0x67, 0xcc, 0xb6, 0x09, 0xd9, 0x64, 0x47, 0xdc, 0xcb, 0xef, 0x9f, 0x1d, 0x4c,
	0xad, 0xee, 0xe3, 0xb8, 0x79, 0x81, 0x9b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x2f, 0xc5,
	0xbb, 0xfe, 0x01, 0x00, 0x00,
}
