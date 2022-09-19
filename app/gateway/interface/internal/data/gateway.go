package data

import (
	"context"

	user "demo/api/user/service/v1"
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

func (r *userRepo) Login(ctx context.Context, req *user.LoginReq) (*user.LoginReply, error) {
	reply, err := r.data.UserSvcCli.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
