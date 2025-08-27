package service

import (
    "errors"

    "github.com/reno-25/redikru-jobs-api/internal/models"
    "github.com/reno-25/redikru-jobs-api/internal/repository"
)

type JobService struct {
    repo repository.Repository
}

func NewJobService(r repository.Repository) *JobService {
    return &JobService{repo: r}
}

func (s *JobService) CreateJob(companyName, title, description string) (*models.Job, error) {
    if companyName == "" || title == "" || description == "" {
        return nil, errors.New("companyName, title and description are required")
    }
    company, err := s.repo.CreateCompanyIfNotExists(companyName)
    if err != nil {
        return nil, err
    }
    job := &models.Job{
        CompanyID:   company.ID,
        Title:       title,
        Description: description,
    }
    if err := s.repo.CreateJob(job); err != nil {
        return nil, err
    }
    job.Company = company
    return job, nil
}

func (s *JobService) ListJobs(filter repository.JobFilter) ([]models.Job, int64, error) {
    return s.repo.ListJobs(filter)
}
