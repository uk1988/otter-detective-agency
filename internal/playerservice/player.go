package playerservice

import (
	"context"
	"oda/api/proto/player"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	player.UnimplementedPlayerServiceServer
	mu      sync.Mutex
	players map[string]*player.Player
}

func NewServer() *Server {
	return &Server{
		players: make(map[string]*player.Player),
	}
}

func (s *Server) CreatePlayer(ctx context.Context, req *player.CreatePlayerRequest) (*player.Player, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	player := &player.Player{
		Id:          id,
		Name:        req.Name,
		CasesSolved: 0,
	}
	s.players[id] = player
	return player, nil
}

func (s *Server) GetPlayer(ctx context.Context, req *player.GetPlayerRequest) (*player.Player, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	player, ok := s.players[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "player not found")
	}
	return player, nil
}

func (s *Server) UpdatePlayerProgress(ctx context.Context, req *player.UpdatePlayerProgressRequest) (*player.Player, error) {
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
