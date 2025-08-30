package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/huseyinbabal/microservices-proto/golang/order"
	"github.com/kinetikpod/gomicro/order/config"
	"github.com/kinetikpod/gomicro/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api                            ports.APIPort
	port                           int
	order.UnimplementedOrderServer // forward compatibility
}

// constructor
func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

// menjalankan gRPC server
func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d, error: %v", a.port, err)
	}
}
