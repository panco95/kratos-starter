package service

import (
	"context"

	pb "demo/api/gateway/interface/v1"
	"demo/app/gateway/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
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
	// if md, ok := metadata.FromServerContext(ctx); ok {
	// 	token := md.Get("x-app-global-token")
	// 	s.log.WithContext(ctx).Infof("Token: %s", token)
	// }

	reply, err := s.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
