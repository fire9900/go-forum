// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"go-forum/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	// Thread -.
	Thread interface {
		Create(ctx context.Context, thread *entity.Thread) error
		GetByID(ctx context.Context, id int) (*entity.Thread, error)
		GetAll(ctx context.Context) ([]entity.Thread, error)
		Update(ctx context.Context, thread *entity.Thread) error
		Delete(ctx context.Context, id int) error
	}
)
