package service

import (
	"context"
	"hifriend/internal/biz"

	pb "hifriend/api/account/v1"
)

type AccountService struct {
	pb.UnimplementedAccountServer

	uc *biz.AccountUsecase
}

func NewAccountService(uc *biz.AccountUsecase) *AccountService {
	return &AccountService{
		uc: uc,
	}
}

func (s *AccountService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	res, err := s.uc.Login(ctx, &biz.Account{
		Username: req.Username,
		Password: req.Password,
		VCode:    req.Vcode,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: res.Password,
	}, nil
}
