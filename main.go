package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Task model structure
type Task struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

var db *gorm.DB
var err error

func main() {
	// Initialize Gin
	r := gin.Default()

	// Connect to the SQLite database
	db, err = gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// AutoMigrate the Task model to create the tasks table
	db.AutoMigrate(&Task{})

	// Define API endpoints
	r.POST("/tasks", createTask)
	r.GET("/tasks/:id", getTask)
	r.PUT("/tasks/:id", updateTask)
	r.DELETE("/tasks/:id", deleteTask)
	r.GET("/tasks", listTasks)

	// Run the server
	r.Run(":8080")
}

// Create a new task
func createTask(c *gin.Context) {
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Generate a unique ID for the task
	// task.ID = uint(db.Table("tasks").RowsAffected + 1)
	task.Status = "pending"

	// Create the task in the database
	db.Create(&task)

	// Return the created task
	c.JSON(201, task)
}

// Retrieve a task by ID
func getTask(c *gin.Context) {
	id := c.Param("id")
	var task Task

	// Find the task in the database
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// Return the task details
	c.JSON(200, task)
}

// Update a task by ID
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task

	// Find the task in the database
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// Bind the updated task details from the request
	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Update the task in the database
	db.Save(&task)

	// Return the updated task
	c.JSON(200, task)
}

// Delete a task by ID
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	var task Task

	// Find the task in the database
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	// Delete the task from the database
	db.Delete(&task)

	// Return success message
	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}

// List all tasks
func listTasks(c *gin.Context) {
	var tasks []Task

	// Retrieve all tasks from the database
	db.Find(&tasks)

	// Return the list of tasks
	c.JSON(200, tasks)
}
