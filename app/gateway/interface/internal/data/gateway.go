package data

import (
	"context"

	user "demo/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GatewayRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayRepo(data *Data, logger log.Logger) *GatewayRepo {
	return &GatewayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *GatewayRepo) Login(ctx context.Context, req *user.LoginReq) (*user.LoginReply, error) {
	reply, err := r.data.UserSvcCli.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (r *GatewayRepo) Register(ctx context.Context, req *user.RegisterReq) (*user.RegisterReply, error) {
	reply, err := r.data.UserSvcCli.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (r *GatewayRepo) Logout(ctx context.Context, req *emptypb.Empty) error {
	_, err := r.data.UserSvcCli.Logout(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
