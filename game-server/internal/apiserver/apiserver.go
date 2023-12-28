package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"
	"v1/internal/apiserver/handlers"
	"v1/internal/config"
	"v1/pkg/storage/postgres"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	logger *logrus.Logger
	config *config.Config
	router *mux.Router
}

func New(config *config.Config, logger *logrus.Logger) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logger,
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {

	connectionStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		s.config.DbPostgres.Username,
		s.config.DbPostgres.Password,
		s.config.DbPostgres.Host,
		s.config.DbPostgres.Port,
		s.config.DbPostgres.DbName,
	)

	store := postgres.NewPostgresDb(connectionStr, s.logger)

	s.configureRouter(store.Db)

	address := fmt.Sprintf("%s:%v", s.config.Server.Host, s.config.Server.Port)

	s.logger.Info("Api server has been started on host: " + string(s.config.Server.Host) + ":" + string(s.config.Server.Port))
	return http.ListenAndServe(address, s.router)

}

func (s *ApiServer) configureRouter(db *sql.DB) {
	handlers.New(db, s.router, s.logger)
}

func baseErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
