package routes

import (
	"github.com/gin-gonic/gin"
	v1 "tsuhaoblog/api/v1"
	"tsuhaoblog/middleware"
	"tsuhaoblog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(middleware.Logger(), middleware.Cors())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块路由接口
		auth.GET("users", v1.GetUserList)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.GET("user/:id", v1.GetUser)

		//文章模块路由接口
		auth.POST("art/add", v1.AddArt)
		auth.PUT("art/:id", v1.EditArt)
		auth.DELETE("art/:id", v1.DeleteArt)

		//分类模块路由接口
		auth.POST("cate/add", v1.AddCate)
		auth.PUT("cate/:id", v1.EditCate)
		auth.DELETE("cate/:id", v1.DeleteCate)

		//上传
		auth.POST("upload", v1.Upload)
	}

	router := r.Group("api/v1")
	{
		//用户模块路由接口
		router.POST("user/add", v1.AddUser)

		//文章模块路由接口
		router.GET("art/:id", v1.GetArt)
		router.GET("arts", v1.GetArtList)

		//分类模块路由接口
		router.GET("cate", v1.GetCateList)
		router.GET("cateArt/:cid", v1.GetCateArt)

		//登陆模块
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
