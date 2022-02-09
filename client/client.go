package client

import (
	"context"
	"github.com/zhengheng7913/grpc-go-starter/filter"
	"github.com/zhengheng7913/grpc-go-starter/naming/discovery"
)

var (
	implementMap = make(map[string]func[T](opt ...Option) Client[T])
)

func init() {
	implementMap[GrpcProtocol] = NewGrpcClient
	//implementMap[HttpProtocol] = NewHttpClient
}

func Get(name string) func(opt ...Option) Client {
	return implementMap[name]
}

func WithServiceName(name string) Option {
	return func(opt *Options) {
		opt.ServiceName = name
	}
}

func WithNamespace(namespace string) Option {
	return func(opt *Options) {
		opt.Namespace = namespace
	}
}

func WithDiscovery(d discovery.Discovery) Option {
	return func(opt *Options) {
		opt.Discovery = d
	}
}

func WithFilter(filters []filter.Filter) Option {
	return func(opt *Options) {
		opt.Filters = filters
	}
}

type Options struct {
	Discovery   discovery.Discovery
	ServiceName string
	Namespace   string
	Filters     []filter.Filter
}

type Option func(opt *Options)

type Client[T interface{}] interface {
	Invoke(context context.Context, method string, req interface{}, opts ...Option) (interface{}, error)

	RealClient() T

	Register(realClient interface{}, opts ...Option)
}

func NewClients() *Clients {
	return &Clients{
		clients: make(map[string]Client),
	}
}

type Clients struct {
	clients map[string]Client
}

func (m *Clients) AddClient(name string, client Client) {
	m.clients[name] = client
}

func (m *Clients) Client(name string) Client {
	return m.clients[name]
}
