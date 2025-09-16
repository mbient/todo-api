package initializers

import (
	"github.com/mbient/todo-api/models"
)

func SyncDb() {
	err := DB.AutoMigrate(&models.Task{}, models.User{})
	if err != nil {
		return
	}
}

func FillDb() {
	var tasks = []models.Task{
		{Title: "Buy groceries", Description: "Buy milk, eggs, and bread"},
		{Title: "Clean the house", Description: "Vacuum the living room and dust the shelves"},
		{Title: "Finish project report", Description: "Complete the final draft of the project report and submit it"},
		{Title: "Schedule dentist appointment", Description: "Call the dentist's office to schedule a check-up"},
		{Title: "Exercise", Description: "Go for a 30-minute run in the park"},
		{Title: "Read a book", Description: "Read at least two chapters of the current book"},
		{Title: "Prepare dinner", Description: "Cook spaghetti with marinara sauce and a side salad"},
		{Title: "Water the plants", Description: "Water all indoor and outdoor plants"},
		{Title: "Organize files", Description: "Sort and organize digital files on the computer"},
		{Title: "Call Mom", Description: "Check in with Mom and see how she's doing"},
	}
	DB.Create(&tasks)
}
