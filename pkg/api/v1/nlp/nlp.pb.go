// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nlp.proto

package nlp

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Corpus struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Corpus) Reset()         { *m = Corpus{} }
func (m *Corpus) String() string { return proto.CompactTextString(m) }
func (*Corpus) ProtoMessage()    {}
func (*Corpus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebd3cd177a18baf, []int{0}
}

func (m *Corpus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Corpus.Unmarshal(m, b)
}
func (m *Corpus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Corpus.Marshal(b, m, deterministic)
}
func (m *Corpus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Corpus.Merge(m, src)
}
func (m *Corpus) XXX_Size() int {
	return xxx_messageInfo_Corpus.Size(m)
}
func (m *Corpus) XXX_DiscardUnknown() {
	xxx_messageInfo_Corpus.DiscardUnknown(m)
}

var xxx_messageInfo_Corpus proto.InternalMessageInfo

func (m *Corpus) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Question struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebd3cd177a18baf, []int{1}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Answer struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebd3cd177a18baf, []int{2}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type AskQuestionRequest struct {
	Corpus               *Corpus   `protobuf:"bytes,1,opt,name=corpus,proto3" json:"corpus,omitempty"`
	Question             *Question `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AskQuestionRequest) Reset()         { *m = AskQuestionRequest{} }
func (m *AskQuestionRequest) String() string { return proto.CompactTextString(m) }
func (*AskQuestionRequest) ProtoMessage()    {}
func (*AskQuestionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebd3cd177a18baf, []int{3}
}

func (m *AskQuestionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskQuestionRequest.Unmarshal(m, b)
}
func (m *AskQuestionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskQuestionRequest.Marshal(b, m, deterministic)
}
func (m *AskQuestionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskQuestionRequest.Merge(m, src)
}
func (m *AskQuestionRequest) XXX_Size() int {
	return xxx_messageInfo_AskQuestionRequest.Size(m)
}
func (m *AskQuestionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AskQuestionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AskQuestionRequest proto.InternalMessageInfo

func (m *AskQuestionRequest) GetCorpus() *Corpus {
	if m != nil {
		return m.Corpus
	}
	return nil
}

func (m *AskQuestionRequest) GetQuestion() *Question {
	if m != nil {
		return m.Question
	}
	return nil
}

type AskQuestionResponse struct {
	Answer               *Answer   `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	Question             *Question `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AskQuestionResponse) Reset()         { *m = AskQuestionResponse{} }
func (m *AskQuestionResponse) String() string { return proto.CompactTextString(m) }
func (*AskQuestionResponse) ProtoMessage()    {}
func (*AskQuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebd3cd177a18baf, []int{4}
}

func (m *AskQuestionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskQuestionResponse.Unmarshal(m, b)
}
func (m *AskQuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskQuestionResponse.Marshal(b, m, deterministic)
}
func (m *AskQuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskQuestionResponse.Merge(m, src)
}
func (m *AskQuestionResponse) XXX_Size() int {
	return xxx_messageInfo_AskQuestionResponse.Size(m)
}
func (m *AskQuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AskQuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AskQuestionResponse proto.InternalMessageInfo

func (m *AskQuestionResponse) GetAnswer() *Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (m *AskQuestionResponse) GetQuestion() *Question {
	if m != nil {
		return m.Question
	}
	return nil
}

type AskQuestionsRequest struct {
	Corpus               *Corpus     `protobuf:"bytes,1,opt,name=corpus,proto3" json:"corpus,omitempty"`
	Questions            []*Question `protobuf:"bytes,2,rep,name=questions,proto3" json:"questions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AskQuestionsRequest) Reset()         { *m = AskQuestionsRequest{} }
func (m *AskQuestionsRequest) String() string { return proto.CompactTextString(m) }
func (*AskQuestionsRequest) ProtoMessage()    {}
func (*AskQuestionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebd3cd177a18baf, []int{5}
}

func (m *AskQuestionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskQuestionsRequest.Unmarshal(m, b)
}
func (m *AskQuestionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskQuestionsRequest.Marshal(b, m, deterministic)
}
func (m *AskQuestionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskQuestionsRequest.Merge(m, src)
}
func (m *AskQuestionsRequest) XXX_Size() int {
	return xxx_messageInfo_AskQuestionsRequest.Size(m)
}
func (m *AskQuestionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AskQuestionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AskQuestionsRequest proto.InternalMessageInfo

func (m *AskQuestionsRequest) GetCorpus() *Corpus {
	if m != nil {
		return m.Corpus
	}
	return nil
}

func (m *AskQuestionsRequest) GetQuestions() []*Question {
	if m != nil {
		return m.Questions
	}
	return nil
}

type AskQuestionsResponse struct {
	Answer               []*Answer `protobuf:"bytes,1,rep,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AskQuestionsResponse) Reset()         { *m = AskQuestionsResponse{} }
func (m *AskQuestionsResponse) String() string { return proto.CompactTextString(m) }
func (*AskQuestionsResponse) ProtoMessage()    {}
func (*AskQuestionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ebd3cd177a18baf, []int{6}
}

func (m *AskQuestionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AskQuestionsResponse.Unmarshal(m, b)
}
func (m *AskQuestionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AskQuestionsResponse.Marshal(b, m, deterministic)
}
func (m *AskQuestionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AskQuestionsResponse.Merge(m, src)
}
func (m *AskQuestionsResponse) XXX_Size() int {
	return xxx_messageInfo_AskQuestionsResponse.Size(m)
}
func (m *AskQuestionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AskQuestionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AskQuestionsResponse proto.InternalMessageInfo

func (m *AskQuestionsResponse) GetAnswer() []*Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

func init() {
	proto.RegisterType((*Corpus)(nil), "nlp.Corpus")
	proto.RegisterType((*Question)(nil), "nlp.Question")
	proto.RegisterType((*Answer)(nil), "nlp.Answer")
	proto.RegisterType((*AskQuestionRequest)(nil), "nlp.AskQuestionRequest")
	proto.RegisterType((*AskQuestionResponse)(nil), "nlp.AskQuestionResponse")
	proto.RegisterType((*AskQuestionsRequest)(nil), "nlp.AskQuestionsRequest")
	proto.RegisterType((*AskQuestionsResponse)(nil), "nlp.AskQuestionsResponse")
}

func init() { proto.RegisterFile("nlp.proto", fileDescriptor_6ebd3cd177a18baf) }

var fileDescriptor_6ebd3cd177a18baf = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xcb, 0x29, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0xcb, 0x29, 0x50, 0x92, 0xe1, 0x62, 0x73, 0xce,
	0x2f, 0x2a, 0x28, 0x2d, 0x16, 0x12, 0xe2, 0x62, 0x49, 0xca, 0x4f, 0xa9, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xe4, 0xb8, 0x38, 0x02, 0x4b, 0x53, 0x8b, 0x4b, 0x32, 0xf3,
	0xf3, 0xb0, 0xca, 0xcb, 0x70, 0xb1, 0x39, 0xe6, 0x15, 0x97, 0xa7, 0x16, 0x61, 0x95, 0x4d, 0xe1,
	0x12, 0x72, 0x2c, 0xce, 0x86, 0x19, 0x10, 0x94, 0x5a, 0x08, 0x62, 0x09, 0x29, 0x73, 0xb1, 0x25,
	0x83, 0x6d, 0x04, 0xab, 0xe5, 0x36, 0xe2, 0xd6, 0x03, 0x39, 0x09, 0xe2, 0x88, 0x20, 0xa8, 0x94,
	0x90, 0x26, 0x17, 0x47, 0x21, 0x54, 0x9f, 0x04, 0x13, 0x58, 0x19, 0x2f, 0x58, 0x19, 0xdc, 0x30,
	0xb8, 0xb4, 0x52, 0x2a, 0x97, 0x30, 0x8a, 0x2d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x20, 0x6b,
	0x12, 0xc1, 0x4e, 0x43, 0xb1, 0x06, 0xe2, 0xda, 0x20, 0xa8, 0x14, 0x29, 0xd6, 0xa4, 0xa3, 0x58,
	0x53, 0x4c, 0x92, 0x6f, 0xb4, 0xb9, 0x38, 0x61, 0xe6, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x63, 0xda,
	0x83, 0x90, 0x57, 0xb2, 0xe6, 0x12, 0x41, 0xb5, 0x08, 0x8b, 0x87, 0x98, 0x71, 0x78, 0xc8, 0x68,
	0x32, 0x23, 0x17, 0x97, 0x9f, 0x4f, 0x40, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x03,
	0x17, 0x37, 0x92, 0x59, 0x42, 0xe2, 0x10, 0x2d, 0x18, 0x71, 0x22, 0x25, 0x81, 0x29, 0x01, 0xb5,
	0xd5, 0x99, 0x8b, 0x07, 0xd9, 0x35, 0x42, 0x18, 0x2a, 0x61, 0x21, 0x21, 0x25, 0x89, 0x45, 0x06,
	0x62, 0x48, 0x12, 0x1b, 0x38, 0xc1, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xd0, 0x6b,
	0x7c, 0x7d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NLPServiceClient is the client API for NLPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NLPServiceClient interface {
	AskQuestion(ctx context.Context, in *AskQuestionRequest, opts ...grpc.CallOption) (*AskQuestionResponse, error)
	AskQuestions(ctx context.Context, in *AskQuestionsRequest, opts ...grpc.CallOption) (*AskQuestionsResponse, error)
}

type nLPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNLPServiceClient(cc grpc.ClientConnInterface) NLPServiceClient {
	return &nLPServiceClient{cc}
}

func (c *nLPServiceClient) AskQuestion(ctx context.Context, in *AskQuestionRequest, opts ...grpc.CallOption) (*AskQuestionResponse, error) {
	out := new(AskQuestionResponse)
	err := c.cc.Invoke(ctx, "/nlp.NLPService/AskQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nLPServiceClient) AskQuestions(ctx context.Context, in *AskQuestionsRequest, opts ...grpc.CallOption) (*AskQuestionsResponse, error) {
	out := new(AskQuestionsResponse)
	err := c.cc.Invoke(ctx, "/nlp.NLPService/AskQuestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NLPServiceServer is the server API for NLPService service.
type NLPServiceServer interface {
	AskQuestion(context.Context, *AskQuestionRequest) (*AskQuestionResponse, error)
	AskQuestions(context.Context, *AskQuestionsRequest) (*AskQuestionsResponse, error)
}

// UnimplementedNLPServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNLPServiceServer struct {
}

func (*UnimplementedNLPServiceServer) AskQuestion(ctx context.Context, req *AskQuestionRequest) (*AskQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskQuestion not implemented")
}
func (*UnimplementedNLPServiceServer) AskQuestions(ctx context.Context, req *AskQuestionsRequest) (*AskQuestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskQuestions not implemented")
}

func RegisterNLPServiceServer(s *grpc.Server, srv NLPServiceServer) {
	s.RegisterService(&_NLPService_serviceDesc, srv)
}

func _NLPService_AskQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NLPServiceServer).AskQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlp.NLPService/AskQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NLPServiceServer).AskQuestion(ctx, req.(*AskQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NLPService_AskQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AskQuestionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NLPServiceServer).AskQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nlp.NLPService/AskQuestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NLPServiceServer).AskQuestions(ctx, req.(*AskQuestionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NLPService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nlp.NLPService",
	HandlerType: (*NLPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskQuestion",
			Handler:    _NLPService_AskQuestion_Handler,
		},
		{
			MethodName: "AskQuestions",
			Handler:    _NLPService_AskQuestions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nlp.proto",
}