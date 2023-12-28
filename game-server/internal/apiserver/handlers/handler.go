package handlers

import (
	"database/sql"
	user_handlers "v1/internal/apiserver/handlers/user"
	user_storage "v1/internal/storage/user"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func New(db *sql.DB, router *mux.Router, logger *logrus.Logger) {

	userStorage := user_storage.New(db, logger)
	user_handlers.Apply(router, userStorage)

}
