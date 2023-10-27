package utils

import (
	"QN/entity"
	"net/http"
)

func CheckReMe(r *http.Request, method string) *entity.ResponseData{
	if r.Method == method{
		return &entity.ResponseData{Status: 200, Mess: ""}
	}else{
		return &entity.ResponseData{Status: 403, Mess: "请求方法错误！"}
	}
}