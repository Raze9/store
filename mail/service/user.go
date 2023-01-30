package service

import (
	"GOproject/GIT/mail/dao"
	"GOproject/GIT/mail/model"
	"GOproject/GIT/mail/pkg/e"
	"GOproject/GIT/mail/pkg/util"
	"GOproject/GIT/mail/serializer"
	"context"
)

type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	if service.Key == "" { //此处无法判断==len(service.key)
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Data:    "密码长度不足",
		}
	}
	//密码加密
	util.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExit
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	user = &model.User{
		NickName: service.NickName,
		UserName: service.UserName,
		Status:   model.Active,
		Avatar:   "knoachan.jpg",
		Money:    util.Encrypt.AesEncoding("10000"),
	}

	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	err = userDao.CreateUser(user)
	if err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
	}
}

func (service *UserService) Login(ctx context.Context) {

}
