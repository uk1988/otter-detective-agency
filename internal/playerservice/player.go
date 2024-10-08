package playerservice

import (
	"context"
	"fmt"
	playerpb "oda/api/proto/player"
	"oda/pkg/dbconnect"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	playerpb.UnimplementedPlayerServiceServer
	pool *pgxpool.Pool
}

func NewServer(connString string) (*Server, error) {
	pool, err := dbconnect.ConnectWithRetry(connString, 5, time.Second*3)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}
	return &Server{pool: pool}, nil
}

func (s *Server) CreatePlayer(ctx context.Context, req *playerpb.CreatePlayerRequest) (*playerpb.Player, error) {
	id := uuid.New()
	query := `INSERT INTO players VALUES ($1, $2, $3) RETURNING id, name, cases_solved`
	row := s.pool.QueryRow(ctx, query, id, req.Name, 0)

	var player playerpb.Player
	var dbId uuid.UUID
	err := row.Scan(&dbId, &player.Name, &player.CasesSolved)
	if err != nil {
		return nil, fmt.Errorf("Unable to create new player: %v", err)
	}
	player.Id = dbId.String()
	return &player, nil
}

func (s *Server) GetPlayer(ctx context.Context, req *playerpb.GetPlayerRequest) (*playerpb.Player, error) {
	query := `SELECT id, name, cases_solved FROM players WHERE id = $1`
	row := s.pool.QueryRow(ctx, query, req.Id)

	var player playerpb.Player
	err := row.Scan(&player.Id, &player.Name, &player.CasesSolved)
	if err != nil {
		return nil, fmt.Errorf("Unable to get player: %v", err)
	}
	return &player, nil
}

func (s *Server) UpdatePlayerProgress(ctx context.Context, req *playerpb.UpdatePlayerProgressRequest) (*playerpb.Player, error) {
	query := `UPDATE players SET cases_solved = cases_solved + 1 WHERE id = $1 RETURNING id, name, cases_solved`
	row := s.pool.QueryRow(ctx, query, req.Id)

	var player playerpb.Player
	err := row.Scan(&player.Id, &player.Name, &player.CasesSolved)
	if err != nil {
		return nil, fmt.Errorf("Unable to update player progress: %v", err)
	}
	return &player, nil
}
