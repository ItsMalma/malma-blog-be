package validator

import (
	"github.com/ItsMalma/gomal"
	"github.com/malma/malma-blog-be/exception"
	"github.com/malma/malma-blog-be/model"
)

type BlogValidator struct {
}

func NewBlogValidator() BlogValidator {
	return BlogValidator{}
}

func (BlogValidator) ValidateCreate(req model.CreateBlogRequest) error {
	results := gomal.Validate(
		gomal.If("title", req.Title).NotNil().NotEmpty().Length(1, 128),
		gomal.If("description", req.Description).NotNil().NotEmpty().Length(1, 256),
		gomal.If("content", req.Content).NotNil().NotEmpty(),
	)
	return exception.TransformValidationResults(results)
}

func (BlogValidator) ValidateUpdate(req model.UpdateBlogRequest) error {
	results := gomal.Validate(
		gomal.If("title", req.Title).NotNil().NotEmpty().Length(1, 128),
		gomal.If("description", req.Description).NotNil().NotEmpty().Length(1, 256),
		gomal.If("content", req.Content).NotNil().NotEmpty(),
	)
	return exception.TransformValidationResults(results)
}

func (BlogValidator) ValidateID(id int64) error {
	if id < 1 {
		return exception.ValidatorError("Blog's id is less than 1")
	}
	return nil
}
