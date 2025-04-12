package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yikuanzz/unitest/entity"
)

//go:generate mockgen -source=./user.go -destination=../mock/user_service_mock.go -package=mock
type UserService interface {
	AddUser(ctx context.Context, username string) (err error)
	GetUser(ctx context.Context, userID int) (user *entity.User, err error)
}

type AddUserRequest struct {
	Username string `json:"username" binding:"required"`
}

type GetUserRequest struct {
	UserID int `form:"user_id" binding:"required"`
}

type GetUserResponse struct {
	Username string `json:"username"`
}

type UserController struct {
	UserService UserService
}

func NewUserController(userService UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) AddUser(ctx *gin.Context) {
	req := &AddUserRequest{}
	if err := ctx.BindJSON(req); err != nil {
		return
	}

	if err := uc.UserService.AddUser(ctx, req.Username); err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	req := &GetUserRequest{}
	if err := ctx.BindQuery(req); err != nil {
		return
	}

	user, err := uc.UserService.GetUser(ctx, req.UserID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, &GetUserResponse{Username: user.Name})
}
