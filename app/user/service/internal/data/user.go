package data

import (
	"context"
	"time"

	"app/app/user/service/internal/biz"
	"app/app/user/service/internal/errors"
	"app/app/user/service/models"
	"app/pkg/utils"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/jameskeane/bcrypt"
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

func (r *userRepo) BuildToken(ctx context.Context, id uint, expire time.Duration) (string, error) {
	token, err := r.data.Jwt.BuildToken(id, expire)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *userRepo) QueryUser(ctx context.Context, user *models.User) (*models.User, error) {
	result := &models.User{}
	err := r.data.MysqlCli.Db().
		Model(&models.User{}).
		Where("`id` = ? OR (`username` <> '' AND `username` = ?) OR (`mobile` <> '' AND `mobile` = ?)", user.ID, user.Username, user.Mobile).
		First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepo) FindUser(ctx context.Context, user *models.User) (*models.User, error) {
	result, err := r.QueryUser(ctx, user)
	if err != nil {
		return nil, err
	}
	if result.ID == 0 {
		return nil, errors.UserNotFound
	}
	return result, nil
}

func (r *userRepo) ExistsUser(ctx context.Context, user *models.User) (bool, error) {
	result, err := r.QueryUser(ctx, user)
	if err != nil {
		return false, err
	}
	if result.ID == 0 {
		return false, nil
	}
	return true, nil
}

func (r *userRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if utils.IsChinese(user.Username) {
		return nil, errors.UsernameHasChinese
	}
	exists, err := r.ExistsUser(ctx, user)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.UserExists
	}

	now := time.Now()
	user.LastLoginTime = &now
	user.LoginTimes = 1
	salt, _ := bcrypt.Salt()
	hash, _ := bcrypt.Hash(user.Password, salt)
	user.PasswordSalt = salt
	user.Password = hash
	if md, ok := metadata.FromServerContext(ctx); ok {
		user.LastLoginIp = md.Get("x-app-global-requestIP")
	}

	err = r.data.MysqlCli.Db().
		Model(&models.User{}).
		Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) Login(ctx context.Context, req *models.User) (*models.User, error) {
	user, err := r.FindUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if !bcrypt.Match(req.Password, user.Password) {
		return nil, errors.PasswordError
	}

	go func() {
		user.LoginTimes++
		if md, ok := metadata.FromServerContext(ctx); ok {
			user.LastLoginIp = md.Get("x-app-global-requestIP")
		}
		now := time.Now()
		user.LastLoginTime = &now
		err = r.data.MysqlCli.Db().
			Model(user).
			Select("login_times", "last_login_time", "last_login_ip").
			Updates(user).Error
		if err != nil {
			r.log.Errorf("login update user %v", err)
		}
	}()

	return user, nil
}
