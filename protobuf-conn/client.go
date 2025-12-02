package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "protobuf-conn/pb"

	"google.golang.org/grpc"
)

func main() {
	// 1. 连接到服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	// 2. 带超时的 ctx（客户端必须带 ctx）
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 3. 调用 RPC 函数
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Sun"})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}

	fmt.Println("Server replied:", resp.Message)
}
