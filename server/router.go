package server

import (
	"DuckyGo/api"
	"DuckyGo/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		r.Use(middleware.SaveLog())
	}

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		v1.GET("me", api.GetComment)

		// 用户注册
		//		v1.POST("user/register", api.UserRegister)

		// 用户登录
		// v1.POST("user/login", api.UserLogin)

	}
	return r
}
