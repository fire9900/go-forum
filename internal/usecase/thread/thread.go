package thread

import (
	"context"
	"github.com/fire9900/go-forum/internal/entity"
	"github.com/fire9900/go-forum/internal/repo"
)

// UseCase -.
type UseCase struct {
	repo repo.ThreadRepo
}

// New -.
func New(r repo.ThreadRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

// Create - создает новый тред через репозиторий.
func (uc *UseCase) Create(ctx context.Context, thread *entity.Thread) error {
	return uc.repo.Create(ctx, thread)
}

// GetByID - получает тред по ID через репозиторий.
func (uc *UseCase) GetByID(ctx context.Context, id int) (*entity.Thread, error) {
	return uc.repo.GetByID(ctx, id)
}

// GetAll - получает все треды через репозиторий.
func (uc *UseCase) GetAll(ctx context.Context) ([]entity.Thread, error) {
	return uc.repo.GetAll(ctx)
}

// Update - обновляет существующий тред через репозиторий.
func (uc *UseCase) Update(ctx context.Context, thread *entity.Thread) error {
	return uc.repo.Update(ctx, thread)
}

// Delete - удаляет тред по ID через репозиторий.
func (uc *UseCase) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, id)
}
