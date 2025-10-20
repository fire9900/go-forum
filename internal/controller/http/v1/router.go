package v1

import (
	"github.com/fire9900/go-forum/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// NewThreadRoutes -.
func NewThreadRoutes(apiV1Group fiber.Router, uc UCs, l logger.Interface) {
	r := &V1{
		uc: uc,
		l:  l,
		v:  validator.New(validator.WithRequiredStructEnabled()),
	}

	translationGroup := apiV1Group.Group("/thread")

	{
		translationGroup.Get("/all", r.getAll)
	}
}
