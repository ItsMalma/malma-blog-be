package controller

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/malma/malma-blog-be/entity"
	"github.com/malma/malma-blog-be/exception"
	"github.com/malma/malma-blog-be/mapper"
	"github.com/malma/malma-blog-be/model"
	"github.com/malma/malma-blog-be/pkg/storage"
	"github.com/malma/malma-blog-be/pkg/utilfiber"
	"github.com/malma/malma-blog-be/pkg/utilfile"
	"github.com/malma/malma-blog-be/repository"
	"github.com/malma/malma-blog-be/validator"
)

type BlogController struct {
	blogValidator  validator.BlogValidator
	blogRepository repository.BlogRepository
	blogMapper     mapper.BlogMapper
	storageManager storage.StorageManager
}

func NewBlogController(
	blogValidator validator.BlogValidator,
	blogRepository repository.BlogRepository,
	blogMapper mapper.BlogMapper,
	storageManager storage.StorageManager,
) BlogController {
	return BlogController{
		blogValidator:  blogValidator,
		blogRepository: blogRepository,
		blogMapper:     blogMapper,
		storageManager: storageManager,
	}
}

func (blogController BlogController) getBlogFromParam(c *fiber.Ctx) (entity.Blog, error) {
	blogId, err := utilfiber.ParamInt64(c, "blogId")
	if err != nil {
		return entity.Blog{}, err
	}

	if err := blogController.blogValidator.ValidateID(blogId); err != nil {
		return entity.Blog{}, err
	}

	blog, err := blogController.blogRepository.FindById(blogId)
	if err != nil {
		return entity.Blog{}, err
	}

	return blog, nil
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
	blog, err := blogController.getBlogFromParam(c)
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
	blog, err := blogController.getBlogFromParam(c)
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
	return c.Status(200).JSON(model.Payload{
		Message: "UPDATED",
		Data:    blogController.blogMapper.EntityToResponse(blog),
		Error:   nil,
	})
}

func (blogController BlogController) DeleteBlogByID(c *fiber.Ctx) error {
	// Get blog by id first
	blog, err := blogController.getBlogFromParam(c)
	if err != nil {
		return err
	}

	// Then, delete the blog
	if err := blogController.blogRepository.Delete(blog); err != nil {
		return err
	}

	// Send response
	return c.Status(200).JSON(model.Payload{
		Message: "DELETED",
		Data:    blogController.blogMapper.EntityToResponse(blog),
		Error:   nil,
	})
}

func (blogController BlogController) UpdateBlogThumbnailByID(c *fiber.Ctx) (err error) {
	// Get blog by id first
	blog, err := blogController.getBlogFromParam(c)
	if err != nil {
		return err
	}

	// Then get thumbnail from form
	thumbnail, err := utilfiber.FormFile(c, "thumbnail")
	if err != nil {
		return err
	}

	if err := blogController.blogValidator.ValidateThumbnail(thumbnail); err != nil {
		return err
	}

	// Save the thumbnail
	var thumbnailFileName string
	if !blog.Thumbnail.Valid {
		thumbnailExtension, err := utilfile.GetExtensionFromBytes(thumbnail)
		if err != nil {
			return err
		}

		thumbnailFileName = uuid.NewString() + thumbnailExtension
	} else {
		thumbnailFileName = blog.Thumbnail.String
	}

	if err := blogController.storageManager.Save(thumbnailFileName, thumbnail); err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if errWhenDeleted := blogController.storageManager.Delete(thumbnailFileName); errWhenDeleted != nil {
				err = errWhenDeleted
			}
		}
	}()

	// Update the blog's thumbnail
	blog.Thumbnail = sql.NullString{
		String: thumbnailFileName,
		Valid:  true,
	}
	blog.UpdatedAt = time.Now()

	// Save the updated blog
	blog, err = blogController.blogRepository.Save(blog)
	if err != nil {
		return err
	}

	// Send response
	return c.Status(200).JSON(model.Payload{
		Message: "UPDATED",
		Data:    blogController.blogMapper.EntityToResponse(blog),
		Error:   nil,
	})
}

func (blogController BlogController) DeleteBlogThumbnailByID(c *fiber.Ctx) (err error) {
	// Get blog by id first
	blog, err := blogController.getBlogFromParam(c)
	if err != nil {
		return err
	}

	// Delete the thumbnail
	if !blog.Thumbnail.Valid {
		return c.Status(404).JSON(model.Payload{
			Message: "NOT_FOUND",
			Data:    nil,
			Error:   "Blog doesn't have a thumbnail",
		})
	}

	thumbnailFileName := blog.Thumbnail.String
	thumbnailContent, err := blogController.storageManager.Load(thumbnailFileName)
	if err != nil {
		return err
	}

	if err := blogController.storageManager.Delete(blog.Thumbnail.String); err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if errWhenSaved := blogController.storageManager.Save(thumbnailFileName, thumbnailContent); errWhenSaved != nil {
				err = errWhenSaved
			}
		}
	}()

	// Update the blog's thumbnail
	blog.Thumbnail = sql.NullString{
		String: "",
		Valid:  false,
	}
	blog.UpdatedAt = time.Now()

	// Save the updated blog
	blog, err = blogController.blogRepository.Save(blog)
	if err != nil {
		return err
	}

	// Send response
	return c.Status(200).JSON(model.Payload{
		Message: "UPDATED",
		Data:    blogController.blogMapper.EntityToResponse(blog),
		Error:   nil,
	})
}

func (blogController BlogController) Routing(router fiber.Router) {
	router.Post("", blogController.CreateBlog)
	router.Get("/:blogId", blogController.GetBlogByID)
	router.Get("", blogController.GetAllBlog)
	router.Put("/:blogId", blogController.UpdateBlogByID)
	router.Delete("/:blogId", blogController.DeleteBlogByID)
	router.Put("/:blogId/thumbnail", blogController.UpdateBlogThumbnailByID)
	router.Delete("/:blogId/thumbnail", blogController.DeleteBlogThumbnailByID)
}
