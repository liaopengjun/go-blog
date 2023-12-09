package middleware

import (
	"gin-blog/common/config"
	"gin-blog/common/response"
	"gin-blog/pkg/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

// jwt 中间件
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token
		token := c.Query("token")
		if token == "" {
			c.Abort() //阻止调用后续处理函数
			response.ResponseError(c, config.NOT_TOKEN)
		}

		// 解析token
		claims, err := jwt.ParseToken(token)

		// token解析失败
		if err != nil {
			c.Abort() //阻止调用后续处理函数
			response.ResponseError(c, config.ERROR_AUTH_CHECK_TOKEN_FAIL)
		}

		// token已过期
		if time.Now().Unix() > claims.ExpiresAt {
			c.Abort()
			response.ResponseError(c, config.ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
		}

		//gin上下文传值
		//c.Set("username", claims.Username)

		//gin 从下文取值
		//c.Get("username")

		c.Next() //调用后续处理函数
	}
}
