package main

import (
	"fmt"
	"go-todo-app/app/controllers"
	"go-todo-app/app/models"
	"log"
)

func main() {
	fmt.Println("Start Rest API")
	fmt.Println(models.Db)

	log.Fatal(controllers.StartMainServer())
}
