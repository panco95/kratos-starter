package biz

import (
	"context"
	"time"

	pb "demo/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserUsecase struct {
	userRepo UserRepo
	authRepo AuthRepo
	log      *log.Helper
}

func NewUserUsecase(userRepo UserRepo, authRepo AuthRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
		authRepo: authRepo,
		log:      log.NewHelper(logger),
	}
}

func (uc *UserUsecase) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	user, err := uc.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	token, err := uc.authRepo.BuildToken(ctx, user.ID, time.Hour*24)
	if err != nil {
		return nil, err
	}

	reply := &pb.LoginReply{
		Token: token,
	}
	return reply, nil
}
