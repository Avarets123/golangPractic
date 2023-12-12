package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *ApiServer {

	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting api server on port" + string(s.config.BindAddr))

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.configureRouter()

	s.logger.SetLevel(level)

	return nil
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())

}

func (s *ApiServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
