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
		router.POST("art/add", v1.AddArt)
		router.GET("art/:id", v1.GetArt)
		router.GET("arts", v1.GetArtList)
		router.PUT("art/:id", v1.EditArt)
		router.DELETE("art/:id", v1.DeleteArt)

		//分类模块路由接口
		router.POST("cate/add", v1.AddCate)
		router.GET("cate", v1.GetCateList)
		router.GET("cateArt/:cid", v1.GetCateArt)
		router.PUT("cate/:id", v1.EditCate)
		router.DELETE("cate/:id", v1.DeleteCate)
	}
	r.Run(utils.HttpPort)
}
