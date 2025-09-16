package initializers

import (
	"github.com/mbient/todo-api/models"
)

func SyncDb() {
	err := DB.AutoMigrate(&models.Task{})
	if err != nil {
		return
	}
}

func FillDb() {
	var tasks = []models.Task{
		{ID: "1", Title: "Buy groceries", Description: "Buy milk, eggs, and bread"},
		{ID: "2", Title: "Clean the house", Description: "Vacuum the living room and dust the shelves"},
		{ID: "3", Title: "Finish project report", Description: "Complete the final draft of the project report and submit it"},
		{ID: "4", Title: "Schedule dentist appointment", Description: "Call the dentist's office to schedule a check-up"},
		{ID: "5", Title: "Exercise", Description: "Go for a 30-minute run in the park"},
		{ID: "6", Title: "Read a book", Description: "Read at least two chapters of the current book"},
		{ID: "7", Title: "Prepare dinner", Description: "Cook spaghetti with marinara sauce and a side salad"},
		{ID: "8", Title: "Water the plants", Description: "Water all indoor and outdoor plants"},
		{ID: "9", Title: "Organize files", Description: "Sort and organize digital files on the computer"},
		{ID: "10", Title: "Call Mom", Description: "Check in with Mom and see how she's doing"},
	}
	DB.Create(&tasks)
}
