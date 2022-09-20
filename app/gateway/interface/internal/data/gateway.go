package data

import (
	"context"

	user "demo/api/user/service/v1"
	"demo/app/gateway/interface/internal/biz"
	"demo/app/gateway/interface/internal/conf"
	"demo/pkg/jwt"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	jwt  *jwt.Jwt
	log  *log.Helper
}

func NewGatewayRepo(data *Data, c *conf.Auth, logger log.Logger) biz.GatewayRepo {
	return &userRepo{
		data: data,
		jwt:  jwt.New([]byte(c.Jwt.Key), c.Jwt.Issue),
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

func (r *userRepo) Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterReply, error) {
	reply, err := r.data.UserSvcCli.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
