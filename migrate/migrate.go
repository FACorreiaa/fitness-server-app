package main

import (
	"fitness-server-app/initializers"
	"fitness-server-app/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&workoutmodel.WorkoutPlan{})
}
