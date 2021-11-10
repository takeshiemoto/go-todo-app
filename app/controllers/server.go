package controllers

import (
	"fmt"
	"go-todo-app/app/models"
	"go-todo-app/config"
	"net/http"

	"github.com/gorilla/mux"
)

func session(w http.ResponseWriter, r *http.Request) (session models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		session = models.Session{UUID: cookie.Value}
		if ok, _ := session.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}

	return session, err
}

func StartMainServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/signup", signup).Methods(http.MethodPost)
	r.HandleFunc("/signin", signin).Methods(http.MethodPost)
	r.HandleFunc("/signout", signout).Methods(http.MethodPost)

	r.HandleFunc("/todos", todoHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/todos/{id}", todoIdHandler).Methods(http.MethodGet, http.MethodPatch, http.MethodDelete)

	r.Handle("/", r)

	return http.ListenAndServe(":"+config.Config.Port, r)
}
