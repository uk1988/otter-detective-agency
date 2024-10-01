package interrogationservice

import (
	"context"
	"fmt"
	interrogationpb "oda/api/proto/interrogation"
	"oda/pkg/dbconnect"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	interrogationpb.UnimplementedInterrogationServiceServer
	pool *pgxpool.Pool
}

func NewServer(connString string) (*Server, error) {
	pool, err := dbconnect.ConnectWithRetry(connString, 5, time.Second*3)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}
	return &Server{pool: pool}, nil
}

func (s *Server) ListSuspects(ctx context.Context, req *interrogationpb.ListSuspectsRequest) (*interrogationpb.SuspectList, error) {
	query := `SELECT id, case_id, name, description FROM suspects WHERE case_id = $1`
	rows, err := s.pool.Query(ctx, query, req.CaseId)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch suspects: %v", err)
	}
	defer rows.Close()

	var suspects []*interrogationpb.Suspect
	for rows.Next() {
		var s interrogationpb.Suspect
		err := rows.Scan(&s.Id, &s.CaseId, &s.Name, &s.Description)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan suspect: %v", err)
		}
		suspects = append(suspects, &s)
	}
	return &interrogationpb.SuspectList{Suspects: suspects}, nil
}

func (s *Server) GetInterrogationQuestions(ctx context.Context, req *interrogationpb.GetInterrogationQuestionsRequest) (*interrogationpb.InterrogationQuestionList, error) {
	query := `SELECT id, suspect_id, question, answer FROM interrogation_questions WHERE suspect_id = $1`
	rows, err := s.pool.Query(ctx, query, req.SuspectId)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch interrogation questions: %v", err)
	}
	defer rows.Close()

	var questions []*interrogationpb.InterrogationQuestion
	for rows.Next() {
		var q interrogationpb.InterrogationQuestion
		err := rows.Scan(&q.Id, &q.SuspectId, &q.Question, &q.Answer)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan interrogation question: %v", err)
		}
		questions = append(questions, &q)
	}
	return &interrogationpb.InterrogationQuestionList{Questions: questions}, nil
}
