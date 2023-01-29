package dao

import (
	"GOproject/GIT/mail/model"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDbClient(ctx)}
}
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// 根据username判断是否有该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Error
	if err == nil || err == gorm.ErrRecordNotFound {
		return user, false, err
	}
	return user, true, nil
}
