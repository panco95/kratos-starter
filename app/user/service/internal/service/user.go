package service

import (
	"context"

	pb "demo/api/user/service/v1"
	"demo/app/user/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	res, err := s.uc.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) Logout(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
