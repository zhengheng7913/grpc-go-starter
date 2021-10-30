package server

import (
	"errors"
	"fmt"
	"github.com/zhengheng7913/grpc-config/config"
	"google.golang.org/grpc"
	"net"
)

func newServiceRegisterAdapter(srv Service) grpc.ServiceRegistrar {
	return &ServiceRegisterAdapter{
		service: srv,
	}
}

type ServiceGRPC struct {
	server *grpc.Server
	cfg    *config.ServiceConfig
	opt    *OptionGRPC
}

func (g *ServiceGRPC) Register(serviceDesc interface{}, serviceImpl interface{}) {
	desc, ok := serviceDesc.(*grpc.ServiceDesc)
	if !ok {
		fmt.Println(errors.New("service desc type invalid"))
		return
	}
	opts := g.opt.Get()
	g.server = grpc.NewServer(opts...)
	g.server.RegisterService(desc, serviceImpl)
}

func (g *ServiceGRPC) Serve() error {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		return fmt.Errorf("Failed to listen: %v ", err)
	}
	//go func() {
	//	err := g.server.Serve(lis)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//}()
	g.server.Serve(lis)
	return nil
}

func (g *ServiceGRPC) Close(c chan struct{}) error {
	panic("implement me")
}

type ServiceRegisterAdapter struct {
	service Service
}

func (s *ServiceRegisterAdapter) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.service.Register(desc, impl)
}

type OptionGRPC struct {
	serviceName string
	opts        []grpc.ServerOption
}

func (o OptionGRPC) ProtocolName() string {
	return ProtocolNameGrpc
}

func (o *OptionGRPC) Apply(inters ...interface{}) {
	gOpts, ok := assertGrpcOptions(inters)
	if !ok {
		panic("unknown service type")
	}
	o.opts = append(o.opts, gOpts...)
}

func (o *OptionGRPC) Get() []grpc.ServerOption {
	return o.opts
}

func assertGrpcOptions(inters ...interface{}) ([]grpc.ServerOption, bool) {
	opts := make([]grpc.ServerOption, len(inters))
	for _, inter := range inters {
		opt, ok := inter.(grpc.ServerOption)
		if !ok {
			return nil, false
		}
		opts = append(opts, opt)
	}
	return opts, true
}

func dessertGrpcOptions(opts ...grpc.ServerOption) []interface{} {
	inters := make([]interface{}, len(opts))
	for _, opt := range opts {
		inter := opt.(interface{})
		inters = append(inters, inter)
	}
	return inters
}