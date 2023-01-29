package service

import (
	"GOproject/GIT/mail/dao"
	"GOproject/GIT/mail/model"
	"GOproject/GIT/mail/pkg/e"
	"GOproject/GIT/mail/pkg/util"
	"GOproject/GIT/mail/serializer"
	"context"
)

type UserRegister struct {
	Nickname string `json:"nickname"form:"nick_name"`
	UserName string `json:"username"form:"user_name"`
	PassWord string `json:"password"form:"password"`
	key      string `json:"key"form:"key"`
}

func (service UserRegister) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if service.key == "" || len(service.key) != 16 {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   "密码长度不足",
		}
	}
	util.Encrypt.SetKey(service.key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{Status: code, Message: e.GetMsg(code)}
	}
	if exist {
		code = e.ErrorExit
		return serializer.Response{Status: code, Message: e.GetMsg(code)}
	}
	user = model.User{UserName: service.UserName, NickName: service.Nickname, Status: model.Active,
		Avatar: "knoachan.jpg", Money: util.Encrypt.AesEncoding("10000")}

	if err = user.SetPassword(service.PassWord); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
	}
}
