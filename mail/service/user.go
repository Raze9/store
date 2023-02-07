package service

import (
	"GOproject/GIT/mail/conf"
	"GOproject/GIT/mail/dao"
	"GOproject/GIT/mail/model"
	"GOproject/GIT/mail/pkg/e"
	"GOproject/GIT/mail/pkg/util"
	"GOproject/GIT/mail/serializer"
	"context"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
)

type ValidEmailService struct {
}

type SendEmailService struct {
	Email         string `json:"email"form:"email"`
	Password      string `json:"password"form:"password"`
	OperationType uint   `json:"operation_Type"form:"operation_Type"` //1绑定 2解绑 3改密
}

type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

// 注册
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

// 登录
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

// 更新
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

// 更新头像
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

// email 发送
func (service *SendEmailService) Send(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice
	token, err := util.GenerateEmailToken(uid, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	address = conf.ValidEmail + token //邮件链接
	maiStr := notice.Text
	maiText := strings.Replace(maiStr, "Email", address, -1)
	//导入mail.v2
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "RAZE")
	m.SetBody("text/html", maiText)
	d := mail.NewDialer(conf.SmtpHost, 467, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendMail
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
	}
}

// 验证email 用token来验证邮箱
func (service *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var (
		userId        uint
		email         string
		password      string
		operationType uint
	)
	code := e.Success
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthTokenTimeout
		} else {
			userId = claims.UserId
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetuserbyId(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	switch operationType {
	case 1:
		user.Email = email
	case 2:
		user.Email = ""
	case 3:
		err = user.SetPassword(password)
		if err != nil {
			return serializer.Response{
				Status:  code,
				Message: e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserbyId(userId, user)
	if err != nil {
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildUser(user),
	}
}
