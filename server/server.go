package server

import (
	"encoding/json"
	"io"
	"kvdb/db"
	"net/http"
)

type Server struct {
	Store *db.KVStore
}

func NewServer(store *db.KVStore) *Server {
	return &Server{Store: store}
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := getKeyFromRequest(r)
	if key == "" {
		s.listKeys(w, r)
		return
	}
	value, exists := s.Store.Get(ctx, key)
	if !exists {
		http.NotFound(w, r)
		return
	}
	io.WriteString(w, value)
}

func (s *Server) PutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := getKeyFromRequest(r)
	value, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	s.Store.Set(ctx, key, string(value))
	w.WriteHeader(http.StatusOK)
}

func (s *Server) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	key := getKeyFromRequest(r)
	deleted := s.Store.Delete(ctx, key)
	if deleted {
		w.WriteHeader(http.StatusOK)
	} else {
		http.NotFound(w, r)
	}
}

func (s *Server) listKeys(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	keys := s.Store.ListKeys(ctx)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(keys)
}

func getKeyFromRequest(r *http.Request) string {
	return r.Context().Value("key").(string)
}
