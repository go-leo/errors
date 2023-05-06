package service

import (
	"github.com/go-leo/errors/example/api"
	"github.com/go-leo/errors/example/code"
	"github.com/go-leo/errors/example/internal/data"
)

type UserSvc struct {
	repo data.UserRepo
}

func NewUserSvc() *UserSvc {
	return &UserSvc{
		repo: data.UserRepo{},
	}
}

func (svc *UserSvc) GetUser(req *api.GetUserReq) (*api.GetUserResp, error) {
	if req.UID == 0 {
		return nil, code.NewErrValidation("uid is required")
	}

	name, err := svc.repo.GetUser(req.UID)
	if err != nil {
		return nil, err
	}

	return &api.GetUserResp{UID: req.UID, Name: name}, nil
}
