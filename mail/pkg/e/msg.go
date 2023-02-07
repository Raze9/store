package e

var MsgFlags = map[int]string{
	Success:               "ok",
	Error:                 "fail",
	InvalidParams:         "参数错误",
	ErrorExit:             "用户存在",
	ErrorFailEncryption:   "密码加密失败",
	ErrorExitNotFound:     "用户不存在",
	ErrorNotCompara:       "密码错误",
	ErrorAuthToken:        "token验证失败",
	ErrorAuthTokenTimeout: "认证超时",
	ErrorUploadErr:        "上传失败",
	ErrorSendMail:         "邮件发送失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]

}
