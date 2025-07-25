package transport

import (
	"context"
	"fmt"
	"github.com/SavinDevelop/techcrm-go/internal/handler/user"
	"github.com/SavinDevelop/techcrm-go/pkg/db"
	"net/http"
	"time"
)

type HTTPServer struct {
	db     *db.Postgres
	mux    *http.ServeMux
	server *http.Server
}

func NewHTTPServer(db *db.Postgres) *HTTPServer {
	mux := http.NewServeMux()
	userHand := user.NewHandler(db)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprint(w, `{"status":"ok"}`); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	})
	mux.HandleFunc("POST /api/v1/user", userHand.Create)

	return &HTTPServer{
		db:  db,
		mux: mux,
		server: &http.Server{
			Addr:         "localhost:8080",
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
	}
}

func (s *HTTPServer) Start() {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Failed to start server: %v", err)
	}
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shutdown server: %v", err)
		return err
	}
	if err := s.db.Close(); err != nil {
		fmt.Printf("Failed to close database: %v", err)
		return err
	}
	return nil
}
