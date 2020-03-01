package util

import (
	"bytes"
	"encoding/json"
	"giler-backend/db"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// SetWXAccessToken 设置微信接口调用凭据
func SetWXAccessToken() {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + os.Getenv("WX_APPID") + "&secret=" + os.Getenv("WX_SECRET"))
	if err != nil {
		log.Println("无法访问外部服务器：", err)
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
			log.Println("转码失败：", err)
		}
	}

	param := make(map[string]interface{})
	_ = json.NewDecoder(result).Decode(&param)

	redis := db.RedisConn
	_, _ = redis.Set("access_token", param["access_token"], time.Duration(param["expires_in"].(float64))*time.Second).Result()
}

// GetWXAccessToken 获取微信接口调用凭据
func GetWXAccessToken() string {
	redis := db.RedisConn
	for {
		val, err := redis.Get("access_token").Result()
		if err != nil {
			SetWXAccessToken()
		} else {
			return val
		}
	}
}

