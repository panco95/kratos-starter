package data

import (
	"context"
	"time"

	"demo/app/user/service/internal/biz"
	"demo/app/user/service/internal/conf"
	"demo/app/user/service/models"
	"demo/pkg/jwt"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	jwt  *jwt.Jwt
	log  *log.Helper
}

func NewUserRepo(data *Data, c *conf.Auth, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		jwt:  jwt.New([]byte(c.Jwt.Key), c.Jwt.Issue),
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
		return nil, biz.ErrUserNotFound
	}
	return user, nil
}

func (r *userRepo) BuildToken(ctx context.Context, id uint, expire time.Duration) (string, error) {
	token, err := r.jwt.BuildToken(id, expire)
	if err != nil {
		return "", err
	}
	return token, nil
}
