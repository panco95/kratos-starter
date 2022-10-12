package errors

import v1 "app/api/gateway/interface/v1"

var (
	UNAUTHORIZED = v1.ErrorUnauthorized("未授权操作")
)
