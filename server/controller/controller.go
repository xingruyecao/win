package controller

import (
	"net/http"
)


func New() *http.ServeMux {
	sm := http.DefaultServeMux
	Register(sm)
	return sm
}

func Register(sm *http.ServeMux) {
	sm.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("首页"))
	}))
	sm.HandleFunc("/login", login)
	sm.HandleFunc("/upload", withSessionCheck(upLoad))
	sm.HandleFunc("/getdownurl", withSessionCheck(getDownUrl))
	sm.HandleFunc("/getfile", withSessionCheck(getAllFile))
}
