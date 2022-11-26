package main

import (
	"log"

	pb "github.com/jintoples/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	calc := pb.NewCalcServiceClient(conn)

	//doCalc(calc)
	//doCalcPrime(calc)
	//doAvg(calc)
	doCalcHighest(calc)
}
