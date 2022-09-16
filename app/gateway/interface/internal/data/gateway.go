package data

import (
	"context"

	"demo/app/gateway/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayRepo(data *Data, logger log.Logger) biz.GatewayRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Login(ctx context.Context, g *biz.User) (*biz.User, error) {
	return g, nil
}
