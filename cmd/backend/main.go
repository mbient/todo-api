package main

import (
	"fmt"

	"github.com/mbient/todo-api/internal/initializers"
	"github.com/mbient/todo-api/internal/routers"
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
