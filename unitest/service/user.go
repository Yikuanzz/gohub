package service

import (
	"context"
	"fmt"

	"github.com/yikuanzz/unitest/entity"
)

//go:generate mockgen -source=./user.go -destination=../mock/user_mock.go -package=mock
type UserRepo interface {
	AddUser(ctx context.Context, user *entity.User) (err error)
	DelUser(ctx context.Context, userID int) (err error)
	GetUser(ctx context.Context, userID int) (user *entity.User, exist bool, err error)
}

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) AddUser(ctx context.Context, username string) (err error) {
	if len(username) == 0 {
		return fmt.Errorf("username not specify")
	}
	return us.userRepo.AddUser(ctx, &entity.User{Name: username})
}

func (us *UserService) GetUser(ctx context.Context, userID int) (user *entity.User, err error) {
	userInfo, exist, err := us.userRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("user %d not found", userID)
	}

	return userInfo, nil
}
