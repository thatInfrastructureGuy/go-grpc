package main

import (
	"github.com/thatInfrastructureGuy/grpc-go-course/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"context"
)

func main() {
	fmt.Println("Client Init!")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Unary RPC")
	req := &greetpb.GreetRequest {
		Greeting: &greetpb.Greeting{
			FirstName: "Ashish",
			LastName: "Kulkarni",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	fmt.Printf("Response from Greet: %v\n",res)
}