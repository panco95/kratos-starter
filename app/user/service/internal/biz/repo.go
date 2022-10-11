package biz

import (
	"app/app/user/service/models"
	"context"
	"time"
)

type UserRepo interface {
	BuildToken(context.Context, uint, time.Duration) (string, error)
	QueryUser(context.Context, *models.User) (*models.User, error)
	FindUser(context.Context, *models.User) (*models.User, error)
	ExistsUser(context.Context, *models.User) (bool, error)
	CreateUser(context.Context, *models.User) (*models.User, error)
	Login(context.Context, *models.User) (*models.User, error)
}
