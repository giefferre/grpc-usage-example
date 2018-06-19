package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterDemoRPCServiceServer(grpcServer, newDemoRPCServer())

	grpcServer.Serve(listener)
}
