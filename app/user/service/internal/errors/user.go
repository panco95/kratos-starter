package errors

import (
	v1 "demo/api/user/service/v1"
)

var (
	ErrUserNotFound = v1.ErrorUserNotFound("用户不存在")
)
