package gameservice

import (
	"context"
	"fmt"
	"log"
	casepb "oda/api/proto/case"
	deductionpb "oda/api/proto/deduction"
	evidencepb "oda/api/proto/evidence"
	interrogationpb "oda/api/proto/interrogation"
	playerpb "oda/api/proto/player"
	"strconv"
	"strings"
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
	playerClient        playerpb.PlayerServiceClient
	caseClient          casepb.CaseServiceClient
	evidenceClient      evidencepb.EvidenceServiceClient
	interrogationClient interrogationpb.InterrogationServiceClient
	deductionClient     deductionpb.DeductionServiceClient
	sessions            map[string]*GameSession
	mu                  sync.Mutex
}

type GameSession struct {
	ID                 string
	Player             *playerpb.Player
	Case               *casepb.Case
	Conn               *websocket.Conn
	CurrentLocation    string
	CurrentSuspect     string
	WaitingForSolution bool
	WaitingForNewCase  bool
}

func NewGameService(
	playerClient playerpb.PlayerServiceClient,
	caseClient casepb.CaseServiceClient,
	evidenceClient evidencepb.EvidenceServiceClient,
	interrogationClient interrogationpb.InterrogationServiceClient,
	deductionClient deductionpb.DeductionServiceClient,
) *GameService {
	return &GameService{
		playerClient:        playerClient,
		caseClient:          caseClient,
		evidenceClient:      evidenceClient,
		interrogationClient: interrogationClient,
		deductionClient:     deductionClient,
		sessions:            make(map[string]*GameSession),
	}
}

func (gs *GameService) CreateNewSession() string {
	sessionID := uuid.New().String()
	session := &GameSession{
		ID:                 sessionID,
		WaitingForSolution: false,
		WaitingForNewCase:  false,
	}

	gs.mu.Lock()
	gs.sessions[sessionID] = session
	gs.mu.Unlock()

	return sessionID
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

	gs.sendMessage(conn, colorGreen+"🔎 Welcome to the Otter Detective Agency! Please enter your name:")

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
	player, err := gs.playerClient.CreatePlayer(context.Background(), &playerpb.CreatePlayerRequest{Name: name})
	if err != nil {
		log.Printf("Error creating player: %v", err)
		gs.sendErrorMessage(session.Conn, "Error creating player. Please try again.")
		return
	}
	session.Player = player
	gs.sendSuccessMessage(session.Conn, fmt.Sprintf("Greetings Detective %s! 🕵️ Let's get started", player.Name))
	gs.assignCase(session)
}

func (gs *GameService) assignCase(session *GameSession) {
	caseList, err := gs.caseClient.ListCases(context.Background(), &casepb.ListCasesRequest{})
	if err != nil || len(caseList.Cases) == 0 {
		log.Printf("Error listing cases: %v", err)
		gs.sendErrorMessage(session.Conn, "No cases available. Please try again later.")
		return
	}

	// For simplicity, we're just getting the first case. In a real game, you'd want to select an unsolved case.
	session.Case = caseList.Cases[0]

	_, err = gs.caseClient.AssignCaseToPlayer(context.Background(), &casepb.AssignCaseRequest{
		PlayerId: session.Player.Id,
		CaseId:   session.Case.Id,
	})
	if err != nil {
		log.Printf("Error assigning case: %v", err)
		gs.sendErrorMessage(session.Conn, "Error assigning case. Please try again later.")
		return
	}

	gs.sendMessage(session.Conn, fmt.Sprintf("📁 You've been assigned the case: %s\n%s", session.Case.Title, session.Case.Description))
	gs.sendGameOptions(session)
}

func (gs *GameService) handleGameCommand(session *GameSession, command string) {
	switch {
	case command == "locations":
		gs.listLocations(session)
	case command == "suspects":
		gs.listSuspects(session)
	case command == "solve":
		gs.promptForSolution(session)
	case strings.HasPrefix(command, "examine "):
		location := strings.TrimPrefix(command, "examine ")
		gs.examineLocation(session, location)
	case strings.HasPrefix(command, "investigate "):
		evidence := strings.TrimPrefix(command, "investigate ")
		gs.investigateEvidence(session, evidence)
	case strings.HasPrefix(command, "interrogate "):
		suspect := strings.TrimPrefix(command, "interrogate ")
		gs.interrogateSuspect(session, suspect)
	case strings.HasPrefix(command, "ask "):
		question := strings.TrimPrefix(command, "ask ")
		gs.askQuestion(session, question)
	default:
		gs.sendMessage(session.Conn, "Invalid command. Please try again.")
		gs.sendGameOptions(session)
	}
}

func (gs *GameService) listLocations(session *GameSession) {
	locations, err := gs.evidenceClient.ListLocations(context.Background(), &evidencepb.ListLocationsRequest{CaseId: session.Case.Id})
	if err != nil {
		log.Printf("Error listing locations: %v", err)
		gs.sendErrorMessage(session.Conn, "Error listing locations. Please try again later.")
		return
	}

	locationsList := "🏠 Locations:\n"
	for _, location := range locations.Locations {
		locationsList += fmt.Sprintf("- %s\n", location.Name)
	}
	gs.sendMessage(session.Conn, locationsList)
	gs.sendMessage(session.Conn, "Type 'examine <location>' to examine a location.")
}

func (gs *GameService) examineLocation(session *GameSession, location string) {
	session.CurrentLocation = location
	evidenceList, err := gs.evidenceClient.ListEvidence(context.Background(), &evidencepb.ListEvidenceRequest{
		CaseId:   session.Case.Id,
		Location: location,
	})
	if err != nil {
		log.Printf("Error listing evidence: %v", err)
		gs.sendErrorMessage(session.Conn, "Error examining location. Please try again.")
		return
	}

	evidenceMsg := fmt.Sprintf("🔍 Evidence at %s:\n", location)
	for _, e := range evidenceList.Evidence {
		evidenceMsg += fmt.Sprintf("- %s\n", e.Name)
	}
	gs.sendMessage(session.Conn, evidenceMsg)
	gs.sendMessage(session.Conn, "Type 'investigate [evidence name]' to examine a piece of evidence.")
}

func (gs *GameService) investigateEvidence(session *GameSession, evidenceName string) {
	evidenceList, err := gs.evidenceClient.ListEvidence(context.Background(), &evidencepb.ListEvidenceRequest{
		CaseId:   session.Case.Id,
		Location: session.CurrentLocation,
	})
	if err != nil {
		log.Printf("Error listing evidence: %v", err)
		gs.sendErrorMessage(session.Conn, "Error investigating evidence. Please try again.")
		return
	}

	var foundEvidence *evidencepb.Evidence
	for _, e := range evidenceList.Evidence {
		if e.Name == evidenceName {
			foundEvidence = e
			break
		}
	}

	if foundEvidence == nil {
		gs.sendErrorMessage(session.Conn, "Evidence not found. Please try again.")
		return
	}

	gs.sendMessage(session.Conn, fmt.Sprintf("🔍 %s: %s", foundEvidence.Name, foundEvidence.Description))
	gs.sendGameOptions(session)
}

func (gs *GameService) listSuspects(session *GameSession) {
	suspects, err := gs.interrogationClient.ListSuspects(context.Background(), &interrogationpb.ListSuspectsRequest{CaseId: session.Case.Id})
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
	gs.sendMessage(session.Conn, "Type 'interrogate [suspect name]' to interrogate a suspect.")
}

func (gs *GameService) interrogateSuspect(session *GameSession, suspectName string) {
	suspects, err := gs.interrogationClient.ListSuspects(context.Background(), &interrogationpb.ListSuspectsRequest{
		CaseId: session.Case.Id,
	})
	if err != nil {
		log.Printf("Error listing suspects: %v", err)
		gs.sendErrorMessage(session.Conn, "Error preparing interrogation. Please try again.")
		return
	}

	var suspectId string
	for _, s := range suspects.Suspects {
		if s.Name == suspectName {
			suspectId = s.Id
			break
		}
	}

	if suspectId == "" {
		gs.sendErrorMessage(session.Conn, "Suspect not found. Please try again.")
		return
	}

	session.CurrentSuspect = suspectId // Store the suspect ID instead of name
	questions, err := gs.interrogationClient.GetInterrogationQuestions(context.Background(), &interrogationpb.GetInterrogationQuestionsRequest{
		SuspectId: suspectId,
	})
	if err != nil {
		log.Printf("Error getting interrogation questions: %v", err)
		gs.sendErrorMessage(session.Conn, "Error preparing interrogation. Please try again.")
		return
	}

	questionList := fmt.Sprintf("❓ Questions for %s:\n", suspectName)
	for i, q := range questions.Questions {
		questionList += fmt.Sprintf("%d. %s\n", i+1, q.Question)
	}
	gs.sendMessage(session.Conn, questionList)
	gs.sendMessage(session.Conn, "Type 'ask [question number]' to ask a question.")
}

func (gs *GameService) askQuestion(session *GameSession, questionNumber string) {
	questions, err := gs.interrogationClient.GetInterrogationQuestions(context.Background(), &interrogationpb.GetInterrogationQuestionsRequest{
		SuspectId: session.CurrentSuspect,
	})
	if err != nil {
		log.Printf("Error getting interrogation questions: %v", err)
		gs.sendErrorMessage(session.Conn, "Error retrieving question. Please try again.")
		return
	}

	qNum, err := strconv.Atoi(questionNumber)
	if err != nil || qNum < 1 || qNum > len(questions.Questions) {
		gs.sendErrorMessage(session.Conn, "Invalid question number. Please try again.")
		return
	}

	question := questions.Questions[qNum-1]
	gs.sendMessage(session.Conn, fmt.Sprintf("💬 Question: %s\nAnswer: %s", question.Question, question.Answer))
	gs.sendGameOptions(session)
}

func (gs *GameService) promptForSolution(session *GameSession) {
	gs.sendMessage(session.Conn, "Who do you think committed the crime? Enter the name of the suspect:")
	// Set a flag in the session to indicate that we're waiting for a solution
	session.WaitingForSolution = true
}

func (gs *GameService) solveCaseAttempt(session *GameSession, solution string) {
	result, err := gs.deductionClient.SolveCase(context.Background(), &deductionpb.SolveCaseRequest{
		PlayerId:         session.Player.Id,
		CaseId:           session.Case.Id,
		ProposedSolution: solution,
	})
	if err != nil {
		log.Printf("Error solving case: %v", err)
		gs.sendErrorMessage(session.Conn, "Error solving case. Please try again.")
		return
	}

	if result.IsCorrect {
		_, err := gs.playerClient.UpdatePlayerProgress(context.Background(), &playerpb.UpdatePlayerProgressRequest{Id: session.Player.Id})
		if err != nil {
			log.Printf("Error updating player progress: %v", err)
			gs.sendErrorMessage(session.Conn, "Error updating player progress. Please try again.")
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
- Type 'examine [location]' to examine a specific location 🔍
- Type 'interrogate [suspect]' to interrogate a specific suspect 💬
- Type 'solve' to attempt to solve the case 🔍
`
	gs.sendMessage(session.Conn, options)
}
