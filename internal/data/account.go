package data

import (
	"context"
	"hifriend/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
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
	if g.Username != "root" {
		return nil, biz.ErrAccountNotFound(map[string]string{
			"username": g.Username,
		})
	}
	return g, nil
}
