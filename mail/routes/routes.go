package routes

import (
	API "GOproject/GIT/mail/api/v1"
	"GOproject/GIT/mail/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		v1.POST("user/register", API.UserRegister)
		v1.POST("user/login", API.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.Jwt())
		{
			authed.PUT("user", API.UserUpdate)
			authed.POST("avatar", API.UploadAvatar)
		}
	}
	return r
}
