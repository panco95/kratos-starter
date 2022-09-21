package data

import (
	"context"
	"time"

	"demo/app/user/service/internal/biz"
	"demo/app/user/service/internal/errors"
	"demo/app/user/service/models"

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

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	db := r.data.MysqlCli.Db().WithContext(ctx)
	err := db.Model(&models.User{}).
		Where("username = ?", username).
		First(user).
		Error
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.ErrUserNotFound
	}
	return user, nil
}

func (r *userRepo) BuildToken(ctx context.Context, id uint, expire time.Duration) (string, error) {
	token, err := r.data.Jwt.BuildToken(id, expire)
	if err != nil {
		return "", err
	}
	return token, nil
}
