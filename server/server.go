package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/itspage/grpc-streaming/hello"
)

const port = ":8000"

type server struct{}

func (s *server) LotsOfReplies(req *pb.HelloRequest, stream pb.HelloService_LotsOfRepliesServer) error {
	for i := 0; i < 10; i++ {
		stream.Send(&pb.HelloResponse{
			Reply: fmt.Sprintf("Hello (%v) %s", i, req.Greeting),
		})
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	s.Serve(lis)
}
