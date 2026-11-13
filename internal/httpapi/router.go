package httpapi

import (
	"github.com/mmorgan00/ledgor/internal/store"
	"net/http"
)

type Server struct {
	store *store.Store
}

func New(st *store.Store) *Server {
	return &Server{store: st}
}

func (s *Server) Build(mux *http.ServeMux) {
	BuildTransactionRouter(s, mux)
}
