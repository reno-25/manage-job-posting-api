package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/microcosm-cc/bluemonday"
    "github.com/reno-25/redikru-jobs-api/internal/service"
    "github.com/reno-25/redikru-jobs-api/internal/repository"
)

type JobHandler struct {
    svc *service.JobService
    p   *bluemonday.Policy
}

func NewJobHandler(svc *service.JobService) *JobHandler {
    return &JobHandler{
        svc: svc,
        p:   bluemonday.UGCPolicy(), // sanitize user generated content
    }
}

type createJobReq struct {
    CompanyName string `json:"companyName" binding:"required"`
    Title       string `json:"title" binding:"required"`
    Description string `json:"description" binding:"required"`
}

func (h *JobHandler) CreateJob(c *gin.Context) {
    var req createJobReq
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // sanitize inputs to prevent stored XSS
    company := h.p.Sanitize(req.CompanyName)
    title := h.p.Sanitize(req.Title)
    desc := h.p.Sanitize(req.Description)

    job, err := h.svc.CreateJob(company, title, desc)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, job)
}

func (h *JobHandler) ListJobs(c *gin.Context) {
    keyword := c.Query("keyword")
    company := c.Query("companyName")

    limitStr := c.DefaultQuery("limit", "20")
    offsetStr := c.DefaultQuery("offset", "0")
    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit <= 0 || limit > 100 {
        limit = 20
    }
    offset, err := strconv.Atoi(offsetStr)
    if err != nil || offset < 0 {
        offset = 0
    }

    filter := repository.JobFilter{
        Keyword:     keyword,
        CompanyName: company,
        Limit:       limit,
        Offset:      offset,
    }

    jobs, total, err := h.svc.ListJobs(filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  jobs,
        "total": total,
        "limit": limit,
        "offset": offset,
    })
}
