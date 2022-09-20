package biz

import (
	"context"

	pb "demo/api/gateway/interface/v1"
	user "demo/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type GatewayRepo interface {
	Login(context.Context, *user.LoginReq) (*user.LoginReply, error)
	Register(context.Context, *user.RegisterReq) (*user.RegisterReply, error)
}

type GatewayUsecase struct {
	repo GatewayRepo
	log  *log.Helper
}

func NewGatewayUsecase(repo GatewayRepo, logger log.Logger) *GatewayUsecase {
	return &GatewayUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GatewayUsecase) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	res, err := uc.repo.Login(ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	reply := &pb.LoginReply{
		Token: res.Token,
	}
	return reply, nil
}

func (uc *GatewayUsecase) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	res, err := uc.repo.Register(ctx, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	reply := &pb.RegisterReply{
		Token: res.Token,
	}
	return reply, nil
}

func (uc *GatewayUsecase) Logout(ctx context.Context) error {
	return nil
}
