package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	account "hifriend/api/account/v1"
)

var (
	ErrAccountNotFound = errors.NotFound(account.ErrorReason_ACCOUNT_NOT_FOUND.String(), "account not found")
)

type Account struct {
	Account  string
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
