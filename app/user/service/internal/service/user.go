package service

import (
	"context"
	"strconv"

	pb "demo/api/user/service/v1"
	"demo/app/user/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
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
	userIdString := "0"
	if md, ok := metadata.FromServerContext(ctx); ok {
		userIdString = md.Get("x-app-global-userId")
	}
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return nil, err
	}
	err = s.uc.Logout(ctx, uint(userId))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
