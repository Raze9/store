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

// 创建
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// 根据username判断是否有该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// 根据id获取user
func (dao *UserDao) GetuserbyId(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

// 根据id更新user
func (dao *UserDao) UpdateUserbyId(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error

}
