package main

import (
	"context"
	"log"
	"net"

	pb "github.com/morimolymoly/grpc-is-fun/helloworld/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8100"
)

// HelloWorldServer ... implements pb.HelloWorldServiceServer
type HelloWorldServer struct {
}

// SayHello ... implements pb.HelloWorldServiceServer.SayHello
func (s *HelloWorldServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to open server: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(server, &HelloWorldServer{})

	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
