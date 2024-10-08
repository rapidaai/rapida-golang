// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: talk-api.proto

package lexatic_backend

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TalkService_CreateAssistantMessage_FullMethodName      = "/talk_api.TalkService/CreateAssistantMessage"
	TalkService_GetAllAssistantConversation_FullMethodName = "/talk_api.TalkService/GetAllAssistantConversation"
	TalkService_GetAllConversationMessage_FullMethodName   = "/talk_api.TalkService/GetAllConversationMessage"
)

// TalkServiceClient is the client API for TalkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TalkServiceClient interface {
	CreateAssistantMessage(ctx context.Context, in *CreateAssistantMessageRequest, opts ...grpc.CallOption) (TalkService_CreateAssistantMessageClient, error)
	GetAllAssistantConversation(ctx context.Context, in *GetAllAssistantConversationRequest, opts ...grpc.CallOption) (*GetAllAssistantConversationResponse, error)
	GetAllConversationMessage(ctx context.Context, in *GetAllConversationMessageRequest, opts ...grpc.CallOption) (*GetAllConversationMessageResponse, error)
}

type talkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTalkServiceClient(cc grpc.ClientConnInterface) TalkServiceClient {
	return &talkServiceClient{cc}
}

func (c *talkServiceClient) CreateAssistantMessage(ctx context.Context, in *CreateAssistantMessageRequest, opts ...grpc.CallOption) (TalkService_CreateAssistantMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &TalkService_ServiceDesc.Streams[0], TalkService_CreateAssistantMessage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &talkServiceCreateAssistantMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TalkService_CreateAssistantMessageClient interface {
	Recv() (*CreateAssistantMessageResponse, error)
	grpc.ClientStream
}

type talkServiceCreateAssistantMessageClient struct {
	grpc.ClientStream
}

func (x *talkServiceCreateAssistantMessageClient) Recv() (*CreateAssistantMessageResponse, error) {
	m := new(CreateAssistantMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *talkServiceClient) GetAllAssistantConversation(ctx context.Context, in *GetAllAssistantConversationRequest, opts ...grpc.CallOption) (*GetAllAssistantConversationResponse, error) {
	out := new(GetAllAssistantConversationResponse)
	err := c.cc.Invoke(ctx, TalkService_GetAllAssistantConversation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *talkServiceClient) GetAllConversationMessage(ctx context.Context, in *GetAllConversationMessageRequest, opts ...grpc.CallOption) (*GetAllConversationMessageResponse, error) {
	out := new(GetAllConversationMessageResponse)
	err := c.cc.Invoke(ctx, TalkService_GetAllConversationMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TalkServiceServer is the server API for TalkService service.
// All implementations should embed UnimplementedTalkServiceServer
// for forward compatibility
type TalkServiceServer interface {
	CreateAssistantMessage(*CreateAssistantMessageRequest, TalkService_CreateAssistantMessageServer) error
	GetAllAssistantConversation(context.Context, *GetAllAssistantConversationRequest) (*GetAllAssistantConversationResponse, error)
	GetAllConversationMessage(context.Context, *GetAllConversationMessageRequest) (*GetAllConversationMessageResponse, error)
}

// UnimplementedTalkServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTalkServiceServer struct {
}

func (UnimplementedTalkServiceServer) CreateAssistantMessage(*CreateAssistantMessageRequest, TalkService_CreateAssistantMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateAssistantMessage not implemented")
}
func (UnimplementedTalkServiceServer) GetAllAssistantConversation(context.Context, *GetAllAssistantConversationRequest) (*GetAllAssistantConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAssistantConversation not implemented")
}
func (UnimplementedTalkServiceServer) GetAllConversationMessage(context.Context, *GetAllConversationMessageRequest) (*GetAllConversationMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConversationMessage not implemented")
}

// UnsafeTalkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TalkServiceServer will
// result in compilation errors.
type UnsafeTalkServiceServer interface {
	mustEmbedUnimplementedTalkServiceServer()
}

func RegisterTalkServiceServer(s grpc.ServiceRegistrar, srv TalkServiceServer) {
	s.RegisterService(&TalkService_ServiceDesc, srv)
}

func _TalkService_CreateAssistantMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateAssistantMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TalkServiceServer).CreateAssistantMessage(m, &talkServiceCreateAssistantMessageServer{stream})
}

type TalkService_CreateAssistantMessageServer interface {
	Send(*CreateAssistantMessageResponse) error
	grpc.ServerStream
}

type talkServiceCreateAssistantMessageServer struct {
	grpc.ServerStream
}

func (x *talkServiceCreateAssistantMessageServer) Send(m *CreateAssistantMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TalkService_GetAllAssistantConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAssistantConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkServiceServer).GetAllAssistantConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TalkService_GetAllAssistantConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkServiceServer).GetAllAssistantConversation(ctx, req.(*GetAllAssistantConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TalkService_GetAllConversationMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllConversationMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TalkServiceServer).GetAllConversationMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TalkService_GetAllConversationMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TalkServiceServer).GetAllConversationMessage(ctx, req.(*GetAllConversationMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TalkService_ServiceDesc is the grpc.ServiceDesc for TalkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TalkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "talk_api.TalkService",
	HandlerType: (*TalkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllAssistantConversation",
			Handler:    _TalkService_GetAllAssistantConversation_Handler,
		},
		{
			MethodName: "GetAllConversationMessage",
			Handler:    _TalkService_GetAllConversationMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateAssistantMessage",
			Handler:       _TalkService_CreateAssistantMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "talk-api.proto",
}
