package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nanfeng/ginchat/internal/model"
	"github.com/nanfeng/ginchat/internal/pkg/utils"
	"github.com/nanfeng/ginchat/internal/pkg/xerr"
)

func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusOK, model.Response{
				Code: xerr.CodeUnauthorized,
				Msg:  "请求头中auth为空",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, model.Response{
				Code: xerr.CodeUnauthorized,
				Msg:  "auth为空",
			})
			c.Abort()
			return
		}

		// 解析token
		fmt.Println(parts[1])
		username, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code: xerr.CodeUnauthorized,
				Msg:  err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("username", username)

		c.Next()
	}
}
