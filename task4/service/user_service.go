package service

import (
	"BlogSystem/dao"
	"BlogSystem/model"
	"BlogSystem/utils"
	"errors"
)

type UserService struct {
	Dao *dao.UserDao
}

var userService *UserService

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}

	return userService
}

func (userServie *UserService) Register(user *model.User) error {
	//对密码加密处理
	passwordEncrypt, err := utils.Encrypt(user.Password)

	if err != nil {
		return err
	}

	user.Password = passwordEncrypt
	return userServie.Dao.Register(user)
}

// 用户登录，返回token
func (userServie *UserService) Login(userParam *model.User) (string, error) {
	//1.验证密码输入是否正确
	userInfo, err := userServie.Dao.GetUserInfoByName(userParam)

	if err != nil { //todo 查不到数据该如何处理
		return "", err
	}

	var flag = utils.CompareHashAndPassword(userInfo.Password, userParam.Password)
	if !flag {
		return "", errors.New("incorrect password")
	}

	//2.根据id和name生成token，并返回
	return utils.GenerateToken(userInfo.ID, userInfo.Username)

}
