package service

import (
	"context"
	v1 "hpp-kratos/api/v1"
	"hpp-kratos/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc *biz.GreeterUsecase
}

func NewUserService(uc *biz.GreeterUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) GetUserInfo(ctx context.Context, in *v1.UserInfoRequest) (*v1.UserInfoReply, error) {
	// logic code...
	return &v1.UserInfoReply{Uid: in.Uid, Name: "test user"}, nil
}
