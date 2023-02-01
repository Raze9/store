package serializer

import (
	"GOproject/GIT/mail/conf"
	"GOproject/GIT/mail/model"
)

type User struct {
	ID       uint   `json:"id" `
	UserName string `json:"user_name" `
	Type     int    `json:"type"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"createAt"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) User {
	return User{
		ID:       user.ID,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
