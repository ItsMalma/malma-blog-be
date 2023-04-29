package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID          int64          `gorm:"column:id;primaryKey"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
	Content     string         `gorm:"column:content"`
	Thumbnail   sql.NullString `gorm:"column:thumbnail"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Blog) TableName() string {
	return "blog"
}
