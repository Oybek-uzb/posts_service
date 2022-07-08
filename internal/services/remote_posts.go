package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Oybek-uzb/posts_service/internal/posts/model"
	"github.com/Oybek-uzb/posts_service/internal/posts/storage"
	pbp "github.com/Oybek-uzb/posts_service/pkg/api/posts_service"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type remotePostsService struct {
	pbp.UnimplementedRemotePostsServiceServer
	repository storage.Repository
}

func NewRemotePostsService(repo storage.Repository) *remotePostsService {
	return &remotePostsService{
		repository: repo,
	}
}

func (r *remotePostsService) GetRemotePosts(ctx context.Context, req *pbp.GetRemotePostsRequest) (*pbp.GetRemotePostsResponse, error) {
	fmt.Println(req)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	var remotePostsChan = make(chan []*model.Post)

	go fetchPosts(remotePostsChan)

	remotePosts := <-remotePostsChan

	for _, rs := range remotePosts {
		err := r.repository.Create(ctx, rs)
		if err != nil {
			return &pbp.GetRemotePostsResponse{IsProcessFinishedSuccessfully: false}, err
		}
	}

	return &pbp.GetRemotePostsResponse{
		IsProcessFinishedSuccessfully: true,
	}, nil
}

func fetchPosts(ch chan []*model.Post) {
	var collectedPosts = make([]*model.Post, 0)
	var mx sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()

			response, err := http.Get(fmt.Sprintf("https://gorest.co.in/public/v1/posts?page=%d", page))
			defer response.Body.Close()
			if err != nil {
				return
			}

			jsonData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return
			}

			var resBody model.Body

			err = json.Unmarshal([]byte(jsonData), &resBody)
			if err != nil {
				return
			}

			for _, post := range resBody.Data {
				mx.Lock()
				collectedPosts = append(collectedPosts, &post)
				mx.Unlock()
			}
		}(i)
	}
	wg.Wait()

	for _, fs := range collectedPosts {
		fmt.Println(fs)
	}

	ch <- collectedPosts
}
