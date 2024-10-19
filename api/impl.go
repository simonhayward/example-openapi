package api

import (
	"encoding/json"
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var _ ServerInterface = (*Server)(nil)

func (s *Server) GetUserById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	resp := User{Id: id}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
