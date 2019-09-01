package http

import (
	"net/http"

	"github.com/gorilla/mux"

	"progo/loom/pkg/endpoint"
)

// NewService wires endpoints to the HTTP transport.
func NewService(svcEndpoints endpoint.Endpoints) http.Handler {

	r := mux.NewRouter()

	r.Methods("GET").Path("/get").Handler(svcEndpoints.GetBuild)

	return r
}
