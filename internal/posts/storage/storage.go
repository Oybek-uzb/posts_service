package storage

import (
	"context"
	"github.com/Oybek-uzb/posts_service/internal/posts/model"
)

type Repository interface {
	Create(ctx context.Context, author *model.Post) error
	FindAll(ctx context.Context) ([]*model.Post, error)
	FindOne(ctx context.Context, id int32) (*model.Post, error)
	Update(ctx context.Context, post model.Post) (*model.Post, error)
	Delete(ctx context.Context, id int32) (*model.Post, error)
}
