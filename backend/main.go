package main

import (
	"fmt"

	"github.com/mbient/todo-api/routers"
)

func main() {
	fmt.Println("TODO LIST API")
	r := routers.TaskRouter()
	r.Run()
}
