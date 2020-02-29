package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// CURD Automatic Generated:1,2,3,4,9,10,11
// Serializer Automatic Generated:1,2,5,6,7,8,9,10,11
// Update Automatic Generated:6,7,8

// User User 模型
type User struct {
	ID        bson.ObjectId   `bson:"_id"`
	CreatedAt time.Time       `bson:"createdAt"`
	UpdatedAt time.Time       `bson:"updatedAt"`
	DeletedAt interface{}     `bson:"deletedAt"`
	OpenID    string          `bson:"openId"`
	Nickname  string          `bson:"nickname"`
	Gender    int             `bson:"gender"`
	Avatar    string          `bson:"avatar"`
	Birthday  time.Time       `bson:"birthday"`
	Tags      []string        `bson:"tags"`
	Friends   []bson.ObjectId `bson:"friends"`
}
