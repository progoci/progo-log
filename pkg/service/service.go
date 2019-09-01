package service

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
)

// Loom describes the loom service.
type Loom interface {
	StartGetBuild(ctx context.Context) http.HandlerFunc
}

// WebSocketReader describes a reader for a websocket connection.
type WebSocketReader func(conn *websocket.Conn)

type loomService struct{}

// NewLoomService creates a new loom service.
func NewLoomService() Loom {
	return &loomService{}
}
