package repository

import (
	"errors"
	"fmt"

	"github.com/malma/malma-blog-be/entity"
	"github.com/malma/malma-blog-be/exception"
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

func (repo BlogRepository) FindById(id int64) (entity.Blog, error) {
	blog := entity.Blog{}
	if err := repo.db.Where("id = ?", id).First(&blog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = exception.NewRepoErr(exception.RepoErrNotFound, fmt.Sprintf("Can't found blog with id %v", id))
		}
		return entity.Blog{}, err
	}

	return blog, nil
}

func (repo BlogRepository) FindAll() ([]entity.Blog, error) {
	blogs := []entity.Blog{}
	if err := repo.db.Find(&blogs).Error; err != nil {
		return nil, err
	}

	return blogs, nil
}
