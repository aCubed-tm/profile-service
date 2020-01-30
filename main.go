package main

import (
	"log"
	"net"

	pb "github.com/acubed-tm/profile-service/protofiles"
	"google.golang.org/grpc"
)

const port = ":50551"

func main() {
	log.Print("Starting server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProfileServiceServer(s, &server{})
	for {
		if err := s.Serve(lis); err != nil {
			log.Printf("Failed to serve with error: %v", err)
		}
	}
}