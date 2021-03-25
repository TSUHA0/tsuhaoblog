package routes

import (
	"github.com/gin-gonic/gin"
	v1 "tsuhaoblog/api/v1"
	"tsuhaoblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{

		//用户模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUser)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		//文章模块路由接口

		//分类模块路由接口

	}
	r.Run(utils.HttpPort)
}
