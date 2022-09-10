package main

import (
	"flag"
	"os"

	"hifriend/internal/conf"

	zapPkg "hifriend/pkg/zap"

	zap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "account"
	// Version is the version of the compiled software.
	Version string = "v1.0.0"
	// flagconf is the config flag.
	flagconf string
	// flaglogpath is the log path.
	flaglogpath string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&flaglogpath, "log", "../../logs", "log path, eg: -log logs")
}

func newApp(logger log.Logger, c *conf.Data, gs *grpc.Server, hs *http.Server) *kratos.App {
	// consul register
	client, err := api.NewClient(&api.Config{
		Address:    c.Consul.Address,
		Scheme:     c.Consul.Scheme,
		PathPrefix: c.Consul.PathPrefix,
		Datacenter: c.Consul.DataCenter,
		WaitTime:   c.Consul.WaitTime.AsDuration(),
		Token:      c.Consul.Token,
		TokenFile:  c.Consul.TokenFile,
		Namespace:  c.Consul.Namespace,
		Partition:  c.Consul.Partition,
	})
	if err != nil {
		panic(err)
	}
	reg := consul.New(client)

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(reg),
	)
}

func main() {
	flag.Parse()
	logger := log.With(zap.NewLogger(zapPkg.NewLogger(flaglogpath, true)),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			env.NewSource("HIFRIENDS_"),
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
