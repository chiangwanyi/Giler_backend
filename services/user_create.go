package services

import (
	"bufio"
	"context"
	"giler-backend/models"
	"giler-backend/serializer"
	"giler-backend/util"
	"github.com/qiniu/api.v7/v7/storage"
	"net/http"
)

// CreateUserService 创建 User 服务
type CreateUserService struct {
	OpenID   string `json:"openId" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Gender   int    `json:"gender"`
	Avatar   string `json:"avatar" binding:"required"`
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
	res, err := http.Get(service.Avatar)
	if err != nil {
		return serializer.Response{
			Status: serializer.InternalServerError,
			Data:   nil,
			Msg:    "无法读取用户头像",
			Error:  err,
		}
	}
	defer res.Body.Close()

	formUploader, upToken := util.GetFormUploader("giler")
	ret := storage.PutRet{}

	reader := bufio.NewReaderSize(res.Body, 32*1024)
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, reader, res.ContentLength, &storage.PutExtra{})
	if err != nil {
		return serializer.Response{
			Status: serializer.InternalServerError,
			Data:   nil,
			Msg:    "上传头像失败",
			Error:  err,
		}
	}

	user := models.User{
		OpenID:   service.OpenID,
		Nickname: service.Nickname,
		Gender:   service.Gender,
		Avatar:   "http://q6ck3a2ag.bkt.clouddn.com/" + ret.Key,
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
