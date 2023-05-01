package validator

import (
	"net/http"

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
		return exception.ValidatorError("ID must greater than 1")
	}
	return nil
}

func (BlogValidator) ValidateThumbnail(thumbnail []byte) error {
	if len(thumbnail) > 1000000 {
		return exception.ValidatorError("Thumbnail's size must less than 1MB")
	}

	contentType := http.DetectContentType(thumbnail)
	if contentType != "image/png" && contentType != "image/jpeg" {
		return exception.ValidatorError("Thumbnail's type must be png or jpeg")
	}

	return nil
}
