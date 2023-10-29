package entity

type UserEntity struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	PassWord string `json:"password"` 
}