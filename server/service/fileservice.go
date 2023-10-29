package service

import (
	"QN/entity"
	"QN/utils"
	"errors"
	"fmt"
	"strings"
    "net/url"
    "context"
    "net/http"
	"github.com/qiniu/go-sdk/v7/storage"
)


func GetFileFromPrefix(prefix string) (map[string][]entity.FileEntity, error){
	var fileEntityMap map[string][]entity.FileEntity
	err := errors.New("")
	fileEntityMap, err = getAllFilesByPrefix(prefix)
	return fileEntityMap, err
}

func getAllFilesByPrefix(prefix string) (map[string][]entity.FileEntity, error) {
    limit := 1000
    delimiter := ""
    marker := ""

    bucketManager := utils.GetBM()
    fileEntitiesByPrefix := make(map[string][]entity.FileEntity)
    err := errors.New("")

    for {
        entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(utils.GetBucket(), prefix, delimiter, marker, limit)
        if err != nil {
            fmt.Println("list error,", err)
            return fileEntitiesByPrefix, err
        }

        for _, entry := range entries {
            fileEntity := processData(entry)
            fileEntitiesByPrefix[fileEntity.Prefix] = append(fileEntitiesByPrefix[fileEntity.Prefix], fileEntity)
        }

        if hasNext {
            marker = nextMarker
        } else {
            // 列举结束
            break
        }
    }
    return fileEntitiesByPrefix, err
}

func UploadFile(file entity.FileEntity) entity.ResponseData{
	localFile := file.UpPath
	if localFile == ""{
		return entity.ResponseData{Mess: "The upload file path is illegal!", Status: http.StatusNotFound}
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
	err := formUploader.PutFile(context.Background(), &ret, utils.GetUptoken(), key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return entity.ResponseData{Mess: err.Error(), Status: http.StatusNotImplemented}
	}
    //TODO 存到数据库     
	fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
	return entity.ResponseData{Mess: fmt.Sprintf("%v %v %v %v %v", ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name), Status: http.StatusOK}
}

func processData(entry storage.ListItem) entity.FileEntity {
    parts := strings.Split(entry.Key, "/")
    key := parts[len(parts)-1]
    prefix := strings.Join(parts[:len(parts)-1], "/")
    
    fileInfo, _ := utils.GetBM().Stat(utils.GetBucket(), entry.Key)
    fmt.Println(fileInfo.Hash)

    result := entity.FileEntity{
        UpPath:  "",
        Key:     key,
        DownUrl: entry.Key,
        Prefix:  prefix,
        Brief: "查了数据库再说",
    }
	utils.GetDownUrl(&result)
    return result
}