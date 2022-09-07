package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"hifriend/internal/conf"
	"hifriend/pkg/database"
	"time"
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

	mysqlCli, err := ConnMysql(c)
	if err != nil {
		return nil, nil, err
	}

	return &Data{
		mysqlCli: mysqlCli,
	}, cleanup, nil
}

func ConnMysql(c *conf.Data) (*database.Client, error) {
	cli, err := database.NewMysql(
		c.Mysql.Source,
		int(c.Mysql.MaxIdleConn),
		int(c.Mysql.MaxOpenConn),
		time.Duration(c.Mysql.ConnLifetime.Seconds)*time.Second,
		int(c.Mysql.LogLevel),
	)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
