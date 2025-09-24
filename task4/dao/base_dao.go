package dao

import (
	"BlogSystem/global"
	"BlogSystem/model"

	"gorm.io/gorm"
)

type BaseDao struct {
	Orm *gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{
		Orm: global.DB,
	}
}

// 通用分页函数
func Paginate(page model.Paginate) func(orm *gorm.DB) *gorm.DB {
	return func(orm *gorm.DB) *gorm.DB {
		return orm.Offset((page.GetPageNo() - 1) * page.GetPageSize()).Limit(page.GetPageSize())
	}
}
