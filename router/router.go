package router

import (
	"comment/api"
	"comment/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		v1.POST("register", api.Register)
		v1.POST("login", api.Login)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("comment", api.CreateComment)
			authed.GET("comment/:article_id", api.CommentList)
			authed.GET("me/info", api.CurrentUserInfo)
			authed.POST("star/create", api.CreateStar)
			authed.POST("star/delete", api.DeleteStar)
			authed.GET("star/list", api.StarList)
			authed.GET("stared/:article_id", api.Stared)
		}
	}

	r.StaticFS("./static", http.Dir("./static"))
	return r
}
