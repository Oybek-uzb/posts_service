package services

import (
	"context"
	"github.com/Oybek-uzb/posts_service/internal/posts/model"
	"github.com/Oybek-uzb/posts_service/internal/posts/storage"
	"time"

	pbp "github.com/Oybek-uzb/posts_service/pkg/api/posts_service"
)

type postsService struct {
	pbp.UnimplementedPostsServiceServer
	repository storage.Repository
}

func NewPostsService(repo storage.Repository) *postsService {
	return &postsService{
		repository: repo,
	}
}

func (s *postsService) GetAllPosts(ctx context.Context, req *pbp.GetAllPostsRequest) (*pbp.GetAllPostsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	allPosts, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var postsLen = len(allPosts)

	if postsLen == 0 {
		return &pbp.GetAllPostsResponse{
			Posts: []*pbp.Post{},
		}, nil
	}

	var resAllPosts []*pbp.Post
	for _, post := range allPosts {
		newPost := pbp.Post{
			Id:     post.Id,
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		}

		resAllPosts = append(resAllPosts, &newPost)
	}

	return &pbp.GetAllPostsResponse{
		Posts: resAllPosts,
	}, nil
}

func (s *postsService) GetPost(ctx context.Context, req *pbp.GetPostRequest) (*pbp.GetPostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	post, err := s.repository.FindOne(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	if post == nil {
		return &pbp.GetPostResponse{
			Post: nil,
		}, nil
	}

	var newPost = &pbp.Post{
		Id:     post.Id,
		UserId: post.UserId,
		Title:  post.Title,
		Body:   post.Body,
	}

	return &pbp.GetPostResponse{
		Post: newPost,
	}, nil
}

func (s *postsService) UpdatePartialPost(ctx context.Context, req *pbp.UpdatePartialPostRequest) (*pbp.UpdatePartialPostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	var updateData = model.Post{
		Id:     req.UpdateData.Id,
		UserId: req.UpdateData.UserId,
		Title:  req.UpdateData.Title,
		Body:   req.UpdateData.Body,
	}

	updatedPost, err := s.repository.Update(ctx, updateData)
	if err != nil {
		return nil, err
	}

	if updatedPost == nil {
		return &pbp.UpdatePartialPostResponse{
			UpdatedPost: nil,
		}, nil
	}

	var newPost = &pbp.Post{
		Id:     updatedPost.Id,
		UserId: updatedPost.UserId,
		Title:  updatedPost.Title,
		Body:   updatedPost.Body,
	}

	return &pbp.UpdatePartialPostResponse{
		UpdatedPost: newPost,
	}, nil
}
func (s *postsService) DeletePost(ctx context.Context, req *pbp.DeletePostRequest) (*pbp.DeletePostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	deletedPost, err := s.repository.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	if deletedPost == nil {
		return &pbp.DeletePostResponse{
			DeletedPost: nil,
		}, nil
	}

	var resDeletedPost = &pbp.Post{
		Id:     deletedPost.Id,
		UserId: deletedPost.UserId,
		Title:  deletedPost.Title,
		Body:   deletedPost.Body,
	}

	return &pbp.DeletePostResponse{
		DeletedPost: resDeletedPost,
	}, nil
}
