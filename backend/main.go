package main

import (
	"fmt"

	"github.com/mbient/todo-api/initializers"
	"github.com/mbient/todo-api/routers"
)

func init() {
	initializers.ConnectToDb()
	initializers.SyncDb()
	initializers.FillDb()
}

func main() {
	fmt.Println("TODO LIST API")
	r := routers.TaskRouter()
	r.Run()
}
