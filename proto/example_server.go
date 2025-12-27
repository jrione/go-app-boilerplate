package proto

import (
	"context"

	"github.com/jrione/go-app-boilerplate/plugin"
)

type ExampleServer struct {
	UnimplementedExampleServiceServer
	logger *plugin.Logger
}

func NewExampleServer(logger *plugin.Logger) *ExampleServer {
	return &ExampleServer{logger: logger}
}

func (s *ExampleServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	s.logger.Info("gRPC SayHello called with name: ", req.Name)
	message := "Hello, " + req.Name + "!"
	return &HelloResponse{Message: message}, nil
}
