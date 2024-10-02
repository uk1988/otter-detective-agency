package csiservice

import (
	"context"
	"fmt"
	"math/rand"
	evidencepb "oda/api/proto/evidence"
	"oda/pkg/dbconnect"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	evidencepb.UnimplementedCSIServiceServer
	pool *pgxpool.Pool
}

func NewServer(connString string) (*Server, error) {
	pool, err := dbconnect.ConnectWithRetry(connString, 5, time.Second*3)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return &Server{pool: pool}, nil
}

func (s *Server) AnalyzeEvidence(ctx context.Context, req *evidencepb.AnalyzeEvidenceRequest) (*evidencepb.AnalysisResult, error) {
	// Generate number between 1 and 10
	randomNumber := rand.Intn(10) + 1

	var result string
	if randomNumber > 3 {
		query := `SELECT analysis_result FROM evidence WHERE id = $1`
		err := s.pool.QueryRow(ctx, query, req.EvidenceId).Scan(&result)
		if err != nil {
			return nil, fmt.Errorf("failed to get analysis result: %v", err)
		}
		if result == "" {
			result = "The analysis was inconclusive"
		}
	} else {
		result = "The analysis was inconclusive"
	}
	return &evidencepb.AnalysisResult{
		EvidenceId: req.EvidenceId,
		Result:     result,
	}, nil
}
