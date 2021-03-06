package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	//code = 1000 User Error
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USER_NOT_EXIT      = 1003
	ERROR_TOKEN_EXIST        = 1004
	ERROR_TOKEN_RUNTIME      = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_TYPE_WRONG   = 1007
	ERROR_USER_NO_PERMISSION = 1008

	//code = 2000 Article Error
	ERROR_ART_NOT_EXIST = 2001

	//code = 3000 Category Wrong
	ERROR_CATENAME_USED = 3001
	ERROR_CATE_NOT_EXIT = 3003
)

var codeMsg = map[int]string{
	SUCCSE:                   "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "用户名已存在",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIT:      "用户不存在",
	ERROR_USER_NO_PERMISSION: "用户无权限",
	ERROR_TOKEN_EXIST:        "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:      "TOKEN过期",
	ERROR_TOKEN_WRONG:        "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:   "TOKEN格式错误",

	ERROR_ART_NOT_EXIST: "文章不存在",

	ERROR_CATENAME_USED: "分类已存在",
	ERROR_CATE_NOT_EXIT: "分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
