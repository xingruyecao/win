package controller

import (
	"QN/entity"
	"QN/utils"
	"fmt"
	"net/http"
)

func upLoad(w http.ResponseWriter, r *http.Request) {
	response := utils.CheckReMe(r, "GET")
	if response.Status != 200{
		utils.SendJSONResponse(w, response)
		return
	}
	// path := "../images/234.jpg"
	path := r.URL.Query().Get("path")
	key := r.URL.Query().Get("key")
	fmt.Println(path, key)
	fileEntity := entity.FileEntity{Path: path, Key: key}
	token := utils.GetToken()
	re, status:= utils.UploadFile(token, fileEntity.Path, fileEntity.Key)
	response.Mess = re
	response.Status = status
	utils.SendJSONResponse(w, response)
}
