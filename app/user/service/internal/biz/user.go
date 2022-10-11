package biz

import (
	"context"
	"time"

	pb "app/api/user/service/v1"
	"app/app/user/service/models"

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
	user, err := uc.userRepo.Login(ctx, &models.User{Username: req.Username, Password: req.Password})
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
	user, err := uc.userRepo.CreateUser(ctx, &models.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	token, err := uc.userRepo.BuildToken(ctx, user.ID, time.Hour*24*31)
	if err != nil {
		return nil, err
	}

	reply := &pb.RegisterReply{
		Token: token,
	}
	return reply, nil
}

func (uc *UserUsecase) Logout(ctx context.Context, userId uint) error {
	user, err := uc.userRepo.FindUser(ctx, &models.User{Model: models.Model{ID: userId}})
	if err != nil {
		return err
	}
	user.UpdatedAt = time.Now()
	return nil
}
