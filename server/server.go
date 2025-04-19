package main

import (
	"context"
	"fmt"
	pb "nbeans/grpchw/proto/generated/hellopb"
	"net"
	"os"

	"google.golang.org/grpc"
)

type greetServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greetServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Println("unable to listen on 3001; exiting")
		os.Exit(1)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greetServer{})
	fmt.Println("Server listening on 3001")
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve", "error", err)
	}
}
