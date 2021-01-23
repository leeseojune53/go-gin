package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"TaskProject/models"
	"net/http"
)

type CreateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task string `json:"task"`
	Deadline string `json:"deadline"`
}

type UpdateTaskInput struct {
	AssignedTo string `json:"assignedTo"`
	Task string `json:"task"`
	Deadline string `json:"deadline"`
}

func FindTasks(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context)  {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{AssignedTo: input.AssignedTo, Task: input.Task}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func FindTask(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input UpdateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var updatedInput models.Task
	updatedInput.AssignedTo = input.AssignedTo
	updatedInput.Task = task.Task

	db.Model(&task).Update(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context)  {
	db := c.MustGet("db").(*gorm.DB)

	var task models.Task

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}