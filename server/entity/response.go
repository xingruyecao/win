package entity

type RequestData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ResponseData struct {
	Status  int    `json:"status"`
	Mess string `json:"mess"`
	Data interface{} `json:"data"`
}
