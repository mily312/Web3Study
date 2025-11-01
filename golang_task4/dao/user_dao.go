package dao

import "BlogSystem/model"

type UserDao struct {
	BaseDao
}

// 单例模式
var userDao *UserDao

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			BaseDao: NewBaseDao(),
		}
	}

	return userDao
}

// 用户注册
func (userDao *UserDao) Register(user *model.User) error {
	return userDao.Orm.Save(user).Error
}

// 查询用户信息
func (userDao *UserDao) GetUserInfoByName(userParam *model.User) (model.User, error) {
	var userResult model.User
	err := userDao.Orm.Where("username = ?", userParam.Username).First(&userResult).Error

	return userResult, err
}
