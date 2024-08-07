// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateUser implements createUser operation.
	//
	// POST /user
	CreateUser(ctx context.Context, req *UserCreate) (CreateUserRes, error)
	// GameTick implements gameTick operation.
	//
	// PUT /games/{game_id}/tick
	GameTick(ctx context.Context, params GameTickParams) (GameTickRes, error)
	// GetBodies implements getBodies operation.
	//
	// GET /games/{game_id}/bodies
	GetBodies(ctx context.Context, params GetBodiesParams) (GetBodiesRes, error)
	// GetGames implements getGames operation.
	//
	// GET /games
	GetGames(ctx context.Context) (GetGamesRes, error)
	// GetSystems implements getSystems operation.
	//
	// GET /games/{game_id}/systems
	GetSystems(ctx context.Context, params GetSystemsParams) (GetSystemsRes, error)
	// GetTypes implements getTypes operation.
	//
	// GET /types
	GetTypes(ctx context.Context) (GetTypesRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
