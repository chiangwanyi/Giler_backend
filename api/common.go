package api

import (
	"bytes"
	"encoding/json"
	"giler-backend/serializer"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

// GetUserOpenID 获取用户 OpenID
func GetUserOpenID(c *gin.Context) {
	client := &http.Client{Timeout: 5 * time.Second}
	code := c.Query("code")
	resp, err := client.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + os.Getenv("WX_APPID") + "&secret=" + os.Getenv("WX_SECRET") + "&js_code=" + code + "&grant_type=authorization_code")
	if err != nil {
		c.JSON(http.StatusInternalServerError, serializer.InternalServerErrorResponse(err, "无法访问外部服务器"))
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, serializer.InternalServerErrorResponse(err, "转码失败"))
		}
	}

	param := make(map[string]interface{})
	_ = json.NewDecoder(result).Decode(&param)

	c.JSON(http.StatusOK, serializer.Response{
		Status: serializer.OK,
		Msg:    "访问外部服务器成功",
		Data:   param,
		Error:  nil,
	})
}
