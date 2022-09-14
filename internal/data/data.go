package data

import (
	"hifriend/internal/conf"
	"hifriend/internal/data/models"
	"hifriend/pkg/database"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	consulApi "github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAccountRepo)

// Data .
type Data struct {
	MysqlCli  *database.Client
	ConsulCli *consulApi.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	consulCli, err := SetupConsul(c)
	if err != nil {
		return nil, nil, err
	}
	mysqlCli, err := SetupMysql(c)
	if err != nil {
		return nil, nil, err
	}

	return &Data{
		MysqlCli:  mysqlCli,
		ConsulCli: consulCli,
	}, cleanup, nil
}

// SetupMysql .
func SetupMysql(c *conf.Data) (*database.Client, error) {
	cli, err := database.NewMysql(
		c.Mysql.Source,
		int(c.Mysql.MaxIdleConn),
		int(c.Mysql.MaxOpenConn),
		c.Mysql.ConnLifetime.AsDuration(),
		int(c.Mysql.LogLevel),
	)
	if err != nil {
		return nil, err
	}
	err = cli.AutoMigrate(
		&models.Account{},
	)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

// SetupConsul .
func SetupConsul(c *conf.Data) (*consulApi.Client, error) {
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
		return nil, err
	}
	return client, nil
}
