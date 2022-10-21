package main

import (
	"fitness-server-app/initializers"
	workouts "fitness-server-app/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	router := gin.Default()
	workouts.WorkoutPlans(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
