package errors

import v1 "app/api/gateway/interface/v1"

var (
	UNAUTHORIZED = v1.ErrorUnauthorized("无权限")
)
