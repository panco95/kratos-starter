package biz

import (
	"context"
	"time"

	pb "demo/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserUsecase struct {
	userRepo UserRepo
	log      *log.Helper
}

func NewUserUsecase(userRepo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		log:      log.NewHelper(logger),
	}
}

func (uc *UserUsecase) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	user, err := uc.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	token, err := uc.userRepo.BuildToken(ctx, user.ID, time.Hour*24)
	if err != nil {
		return nil, err
	}

	reply := &pb.LoginReply{
		Token: token,
	}
	return reply, nil
}

func (uc *UserUsecase) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	user, err := uc.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	token, err := uc.userRepo.BuildToken(ctx, user.ID, time.Hour*24)
	if err != nil {
		return nil, err
	}

	reply := &pb.RegisterReply{
		Token: token,
	}
	return reply, nil
}

func (uc *UserUsecase) Logout(ctx context.Context, userId uint) error {
	user, err := uc.userRepo.FindByUserId(ctx, uint(userId))
	if err != nil {
		return err
	}
	user.UpdatedAt = time.Now()
	return nil
}
