package utils

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
    "github.com/qiniu/go-sdk/v7/storage"
	"fmt"
)

func GetToken() (value string){
	accessKey := "nFRUzQdAMrmIhX9Nkqtjy9B2HmKfiKjHrzwNxTs1"
	secretKey := "zLFsST97cMHttlpoaa0EOY3O74E7lnk6GqvS7KD5"
	mac := qbox.NewMac(accessKey, secretKey)
	
	bucket:="winmax"
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	// putPolicy.Expires = 7200 //示例2小时有效期
	upToken := putPolicy.UploadToken(mac)
	fmt.Println("获取到token!!")
	return upToken
}
