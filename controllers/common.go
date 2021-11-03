package controllers

// Predefined const error strings.
const (
	ErrInputData    = "数据输入错误"
	ErrDatabase     = "数据库操作错误"
	ErrDupUser      = "用户信息已存在"
	ErrNoUser       = "用户信息不存在"
	ErrPass         = "密码不正确"
	ErrNoUserPass   = "用户信息不存在或密码不正确"
	ErrNoUserChange = "用户信息不存在或数据未改变"
	ErrInvalidUser  = "用户信息不正确"
	ErrOpenFile     = "打开文件出错"
	ErrWriteFile    = "写文件出错"
	ErrSystem       = "操作系统错误"
)

// User Data definition
type UserSuccessLoginData struct {
	Id int `json:"id"`
	AccessToken string `json:"access_token"`
	UserName string `json:"user_name"`
}


var (
	errUserToken = &Response{500, 10002, "服务器错误", "令牌操作错误"}
	errNoUser = &Response{400, 10004, "用户信息不存在", "数据库记录不存在"}
)