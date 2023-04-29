package controller

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/malma/malma-blog-be/entity"
	"github.com/malma/malma-blog-be/exception"
	"github.com/malma/malma-blog-be/mapper"
	"github.com/malma/malma-blog-be/model"
	"github.com/malma/malma-blog-be/repository"
	"github.com/malma/malma-blog-be/validator"
)

type BlogController struct {
	blogValidator  validator.BlogValidator
	blogRepository repository.BlogRepository
	blogMapper     mapper.BlogMapper
}

func NewBlogController(
	blogValidator validator.BlogValidator,
	blogRepository repository.BlogRepository,
	blogMapper mapper.BlogMapper,
) BlogController {
	return BlogController{
		blogValidator:  blogValidator,
		blogRepository: blogRepository,
		blogMapper:     blogMapper,
	}
}

func (blogController BlogController) CreateBlog(c *fiber.Ctx) error {
	if contentType := c.Get(fiber.HeaderContentType, "empty"); contentType != fiber.MIMEApplicationJSON {
		return exception.ErrContentType(fiber.MIMEApplicationJSON, contentType)
	}

	req := model.CreateBlogRequest{}
	if err := c.BodyParser(&req); err != nil {
		return exception.ErrParseRequest(err.Error())
	}

	if err := blogController.blogValidator.ValidateCreate(req); err != nil {
		return err
	}

	blog := entity.Blog{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		Thumbnail:   sql.NullString{String: "", Valid: false},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	blog, err := blogController.blogRepository.Save(blog)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(model.Payload{
		Message: "CREATED",
		Data:    blogController.blogMapper.EntityToResponse(blog),
		Error:   nil,
	})
}

func (blogController BlogController) Routing(router fiber.Router) {
	router.Post("", blogController.CreateBlog)
}
