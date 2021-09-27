package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"priviledge/common"
	"priviledge/model"
	"priviledge/response"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 获取验证后的userid

		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户
		if user.ID == 0 {
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足")
			c.Abort()
			return
		}

		// 用户存在 将User信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
