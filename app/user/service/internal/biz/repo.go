package biz

import (
	"context"
	"demo/app/user/service/models"
	"time"
)

type UserRepo interface {
	BuildToken(context.Context, uint, time.Duration) (string, error)
	QueryUser(context.Context, *models.User) (*models.User, error)
	FindUser(context.Context, *models.User) (*models.User, error)
	ExistsUser(context.Context, *models.User) (bool, error)
	CreateUser(context.Context, *models.User) (*models.User, error)
}
