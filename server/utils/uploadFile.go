package utils

import (
	// "github.com/qiniu/go-sdk/v7/auth/qbox"
    "github.com/qiniu/go-sdk/v7/storage"
	"fmt"
	"QN/entity"
	"context"
)

func UploadFile(upToken string, localFile string, key string) (string, int){
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := entity.PutpetEntity{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "XTU",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err.Error(), 403
	}
	fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
	return fmt.Sprintf("%v %v %v %v %v", ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name), 200
}