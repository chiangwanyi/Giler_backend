package services

import (
	"giler-backend/models"
	"giler-backend/serializer"
	"github.com/globalsign/mgo/bson"
)

// DeleteUserService 删除 User 服务
type DeleteUserService struct {
}

// Delete 删除 User
func (service *DeleteUserService) Delete(id string) serializer.Response {
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
		if err := models.DeleteUser(user); err != nil {
			return serializer.InternalServerErrorResponse(err, "删除失败")
		}
		return serializer.Response{
			Status: serializer.OK,
			Data:   nil,
			Msg:    "删除成功",
			Error:  nil,
		}
	}
	return serializer.BadParameterErrorResponse(nil)
}
