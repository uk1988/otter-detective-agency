package main

import (
	"fmt"
	"log"
	"net"
	"oda/internal/evidenceservice"
	"os"

	evidencepb "oda/api/proto/evidence"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50053"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	evidenceServer, err := evidenceservice.NewServer(connString)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	grpcServer := grpc.NewServer()
	evidencepb.RegisterEvidenceServiceServer(grpcServer, evidenceServer)

	log.Printf("Evidence service listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
