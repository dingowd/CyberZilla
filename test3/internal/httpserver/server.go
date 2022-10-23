package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/dingowd/CyberZilla/test3/internal/app"
	"github.com/dingowd/CyberZilla/test3/models"
	"github.com/dingowd/CyberZilla/test3/utils"
	"io/ioutil"
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
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Method isn`t GET")
		return
	}
	admin := r.URL.Query().Get("admin")
	if len(admin) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "admin user missing")
		return
	}
	//var u []byte
	u, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	var user models.User
	if err := json.Unmarshal(u, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	if err := s.App.Storage.CreateUser(admin, user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Write(utils.ReturnStatus("Ok"))
}

func (s *Server) ViewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Method isn`t GET")
		return
	}
	admin := r.URL.Query().Get("admin")
	user := r.URL.Query().Get("user")
	if len(admin) == 0 || len(user) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Check admin and user in URL")
		return
	}
	u, err := s.App.Storage.ViewUser(admin, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	b, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Method isn`t PUT")
		return
	}
	admin := r.URL.Query().Get("admin")
	if len(admin) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Check admin in URL")
		return
	}
	u, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	var user models.User
	if err := json.Unmarshal(u, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	if err := s.App.Storage.UpdateUser(admin, user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Write(utils.ReturnStatus("Ok"))
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Method isn`t DELETE")
		return
	}
	admin := r.URL.Query().Get("admin")
	user := r.URL.Query().Get("user")
	if len(admin) == 0 || len(user) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Check admin and user in URL")
		return
	}
	if err := s.App.Storage.DeleteUser(admin, user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.ReturnError(err.Error()))
		return
	}
	w.Write(utils.ReturnStatus("Ok"))
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
