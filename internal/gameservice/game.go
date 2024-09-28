package gameservice

import (
	"context"
	"fmt"
	"log"
	databasepb "oda/api/proto/database"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	clearScreen = "\033[2J\033[H"
)

type GameService struct {
	dbClient databasepb.DatabaseServiceClient
	sessions map[string]*GameSession
	mu       sync.Mutex
}

type GameSession struct {
	ID                 string
	Player             *databasepb.Player
	Case               *databasepb.Case
	Conn               *websocket.Conn
	PlayerCaseId       string
	WaitingForSolution bool
	WaitingForNewCase  bool
}

func NewGameService(dbClient databasepb.DatabaseServiceClient) *GameService {
	return &GameService{
		dbClient: dbClient,
		sessions: make(map[string]*GameSession),
	}
}

func (gs *GameService) HandleConnection(conn *websocket.Conn) {
	sessionID := uuid.New().String()
	session := &GameSession{
		ID:                 sessionID,
		Conn:               conn,
		WaitingForSolution: false,
		WaitingForNewCase:  false,
	}

	gs.mu.Lock()
	gs.sessions[sessionID] = session
	gs.mu.Unlock()

	defer func() {
		gs.mu.Lock()
		delete(gs.sessions, sessionID)
		gs.mu.Unlock()
		conn.Close()
	}()

	gs.sendMessage(conn, "Welcome to the Otter Detective Agency! Please enter your name:")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		gs.handleMessage(session, string(message))
	}
}

func (gs *GameService) handleMessage(session *GameSession, message string) {
	if session.Player == nil {
		gs.createPlayer(session, message)
	} else if session.Case == nil {
		gs.assignCase(session)
	} else if session.WaitingForSolution {
		gs.solveCaseAttempt(session, message)
		session.WaitingForSolution = false
	} else if session.WaitingForNewCase {
		gs.handleNewCaseResponse(session, message)
	} else {
		gs.handleGameCommand(session, message)
	}
}

func (gs *GameService) createPlayer(session *GameSession, name string) {
	player, err := gs.dbClient.CreatePlayer(context.Background(), &databasepb.CreatePlayerRequest{Name: name})
	if err != nil {
		log.Printf("Error creating player: %v", err)
		gs.sendErrorMessage(session.Conn, "Error creating player. Please try again.")
		return
	}
	session.Player = player
	gs.sendMessage(session.Conn, "Welcome, "+player.Name+"! Let's get started.")
}

func (gs *GameService) assignCase(session *GameSession) {
	// For now, just getting the first case in the database
	case_, err := gs.dbClient.ListCases(context.Background(), &databasepb.ListCasesRequest{})
	if err != nil || len(case_.Cases) == 0 {
		log.Printf("Error listing cases: %v", err)
		gs.sendMessage(session.Conn, "No cases available. Please try again later.")
		return
	}

	session.Case = case_.Cases[0]

	playerCase, err := gs.dbClient.AssignCaseToPlayer(context.Background(), &databasepb.AssignCaseRequest{
		PlayerId: session.Player.Id,
		CaseId:   session.Case.Id,
	})
	if err != nil {
		log.Printf("Error assigning case: %v", err)
		gs.sendMessage(session.Conn, "Error assigning case. Please try again later.")
		return
	}
	session.PlayerCaseId = playerCase.Id

	gs.sendMessage(session.Conn, fmt.Sprintf("📁 You've been assigned the case: %s\n%s", session.Case.Title, session.Case.Description))
	gs.sendGameOptions(session)
}

func (gs *GameService) handleGameCommand(session *GameSession, command string) {
	switch command {
	case "locations":
		gs.listLocations(session)
	case "suspects":
		gs.listSuspects(session)
	case "solve":
		gs.promptForSolution(session)
	default:
		gs.sendMessage(session.Conn, "Invalid command. Please try again.")
		gs.sendGameOptions(session)
	}
}

func (gs *GameService) listLocations(session *GameSession) {
	locations, err := gs.dbClient.ListLocations(context.Background(), &databasepb.ListLocationsRequest{CaseId: session.Case.Id})
	if err != nil {
		log.Printf("Error listing locations: %v", err)
		gs.sendMessage(session.Conn, "Error listing locations. Please try again later.")
		return
	}

	locationsList := "🏠 Locations:\n"
	for _, location := range locations.Locations {
		locationsList += fmt.Sprintf("- %s\n", location.Name)
	}
	gs.sendMessage(session.Conn, locationsList)
	gs.sendGameOptions(session)
}

func (gs *GameService) listSuspects(session *GameSession) {
	suspects, err := gs.dbClient.ListSuspects(context.Background(), &databasepb.ListSuspectsRequest{CaseId: session.Case.Id})
	if err != nil {
		log.Printf("Error listing suspects: %v", err)
		gs.sendMessage(session.Conn, "Error listing suspects. Please try again later.")
		return
	}

	suspectsList := "👥 Suspects:\n"
	for _, suspect := range suspects.Suspects {
		suspectsList += fmt.Sprintf("- %s:%s\n", suspect.Name, suspect.Description)
	}
	gs.sendMessage(session.Conn, suspectsList)
	gs.sendGameOptions(session)
}

func (gs *GameService) promptForSolution(session *GameSession) {
	gs.sendMessage(session.Conn, "Who do you think committed the crime? Enter the name of the suspect:")
	// Set a flag in the session to indicate that we're waiting for a solution
	session.WaitingForSolution = true
}

func (gs *GameService) solveCaseAttempt(session *GameSession, solution string) {
	result, err := gs.dbClient.SolveCase(context.Background(), &databasepb.SolveCaseRequest{
		PlayerCaseId:     session.PlayerCaseId,
		ProposedSolution: solution,
	})
	if err != nil {
		log.Printf("Error solving case: %v", err)
		gs.sendMessage(session.Conn, "Error solving case. Please try again.")
		return
	}

	if result.IsCorrect {
		_, err := gs.dbClient.UpdatePlayerProgress(context.Background(), &databasepb.UpdatePlayerProgressRequest{Id: session.Player.Id})
		if err != nil {
			log.Printf("Error updating player progress: %v", err)
			gs.sendMessage(session.Conn, "Error updating player progress. Please try again.")
			return
		}
		gs.sendSuccessMessage(session.Conn, "🎉 Congratulations! You've solved the case. 🕵️")
		gs.sendMessage(session.Conn, "Would you like to solve another case? (yes/no)")
		session.WaitingForNewCase = true
	} else {
		gs.sendErrorMessage(session.Conn, "❌ Sorry, that's not the correct solution. Try again! 🔍")
		gs.sendGameOptions(session)
	}
}

func (gs *GameService) handleNewCaseResponse(session *GameSession, response string) {
	session.WaitingForNewCase = false
	switch response {
	case "yes":
		session.Case = nil
		session.PlayerCaseId = ""
		gs.assignCase(session)
	case "no":
		gs.sendMessage(session.Conn, "Thank you for playing! 👋")
		session.Conn.Close()
	default:
		gs.sendErrorMessage(session.Conn, "Invalid response. Please answer 'yes' or 'no'.")
		session.WaitingForNewCase = true
		gs.sendMessage(session.Conn, "Would you like to solve another case? (yes/no)")
	}
}

func (gs *GameService) sendMessage(conn *websocket.Conn, message string) {
	fullMessage := clearScreen + colorGreen + message + colorReset
	if err := conn.WriteMessage(websocket.TextMessage, []byte(fullMessage)); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func (gs *GameService) sendSuccessMessage(conn *websocket.Conn, message string) {
	fullMessage := clearScreen + colorGreen + message + colorReset
	if err := conn.WriteMessage(websocket.TextMessage, []byte(fullMessage)); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func (gs *GameService) sendErrorMessage(conn *websocket.Conn, message string) {
	fullMessage := clearScreen + colorRed + message + colorReset
	if err := conn.WriteMessage(websocket.TextMessage, []byte(fullMessage)); err != nil {
		log.Printf("Error sending error message: %v", err)
	}
}

func (gs *GameService) sendGameOptions(session *GameSession) {
	options := `
🕵️ What would you like to do?
- Type 'locations' to list locations 🏠
- Type 'suspects' to list suspects 👥
- Type 'solve' to attempt to solve the case 🔍
`
	gs.sendMessage(session.Conn, options)
}
