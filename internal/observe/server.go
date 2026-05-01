package observe

import (
	"context"
	"errors"
	"net/http"
)

type SnapshotProvider interface {
	Snapshot() Snapshot
}

type Server struct {
	server *http.Server
}

func NewServer(addr string, provider SnapshotProvider) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/plain; version=0.0.4; charset=utf-8")
		_, _ = w.Write([]byte(RenderPrometheus(provider.Snapshot())))
	})

	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

func (s *Server) ListenAndServe() error {
	err := s.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
