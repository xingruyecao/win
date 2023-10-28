package service

import (
	"QN/entity"
	"QN/utils"
	"errors"
	"fmt"
	"strings"

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

func processData(entry storage.ListItem) entity.FileEntity {
    parts := strings.Split(entry.Key, "/")
    key := parts[len(parts)-1]
    prefix := strings.Join(parts[:len(parts)-1], "/")
    result := entity.FileEntity{
        UpPath:  "",
        Key:     key,
        DownUrl: entry.Key,
        Prefix:  prefix,
    }
	utils.GetDownUrl(&result)

    return result
}
