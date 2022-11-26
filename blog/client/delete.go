package main

import (
	"context"
	"log"

	pb "github.com/jintoples/grpc-go-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Printf("Error happend while delete: %v\n", err)
	}

	log.Println("Blog was deleted")
}
