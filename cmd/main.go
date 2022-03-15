package main

import (
	"fmt"
	"log"
	"logic_grpc/internal/config"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c := config.ViperConfigInit()
	srvConfig := c.GetServerConfig()
	listener, err := net.Listen(srvConfig.GrpcProtocol, fmt.Sprintf("%s:%s", srvConfig.Host, srvConfig.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	fmt.Println("Server is running on port:", srvConfig.GrpcPort)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
