package middleware

import (
	"GOproject/GIT/mail/pkg/e"
	"GOproject/GIT/mail/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthTokenTimeout
			}
		}
		if code != e.Success {
			c.JSON(200, gin.H{
				"status":  code,
				"message": e.GetMsg(code),
				"data":    data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
