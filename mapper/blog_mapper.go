package mapper

import (
	"github.com/malma/malma-blog-be/entity"
	"github.com/malma/malma-blog-be/model"
	"gopkg.in/guregu/null.v4"
)

type BlogMapper struct {
}

func NewBlogMapper() BlogMapper {
	return BlogMapper{}
}

func (mpr BlogMapper) EntityToResponse(blog entity.Blog) model.BlogResponse {
	res := model.BlogResponse{
		ID:          blog.ID,
		Title:       blog.Title,
		Description: blog.Description,
		Content:     blog.Content,
		Thumbnail:   null.NewString(blog.Thumbnail.String, blog.Thumbnail.Valid),
		CreatedAt:   blog.CreatedAt.Unix(),
		UpdatedAt:   blog.UpdatedAt.Unix(),
	}

	return res
}

func (mpr BlogMapper) EntitiesToResponses(blogs []entity.Blog) []model.BlogResponse {
	ress := []model.BlogResponse{}
	for _, blog := range blogs {
		ress = append(ress, mpr.EntityToResponse(blog))
	}

	return ress
}
