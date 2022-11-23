package main

import (
	"context"
	"log"

	pb "github.com/jintoples/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.CalcResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
