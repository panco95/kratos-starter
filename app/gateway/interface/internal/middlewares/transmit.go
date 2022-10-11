package middlewares

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func Transmit() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				ctx = metadata.AppendToClientContext(ctx, "x-app-global-requestIP", tr.RequestHeader().Get("x-app-global-requestIP"))
				defer func() {
				}()
			}
			return handler(ctx, req)
		}
	}
}
