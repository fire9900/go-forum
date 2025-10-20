package v1

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (r *V1) getAll(ctx *fiber.Ctx) error {

	//TODO implement me
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"123": 123})
}
