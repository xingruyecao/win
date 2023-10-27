package service

import (
	"QN/entity"
	"QN/utils"
	"errors"
	"fmt"
	"strings"

	"github.com/qiniu/go-sdk/v7/storage"
)


func GetFileFromPrefix(prefix string) ([]entity.FileEntity, error){
	var fileEntity []entity.FileEntity
	err := errors.New("")
	fileEntity, err = getAllFile(prefix)
	return fileEntity[:], err
}

func getFileFromPrefix(prefix string) ([]entity.FileEntity, error){
	panic("")
}

func getAllFile(prefix string) ([]entity.FileEntity, error){
	limit := 1000
	// prefix := "qshell/"
	delimiter := ""
	//初始列举marker为空
	marker := ""

	bucketManager := utils.GetBM()
	var fileEntities []entity.FileEntity
	err := errors.New("")

	for {
		entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(utils.GetBucket(), prefix, delimiter, marker, limit)
		if err != nil {
			fmt.Println("list error,", err)
			return fileEntities, err
		}
		//print entries
		for _, entry := range entries {
			fileEntity := processData(entry)
			fileEntities = append(fileEntities, fileEntity)
			// fmt.Println(fileEntity)
		}
		if hasNext {
			marker = nextMarker
		} else {
			//list end
			break
		}
	}
	return fileEntities, err
}

func processData(entry storage.ListItem) entity.FileEntity {
    // 根据"/"分割Key
    parts := strings.Split(entry.Key, "/")

    // 提取最后一部分作为key
    key := parts[len(parts)-1]

    // 剩下的部分作为Prefix
    prefix := strings.Join(parts[:len(parts)-1], "/")

    // 创建FileEntity对象并返回
    result := entity.FileEntity{
        UpPath:  "",
        Key:     key,
        DownUrl: entry.Key,
        Prefix:  prefix,
    }

	utils.GetDownUrl(&result)

    return result
}
