package persistent

import (
	"context"
	"fmt"
	"github.com/fire9900/go-forum/internal/entity"
	"github.com/fire9900/go-forum/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

// PostRepo -.
type PostRepo struct {
	*postgres.Postgres
}

// NewPostRepo -.
func NewPostRepo(pg *postgres.Postgres) *PostRepo {
	return &PostRepo{pg}
}

// Create - создает новый пост в базе данных.
func (r *PostRepo) Create(ctx context.Context, post *entity.Post) error {
	query, args, err := r.Builder.
		Insert("posts").
		Columns("content", "create_at", "thread_id", "user_id", "parent_id").
		Values(post.Content, post.CreateAt, post.ThreadID, post.UserID, post.ParentID).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return err
	}

	err = r.Pool.QueryRow(ctx, query, args...).Scan(&post.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetByID - получает пост по ID.
func (r *PostRepo) GetByID(ctx context.Context, id int) (*entity.Post, error) {
	query, args, err := r.Builder.
		Select("id", "content", "create_at", "thread_id", "user_id", "parent_id").
		From("posts").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	post := &entity.Post{}
	err = r.Pool.QueryRow(ctx, query, args...).Scan(
		&post.ID,
		&post.Content,
		&post.CreateAt,
		&post.ThreadID,
		&post.UserID,
		&post.ParentID,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// GetAll - получает все посты из базы данных.
func (r *PostRepo) GetAll(ctx context.Context) ([]entity.Post, error) {
	query, args, err := r.Builder.
		Select("id", "content", "create_at", "thread_id", "user_id", "parent_id").
		From("posts").
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

	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(
			&post.ID,
			&post.Content,
			&post.CreateAt,
			&post.ThreadID,
			&post.UserID,
			&post.ParentID,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// GetByThreadID - получает все посты определенного треда.
func (r *PostRepo) GetByThreadID(ctx context.Context, threadID int) ([]entity.Post, error) {
	query, args, err := r.Builder.
		Select("id", "content", "create_at", "thread_id", "user_id", "parent_id").
		From("posts").
		Where(squirrel.Eq{"thread_id": threadID}).
		OrderBy("create_at ASC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entity.Post
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(
			&post.ID,
			&post.Content,
			&post.CreateAt,
			&post.ThreadID,
			&post.UserID,
			&post.ParentID,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// Update - обновляет существующий пост.
func (r *PostRepo) Update(ctx context.Context, post *entity.Post) error {
	query, args, err := r.Builder.
		Update("posts").
		Set("content", post.Content).
		Where(squirrel.Eq{"id": post.ID}).
		ToSql()
	if err != nil {
		return err
	}

	result, err := r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("post with id %d not found", post.ID)
	}

	return nil
}

// Delete - удаляет пост по ID.
func (r *PostRepo) Delete(ctx context.Context, id int) error {
	query, args, err := r.Builder.
		Delete("posts").
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
		return fmt.Errorf("post with id %d not found", id)
	}

	return nil
}
