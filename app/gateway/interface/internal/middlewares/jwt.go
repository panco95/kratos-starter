package middlewares

import (
	"app/app/gateway/interface/internal/errors"
	"app/pkg/jwt"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func CheckToken(jwt *jwt.Jwt) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				token := tr.RequestHeader().Get("x-app-global-token")
				id := jwt.ParseToken(token)
				if id == 0 {
					return nil, errors.UNAUTHORIZED
				}
				ctx = metadata.AppendToClientContext(ctx, "x-app-global-userId", fmt.Sprintf("%d", id))

				defer func() {
					refreshToken, _ := jwt.BuildToken(id, 3600)
					tr.ReplyHeader().Set("x-app-global-refreshToken", refreshToken)
				}()
			}
			return handler(ctx, req)
		}
	}
}

func CheckTokenRoute(project string) func(context.Context, string) bool {
	return func(ctx context.Context, operation string) bool {
		prefix := "/" + project + "."
		if operation == prefix+"gateway.v1.GatewayInterface/Login" ||
			operation == prefix+"gateway.v1.GatewayInterface/Register" {
			return false
		} else {
			return true
		}
	}
}
