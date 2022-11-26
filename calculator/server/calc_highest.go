package main

import (
	"io"
	"log"

	pb "github.com/jintoples/grpc-go-course/calculator/proto"
)

func (s *Server) Highest(stream pb.CalcService_HighestServer) error {
	log.Printf("Highest function was invoked")

	var max int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if req.Number > max {
			max = req.Number

			err = stream.Send(&pb.HighResponse{
				Result: max,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
