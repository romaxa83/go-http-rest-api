package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/romaxa83/go-http-rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	logger *logrus.Logger
	router *mux.Router
	store store.Store
}

func newServer(store store.Store) *server {
	s := &server {
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// функция для конфигурирования роутера
func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
