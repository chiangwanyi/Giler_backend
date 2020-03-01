package util

import (
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"os"
)

// GetFormUploader 获取七牛云文件上传模块
func GetFormUploader(bucket string) (*storage.FormUploader, string){
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(os.Getenv("QINIUC_ACCESSKEY"), os.Getenv("QINIUC_SECRETKEY"))
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	return storage.NewFormUploader(&cfg), upToken
}
