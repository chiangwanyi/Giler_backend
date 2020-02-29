package services

import (
	"giler-backend/models"
	"giler-backend/serializer"
	"github.com/globalsign/mgo/bson"
	"time"
)

// UpdateUserService 修改 User 服务
type UpdateUserService struct {
	Nickname string    `json:"nickname"`
	Gender   int       `json:"gender"`
	Birthday time.Time `json:"birthday"`
}

// UpdateUserField 修改 User
func (service *UpdateUserService) UpdateUserField(id string) serializer.Response {
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

		if service.Nickname != "" {
			user.Nickname = service.Nickname
		}
		if service.Gender != -1 {
			user.Gender = service.Gender
		}
		// if service.Birthday != nil {
		// 	user.Birthday = service.Birthday
		// }

		if err := models.UpdateUser(user); err != nil {
			return serializer.InternalServerErrorResponse(err, "修改失败")
		}
		return serializer.Response{
			Status: serializer.OK,
			Data:   nil,
			Msg:    "修改成功",
			Error:  nil,
		}
	}
	return serializer.BadParameterErrorResponse(nil)
}
