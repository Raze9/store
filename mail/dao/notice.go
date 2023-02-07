package dao

import (
	"GOproject/GIT/mail/model"
	"context"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDbClient(ctx)}
}

func NoticeDaoByDb(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
