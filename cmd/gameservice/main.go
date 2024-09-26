package main

import (
	"fmt"
	"log"
	"net/http"
	"oda/internal/gameservice"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}, // Allow all connections by default
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/startnewgame", handleStartNewGame).Methods("GET")
	r.HandleFunc("/game/{sessionID}", handleGameWebSocket)

	fmt.Println("Game service starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleStartNewGame(w http.ResponseWriter, r *http.Request) {
	// Start a new game
	sessionID := gameservice.CreateNewSession()
	fmt.Fprintf(w, "Please connect to ws://localhost:8080/game/%s", sessionID)
}

func handleGameWebSocket(w http.ResponseWriter, r *http.Request) {
	// Handle game websocket
	vars := mux.Vars(r)
	sessionID := vars["sessionID"]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	gameservice.HandleGameWebSocket(sessionID, w, r)
}
