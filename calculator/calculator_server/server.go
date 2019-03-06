package main

import (
	"github.com/thatInfrastructureGuy/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"fmt"
	"net"
	"log"
	"context"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v\n",req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
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
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}