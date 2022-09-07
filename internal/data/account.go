package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"hifriend/internal/biz"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *accountRepo) Login(ctx context.Context, g *biz.Account) (*biz.Account, error) {
	if g.Account != "root" {
		return nil, biz.ErrAccountNotFound
	}
	return g, nil
}
