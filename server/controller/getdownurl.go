package controller

import (
	"QN/entity"
	"QN/utils"
	"encoding/json"
	"errors"
	"net/http"
)

func getDownUrl(w http.ResponseWriter, r *http.Request){
	response := utils.CheckReMe(r, "POST")
	if response.Status != 200{
		utils.SendJSONResponse(w, response)
		return
	}
	decoder := json.NewDecoder(r.Body)
    var requestData entity.FileEntity
    err := decoder.Decode(&requestData)
	if requestData.Key=="" || requestData.Prefix==""{
		err = errors.New("key or prefix is null")
	}
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		utils.SendJSONResponse(w, entity.ResponseData{Mess: err.Error(), Status: 403})
        return
    }
	utils.GetDownUrl(&requestData)
	utils.SendJSONResponse(w, entity.ResponseData{Mess: "成功获取下载路径！", Status: 200, Data: requestData})
}