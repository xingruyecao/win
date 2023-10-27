package utils

import (
	// "github.com/qiniu/go-sdk/v7/auth/qbox"
	"QN/entity"
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadFile(file entity.FileEntity) (string, int){
	localFile := file.UpPath
	if localFile == ""{
		return errors.New("The upload file path is illegal!").Error(), 404
	}
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := entity.PutpetEntity{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "XTU",
		},
	}
	key, _ := url.JoinPath(file.Prefix, file.Key)
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err.Error(), 403
	}
	fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
	return fmt.Sprintf("%v %v %v %v %v", ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name), 200
}