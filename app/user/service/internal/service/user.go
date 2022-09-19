package service

import (
	"context"

	pb "demo/api/user/service/v1"
	"demo/app/user/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServiceServer

	log *log.Helper
	uc  *biz.UserUsecase
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	res, err := s.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
