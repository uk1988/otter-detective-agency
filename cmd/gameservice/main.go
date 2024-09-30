package main

import (
	"fmt"
	"log"
	"net/http"
	"oda/internal/gameservice"
	"os"

	casepb "oda/api/proto/case"
	deductionpb "oda/api/proto/deduction"
	evidencepb "oda/api/proto/evidence"
	interrogationpb "oda/api/proto/interrogation"
	playerpb "oda/api/proto/player"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections for this prototype
	},
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Set up gRPC connections to each service
	playerConn, err := grpc.Dial("player-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to PlayerService: %v", err)
	}
	defer playerConn.Close()
	playerClient := playerpb.NewPlayerServiceClient(playerConn)

	caseConn, err := grpc.Dial("case-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to CaseService: %v", err)
	}
	defer caseConn.Close()
	caseClient := casepb.NewCaseServiceClient(caseConn)

	evidenceConn, err := grpc.Dial("evidence-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to EvidenceService: %v", err)
	}
	defer evidenceConn.Close()
	evidenceClient := evidencepb.NewEvidenceServiceClient(evidenceConn)

	interrogationConn, err := grpc.Dial("interrogation-service:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to InterrogationService: %v", err)
	}
	defer interrogationConn.Close()
	interrogationClient := interrogationpb.NewInterrogationServiceClient(interrogationConn)

	deductionConn, err := grpc.Dial("deduction-service:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to DeductionService: %v", err)
	}
	defer deductionConn.Close()
	deductionClient := deductionpb.NewDeductionServiceClient(deductionConn)

	gameService := gameservice.NewGameService(
		playerClient,
		caseClient,
		evidenceClient,
		interrogationClient,
		deductionClient,
	)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		gameService.HandleConnection(conn)
	})

	log.Printf("Game service listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
