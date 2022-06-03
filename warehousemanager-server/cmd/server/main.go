package main

import (
	"flag"
	"fmt"
	"github.com/martinkaburu/warehouse-manager/pkg/errors"
	"github.com/martinkaburu/warehouse-manager/warehousemanager-server/internal/server"
	"github.com/martinkaburu/warehouse-manager/warehousemanager-server/internal/service"
	consumer "github.com/martinkaburu/warehouse-manager/warehouseproto"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	var (
		grpcPort = flag.Int("grpc-port", 2000, "Port on which gRPC server should listen TCP conn.")
		httpPort = flag.Int("server-port", 8000, "Port on which the cargo API will be exposed")
	)
	flag.Parse()

	go func() {
		err := server.RunServer(*httpPort)
		if err != nil {
			log.Println(errors.ServerError{Message: "Unable to start server", Code: "500"}, time.Now().String())
		}
	}()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Println(errors.ServerError{Message: "Unable to start gRPC server", Code: err.Error()})
	}
	grpcServer := grpc.NewServer()

	consumer.RegisterOrderConsumerServer(grpcServer, &service.OrderConsumerServer{})

	log.Printf("Initializing gRPC server on port %d", *grpcPort)

	grpcServer.Serve(listen)

}
