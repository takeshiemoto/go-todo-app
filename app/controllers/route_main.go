package controllers

import (
	"encoding/json"
	"fmt"
	"go-todo-app/app/dto"
	"go-todo-app/app/models"
	"log"
	"net/http"
)

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		postTodo(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	u, err := s.GetUserBySession()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	var todos = []models.Todo{}

	t, err := u.GetTodosByUser()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if t != nil {
		todos = t
	}

	todosResponseDto := dto.TodosResponseDto{
		Data: todos,
	}

	j, err := json.Marshal(todosResponseDto)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(j))

	return
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var requestDto dto.TodoCreateRequestDto

	err = json.Unmarshal(body, &requestDto)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	u, err := s.GetUserBySession()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = u.CreateTodo(requestDto.Content)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	return
}

func updateTodo(w http.ResponseWriter, r *http.Request) {

}
