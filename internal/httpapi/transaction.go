package httpapi

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mmorgan00/ledgor/internal/store"
	"net/http"
)

type ErrorResponse struct {
	Message string
}

func BuildTransactionRouter(s *Server, mux *http.ServeMux) {
	mux.HandleFunc("POST /transaction", s.create)
	mux.HandleFunc("GET /transaction/{id}", s.get)
	mux.HandleFunc("POST /transaction/reverse", s.reverse)
}

func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	txn := store.NewTransaction(10, nil, "idemp_key")
	s.store.TransactionInsert(&txn)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(txn)

}
func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "invalid transaction id"})
		return
	}

	txn, err := s.store.TransactionGetById(id)
	// Source - https://stackoverflow.com/a/31622112
	// Posted by dm03514, modified by community. See post 'Timeline' for change history
	// Retrieved 2026-11-12, License - CC BY-SA 4.0
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(txn)
}
func (s *Server) reverse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reverse Transaction")
}
