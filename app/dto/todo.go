package dto

import "go-todo-app/app/models"

type TodosResponseDto struct {
	Data []models.Todo `json:"data"`
}

type TodoCreateRequestDto struct {
	Content string `json:"content"`
}

type TodoResponseDto struct {
	Data models.Todo `json:"data"`
}
