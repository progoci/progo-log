package endpoint

import (
	"context"
	"net/http"
	"progo/loom/pkg/service"
)

// Endpoints holds all the endpoints for the loom service.
type Endpoints struct {
	GetBuild http.HandlerFunc
}

// MakeEndpoints initializes all endpoints for the loom service.
func MakeEndpoints(s service.Loom) Endpoints {
	return Endpoints{
		GetBuild: makeGetBuildEndpoint(s),
	}
}

func makeGetBuildEndpoint(s service.Loom) http.HandlerFunc {
	return s.StartGetBuild(context.Background())
}
