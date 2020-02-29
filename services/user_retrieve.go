package services

import (
	"giler-backend/models"
	"giler-backend/serializer"
	"github.com/globalsign/mgo/bson"
)

// QueryUserService 查找 User 服务
type QueryUserService struct {
}

// QueryUserByOpenIDService 通过 OpenID 查找 User 服务
type QueryUserByOpenIDService struct {
}

// ListUserService 查找所有 User 服务
type ListUserService struct {
}


// Query 通过 ID 查找 User
func (service *QueryUserService) Query(id string) serializer.Response {
	if bson.IsObjectIdHex(id) {
		user, err := models.QueryUserByID(bson.ObjectIdHex(id))
		if err != nil {
			return serializer.Response{
				Status: serializer.OK,
				Data:   nil,
				Msg:    "查询结果为空",
				Error:  nil,
			}
		}
		return serializer.Response{
			Status: serializer.OK,
			Data:   serializer.BuildUser(user),
			Msg:    "查找成功",
			Error:  "",
		}
	}
	return serializer.BadParameterErrorResponse(nil)
}

// Query 通过 OpenID 查找 User
func (service *QueryUserByOpenIDService) Query(oid string) serializer.Response {
	user, err := models.QueryUserByOpenID(oid)
	if err != nil {
		return serializer.Response{
			Status: serializer.OK,
			Data:   nil,
			Msg:    "查询结果为空",
			Error:  nil,
		}
	}
	return serializer.Response{
		Status: serializer.OK,
		Data:   serializer.BuildUser(user),
		Msg:    "查找成功",
		Error:  "",
	}
}

// List 查找所有 User
func (service *ListUserService) List() serializer.Response {
	list, err := models.QueryAllUser()
	if err != nil {
		return serializer.Response{
			Status: serializer.OK,
			Data:   nil,
			Msg:    "查询结果为空",
			Error:  nil,
		}
	}
	return serializer.Response{
		Status: serializer.OK,
		Data:   serializer.BuildUserList(list),
		Msg:    "查找成功",
		Error:  "",
	}
}
