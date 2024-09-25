package main

import (
	"log"
	"net"
	"oda/api/proto/player"
	"oda/internal/playerservice"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Initialize and register the player service
	playerServer := playerservice.NewServer()
	player.RegisterPlayerServiceServer(s, playerServer)

	log.Println("Server starting on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
