package utils

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

var mac *auth.Credentials
var upToken string
var timeLimit int = 3600
var bucket string = "winmax"
var domain string = "s34kk0k9i.hn-bkt.clouddn.com"
var bucketManager *storage.BucketManager 

func GetToken() (value string) {
	accessKey := "nFRUzQdAMrmIhX9Nkqtjy9B2HmKfiKjHrzwNxTs1"
	secretKey := "zLFsST97cMHttlpoaa0EOY3O74E7lnk6GqvS7KD5"
	mac = qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager = storage.NewBucketManager(mac, &cfg)
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	putPolicy.Expires = uint64(timeLimit) //示例2小时有效期
	upToken = putPolicy.UploadToken(mac)
	fmt.Println("获取到token!!")
	return upToken
}

func GetBM() *storage.BucketManager{
	return bucketManager
}

func GetBucket() string{
	return bucket
}

func GetUptoken() string{
	return upToken
}