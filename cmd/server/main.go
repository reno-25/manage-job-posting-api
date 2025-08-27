package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/reno-25/redikru-jobs-api/internal/handlers"
	"github.com/reno-25/redikru-jobs-api/internal/repository"
	"github.com/reno-25/redikru-jobs-api/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        // fallback build DSN from parts
        user := os.Getenv("DB_USER")
        pass := os.Getenv("DB_PASS")
        host := os.Getenv("DB_HOST")
        port := os.Getenv("DB_PORT")
        name := os.Getenv("DB_NAME")
        dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
    }

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect db: %v", err)
    }

    // Migrate basic models (GORM)
    if err := repository.AutoMigrate(db); err != nil {
        log.Fatalf("migrate error: %v", err)
    }

    repo := repository.NewMySQLRepository(db)
    srv := service.NewJobService(repo)
    h := handlers.NewJobHandler(srv)

    r := gin.Default()
    r.POST("/jobs", h.CreateJob)
    r.GET("/jobs", h.ListJobs)

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8081"
    }
    r.Run(":" + port)
}
