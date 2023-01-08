package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	// ErrorExistUser 成员错误
	ErrorExistUser      = 102
	ErrorNotExistUser   = 103
	ErrorFailEncryption = 106
	ErrorNotCompare     = 107

	ErrorAuthCheckTokenFail    = 301 //token 错误
	ErrorAuthCheckTokenTimeout = 302 //token 过期
	ErrorAuthToken             = 303
	ErrorAuth                  = 304
	ErrorDatabase              = 401
	ErrorNotExistData          = 402
)
