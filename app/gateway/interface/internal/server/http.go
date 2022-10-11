package server

import (
	"context"

	pb "app/api/gateway/interface/v1"
	"app/app/gateway/interface/internal/conf"
	"app/app/gateway/interface/internal/data"
	"app/app/gateway/interface/internal/middlewares"
	"app/app/gateway/interface/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, data *data.Data, gatewaySvc *service.GatewayService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.RequestDecoder(transmitRequestDecoder),
		http.Middleware(
			recovery.Recovery(
				recovery.WithLogger(logger),
				recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {
					return nil
				}),
			),
			tracing.Server(),
			logging.Server(logger),
			metadata.Server(
				metadata.WithPropagatedPrefix("x-app"),
			),
			middlewares.Transmit(),
			selector.Server(
				recovery.Recovery(),
				tracing.Server(),
				middlewares.CheckToken(data.Jwt),
			).Match(middlewares.CheckTokenRoute(c.Info.Project)).Build(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	pb.RegisterGatewayInterfaceHTTPServer(srv, gatewaySvc)
	return srv
}
