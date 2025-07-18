package postgres

import (
	"fitgym/backend/repository/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkoutRepository struct {
	db *gorm.DB
}

func NewWorkoutRepository(db *gorm.DB) *WorkoutRepository {
	return &WorkoutRepository{db: db}
}

func (r *WorkoutRepository) CreateWorkout(workout *model.Workout) error {
	workout.ID = uuid.New()
	return r.db.Create(workout).Error
}

func (r *WorkoutRepository) GetWorkoutsByUserID(id uuid.UUID) ([]model.Workout, error) {
	var workouts []model.Workout
	if err := r.db.Preload("Exercises").Where("user_id = ?", id).Find(&workouts).Error; err != nil {
		return nil, err
	}
	return workouts, nil
}

func (r *WorkoutRepository) GetWorkoutByID(id uuid.UUID) (*model.Workout, error) {
	var workout model.Workout
	if err := r.db.Where("id = ?", id).First(&workout).Error; err != nil {
		return nil, err
	}
	return &workout, nil
}

func (r *WorkoutRepository) UpdateWorkout(workout *model.Workout) error {
	return r.db.Save(workout).Error
}

func (r *WorkoutRepository) DeleteWorkout(workout *model.Workout) error {
	return r.db.Delete(workout).Error
}
