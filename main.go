package main

import "go-todo-app/app/models"

func main() {
	//fmt.Println(config.Config.Port)
	//fmt.Println(config.Config.SQLDriver)
	//fmt.Println(config.Config.DbName)
	//fmt.Println(config.Config.LogFile)
	//
	//log.Println("test")
	//fmt.Println(models.Db)
	//
	//u := &models.User{}
	//u.Name = "test"
	//u.Email = "test@example.com"
	//u.PassWord = "testtest"
	//fmt.Println(u)

	//u.CreateUser()

	u, _ := models.GetUser(2)
	//fmt.Printf("%+v\n", u)
	//
	//u.Name = "Takeshi Emoto"
	//u.Email = "test3@example.com"
	//u.UpdateUser()
	//u, _ = models.GetUser(2)
	//fmt.Printf("%+v\n", u)
	u.DeleteUser()
}
