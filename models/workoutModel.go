package workoutmodel

import "gorm.io/gorm"

// todo all structs
type WorkoutPlan struct {
	gorm.Model
	Title     string
	Content   string
	Exercices []string
}
