package controller

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yikuanzz/unitest/entity"
	"github.com/yikuanzz/unitest/mock"
	"github.com/yikuanzz/unitest/utils"
	"go.uber.org/mock/gomock"
)

func TestUserController_AddUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	req := &AddUserRequest{Username: "test"}
	mockUserService := mock.NewMockUserService(ctl)
	mockUserService.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(nil)

	userController := NewUserController(mockUserService)

	success, resp := utils.CreatePostReqCtx(req, userController.AddUser)
	assert.True(t, success)
	fmt.Println(resp)
}

func TestUserController_GetUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	req := &GetUserRequest{UserID: 1}
	user := &entity.User{Name: "test"}
	mockUserService := mock.NewMockUserService(ctl)
	mockUserService.EXPECT().GetUser(gomock.Any(), req.UserID).Return(user, nil)

	userController := NewUserController(mockUserService)

	success, resp := utils.CreateGetReqCtx(req, userController.GetUser)
	assert.True(t, success)
	fmt.Println(resp)
}
