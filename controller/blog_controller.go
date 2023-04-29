package controller

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/malma/malma-blog-be/entity"
	"github.com/malma/malma-blog-be/exception"
	"github.com/malma/malma-blog-be/mapper"
	"github.com/malma/malma-blog-be/model"
	"github.com/malma/malma-blog-be/pkg/utilfiber"
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

func (blogController BlogController) GetBlogByID(c *fiber.Ctx) error {
	blogId, err := utilfiber.ParamInt64(c, "blogId")
	if err != nil {
		return err
	}

	if err := blogController.blogValidator.ValidateID(blogId); err != nil {
		return err
	}

	blog, err := blogController.blogRepository.FindById(blogId)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(model.Payload{
		Message: "FOUND",
		Data:    blogController.blogMapper.EntityToResponse(blog),
		Error:   nil,
	})
}

func (blogController BlogController) GetAllBlog(c *fiber.Ctx) error {
	blogs, err := blogController.blogRepository.FindAll()
	if err != nil {
		return err
	}

	return c.Status(200).JSON(model.Payload{
		Message: "FOUND",
		Data:    blogController.blogMapper.EntitiesToResponses(blogs),
		Error:   nil,
	})
}

func (blogController BlogController) UpdateBlogByID(c *fiber.Ctx) error {
	// Get blog by id first
	blogId, err := utilfiber.ParamInt64(c, "blogId")
	if err != nil {
		return err
	}

	if err := blogController.blogValidator.ValidateID(blogId); err != nil {
		return err
	}

	blog, err := blogController.blogRepository.FindById(blogId)
	if err != nil {
		return err
	}

	// Then, get request body
	if contentType := c.Get(fiber.HeaderContentType, "empty"); contentType != fiber.MIMEApplicationJSON {
		return exception.ErrContentType(fiber.MIMEApplicationJSON, contentType)
	}

	req := model.UpdateBlogRequest{}
	if err := c.BodyParser(&req); err != nil {
		return exception.ErrParseRequest(err.Error())
	}

	if err := blogController.blogValidator.ValidateUpdate(req); err != nil {
		return err
	}

	// Update the blog by request body
	blog.Title = req.Title
	blog.Description = req.Description
	blog.Content = req.Content
	blog.UpdatedAt = time.Now()

	// Save the updated blog
	blog, err = blogController.blogRepository.Save(blog)
	if err != nil {
		return err
	}

	// Send response
	return c.Status(201).JSON(model.Payload{
		Message: "CREATED",
		Data:    blogController.blogMapper.EntityToResponse(blog),
		Error:   nil,
	})
}

func (blogController BlogController) Routing(router fiber.Router) {
	router.Post("", blogController.CreateBlog)
	router.Get("/:blogId", blogController.GetBlogByID)
	router.Get("", blogController.GetAllBlog)
	router.Put("/:blogId", blogController.UpdateBlogByID)
}
