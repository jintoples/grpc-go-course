package main

import (
	"context"
	"log"

	pb "github.com/jintoples/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Ikhwal",
		Title:    "New title",
		Content:  "New content update",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happend while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
