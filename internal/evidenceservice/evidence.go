package evidenceservice

import (
	"context"
	"fmt"
	evidencepb "oda/api/proto/evidence"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	evidencepb.UnimplementedEvidenceServiceServer
	pool *pgxpool.Pool
}

func NewServer(connString string) (*Server, error) {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}
	return &Server{pool: pool}, nil
}

func (s *Server) ListEvidence(ctx context.Context, req *evidencepb.ListEvidenceRequest) (*evidencepb.EvidenceList, error) {
	query := `SELECT id, case_id, description, location FROM evidence WHERE case_id = $1`
	rows, err := s.pool.Query(ctx, query, req.CaseId)
	if err != nil {
		return nil, fmt.Errorf("Unable to list evidence: %v", err)
	}
	defer rows.Close()

	var evidenceList []*evidencepb.Evidence
	for rows.Next() {
		var e evidencepb.Evidence
		err := rows.Scan(&e.Id, &e.CaseId, &e.Description, &e.Location)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan row: %v", err)
		}
		evidenceList = append(evidenceList, &e)
	}
	return &evidencepb.EvidenceList{Evidence: evidenceList}, nil
}

func (s *Server) ListLocations(ctx context.Context, req *evidencepb.ListLocationsRequest) (*evidencepb.LocationList, error) {
	query := `SELECT DISCTINCT location FROM evidence WHERE case_id = $1`
	rows, err := s.pool.Query(ctx, query, req.CaseId)
	if err != nil {
		return nil, fmt.Errorf("Unable to list locations: %v", err)
	}
	defer rows.Close()

	var locationList []*evidencepb.Location
	for rows.Next() {
		var l evidencepb.Location
		err := rows.Scan(&l.Name)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan row: %v", err)
		}
		locationList = append(locationList, &l)
	}
	return &evidencepb.LocationList{Locations: locationList}, nil
}
