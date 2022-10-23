package httpserver

import (
	"fmt"
	"github.com/dingowd/CyberZilla/test3/internal/app"
	"net/http"
)

type Server struct {
	App  *app.App
	Addr string
	Srv  *http.Server
}

func NewServer(app *app.App, addr string) *Server {
	return &Server{App: app, Addr: addr}
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("Hello from server"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Method isn`t GET")
}

func (s *Server) ViewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("Hello from server"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Method isn`t GET")
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("Hello from server"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Method isn`t GET")
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Write([]byte("Hello from server"))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Method isn`t GET")
}

func (s *Server) Start() error {
	s.App.Logg.Info("http server starting")
	mux := http.NewServeMux()
	s.Srv = &http.Server{Addr: s.Addr, Handler: mux}
	mux.HandleFunc("/create", loggingMiddleware(s.CreateUser, s.App.Logg))
	mux.HandleFunc("/view", loggingMiddleware(s.ViewUser, s.App.Logg))
	mux.HandleFunc("/update", loggingMiddleware(s.UpdateUser, s.App.Logg))
	mux.HandleFunc("/delete", loggingMiddleware(s.DeleteUser, s.App.Logg))
	s.Srv.ListenAndServe()
	return nil
}

func (s *Server) Stop() error {
	s.App.Logg.Info("Stop http server")
	return s.Srv.Close()
}
