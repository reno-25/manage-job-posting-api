package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reno-25/redikru-jobs-api/internal/handlers"
	"github.com/reno-25/redikru-jobs-api/internal/models"
	"github.com/reno-25/redikru-jobs-api/internal/repository"
	"github.com/reno-25/redikru-jobs-api/internal/service"
)

type mockRepo struct{}

func (m *mockRepo) CreateCompanyIfNotExists(name string) (*models.Company, error) {
    return &models.Company{ID: "comp-1", Name: name}, nil
}
func (m *mockRepo) CreateJob(job *models.Job) error {
    job.ID = "job-1"
    return nil
}
func (m *mockRepo) ListJobs(filter repository.JobFilter) ([]models.Job, int64, error) {
    jobs := []models.Job{
        {ID: "job-1", CompanyID: "comp-1", Title: "Backend Engineer", Description: "desc", CreatedAt: jobsTime()},
    }
    return jobs, int64(len(jobs)), nil
}
func jobsTime() (t time.Time) {
    return time.Now()
}

func TestCreateJobHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    svc := service.NewJobService(&mockRepo{})
    h := handlers.NewJobHandler(svc)

    router := gin.Default()
    router.POST("/jobs", h.CreateJob)

    reqBody := map[string]string{
        "companyName": "Acme",
        "title":       "Backend",
        "description": "desc",
    }
    b, _ := json.Marshal(reqBody)
    req := httptest.NewRequest(http.MethodPost, "/jobs", bytes.NewBuffer(b))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    if w.Code != http.StatusCreated {
        t.Fatalf("expected 201 got %d body: %s", w.Code, w.Body.String())
    }
}

func TestListJobsHandler(t *testing.T) {
    gin.SetMode(gin.TestMode)
    svc := service.NewJobService(&mockRepo{})
    h := handlers.NewJobHandler(svc)

    router := gin.Default()
    router.GET("/jobs", h.ListJobs)

    req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Fatalf("expected 200 got %d body: %s", w.Code, w.Body.String())
    }
}
