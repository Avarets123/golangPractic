package user_handlers

import (
	"net/http"
	user_storage "v1/internal/storage/user"

	"github.com/gorilla/mux"
)

type userHandlers struct {
	router  *mux.Router
	storage *user_storage.UserStorage
}

func Apply(router *mux.Router, userStorage *user_storage.UserStorage) {

	userHandler := &userHandlers{
		router:  router,
		storage: userStorage,
	}

	router.HandleFunc("/users", userHandler.userSave)

}

func (uh *userHandlers) userSave(w http.ResponseWriter, r *http.Request) {}
