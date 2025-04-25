package controllers

import (
	"motey-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (tc *TaskController) RegisterRoutes(c *TaskController, router *gin.RouterGroup) {
	tasks := router.Group("/tasks")
	{
		tasks.GET(":id", tc.GetTask)
		tasks.GET("", tc.GetTasks)
		tasks.POST("", tc.CreateTask)
	}
}


func (tc *TaskController) GetTask(c *gin.Context) {
	uuid := c.Param("id")

	task, err := tc.taskService.GetTaskByID(c. Request.Context(), uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskService.GetTasks(c.Request.Context(), c.Query("userid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var request services.CreateTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.taskService.CreateTask(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}
