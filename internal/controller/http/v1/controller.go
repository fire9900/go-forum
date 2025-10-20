package v1

import (
	"github.com/go-playground/validator/v10"
	"go-forum/internal/usecase"
	"go-forum/pkg/logger"
)

// V1 -.
type V1 struct {
	uc UCs
	l  logger.Interface
	v  *validator.Validate
}

type UCs struct {
	T usecase.Thread
}
