package controllers

import (
	"net/http"
	"strconv"
	"61HW/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get all tasks
// @Description Get all tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	tasks := []models.Task{
		{ID: 1, Title: "Buy groceries", Description: "Buy milk, bread, and cheese"},
		{ID: 2, Title: "Read a book", Description: "Read 'Golang for Beginners'"},
	}
	c.JSON(http.StatusOK, tasks)
}

// @Summary Create a new task
// @Description Create a new task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body models.Task true "Task to create"
// @Success 201 {object} models.Task
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = 3 // Example ID
	c.JSON(http.StatusCreated, task)
}

// @Summary Get a task by ID
// @Description Get a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [get]
func GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task := models.Task{ID: uint(id), Title: "Sample Task", Description: "This is a sample task"}
	c.JSON(http.StatusOK, task)
}
