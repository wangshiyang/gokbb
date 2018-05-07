package jwt

import (
	"github.com/gin-gonic/gin"
	"shawn/gokbb/common/exception"
	"shawn/gokbb/common/util"
	"time"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = exception.SUCCESS
		token := c.Query("token")
		if token == "" {
			token = c.GetHeader("token")
		}

		if token == "" {
			code = exception.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = exception.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = exception.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != exception.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  exception.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
