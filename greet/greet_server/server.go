package main

import (
	"github.com/thatInfrastructureGuy/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"fmt"
	"net"
	"log"
	"context"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("GRPC Server Init!")
	lis, err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}