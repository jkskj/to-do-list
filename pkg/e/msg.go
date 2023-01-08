package e

var MsgFlags = map[int]string{
	SUCCESS:       "成功",
	ERROR:         "失败!!!!",
	InvalidParams: "请求参数错误",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorNotExistUser:          "用户不存在，请先注册",
	ErrorNotCompare:            "密码不匹配",
	ErrorDatabase:              "数据库操作出错,请重试",
	ErrorExistUser:             "用户已存在",
	ErrorFailEncryption:        "加密密码失败",
	ErrorNotExistData:          "数据不存在",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
