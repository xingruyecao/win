package controller

import (
	"QN/entity"
	"QN/service"
	"QN/utils"
	"errors"
	"net/http"
)

func getAllFile(w http.ResponseWriter, r *http.Request){
	response := utils.CheckReMe(r, "GET")
	if response.Status != 200{
		utils.SendJSONResponse(w, response)
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