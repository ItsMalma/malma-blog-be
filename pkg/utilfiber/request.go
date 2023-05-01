package utilfiber

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/malma/malma-blog-be/exception"
	"github.com/valyala/fasthttp"
)

func ParamInt64(c *fiber.Ctx, name string) (int64, error) {
	p := c.Params(name)
	pi, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return 0, fiber.NewError(400, fmt.Sprintf("Parameter %v is not integer", name))
	}
	return pi, nil
}

func FormFile(c *fiber.Ctx, name string) ([]byte, error) {
	fileHeader, err := c.FormFile(name)
	if err != nil {
		if errors.Is(err, fasthttp.ErrNoMultipartForm) {
			return nil, exception.ErrContentType("multipart/form-data", c.Get(fiber.HeaderContentType))
		} else if errors.Is(err, fasthttp.ErrMissingFile) {
			return nil, exception.ControllerError{
				Msg:  "FORM_FIELD_NOT_FOUND",
				Err:  fmt.Sprintf("Missing form file with field %v", name),
				Code: 400,
			}
		}
		return nil, err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}
