package main

import (
	"fmt"

	"github.com/soutaschool/todo-app/app/controllers"
	"github.com/soutaschool/todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
