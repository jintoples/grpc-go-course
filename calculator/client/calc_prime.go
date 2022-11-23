package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jintoples/grpc-go-course/calculator/proto"
)

func doCalcPrime(c pb.CalcServiceClient) {
	log.Println("doCalcPrime was invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Prime: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Prime: %d\n", msg.Result)
	}
}
