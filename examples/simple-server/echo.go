package main

import (
	"context"
	"errors"

	"github.com/zhengheng7913/grpc-go-starter/examples/simple-server/proto"
)

var (
	ExampleService     = "grpc.one.user_server.UserService"
	ExampleServiceHTTP = "grpc.one.user_server.UserServiceHTTP"
)

type EchoServiceImpl struct {
	proto.UnimplementedEchoServiceServer
}

func (e EchoServiceImpl) Echo(ctx context.Context, request *proto.EchoRequest) (*proto.EchoReply, error) {

	return &proto.EchoReply{
		Message: "hello world",
	}, errors.New("333")
}
