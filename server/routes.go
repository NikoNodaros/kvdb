package server

import (
	"context"
	"net/http"
	"path"
	"strings"
)

func (s *Server) Route() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.keyMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.GetHandler(w, r)
		case http.MethodPut:
			s.PutHandler(w, r)
		case http.MethodDelete:
			s.DeleteHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}))

	return mux
}

func (s *Server) keyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.Trim(path.Clean(r.URL.Path), "/")
		ctx := context.WithValue(r.Context(), "key", key)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
