package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "protobuf-conn/pb"

	"google.golang.org/grpc"
)

type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

// 真正的业务逻辑（你实现的函数）
func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello, " + req.Name,
	}, nil
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// 注册服务
	pb.RegisterHelloServiceServer(grpcServer, &HelloServer{})

	fmt.Println("Server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
