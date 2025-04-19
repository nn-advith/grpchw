package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "nbeans/grpchw/proto/generated/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error creating connection; ", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Piccolo"})
	if err != nil {
		fmt.Println("error during rpc call;", err)
	}
	fmt.Println(res.Message)
}
