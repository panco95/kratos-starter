package biz

import (
	"context"
	"demo/app/user/service/models"
	"time"
)

type AuthRepo interface {
	BuildToken(ctx context.Context, id uint, expire time.Duration) (string, error)
}

type UserRepo interface {
	FindByUsername(context.Context, string) (*models.User, error)
}
