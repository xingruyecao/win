package utils

import (
	"net/http"
	"encoding/json"
)

func SendJSONResponse(w http.ResponseWriter, data interface{}) {
    jsonResponse, err := json.Marshal(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}