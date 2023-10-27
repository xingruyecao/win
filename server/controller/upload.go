package controller

import (
	"QN/entity"
	"QN/utils"
	"encoding/json"
	"errors"

	// "fmt"
	"net/http"
)

func upLoad(w http.ResponseWriter, r *http.Request) {
	response := utils.CheckReMe(r, "POST")
	if response.Status != 200{
		utils.SendJSONResponse(w, response)
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
	re, status:= utils.UploadFile(requestData)
	response.Mess = re
	response.Status = status
	utils.SendJSONResponse(w, response)
}
