package service

import (
	"context"

	pb "demo/api/gateway/interface/v1"
	"demo/app/gateway/interface/internal/biz"
)

type GatewayService struct {
	pb.UnimplementedGatewayInterfaceServer

	uc *biz.GatewayUsecase
}

func NewGatewayService(uc *biz.GatewayUsecase) *GatewayService {
	return &GatewayService{
		uc: uc,
	}
}

func (s *GatewayService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	// if md, ok := metadata.FromServerContext(ctx); ok {
	// 	token := md.Get("x-app-global-token")
	// 	log.Print(token)
	// }

	_, err := s.uc.Login(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: "token",
	}, nil
}
