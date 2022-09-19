package server

import (
	"context"

	pb "demo/api/user/service/v1"
	"demo/app/user/service/internal/conf"
	"demo/app/user/service/internal/service"
	"demo/pkg/tracer"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, userSvc *service.UserService, logger log.Logger) *grpc.Server {
	err := tracer.InitJaegerTracer(c.Tracer.Jaeger.Endpoint, "user", "grpc server")
	if err != nil {
		log.NewHelper(logger).Errorf("InitJaegerTracer %v", err)
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(
				recovery.WithLogger(logger),
				recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {
					return nil
				}),
			),
			metadata.Server(
				metadata.WithPropagatedPrefix("x-app"),
			),
			logging.Server(logger),
			validate.Validator(),
			tracing.Server(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(srv, userSvc)
	return srv
}
