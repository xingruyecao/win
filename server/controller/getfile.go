package controller

import (
	"QN/entity"
	"QN/service"
	"QN/utils"
	"net/http"
)

func getAllFile(w http.ResponseWriter, r *http.Request){
	response := utils.CheckReMe(r, "GET")
	if response.Status != 200{
		utils.SendJSONResponse(w, response)
		return
	}
	prefix := r.URL.Query().Get("prefix")
	fileEntityList, err := service.GetFileFromPrefix(prefix)
	status := 200
	if len(fileEntityList) == 0 {
		status = 403
	}
	utils.SendJSONResponse(w, entity.ResponseData{Mess: err.Error(), Status: status, Data: fileEntityList})
}