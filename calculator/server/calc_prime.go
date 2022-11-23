package main

import (
	"log"

	pb "github.com/jintoples/grpc-go-course/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalcService_PrimeServer) error {
	log.Printf("Calculate prime function was invoked with: %v\n", in)

	var k int32
	k = 2
	n := in.Number

	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			n = n / k
		} else {
			k = k + 1
		}
	}

	return nil
}
