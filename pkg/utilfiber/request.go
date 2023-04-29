package utilfiber

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ParamInt64(c *fiber.Ctx, name string) (int64, error) {
	p := c.Params(name)
	pi, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return 0, fiber.NewError(400, fmt.Sprintf("Parameter %v is not integer", name))
	}
	return pi, nil
}
