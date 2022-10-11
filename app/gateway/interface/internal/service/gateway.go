package service

import (
	"context"

	pb "app/api/gateway/interface/v1"
	user "app/api/user/service/v1"
	"app/app/gateway/interface/internal/data"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GatewayService struct {
	pb.UnimplementedGatewayInterfaceServer

	log  *log.Helper
	repo *data.GatewayRepo
}

func NewGatewayService(repo *data.GatewayRepo, logger log.Logger) *GatewayService {
	return &GatewayService{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (s *GatewayService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	reply, err := s.repo.Login(ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Token: reply.Token,
	}, nil
}

func (s *GatewayService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	reply, err := s.repo.Register(ctx, &user.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.RegisterReply{
		Token: reply.Token,
	}, nil
}

func (s *GatewayService) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.repo.Logout(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
