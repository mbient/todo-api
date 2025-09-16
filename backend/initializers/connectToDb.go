package initializers

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDb() {
	_, err := gorm.Open(sqlite.Open("todo-list.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("successfully connected to the database")
}
