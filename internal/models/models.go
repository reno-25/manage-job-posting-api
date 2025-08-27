package models

import "time"

type Company struct {
    ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
    Name      string    `gorm:"type:text;not null" json:"name"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Job struct {
    ID          string    `gorm:"type:char(36);primaryKey" json:"id"`
    CompanyID   string    `gorm:"type:char(36);index" json:"company_id"`
    Company     *Company  `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
    Title       string    `gorm:"type:text;not null" json:"title"`
    Description string    `gorm:"type:text;not null" json:"description"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
