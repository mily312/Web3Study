package dao

import (
	"BlogSystem/model"
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
