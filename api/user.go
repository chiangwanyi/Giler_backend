package api

import (
	"giler-backend/serializer"
	"giler-backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser 创建 User
func CreateUser(c *gin.Context) {
	s := services.CreateUserService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(http.StatusBadRequest, serializer.BadParameterErrorResponse(err))
	} else {
		c.JSON(http.StatusOK, s.Create())
	}
}

// QueryUser 查找 User
func QueryUser(c *gin.Context) {
	id := c.Query("id")
	oid := c.Query("openid")
	if id != "" {
		s := services.QueryUserService{}
		c.JSON(http.StatusOK, s.Query(id))
	} else {
		s := services.QueryUserByOpenIDService{}
		c.JSON(http.StatusOK, s.Query(oid))
	}
}

// ListUser 查找所有 User
func ListUser(c *gin.Context) {
	s := services.ListUserService{}
	c.JSON(http.StatusOK, s.List())
}

// UpdateUser 更新 User
func UpdateUser(c *gin.Context) {
	s := services.UpdateUserService{}
	if err := c.ShouldBind(&s); err != nil {
		c.JSON(http.StatusBadRequest, serializer.BadParameterErrorResponse(err))
	} else {
		c.JSON(http.StatusOK, s.UpdateUserField(c.Param("id")))
	}
}

// DeleteUser 删除 User
func DeleteUser(c *gin.Context) {
	s := services.DeleteUserService{}
	c.JSON(http.StatusOK, s.Delete(c.Param("id")))
}

