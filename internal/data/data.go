package data

import (
	"hifriend/internal/conf"
	"hifriend/internal/data/models"
	"hifriend/pkg/database"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAccountRepo)

// Data .
type Data struct {
	mysqlCli *database.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	mysqlCli, err := SetupMysql(c)
	if err != nil {
		return nil, nil, err
	}

	return &Data{
		mysqlCli: mysqlCli,
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
