package conf_test

import (
	"GOproject/GIT/mail/conf"
	"testing"
)

func TestInit(t *testing.T) {
	conf.Init()

	if conf.AppMode == "" {
		t.Error("AppMode is empty")
	}
	if conf.HttpPort == "" {
		t.Error("HttpPort is empty")
	}
	if conf.Db == "" {
		t.Error("Db is empty")
	}
	if conf.DbName == "" {
		t.Error("DbName is empty")
	}
	if conf.DbPort == "" {
		t.Error("DbPort is empty")
	}
	if conf.DbHost == "" {
		t.Error("DbHost is empty")
	}
	if conf.DbUser == "" {
		t.Error("DbUser is empty")
	}
	if conf.DbPassWord == "" {
		t.Error("DbPassWord is empty")
	}
	if conf.RedisDb == "" {
		t.Error("RedisDb is empty")
	}
	if conf.RedisAddr == "" {
		t.Error("RedisAddr is empty")
	}
	if conf.RedisPw == "" {
		t.Error("RedisPw is empty")
	}
	if conf.RedisDbName == "" {
		t.Error("RedisDbName is empty")
	}
	if conf.ValidEmail == "" {
		t.Error("ValidEmail is empty")
	}
	if conf.SmtpHost == "" {
		t.Error("SmtpHost is empty")
	}
	if conf.SmtpEmail == "" {
		t.Error("SmtpEmail is empty")
	}
	if conf.SmtpPass == "" {
		t.Error("SmtpPass is empty")
	}
	if conf.Host == "" {
		t.Error("Host is empty")
	}
	if conf.ProductPath == "" {
		t.Error("ProductPath is empty")
	}
	if conf.AvatarPath == "" {
		t.Error("AvatarPath is empty")
	}
}
