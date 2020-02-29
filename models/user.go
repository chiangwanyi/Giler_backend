package models

import (
	"giler-backend/db"
	"time"

	"github.com/globalsign/mgo/bson"
)

// CURD Automatic Generated:1,2,3,4,9,10,11
// Serializer Automatic Generated:1,2,6,7,8,9,10,11
// Update Automatic Generated:6,7,9

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

/* Automatic Generated */

// CreateUser 创建 User
func (user *User) CreateUser() error {
	session := db.MongoDBConn.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.DeletedAt = nil
	user.Birthday = time.Date(1970, 1, 1, 12, 0, 0, 0, time.Local)
	user.Tags = []string{}
	user.Friends = []bson.ObjectId{}

	return client.Insert(user)
}

// QueryUserByID 通过 ID 查找 User
func QueryUserByID(id bson.ObjectId) (user User, err error) {
	session := db.MongoDBConn.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	if err = client.Find(bson.M{"$and": []bson.M{bson.M{"_id": id}, bson.M{"deletedAt": nil}}}).One(&user); err != nil {
		return User{}, err
	}
	return user, nil
}

// QueryUserByOpenID 通过 OpenID 查找 User
func QueryUserByOpenID(openID string) (user User, err error) {
	session := db.MongoDBConn.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	if err = client.Find(bson.M{"$and": []bson.M{bson.M{"openId": openID}, bson.M{"deletedAt": nil}}}).One(&user); err != nil {
		return User{}, err
	}
	return user, nil
}

// QueryAllUser 查找所有 User
func QueryAllUser() (list []User, err error) {
	session := db.MongoDBConn.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	if err = client.Find(bson.M{"deletedAt": nil}).All(&list); err == nil {
		return list, nil
	}
	return []User{}, err
}

// UpdateUser 修改 User
func UpdateUser(user User) error {
	session := db.MongoDBConn.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	selector := bson.M{"$and": []bson.M{bson.M{"_id": user.ID}, bson.M{"deletedAt": nil}}}
	return client.Update(selector, user)
}

// DeleteUser 删除 User
func DeleteUser(user User) error {
	session := db.MongoDBConn.Copy()
	defer session.Close()
	client := session.DB("").C("user")

	selector := bson.M{"$and": []bson.M{bson.M{"_id": user.ID}, bson.M{"deletedAt": nil}}}
	update := bson.M{"$set": bson.M{"deletedAt": time.Now()}}
	return client.Update(selector, update)
}