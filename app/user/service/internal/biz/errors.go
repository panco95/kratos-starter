package biz

import (
	pb "demo/api/user/service/v1"
)

var (
	ErrUserNotFound = pb.ErrorUserNotFound("用户不存在")
)
