package controller

import "net/http"

type MyServer struct{}

func New() *http.ServeMux {
	sm := http.DefaultServeMux
	Register(sm)
	return sm
}

func Register(sm *http.ServeMux) {
	sm.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("首页"))
	}))
	sm.HandleFunc("/upload", upLoad)
	sm.HandleFunc("/test", test)
	sm.HandleFunc("/getdownurl", getDownUrl)
	sm.HandleFunc("/getfile", getAllFile)
}
