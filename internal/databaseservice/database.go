package databaseservice

import (
	"context"
	"fmt"
	databasepb "oda/api/proto/database"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	databasepb.UnimplementedDatabaseServiceServer
	pool *pgxpool.Pool
}

func NewServer(connString string) (*Server, error) {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return &Server{
		pool: pool,
	}, nil
}

func (s *Server) CreatePlayer(ctx context.Context, req *databasepb.CreatePlayerRequest) (*databasepb.Player, error) {
	id := uuid.New().String()
	query := `INSERT INTO players (id, name, cases_solved) VALUES ($1, $2, $3) RETURNING id, name, cases_solved`
	row := s.pool.QueryRow(ctx, query, id, req.Name, 0)

	var player databasepb.Player
	err := row.Scan(&player.Id, &player.Name, &player.CasesSolved)
	if err != nil {
		return nil, fmt.Errorf("unable to create player: %v", err)
	}
	return &player, nil
}

func (s *Server) GetPlayer(ctx context.Context, req *databasepb.GetPlayerRequest) (*databasepb.Player, error) {
	query := `SELECT id, name, cases_solved FROM players WHERE id = $1`
	row := s.pool.QueryRow(ctx, query, req.Id)

	var player databasepb.Player
	err := row.Scan(&player.Id, &player.Name, &player.CasesSolved)
	if err != nil {
		return nil, fmt.Errorf("unable to get player: %v", err)
	}
	return &player, nil
}

func (s *Server) UpdatePlayerProgress(ctx context.Context, req *databasepb.UpdatePlayerProgressRequest) (*databasepb.Player, error) {
	query := `UPDATE players SET cases_solved = cases_solved + 1 WHERE id = $1 RETURNING id, name, cases_solved`
	row := s.pool.QueryRow(ctx, query, req.Id)

	var player databasepb.Player
	err := row.Scan(&player.Id, &player.Name, &player.CasesSolved)
	if err != nil {
		return nil, fmt.Errorf("unable to update player progress: %v", err)
	}
	return &player, nil
}

func (s *Server) CreateCase(ctx context.Context, req *databasepb.CreateCaseRequest) (*databasepb.Case, error) {
	id := uuid.New().String()
	query := `INSERT INTO cases (id, title, description) VALUES ($1, $2, $3) RETURNING id, title, description`
	row := s.pool.QueryRow(ctx, query, id, req.Title, req.Description)

	var case_ databasepb.Case
	err := row.Scan(&case_.Id, &case_.Title, &case_.Description)
	if err != nil {
		return nil, fmt.Errorf("unable to create case: %v", err)
	}
	return &case_, nil
}

func (s *Server) GetCase(ctx context.Context, req *databasepb.GetCaseRequest) (*databasepb.Case, error) {
	query := `SELECT id, title, description FROM cases WHERE id = $1`
	row := s.pool.QueryRow(ctx, query, req.Id)

	var case_ databasepb.Case
	err := row.Scan(&case_.Id, &case_.Title, &case_.Description)
	if err != nil {
		return nil, fmt.Errorf("unable to get case: %v", err)
	}
	return &case_, nil
}

func (s *Server) ListCases(ctx context.Context, req *databasepb.ListCasesRequest) (*databasepb.CaseList, error) {
	query := `SELECT id, title, description FROM cases`
	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to list cases: %v", err)
	}
	defer rows.Close()

	var cases []*databasepb.Case
	for rows.Next() {
		var c databasepb.Case
		err := rows.Scan(&c.Id, &c.Title, &c.Description)
		if err != nil {
			return nil, fmt.Errorf("unable to list cases: %v", err)
		}
		cases = append(cases, &c)
	}
	return &databasepb.CaseList{Cases: cases}, nil
}

func (s *Server) AssignCaseToPlayer(ctx context.Context, req *databasepb.AssignCaseRequest) (*databasepb.PlayerCase, error) {
	id := uuid.New().String()
	query := `INSERT INTO player_cases (id, player_id, case_id, status) VALUES ($1, $2, $3, $4) RETURNING id, player_id, case_id, status`
	row := s.pool.QueryRow(ctx, query, id, req.PlayerId, req.CaseId, "in_progress")

	var playerCase databasepb.PlayerCase
	err := row.Scan(&playerCase.Id, &playerCase.PlayerId, &playerCase.CaseId, &playerCase.Status)
	if err != nil {
		return nil, fmt.Errorf("unable to assign case to player: %v", err)
	}
	return &playerCase, nil
}

func (s *Server) GetPlayerCase(ctx context.Context, req *databasepb.GetPlayerCaseRequest) (*databasepb.PlayerCase, error) {
	query := `SELECT id, player_id, case_id, status FROM player_cases WHERE id = $1 AND case_id = $2`
	row := s.pool.QueryRow(ctx, query, req.PlayerId, req.CaseId)

	var playerCase databasepb.PlayerCase
	err := row.Scan(&playerCase.Id, &playerCase.PlayerId, &playerCase.CaseId, &playerCase.Status)
	if err != nil {
		return nil, fmt.Errorf("unable to get player case: %v", err)
	}
	return &playerCase, nil
}

func (s *Server) ListEvidence(ctx context.Context, req *databasepb.ListEvidenceRequest) (*databasepb.EvidenceList, error) {
	query := `SELECT id, case_id, description, location FROM evidence WHERE case_id = $1`
	rows, err := s.pool.Query(ctx, query, req.CaseId)
	if err != nil {
		return nil, fmt.Errorf("unable to list evidence: %v", err)
	}
	defer rows.Close()

	var evidenceList []*databasepb.Evidence
	for rows.Next() {
		var e databasepb.Evidence
		err := rows.Scan(&e.Id, &e.CaseId, &e.Description, &e.Location)
		if err != nil {
			return nil, fmt.Errorf("unable to list evidence: %v", err)
		}
		evidenceList = append(evidenceList, &e)
	}
	return &databasepb.EvidenceList{Evidence: evidenceList}, nil
}

func (s *Server) ListSuspects(ctx context.Context, req *databasepb.ListSuspectsRequest) (*databasepb.SuspectList, error) {
	query := `SELECT id, case_id, name, description FROM suspects WHERE case_id = $1`
	rows, err := s.pool.Query(ctx, query, req.CaseId)
	if err != nil {
		return nil, fmt.Errorf("unable to list suspects: %v", err)
	}
	defer rows.Close()

	var suspects []*databasepb.Suspect
	for rows.Next() {
		var s databasepb.Suspect
		err := rows.Scan(&s.Id, &s.CaseId, &s.Name, &s.Description)
		if err != nil {
			return nil, fmt.Errorf("unable to list suspects: %v", err)
		}
		suspects = append(suspects, &s)
	}
	return &databasepb.SuspectList{Suspects: suspects}, nil
}

func (s *Server) GetInterrogationQuestions(ctx context.Context, req *databasepb.GetInterrogationQuestionsRequest) (*databasepb.InterrogationQuestionList, error) {
	query := `SELECT id, case_id, question, answer FROM interrogation_questions WHERE suspect_id = $1`
	rows, err := s.pool.Query(ctx, query, req.SuspectId)
	if err != nil {
		return nil, fmt.Errorf("unable to get interrogation questions: %v", err)
	}
	defer rows.Close()

	var questions []*databasepb.InterrogationQuestion
	for rows.Next() {
		var q databasepb.InterrogationQuestion
		err := rows.Scan(&q.Id, &q.SuspectId, &q.Question, &q.Answer)
		if err != nil {
			return nil, fmt.Errorf("unable to get interrogation questions: %v", err)
		}
		questions = append(questions, &q)
	}
	return &databasepb.InterrogationQuestionList{Questions: questions}, nil
}

func (s *Server) SolveCase(ctx context.Context, req *databasepb.SolveCaseRequest) (*databasepb.SolutionResult, error) {
	query := `
        UPDATE player_cases 
        SET status = CASE WHEN $2 = (SELECT solution FROM cases WHERE id = player_cases.case_id) THEN 'solved' ELSE 'failed' END
        WHERE id = $1
        RETURNING (status = 'solved') as is_correct
    `
	row := s.pool.QueryRow(ctx, query, req.PlayerCaseId, req.ProposedSolution)

	var result databasepb.SolutionResult
	err := row.Scan(&result.IsCorrect)
	if err != nil {
		return nil, fmt.Errorf("error solving case: %v", err)
	}

	if result.IsCorrect {
		result.Feedback = "Congratulations! You solved the case."
	} else {
		result.Feedback = "Sorry, that's not the correct solution. Try again!"
	}

	return &result, nil
}

func (s *Server) ListLocations(ctx context.Context, req *databasepb.ListLocationsRequest) (*databasepb.LocationList, error) {
	query := `SELECT DISTINCT location FROM evidence WHERE case_id = $1`
	rows, err := s.pool.Query(ctx, query, req.CaseId)
	if err != nil {
		return nil, fmt.Errorf("unable to list locations: %v", err)
	}
	defer rows.Close()

	var locations []*databasepb.Location
	for rows.Next() {
		var locationName string
		err := rows.Scan(&locationName)
		if err != nil {
			return nil, fmt.Errorf("unable to list locations: %v", err)
		}
		locations = append(locations, &databasepb.Location{Name: locationName})
	}
	return &databasepb.LocationList{Locations: locations}, nil
}
