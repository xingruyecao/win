package utils

import (
	"QN/entity"
	// "fmt"
	"net/url"
	// "time"

	"github.com/qiniu/go-sdk/v7/storage"
)

func GetDownUrl(file *entity.FileEntity){
	// deadline := time.Now().Add(time.Second * time.Duration(timeLimit)).Unix() //有效期
	key, _ := url.JoinPath(file.Prefix, file.Key)
	// fmt.Println(key)
	// file.DownUrl = storage.MakePrivateURL(mac, domain, key, deadline)
	file.DownUrl = storage.MakePublicURL(domain, key)
}