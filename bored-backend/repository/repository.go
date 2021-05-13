package repository

import (
	"fmt"

	"github.com/silvagpmiguel/bored/bored-backend/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Repository gorm structure
type Repository struct {
	db *gorm.DB
}

// InitRepository open sqlite connection and set auto migration to create/update table data
func InitRepository(name string) (repo *Repository, err error) {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database : %v", err)
	}
	db.AutoMigrate(&model.Activity{})
	return &Repository{db}, nil
}

// CreateActivity create a new activity and persist it to bd
func (repo *Repository) CreateActivity(activity *model.Activity) bool {
	result := repo.db.Create(activity)
	return result.RowsAffected > 0
}

// GetActivity get a list of activities from an activity structure
func (repo *Repository) GetActivity(activity *model.Activity) ([]model.Activity, bool) {
	var activities []model.Activity
	result := repo.db.Where(activity).Find(&activities)
	return activities, result.RowsAffected > 0
}

// UpdateActivity find all activities from a predicateStruct and update them with the values of newActivity
func (repo *Repository) UpdateActivity(predicateStruct *model.Activity, newActivity model.Activity) bool {
	result := repo.db.Where(predicateStruct).Updates(newActivity)
	return result.RowsAffected > 0
}
