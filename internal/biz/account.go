package biz

import (
	"context"

	account "hifriend/api/account/v1"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrAccountNotFound = func(metadata map[string]string) error {
		return account.ErrorAccountNotFound("account not found").WithMetadata(metadata)
	}
)

type Account struct {
	Username string
	Password string
	VCode    string
}

type AccountRepo interface {
	Login(context.Context, *Account) (*Account, error)
}

type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewAccountUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AccountUsecase) Login(ctx context.Context, g *Account) (*Account, error) {
	uc.log.WithContext(ctx).Infof("Account: %v", g)
	return uc.repo.Login(ctx, g)
}
