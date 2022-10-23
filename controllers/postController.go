package controllers

import (
	"fitness-server-app/initializers"
	workoutmodel "fitness-server-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWorkoutPlan(c *gin.Context) {
	//Get data off req body
	var body struct {
		Title     string
		Content   string
		Exercices []string
	}
	c.Bind(&body)

	//Create workoutplan
	workoutplan := workoutmodel.WorkoutPlan{Title: body.Title, Content: body.Content, Exercices: body.Exercices}

	result := initializers.DB.Create(&workoutplan)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "workoutplan Done",
	})
}

func GetWorkoutPlanById(c *gin.Context) {
	//get the workoutplans
	var workoutplans []workoutmodel.WorkoutPlan
	initializers.DB.Find(&workoutplans)

	//respond with workoutplans
	c.JSON(http.StatusOK, gin.H{
		"workoutplans": workoutplans,
	})
}

func GetAllWorkoutPlans(c *gin.Context) {
	//get id
	id := c.Param("id")

	//get the workoutplans
	var workoutplan workoutmodel.WorkoutPlan
	initializers.DB.First(&workoutplan, id)

	//respond with workoutplans
	c.JSON(http.StatusOK, gin.H{
		"workoutplans": workoutplan,
	})
}
func UpdateWorkoutPlan(c *gin.Context) {
	//get id of url
	id := c.Param("id")

	var body struct {
		Title     string
		Content   string
		Exercices []string
	}

	c.Bind(&body)

	var workoutplan workoutmodel.WorkoutPlan
	initializers.DB.First(&workoutplan, id)

	//update it
	initializers.DB.Model(&workoutplan).Updates(workoutmodel.WorkoutPlan{
		Title:   body.Title,
		Content: body.Content, Exercices: body.Exercices,
	})

	//respond
	c.JSON(http.StatusOK, gin.H{
		"workoutplans": workoutplan,
	})
}

func DeleteWorkoutPlan(c *gin.Context) {
	//get the id
	id := c.Param("id")

	//delete the workoutplans
	initializers.DB.Delete(&workoutmodel.WorkoutPlan{}, id)
	//respond
	c.Status(http.StatusOK)
}
