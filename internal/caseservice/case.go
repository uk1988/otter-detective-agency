package caseservice

import (
	"context"
	"fmt"
	casepb "oda/api/proto/case"
	"oda/pkg/dbconnect"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	casepb.UnimplementedCaseServiceServer
	pool *pgxpool.Pool
}

func NewServer(connString string) (*Server, error) {
	pool, err := dbconnect.ConnectWithRetry(connString, 5, time.Second*3)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}
	return &Server{pool: pool}, nil
}

func (s *Server) CreateCase(ctx context.Context, req *casepb.CreateCaseRequest) (*casepb.Case, error) {
	id := uuid.New()
	query := `INSERT INTO cases VALUES ($1, $2, $3) RETURNING id, title, description`
	row := s.pool.QueryRow(ctx, query, id, req.Title, req.Description)

	var c casepb.Case
	var dbId uuid.UUID
	err := row.Scan(&dbId, &c.Title, &c.Description)
	if err != nil {
		return nil, fmt.Errorf("Unable to create new case: %v", err)
	}
	c.Id = dbId.String()
	return &c, nil
}

func (s *Server) GetCase(ctx context.Context, req *casepb.GetCaseRequest) (*casepb.Case, error) {
	query := `SELECT id, title, description FROM cases WHERE id = $1`
	row := s.pool.QueryRow(ctx, query, req.Id)

	var c casepb.Case
	err := row.Scan(&c.Id, &c.Title, &c.Description)
	if err != nil {
		return nil, fmt.Errorf("Unable to get case: %v", err)
	}
	return &c, nil
}

func (s *Server) ListCases(ctx context.Context, req *casepb.ListCasesRequest) (*casepb.CaseList, error) {
	query := `SELECT id, title, description FROM cases`
	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Unable to list cases: %v", err)
	}
	defer rows.Close()

	var cases casepb.CaseList
	for rows.Next() {
		var c casepb.Case
		err := rows.Scan(&c.Id, &c.Title, &c.Description)
		if err != nil {
			return nil, fmt.Errorf("Unable to list cases: %v", err)
		}
		cases.Cases = append(cases.Cases, &c)
	}
	return &cases, nil
}

func (s *Server) AssignCaseToPlayer(ctx context.Context, req *casepb.AssignCaseRequest) (*casepb.PlayerCase, error) {
	query := `INSERT INTO player_cases VALUES ($1, $2, $3) RETURNING player_id, case_id, status`
	row := s.pool.QueryRow(ctx, query, req.PlayerId, req.CaseId, "in_progress")

	var pc casepb.PlayerCase
	err := row.Scan(&pc.PlayerId, &pc.CaseId, &pc.Status)
	if err != nil {
		return nil, fmt.Errorf("Unable to assign case to player: %v", err)
	}
	return &pc, nil
}

func (s *Server) GetPlayerCase(ctx context.Context, req *casepb.GetPlayerCaseRequest) (*casepb.GetPlayerCaseResponse, error) {
	query := `SELECT pc.player_id, pc.case_id, pc.status, c.title, c.description FROM player_cases pc JOIN cases c ON pc.case_id = c.id WHERE pc.player_id = $1 AND pc.status = 'in_progress'`
	row := s.pool.QueryRow(ctx, query, req.PlayerId)

	var pc casepb.PlayerCase
	var c casepb.Case
	err := row.Scan(&pc.PlayerId, &pc.CaseId, &pc.Status, &c.Title, &c.Description)
	if err != nil {
		return nil, fmt.Errorf("Unable to get player case: %v", err)
	}
	return &casepb.GetPlayerCaseResponse{PlayerCase: &pc, Case: &c}, nil
}
