package middleware

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"bookkeeping/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// GetUserInfo 中间件
func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := getMe(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			return
		}

		// Set example variable
		c.Set("me", user)

		// before request
		c.Next()
	}
}

func getMe(c *gin.Context) (model.User, error) {
	user := model.User{}
	// 客户端发送请求携带 http header Authorization = Bearer <jwt>
	auth := c.GetHeader("Authorization")
	jwtString := auth[7:]

	if auth == "" {
		return user, fmt.Errorf("token empty")
	}

	// 解析 jwtString 获取 user_id
	token, _ := utils.Parse(jwtString)

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 通过 user_id 获取 user
		tx := database.DB.Model(&model.User{}).First(&user, "id = ?", claims["user_id"])
		if tx.Error == nil {
			return user, nil
		}
	} else {
		return user, fmt.Errorf("invalid jwt")
	}

	return user, nil
}
