package data

import (
	"context"
	"time"

	pb "demo/api/user/service/v1"
	"demo/app/user/service/internal/biz"
	"demo/pkg/jwt"

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

func (r *userRepo) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	jwtClass := jwt.New([]byte("demo"), "panco")
	token, err := jwtClass.BuildToken(1, time.Hour*24)
	if err != nil {
		return nil, err
	}
	reply := &pb.LoginReply{
		Token: token,
	}
	return reply, nil
}
