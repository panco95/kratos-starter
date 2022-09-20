package service

import (
	"context"

	pb "demo/api/gateway/interface/v1"
	"demo/app/gateway/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GatewayService struct {
	pb.UnimplementedGatewayInterfaceServer

	log *log.Helper
	uc  *biz.GatewayUsecase
}

func NewGatewayService(uc *biz.GatewayUsecase, logger log.Logger) *GatewayService {
	return &GatewayService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *GatewayService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	reply, err := s.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *GatewayService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	reply, err := s.uc.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *GatewayService) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.uc.Logout(ctx)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
