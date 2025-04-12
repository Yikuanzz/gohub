package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yikuanzz/unitest/entity"
	"github.com/yikuanzz/unitest/mock"
	"go.uber.org/mock/gomock"
)

func TestUserService_AddUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepo := mock.NewMockUserRepo(ctl)
	userInfo := &entity.User{Name: "test"}

	mockUserRepo.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(nil)

	userService := NewUserService(mockUserRepo)
	err := userService.AddUser(context.TODO(), userInfo.Name)
	assert.NoError(t, err)
}

func TestUserService_GetUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userID := 1
	username := "test"

	mockUserRepo := mock.NewMockUserRepo(ctl)

	// return user when userID is 1
	mockUserRepo.EXPECT().GetUser(context.TODO(), userID).Return(&entity.User{ID: userID, Name: username}, true, nil)

	userService := NewUserService(mockUserRepo)
	userInfo, err := userService.GetUser(context.TODO(), userID)
	assert.NoError(t, err)
	assert.Equal(t, userID, userInfo.ID)
}
