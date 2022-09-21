package data

import (
	"context"

	userPB "demo/api/user/service/v1"
	"demo/app/gateway/interface/internal/conf"
	"demo/pkg/jwt"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	consulApi "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGatewayRepo)

// Data .
type Data struct {
	ConsulCli  *consulApi.Client
	UserSvcCli userPB.UserServiceClient
	Jwt        *jwt.Jwt
}

// NewData .
func NewData(dataConf *conf.Data, authConf *conf.Auth, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	var (
		err  error
		data = &Data{
			Jwt: jwt.New([]byte(authConf.Jwt.Key), authConf.Jwt.Issue),
		}
	)

	err = data.SetupConsul(dataConf)
	if err != nil {
		return nil, nil, err
	}
	err = data.SetupGRPCSvcCli(logger)
	if err != nil {
		return nil, nil, err
	}

	return data, cleanup, nil
}

// SetupConsul .
func (data *Data) SetupConsul(c *conf.Data) error {
	client, err := consulApi.NewClient(&consulApi.Config{
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
		return err
	}
	data.ConsulCli = client
	return nil
}

// SetupGRPCSvcCli .
func (data *Data) SetupGRPCSvcCli(logger log.Logger) error {
	selector.SetGlobalSelector(wrr.NewBuilder())
	endpoint := "discovery:///demo.user.service"
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithMiddleware(
			logging.Client(logger),
			tracing.Client(),
			metadata.Client(),
		),
		grpc.WithDiscovery(
			consul.New(data.ConsulCli),
		),
	)
	if err != nil {
		return err
	}

	data.UserSvcCli = userPB.NewUserServiceClient(conn)
	return nil
}
