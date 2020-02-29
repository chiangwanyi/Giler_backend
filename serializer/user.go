package serializer

import (
	"giler-backend/models"
	"github.com/globalsign/mgo/bson"
	"time"
)

// User  User 序列化器
type User struct {
	ID        string          `json:"_id"`
	CreatedAt time.Time       `json:"createdAt"`
	Nickname  string          `json:"nickname"`
	Gender    int             `json:"gender"`
	Avatar    string          `json:"avatar"`
	Birthday  time.Time       `json:"birthday"`
	Tags      []string        `json:"tags"`
	Friends   []bson.ObjectId `json:"friends"`
}

// BuildUser 序列化 User
func BuildUser(item models.User) User {
	return User{
		ID:        item.ID.Hex(),
		CreatedAt: item.CreatedAt,
		Nickname:  item.Nickname,
		Gender:    item.Gender,
		Avatar:    item.Avatar,
		Birthday:  item.Birthday,
		Tags:      item.Tags,
		Friends:   item.Friends,
	}
}

// BuildUserList 序列化 User 列表
func BuildUserList(items []models.User) (list []User) {
	for _, v := range items {
		list = append(list, BuildUser(v))
	}
	return list
}
