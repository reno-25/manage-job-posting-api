package repository

import (
    "fmt"
    "strings"

    "github.com/google/uuid"
    "github.com/reno-25/redikru-jobs-api/internal/models"
    "gorm.io/gorm"
)

type mysqlRepo struct {
    db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) Repository {
    return &mysqlRepo{db: db}
}

func (m *mysqlRepo) CreateCompanyIfNotExists(name string) (*models.Company, error) {
    var c models.Company
    if err := m.db.Where("LOWER(name) = ?", strings.ToLower(name)).First(&c).Error; err == nil {
        return &c, nil
    } else if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }
    c = models.Company{
        ID:   uuid.NewString(),
        Name: name,
    }
    if err := m.db.Create(&c).Error; err != nil {
        return nil, err
    }
    return &c, nil
}

func (m *mysqlRepo) CreateJob(job *models.Job) error {
    job.ID = uuid.NewString()
    return m.db.Create(job).Error
}

func (m *mysqlRepo) ListJobs(filter JobFilter) ([]models.Job, int64, error) {
    var jobs []models.Job
    var total int64

    q := m.db.Model(&models.Job{}).Preload("Company")
    if filter.Keyword != "" {
        // simple LIKE search - note: for large dataset consider FULLTEXT or external search
        like := fmt.Sprintf("%%%s%%", filter.Keyword)
        q = q.Where("title LIKE ? OR description LIKE ?", like, like)
    }
    if filter.CompanyName != "" {
        q = q.Joins("JOIN companies ON companies.id = jobs.company_id").
            Where("LOWER(companies.name) = ?", strings.ToLower(filter.CompanyName))
    }

    q.Count(&total)
    q = q.Order("created_at DESC").Limit(filter.Limit).Offset(filter.Offset)
    if err := q.Find(&jobs).Error; err != nil {
        return nil, 0, err
    }
    return jobs, total, nil
}
