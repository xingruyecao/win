package entity

type ResponseData struct {
	Status  int    `json:"status"`
	Mess string `json:"mess"`
	Data interface{} `json:"data"`
}

type SessionData struct{
	Username string
}
