package biz

import (
	"context"

	pb "demo/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = pb.ErrorUserNotFound("user not found")
)

type User struct {
	Username string
	Password string
	VCode    string
}

type UserRepo interface {
	Login(context.Context, *User) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, g *User) (*User, error) {
	return uc.repo.Login(ctx, g)
}
