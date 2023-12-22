package apiserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"v1/internal/app/users"
	"v1/internal/storage"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	logger *logrus.Logger
	config *Config
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

	err := s.configureLogger()

	if err != nil {
		return err
	}

	store, err := storage.New(s.logger)

	if err != nil {
		return err
	}

	s.configureRouter(store.Db)

	address := fmt.Sprintf("%s:%v", s.config.Server.Host, s.config.Server.Port)

	fmt.Println(address)

	s.logger.Info("Api server has been started on host: " + string(s.config.Server.Host) + ":" + string(s.config.Server.Port))
	return http.ListenAndServe(address, s.router)

}

func (s *ApiServer) configureLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil

}

type TestStruct struct {
	Router    string
	IsSuccess bool
}

func (s *ApiServer) configureRouter(db *sql.DB) {

	s.router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		jsData, _ := json.Marshal(TestStruct{Router: "/hello", IsSuccess: true})

		w.Write(jsData)

	})

	s.router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		userStorage := users.New(db)

		userId := userStorage.Save("simple", "123")

		s.logger.Info("new userId is" + string(userId))

		w.Write([]byte(userId))

	})

}

func baseErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
