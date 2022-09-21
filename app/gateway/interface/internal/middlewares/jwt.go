package middlewares

import (
	"context"
	"demo/app/gateway/interface/internal/errors"
	"demo/pkg/jwt"
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
				}()
			}
			return handler(ctx, req)
		}
	}
}
