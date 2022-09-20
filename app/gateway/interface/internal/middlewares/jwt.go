package middlewares

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func CheckToken() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				token := tr.RequestHeader().Get("x-app-global-token")
				log.Print(token)
			}
			return handler(ctx, req)
		}
	}
}
