package controller

import (
	"QN/entity"
	"QN/utils"
	"encoding/json"
	"QN/service"
	"errors"
	"net/http"
)

func getDownUrl(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		utils.SendJSONResponse(w, entity.ResponseData{Mess: "Method not allowed!", Status: http.StatusMethodNotAllowed})
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
	utils.SendJSONResponse(w, entity.ResponseData{Mess: "SUCCESS!", Status: 200, Data: requestData})
}

func getAllFile(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		utils.SendJSONResponse(w, entity.ResponseData{Mess: "Method not allowed!", Status: http.StatusMethodNotAllowed})
        return
    }
	prefix := r.URL.Query().Get("prefix")
	fileEntityListMap, err := service.GetFileFromPrefix(prefix)
	status := 200
	err = errors.New("SUCCESS")
	if len(fileEntityListMap) == 0 {
		status = 403
		err = errors.New("The data queried is empty!")
	}
	utils.SendJSONResponse(w, entity.ResponseData{Mess: err.Error(), Status: status, Data: fileEntityListMap})
}

func upLoad(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendJSONResponse(w, entity.ResponseData{Mess: "Method not allowed!", Status: http.StatusMethodNotAllowed})
        return
    }
	decoder := json.NewDecoder(r.Body)
    var requestData entity.FileEntity
    err := decoder.Decode(&requestData)
	if requestData.UpPath == "" || requestData.Key == "" || requestData.Prefix == ""{
		err = errors.New("Incomplete attribute values!")
	}
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		utils.SendJSONResponse(w, entity.ResponseData{Mess: err.Error(), Status: 403})
        return
    }
	utils.SendJSONResponse(w, service.UploadFile(requestData))
}

