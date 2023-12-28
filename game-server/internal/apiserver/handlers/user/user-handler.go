package user_handlers

import (
	"encoding/json"
	"fmt"
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

	router.HandleFunc("/users", userHandler.userSave).Methods("POST")
	router.HandleFunc("/users", userHandler.findMany).Methods("GET")
	router.HandleFunc("/users/:id", userHandler.findOne).Methods("GET")
	router.HandleFunc("/users/:id", userHandler.Delete).Methods("DELETE")

}

func (uh *userHandlers) userSave(w http.ResponseWriter, r *http.Request) {}

func (uh *userHandlers) findMany(w http.ResponseWriter, r *http.Request) {

	users, err := uh.storage.FindMany(10, 0)

	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
	}

	jData, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")

	w.Write(jData)

}

func (uh *userHandlers) findOne(w http.ResponseWriter, r *http.Request) {}

func (uh *userHandlers) Delete(w http.ResponseWriter, r *http.Request) {}
