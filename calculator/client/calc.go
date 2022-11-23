package main

import (
	"context"
	"log"

	pb "github.com/jintoples/grpc-go-course/calculator/proto"
)

func doCalc(c pb.CalcServiceClient) {
	log.Printf("doCalc was invoked")
	res, err := c.Sum(context.Background(), &pb.CalcRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("Sum = %d\n", res.Result)
}
