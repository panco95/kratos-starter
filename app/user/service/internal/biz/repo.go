package biz

import (
	"context"
	"demo/app/user/service/models"
	"time"
)

type UserRepo interface {
	FindByUsername(context.Context, string) (*models.User, error)
	FindByUserId(context.Context, uint) (*models.User, error)
	BuildToken(ctx context.Context, id uint, expire time.Duration) (string, error)
}
