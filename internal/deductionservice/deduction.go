package deductionservice

import (
	"context"
	"fmt"
	deductionpb "oda/api/proto/deduction"
	"oda/pkg/dbconnect"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	deductionpb.UnimplementedDeductionServiceServer
	pool *pgxpool.Pool
}

func NewServer(connString string) (*Server, error) {
	pool, err := dbconnect.ConnectWithRetry(connString, 5, time.Second*3)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}
	return &Server{pool: pool}, nil
}

func (s *Server) SolveCase(ctx context.Context, req *deductionpb.SolveCaseRequest) (*deductionpb.SolutionResult, error) {
	query := `UPDATE player_cases SET status = CASE WHEN $3 = (SELECT solution FROM cases WHERE id = player_cases.case_id) THEN 'solved' ELSE 'failed' END
  WHERE player_id = $1 AND case_id = $2
  RETURNING (status = 'solved') as is_correct`
	row := s.pool.QueryRow(ctx, query, req.PlayerId, req.CaseId, req.ProposedSolution)

	var result deductionpb.SolutionResult
	err := row.Scan(&result.IsCorrect)
	if err != nil {
		return nil, fmt.Errorf("Unable to solve case: %v", err)
	}

	if result.IsCorrect {
		result.Feedback = "Congratulations! 🦦 You solved the case! 🎉"
	} else {
		result.Feedback = "Sorry, that's not the correct solution. 😢 Try again!"
	}
	return &result, nil
}
