package routers

import (
	v1 "gin-blog/controller/v1"
	"gin-blog/global"
	"gin-blog/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())   //gin日志中间件
	r.Use(gin.Recovery()) // gin记录错误中间件

	//  debug release
	gin.SetMode(global.Config.AppConfig.RunMode)

	// 获取令牌
	r.GET("/auth", v1.GetAuth)
	r.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "编译成功",
		})
	})

	// 路由
	apiv1 := r.Group("/api/v1")

	// 中间件路由
	apiv1.Use(middleware.Jwt())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
