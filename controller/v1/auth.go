package v1

import (
	"gin-blog/common/config"
	"gin-blog/common/response"
	"gin-blog/models"
	"gin-blog/pkg/jwt"
	"gin-blog/pkg/logging"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// 获取Token
func GetAuth(c *gin.Context) {

	// 获取参数
	username := c.Query("username")
	password := c.Query("password")

	logging.Info(username, password)

	// 校验参数
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	if !ok {
		response.ResponseError(c, config.INVALID_PARAMS)
		return
	}

	// 用户不存在
	isExist := models.CheckAuth(username, password)
	if !isExist {
		response.ResponseError(c, config.ERROR_AUTH)
		return
	}

	// 生成token
	token, err := jwt.GenerateToken(username, password)
	if err != nil {
		response.ResponseError(c, config.ERROR_AUTH_TOKEN)
		return
	}
	// 返回数据
	data := map[string]string{
		"token": token,
	}
	response.ResponseSuccess(c, data)
}
