package service

import (
	"BlogSystem/dao"
	"BlogSystem/model"
)

type PostService struct {
	Dao *dao.PostDao
}

var postService *PostService

func NewPostService() *PostService {
	if postService == nil {
		postService = &PostService{
			Dao: dao.NewPostDao(),
		}
	}

	return postService
}

// 添加文章
func (postService *PostService) AddPost(post *model.Post) error {
	
	return postService.Dao.AddPost(post)
}
