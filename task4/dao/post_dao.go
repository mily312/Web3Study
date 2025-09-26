package dao

import (
	"BlogSystem/model"
	"BlogSystem/service/dto"
)

type PostDao struct {
	BaseDao
}

var postDao *PostDao

func NewPostDao() *PostDao {
	if postDao == nil {
		postDao = &PostDao{
			BaseDao: NewBaseDao(),
		}
	}

	return postDao
}

// 添加文章
func (postDao *PostDao) AddPost(post *model.Post) error {
	//return postDao.Orm.Omit(clause.Associations).Create(post).Error
	return postDao.Orm.Create(post).Error
}

// 文章列表分页展示
func (postDao *PostDao) PostList(postPaginate dto.PostPaginateDto) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	err := postDao.Orm.Model(&model.Post{}).Scopes(Paginate(postPaginate.Paginate)).
		Find(&posts).Offset(-1).Limit(-1).Count(&total).Error

	return posts, total, err
}

// 单个文章详情
func (postDao *PostDao) PostDetail(postId uint) (model.Post, error) {
	var postInfo model.Post
	err := postDao.Orm.Model(&model.Post{}).First(&postInfo, postId).Error

	return postInfo, err
}

// 文章修改
func (postDao *PostDao) UpdatePost(post model.Post) error {
	return postDao.Orm.Model(&post).Updates(post).Error
}

// 文章删除
func (post *PostDao) DelPost(postId uint) error {
	return postDao.Orm.Delete(&model.Post{}, postId).Error
}
