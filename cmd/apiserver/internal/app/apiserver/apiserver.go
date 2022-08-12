package apiserver

import (
	"io"
	"net/http"

	"github.com/MaksimUlitin/api-server/cmd/apiserver/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//apiserver
type Apiserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *Apiserver {
	return &Apiserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//start...
func (s *Apiserver) Start() error {

	if err := s.configureStore(); err != nil {
		return err
	}
	s.configureRouter()
	if err := s.configureLoger(); err != nil {
		return err
	}
	s.logger.Info("starting api server!")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Apiserver) configureLoger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *Apiserver) configureRouter() {
	s.router.HandleFunc("/hello", s.hendleHello())
}

func (s *Apiserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *Apiserver) hendleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}

}
