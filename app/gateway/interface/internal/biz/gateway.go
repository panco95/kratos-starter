package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string
	Password string
	VCode    string
}

type GatewayRepo interface {
	Login(context.Context, *User) (*User, error)
}

type GatewayUsecase struct {
	repo GatewayRepo
	log  *log.Helper
}

func NewGatewayUsecase(repo GatewayRepo, logger log.Logger) *GatewayUsecase {
	return &GatewayUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GatewayUsecase) Login(ctx context.Context, g *User) (*User, error) {
	return uc.repo.Login(ctx, g)
}
