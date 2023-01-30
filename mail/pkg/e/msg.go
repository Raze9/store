package e

var MsgFlags = map[int]string{
	Success:             "ok",
	Error:               "fail",
	InvalidParams:       "参数错误",
	ErrorExit:           "用户存在",
	ErrorFailEncryption: "密码加密失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]

}
