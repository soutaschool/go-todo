package main

import (
	"fmt"

	"github.com/soutaschool/todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(2)
	// user.CreateTodo("Second todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user, _ := models.GetUser(2)
	// todos, _ := user.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(1)
	t.Content = "Update todo"
	t.UpdateTodo()
}
