package service

func CheckValidUser(username, password string) bool {
	return (username == "user" && password == "123456") 	
}