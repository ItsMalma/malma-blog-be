package repository

import (
	"github.com/malma/malma-blog-be/entity"
	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return BlogRepository{db}
}

func (repo BlogRepository) Save(blog entity.Blog) (entity.Blog, error) {
	if err := repo.db.Save(&blog).Error; err != nil {
		return entity.Blog{}, err
	}

	return blog, nil
}
