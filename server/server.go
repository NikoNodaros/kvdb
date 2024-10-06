package server

import (
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
}

func (s *Server) PutHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) DeleteHandler(w http.ResponseWriter, r *http.Request) {

}
