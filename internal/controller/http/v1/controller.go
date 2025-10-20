package v1

import (
	"github.com/fire9900/go-forum/internal/usecase"
	"github.com/fire9900/go-forum/pkg/logger"
	"github.com/go-playground/validator/v10"
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
