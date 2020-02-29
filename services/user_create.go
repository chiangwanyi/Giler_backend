package services

import (
	"giler-backend/models"
	"giler-backend/serializer"
)

// CreateUserService 创建 User 服务
type CreateUserService struct {
	OpenID   string `json:"openId" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Gender   int    `json:"gender"`
	Avatar   string `json:"avatar"`
}

// Valid 校验字段
func (service *CreateUserService) Valid() *serializer.Response {
	if user, err := models.QueryUserByOpenID(service.OpenID); err == nil {
		return &serializer.Response{
			Status: serializer.ExistingError,
			Data:   serializer.BuildUser(user),
			Msg:    "该用户已经注册",
			Error:  nil,
		}
	}
	return nil
}

// Create 创建 User
func (service *CreateUserService) Create() serializer.Response {
	user := models.User{
		OpenID:   service.OpenID,
		Nickname: service.Nickname,
		Gender:   service.Gender,
		Avatar:   service.Avatar,
	}

	if err := service.Valid(); err != nil {
		return *err
	}

	if err := user.CreateUser(); err != nil {
		return serializer.InternalServerErrorResponse(err, "创建失败")
	}

	return serializer.Response{
		Status: serializer.OK,
		Data:   serializer.BuildUser(user),
		Msg:    "创建成功",
		Error:  nil,
	}
}
