package service

import (
	"GOproject/GIT/mail/dao"
	"GOproject/GIT/mail/model"
	"GOproject/GIT/mail/pkg/e"
	"GOproject/GIT/mail/pkg/util"
	"GOproject/GIT/mail/serializer"
	"context"
	"mime/multipart"
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
			Data:    "用户不存在",
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
		Avatar:   "ava.jpg",
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

func (service *UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist || err != nil {
		code = e.ErrorExitNotFound
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompara
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Data:    "密码错误",
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code := e.ErrorAuthToken
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Data:    "token验证失败",
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}

func (service UserService) Update(ctx context.Context, uid uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetuserbyId(uid)
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserbyId(uid, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildUser(user),
	}
}
func (service *UserService) Post(ctx context.Context, uid uint, file multipart.File, filesize int64) serializer.Response {
	code := e.Success
	var (
		user *model.User
		err  error
	)
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetuserbyId(uid)
	if err != nil {
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	path, err := UploadStatic(file, uid, user.UserName)
	if err != nil {
		code = e.ErrorUploadErr
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserbyId(uid, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildUser(user),
	}

}
