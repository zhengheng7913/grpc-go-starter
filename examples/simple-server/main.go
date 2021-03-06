package main

import (
	starter "github.com/zhengheng7913/grpc-go-starter"
	"github.com/zhengheng7913/grpc-go-starter/examples/simple-server/proto"
	"github.com/zhengheng7913/grpc-go-starter/server"
	_ "github.com/zhengheng7913/grpc-polaris-plugin/registry"
)

func main() {
	s := starter.NewServer()

	proto.RegisterEchoServiceServer(
		server.WithServiceRegisterAdapter(s.Service(ExampleService)),
		&EchoServiceImpl{},
	)

	s.Service(ExampleServiceHTTP).Register(
		server.NewHttpServiceDesc(proto.RegisterEchoServiceHandler),
		nil,
	)
	s.Serve()
}
