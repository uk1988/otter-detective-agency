package main

import (
	"fmt"
	"log"
	"net"
	"oda/internal/interrogationservice"
	"os"

	interrogationpb "oda/api/proto/interrogation"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50054"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	interrogationServer, err := interrogationservice.NewServer(connString)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	grpcServer := grpc.NewServer()
	interrogationpb.RegisterInterrogationServiceServer(grpcServer, interrogationServer)

	log.Printf("Interrogation service listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
