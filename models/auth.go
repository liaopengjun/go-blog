package models

import "gin-blog/global"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth 获取用户
func CheckAuth(username, password string) bool {
	var auth Auth
	global.Db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
