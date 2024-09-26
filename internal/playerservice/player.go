package playerservice

import (
	"context"
	playerpb "oda/api/proto/player"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	playerpb.UnimplementedPlayerServiceServer
	mu      sync.Mutex
	players map[string]*playerpb.Player
}

func NewServer() *Server {
	return &Server{
		players: make(map[string]*playerpb.Player),
	}
}

func (s *Server) CreatePlayer(ctx context.Context, req *playerpb.CreatePlayerRequest) (*playerpb.Player, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	player := &playerpb.Player{
		Id:          id,
		Name:        req.Name,
		CasesSolved: 0,
	}
	s.players[id] = player
	return player, nil
}

func (s *Server) GetPlayer(ctx context.Context, req *playerpb.GetPlayerRequest) (*playerpb.Player, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	player, ok := s.players[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "player not found")
	}
	return player, nil
}

func (s *Server) UpdatePlayerProgress(ctx context.Context, req *playerpb.UpdatePlayerProgressRequest) (*playerpb.Player, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	player, ok := s.players[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "player not found")
	}

	if req.CaseSolved {
		player.CasesSolved++
	}
	return player, nil
}
