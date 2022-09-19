package data

import (
	"context"

	v1 "demo/api/user/service/v1"
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
	reply, err := r.data.UserSvcCli.Login(ctx, &v1.LoginReq{Username: g.Username, Password: g.Password})
	if err != nil {
		return nil, err
	}
	r.log.WithContext(ctx).Infof("Reply: %v", reply)
	return g, nil
}
