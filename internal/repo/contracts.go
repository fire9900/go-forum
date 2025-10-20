package repo

import (
	"context"
	"github.com/fire9900/go-forum/internal/entity"
)

type (
	ThreadRepo interface {
		Create(ctx context.Context, thread *entity.Thread) error
		GetByID(ctx context.Context, id int) (*entity.Thread, error)
		GetAll(ctx context.Context) ([]entity.Thread, error)
		Update(ctx context.Context, thread *entity.Thread) error
		Delete(ctx context.Context, id int) error
	}

	PostRepo interface {
		Create(ctx context.Context, post *entity.Post) error
		GetByID(ctx context.Context, id int) (*entity.Post, error)
		GetAll(ctx context.Context) ([]entity.Post, error)
		GetByThreadID(ctx context.Context, threadID int) ([]entity.Post, error)
		Update(ctx context.Context, post *entity.Post) error
		Delete(ctx context.Context, id int) error
	}
)
