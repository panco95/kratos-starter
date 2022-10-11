package errors

import (
	v1 "app/api/user/service/v1"
)

var (
	UserNotFound       = v1.ErrorUserNotFound("用户不存在")
	UserExists         = v1.ErrorUserExists("用户已存在")
	UsernameHasChinese = v1.ErrorUsernameHasChinese("用户名不允许中文")
	PasswordError      = v1.ErrorPasswordError("密码错误")
)
