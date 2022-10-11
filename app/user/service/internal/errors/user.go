package errors

import (
	v1 "demo/api/user/service/v1"
)

var (
	UserNotFound       = v1.ErrorUserNotFound("用户不存在")
	UserExists         = v1.ErrorUserExists("用户已存在")
	UsernameHasChinese = v1.ErrorUsernameHasChinese("用户名不允许中文")
)
