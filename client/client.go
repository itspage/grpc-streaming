package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/itspage/grpc-streaming/hello"
)

const serverAddress = "localhost:8000"

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewHelloServiceClient(conn)

	// Call the streaming API
	req := &pb.HelloRequest{Greeting: "Dave"}
	stream, err := client.LotsOfReplies(context.Background(), req)
	if err != nil {
		log.Fatalf("Error getting replies: %v", err)
	}
	for {
		greeting, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("streaming error: %v", err)
		}
		log.Printf("Greeting: %v", greeting.Reply)
	}
}
