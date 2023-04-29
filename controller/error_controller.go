package controller

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/malma/malma-blog-be/exception"
	"github.com/malma/malma-blog-be/model"
)

func ErrorController() fiber.ErrorHandler {
	repoErrTypeCodes := map[exception.RepositoryErrorType]int{
		exception.RepoErrNotFound: 404,
	}

	return func(c *fiber.Ctx, err error) error {
		switch err := err.(type) {
		case *fiber.Error:
			return c.Status(err.Code).JSON(model.Payload{
				Message: model.ToPayloadMsg(http.StatusText(err.Code)),
				Data:    nil,
				Error:   err.Message,
			})
		case exception.ControllerError:
			return c.Status(err.Code).JSON(model.Payload{
				Message: err.Msg,
				Data:    nil,
				Error:   err.Err,
			})
		case exception.ValidatorErrors:
			return c.Status(400).JSON(model.Payload{
				Message: "BAD_REQUEST",
				Data:    nil,
				Error:   err,
			})
		case exception.ValidatorError:
			return c.Status(400).JSON(model.Payload{
				Message: "BAD_REQUEST",
				Data:    nil,
				Error:   err,
			})
		case exception.RepositoryError:
			return c.Status(repoErrTypeCodes[err.Type]).JSON(model.Payload{
				Message: string(err.Type),
				Data:    nil,
				Error:   err.Err,
			})
		default:
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(model.Payload{
				Message: "INTERNAL_SERVER_ERROR",
				Data:    nil,
				Error:   "Internal Server Error",
			})
		}
	}
}
