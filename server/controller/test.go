package controller

import (
	"QN/entity"
	"QN/utils"
	"encoding/json"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request){
	response := utils.CheckReMe(r, "POST")
	if response.Status != 200{
		utils.SendJSONResponse(w, response)
		return
	}
	decoder := json.NewDecoder(r.Body)
    var requestData entity.RequestData
    err := decoder.Decode(&requestData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		utils.SendJSONResponse(w, entity.ResponseData{Mess: err.Error(), Status: 403})
        return
    }
    utils.SendJSONResponse(w, entity.ResponseData{Mess: "POST成功!", Status: 200})
}