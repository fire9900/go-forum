package persistent

import (
	"context"
	"fmt"
	"go-forum/internal/entity"
	"go-forum/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

// ThreadRepo -.
type ThreadRepo struct {
	*postgres.Postgres
}

// NewThreadRepo -.
func NewThreadRepo(pg *postgres.Postgres) *ThreadRepo {
	return &ThreadRepo{pg}
}

// Create - создает новый тред в базе данных.
func (r *ThreadRepo) Create(ctx context.Context, thread *entity.Thread) error {
	query, args, err := r.Builder.
		Insert("threads").
		Columns("title", "content", "create_at", "user_id").
		Values(thread.Title, thread.Content, thread.CreateAt, thread.UserID).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return err
	}

	err = r.Pool.QueryRow(ctx, query, args...).Scan(&thread.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetByID - получает тред по ID.
func (r *ThreadRepo) GetByID(ctx context.Context, id int) (*entity.Thread, error) {
	query, args, err := r.Builder.
		Select("id", "title", "content", "create_at", "user_id").
		From("threads").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	thread := &entity.Thread{}
	err = r.Pool.QueryRow(ctx, query, args...).Scan(
		&thread.ID,
		&thread.Title,
		&thread.Content,
		&thread.CreateAt,
		&thread.UserID,
	)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

// GetAll - получает все треды из базы данных.
func (r *ThreadRepo) GetAll(ctx context.Context) ([]entity.Thread, error) {
	query, args, err := r.Builder.
		Select("id", "title", "content", "create_at", "user_id").
		From("threads").
		OrderBy("create_at DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []entity.Thread
	for rows.Next() {
		var thread entity.Thread
		err := rows.Scan(
			&thread.ID,
			&thread.Title,
			&thread.Content,
			&thread.CreateAt,
			&thread.UserID,
		)
		if err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return threads, nil
}

// Update - обновляет существующий тред.
func (r *ThreadRepo) Update(ctx context.Context, thread *entity.Thread) error {
	query, args, err := r.Builder.
		Update("threads").
		Set("title", thread.Title).
		Set("content", thread.Content).
		Where(squirrel.Eq{"id": thread.ID}).
		ToSql()
	if err != nil {
		return err
	}

	result, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("thread with id %d not found", thread.ID)
	}

	return nil
}

// Delete - удаляет тред по ID.
func (r *ThreadRepo) Delete(ctx context.Context, id int) error {
	query, args, err := r.Builder.
		Delete("threads").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	result, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("thread with id %d not found", id)
	}

	return nil
}
