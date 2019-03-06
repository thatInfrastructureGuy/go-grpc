package main

import (
	"github.com/thatInfrastructureGuy/grpc-go-course/calculator/calculatorpb"
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

	c := calculatorpb.NewCalculatorServiceClient(conn)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting Sum Unary RPC")
	req := &calculatorpb.SumRequest{
		FirstNumber: 32,
		SecondNumber: 14,
	}
	res, err := c.Sum(context.Background(),req)
	if err != nil {
		log.Fatalf("Error calling server in Unary Sum RPC: %v", err)
	}
	fmt.Printf("Response from Sum: %v\n",res.SumResult)
}