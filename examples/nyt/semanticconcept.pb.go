// Code generated by protoc-gen-go.
// source: examples/nyt/semanticconcept.proto
// DO NOT EDIT!

package nyt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SemanticConceptResponse struct {
	Status     string                   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Copyright  string                   `protobuf:"bytes,2,opt,name=copyright" json:"copyright,omitempty"`
	NumResults uint32                   `protobuf:"varint,3,opt,name=num_results" json:"num_results,omitempty"`
	Results    []*SemanticConceptResult `protobuf:"bytes,4,rep,name=results" json:"results,omitempty"`
}

func (m *SemanticConceptResponse) Reset()                    { *m = SemanticConceptResponse{} }
func (m *SemanticConceptResponse) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptResponse) ProtoMessage()               {}
func (*SemanticConceptResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SemanticConceptResponse) GetResults() []*SemanticConceptResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type SemanticConceptResult struct {
	ArticleList *SemanticConceptArticleList `protobuf:"bytes,1,opt,name=article_list" json:"article_list,omitempty"`
}

func (m *SemanticConceptResult) Reset()                    { *m = SemanticConceptResult{} }
func (m *SemanticConceptResult) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptResult) ProtoMessage()               {}
func (*SemanticConceptResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SemanticConceptResult) GetArticleList() *SemanticConceptArticleList {
	if m != nil {
		return m.ArticleList
	}
	return nil
}

type SemanticConceptArticleList struct {
	Results []*SemanticConceptArticle `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	Total   uint32                    `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
}

func (m *SemanticConceptArticleList) Reset()                    { *m = SemanticConceptArticleList{} }
func (m *SemanticConceptArticleList) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptArticleList) ProtoMessage()               {}
func (*SemanticConceptArticleList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SemanticConceptArticleList) GetResults() []*SemanticConceptArticle {
	if m != nil {
		return m.Results
	}
	return nil
}

type SemanticConceptArticle struct {
	Body   string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
	Byline string `protobuf:"bytes,2,opt,name=byline" json:"byline,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Url    string `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
}

func (m *SemanticConceptArticle) Reset()                    { *m = SemanticConceptArticle{} }
func (m *SemanticConceptArticle) String() string            { return proto.CompactTextString(m) }
func (*SemanticConceptArticle) ProtoMessage()               {}
func (*SemanticConceptArticle) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*SemanticConceptResponse)(nil), "nyt.SemanticConceptResponse")
	proto.RegisterType((*SemanticConceptResult)(nil), "nyt.SemanticConceptResult")
	proto.RegisterType((*SemanticConceptArticleList)(nil), "nyt.SemanticConceptArticleList")
	proto.RegisterType((*SemanticConceptArticle)(nil), "nyt.SemanticConceptArticle")
}

var fileDescriptor1 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0xa9, 0x9d, 0x93, 0x7e, 0x5d, 0x05, 0x23, 0x6a, 0x99, 0x07, 0x47, 0x4e, 0x03, 0xa5,
	0x83, 0x89, 0x7f, 0x80, 0x78, 0x15, 0x0f, 0xf3, 0xe4, 0x69, 0xa4, 0xf5, 0x43, 0x03, 0x69, 0x12,
	0x92, 0xaf, 0x60, 0x2f, 0xfe, 0xed, 0xc6, 0xac, 0xe2, 0xc1, 0xd5, 0x63, 0xdf, 0x7b, 0xbc, 0xef,
	0xf5, 0x17, 0xe0, 0xf8, 0x21, 0x5a, 0xab, 0xd0, 0xaf, 0x74, 0x4f, 0x2b, 0x8f, 0xad, 0xd0, 0x24,
	0x9b, 0xc6, 0xe8, 0x06, 0x2d, 0x55, 0xd6, 0x19, 0x32, 0x2c, 0x0d, 0x16, 0xff, 0x84, 0x8b, 0xe7,
	0xc1, 0x7d, 0xd8, 0xb9, 0x1b, 0xf4, 0xd6, 0x68, 0x8f, 0xec, 0x18, 0xa6, 0x9e, 0x04, 0x75, 0xbe,
	0x4c, 0x16, 0xc9, 0x32, 0x63, 0x27, 0x90, 0x35, 0xc6, 0xf6, 0x4e, 0xbe, 0xbd, 0x53, 0x79, 0x10,
	0xa5, 0x53, 0xc8, 0x75, 0xd7, 0x6e, 0x1d, 0xfa, 0x4e, 0x91, 0x2f, 0xd3, 0x20, 0x16, 0xec, 0x1a,
	0x8e, 0x7e, 0x84, 0xc9, 0x22, 0x5d, 0xe6, 0xeb, 0x79, 0x15, 0x2e, 0x55, 0x7f, 0xcf, 0x84, 0x08,
	0x7f, 0x82, 0xb3, 0xbd, 0x06, 0xbb, 0x83, 0x99, 0x70, 0x41, 0x55, 0xb8, 0x55, 0xd2, 0x53, 0xdc,
	0x90, 0xaf, 0xaf, 0xf6, 0x55, 0xdd, 0xef, 0x72, 0x8f, 0x21, 0xc6, 0x5f, 0x60, 0x3e, 0xee, 0xb2,
	0x9b, 0xdf, 0x69, 0x49, 0x9c, 0x76, 0xf9, 0x4f, 0x1f, 0x2b, 0xe0, 0x90, 0x0c, 0x09, 0x15, 0x7f,
	0xb6, 0xe0, 0x1b, 0x38, 0x1f, 0x09, 0xce, 0x60, 0x52, 0x9b, 0xd7, 0x7e, 0xe0, 0x14, 0xb8, 0xd5,
	0xbd, 0x92, 0x1a, 0x07, 0x48, 0xdf, 0x35, 0x92, 0x14, 0x46, 0x3c, 0x19, 0xcb, 0x21, 0xed, 0x9c,
	0x0a, 0x68, 0xc2, 0x47, 0x3d, 0x8d, 0x4f, 0x71, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x16, 0x74,
	0xc9, 0xcc, 0xb0, 0x01, 0x00, 0x00,
}
