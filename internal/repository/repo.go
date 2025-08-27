package repository

import (
    "github.com/reno-25/redikru-jobs-api/internal/models"
    "gorm.io/gorm"
)

type Repository interface {
    CreateCompanyIfNotExists(name string) (*models.Company, error)
    CreateJob(job *models.Job) error
    ListJobs(filter JobFilter) ([]models.Job, int64, error)
}

type JobFilter struct {
    Keyword     string
    CompanyName string
    Limit       int
    Offset      int
}

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(&models.Company{}, &models.Job{})
}
