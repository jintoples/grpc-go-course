package main

import (
	"io"
	"log"

	pb "github.com/jintoples/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalcService_AvgServer) error {
	log.Printf("Average function was invoked")

	var total int32 = 0
	var count int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(total) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)

		total += req.Number
		count++
	}
}
