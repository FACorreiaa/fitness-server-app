package workouts

import (
	"fitness-server-app/controllers"

	"github.com/gin-gonic/gin"
)

func WorkoutPlans(route *gin.Engine) {
	workouts := route.Group("/workoutplans")
	workouts.POST("/posts", controllers.CreateWorkoutPlan)
	workouts.PUT("/posts/:id", controllers.UpdateWorkoutPlan)
	workouts.GET("/posts", controllers.GetWorkoutPlanById)
	workouts.GET("/posts/:id", controllers.GetAllWorkoutPlans)
	workouts.DELETE("/posts/:id", controllers.DeleteWorkoutPlan)
}
