package entity

type FileEntity struct {
	UpPath string `json:"uppath"`
	Key string `json:"key"`
	DownUrl string `json:"downurl"`
	Prefix string `json:"prefix"`
	Brief string `json:"brief"`
}