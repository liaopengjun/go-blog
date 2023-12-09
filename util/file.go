package util

import (
	"os"
)

// MkDir 创建目录
func MkDir(filePath string) error {
	// 获取根路径
	dir, _ := os.Getwd()
	// MkdirAll 创建目录赋权限0777
	return os.MkdirAll(dir+"/"+filePath, os.ModePerm)
}
