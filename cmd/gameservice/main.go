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

	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	r.HandleFunc("/startnewgame", func(w http.ResponseWriter, r *http.Request) {
		handleStartNewGame(w, r, gameService)
	}).Methods("GET")
	r.HandleFunc("/game/{SessionID}", func(w http.ResponseWriter, r *http.Request) {
		handleGameWebSocket(w, r, gameService)
	})

	log.Printf("Game Service listening on port :%s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}

func handleStartNewGame(w http.ResponseWriter, r *http.Request, gs *gameservice.GameService) {
	sessionID := gs.CreateNewSession()
	fmt.Fprintf(w, "Please connect to ws://localhost:8080/game/%s", sessionID)
}

func handleGameWebSocket(w http.ResponseWriter, r *http.Request, gs *gameservice.GameService) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	gs.HandleConnection(conn)
}
