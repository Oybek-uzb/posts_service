package postgres

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/Oybek-uzb/posts_service/internal/posts/model"
	"github.com/Oybek-uzb/posts_service/internal/posts/storage"
	"github.com/Oybek-uzb/posts_service/pkg/client/postgres"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgres.Client
}

func (r *repository) Create(ctx context.Context, post *model.Post) error {
	q := `INSERT INTO posts (id, user_id, title, body) 
		  VALUES ($1, $2, $3, $4)
		  RETURNING id
    	  `

	err := r.client.QueryRow(ctx, q, post.Id, post.UserId, post.Title, post.Body).Scan(&post.Id)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			sqlErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return sqlErr
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) ([]*model.Post, error) {
	qb := sq.Select("id, user_id, title, body").From("public.posts")

	sql, i, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.client.Query(ctx, sql, i...)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			sqlErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return nil, sqlErr
		}
		return nil, err
	}

	posts := make([]*model.Post, 0)

	for rows.Next() {
		var ps model.Post

		err = rows.Scan(&ps.Id, &ps.UserId, &ps.Title, &ps.Body)
		if err != nil {
			if errors.Is(err, pgErr) {
				pgErr = err.(*pgconn.PgError)
				sqlErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
				return nil, sqlErr
			}
			return nil, err
		}

		posts = append(posts, &ps)
	}

	err = rows.Err()
	if err != nil {
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			sqlErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return nil, sqlErr
		}
		return nil, err
	}

	return posts, nil
}

func (r *repository) FindOne(ctx context.Context, id int32) (*model.Post, error) {
	q := `
		SELECT id, user_id, title, body FROM public.posts WHERE id=$1 LIMIT 1
	`

	var ps model.Post
	err := r.client.QueryRow(ctx, q, id).Scan(&ps.Id, &ps.UserId, &ps.Title, &ps.Body)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			sqlErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return nil, sqlErr
		}
		return nil, err
	}

	return &ps, nil
}

func (r *repository) Update(ctx context.Context, post model.Post) (*model.Post, error) {
	q := `
		WITH updated AS (
			UPDATE public.posts
			SET user_id = $1, title = $2, body = $3
			WHERE id = $4
			RETURNING *)
		SELECT updated.* FROM updated
		WHERE id=$4
		LIMIT 1;
	`

	var ps model.Post
	err := r.client.QueryRow(ctx, q, post.UserId, post.Title, post.Body, post.Id).Scan(&ps.Id, &ps.UserId, &ps.Title, &ps.Body)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			sqlErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return nil, sqlErr
		}
		return nil, err
	}

	return &ps, nil
}

func (r *repository) Delete(ctx context.Context, id int32) (*model.Post, error) {
	q := `
		WITH deleted AS (
    		DELETE FROM public.posts
    		WHERE id=$1
    		RETURNING *)
		SELECT deleted.* FROM deleted
		WHERE id=$1
		LIMIT 1;
	`

	var ps model.Post
	err := r.client.QueryRow(ctx, q, id).Scan(&ps.Id, &ps.UserId, &ps.Title, &ps.Body)

	var pgErr *pgconn.PgError
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			sqlErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			return nil, sqlErr
		}
		return nil, err
	}

	return &ps, nil
}

func NewRepository(client postgres.Client) storage.Repository {
	return &repository{
		client: client,
	}
}
