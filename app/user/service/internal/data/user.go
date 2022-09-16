package data

import (
	"context"

	"demo/app/user/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Login(ctx context.Context, g *biz.User) (*biz.User, error) {
	if g.Username != "root" {
		return nil, biz.ErrUserNotFound
	}
	return g, nil
}
