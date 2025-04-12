package repo

import (
	"context"

	"github.com/yikuanzz/unitest/entity"
	"xorm.io/xorm"
)

type UserRepo interface {
	AddUser(ctx context.Context, user *entity.User) (err error)
	DelUser(ctx context.Context, userID int) (err error)
	GetUser(ctx context.Context, userID int) (user *entity.User, exist bool, err error)
}

type userRepo struct {
	db *xorm.Engine
}

func NewUserRepo(db *xorm.Engine) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) AddUser(ctx context.Context, user *entity.User) (err error) {
	_, err = r.db.Insert(user)
	return err
}

func (r *userRepo) DelUser(ctx context.Context, userID int) (err error) {
	_, err = r.db.ID(userID).Delete(&entity.User{})
	return err
}

func (r *userRepo) GetUser(ctx context.Context, userID int) (user *entity.User, exist bool, err error) {
	user = &entity.User{ID: userID}
	exist, err = r.db.Get(user)
	return user, exist, err
}
