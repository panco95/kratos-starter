package data

import (
	"context"
	"demo/app/user/service/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *authRepo) BuildToken(ctx context.Context, id uint, expire time.Duration) (string, error) {
	token, err := r.data.Jwt.BuildToken(id, expire)
	if err != nil {
		return "", err
	}
	return token, nil
}
