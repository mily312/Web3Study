package service

import (
	"BlogSystem/dao"
	"BlogSystem/model"
	"BlogSystem/service/dto"
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

// 分页展示文章信息
func (PostService *PostService) PostListPaginate(postPaginate dto.PostPaginateDto) ([]model.Post, int64, error) {
	return postService.Dao.PostList(postPaginate)
}

// 单个文章详情信息
func (postService *PostService) PostDetail(postId uint) (model.Post, error) {
	return postService.Dao.PostDetail(postId)
}

// 文章修改
func (postService *PostService) UpdatePost(updatePost model.UpdatePost) error {

	var post model.Post
	post.ID = updatePost.ID
	post.Title = updatePost.Title
	post.Content = updatePost.Content

	return postService.Dao.UpdatePost(post)
}

// 文章删除
func (postService *PostService) DelPost(postId uint) error {
	return postService.Dao.DelPost(postId)
}
