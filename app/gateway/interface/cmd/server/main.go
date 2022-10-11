package main

import (
	"flag"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	"app/app/gateway/interface/internal/conf"
	"app/app/gateway/interface/internal/data"
	"app/pkg/tracer"
	zapPkg "app/pkg/zap"

	zap "github.com/go-kratos/kratos/contrib/log/zap/v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Project     string = ""
	Name        string = ""
	Version     string = "v1.0.0"
	flagconf    string
	flaglogpath string
	id          string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&flaglogpath, "log", "../../logs", "log path, eg: -log logs")
}

func newApp(logger log.Logger, c *conf.Server, data *data.Data, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
		kratos.Registrar(consul.New(data.ConsulCli)),
	)
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
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

	Project = bc.Server.Info.Project
	Name = Project + "." + bc.Server.Info.Name
	id = Name + "#" + bc.Server.Http.Addr
	logger := log.With(zap.NewLogger(zapPkg.NewLogger(flaglogpath, true)),
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		// "trace.id", tracing.TraceID(),
		// "span.id", tracing.SpanID(),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Auth, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	err = tracer.InitJaegerTracer(bc.Server.Tracer.Jaeger.Endpoint, Name)
	if err != nil {
		log.NewHelper(logger).Errorf("InitJaegerTracer %v", err)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
