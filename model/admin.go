package model

type Admin struct {
	UserName       string
	PasswordDigest string
	Avatar         string `gorm:"size:1000"`
}
