package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"bookkeeping/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
)

type MeController struct{}

func (m *MeController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	{
		v1.GET("/me", m.Get)
	}
}

func (m *MeController) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (m *MeController) Destroy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (m *MeController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (m *MeController) Get(c *gin.Context) {
	user := model.User{}
	// 客户端发送请求携带 http header Authorization = Bearer <jwt>
	auth := c.GetHeader("Authorization")
	jwtString := auth[7:]
	// 解析 jwtString 获取 user_id
	t, err := utils.Parse(jwtString)
	if err != nil {
		c.String(http.StatusUnauthorized, "invalid jwt")
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		if userId, ok := claims["user_id"].(string); !ok {
			c.String(http.StatusUnauthorized, "invalid jwt")
			return
		} else {
			if userIdInt, err := strconv.Atoi(userId); err != nil {
				c.String(http.StatusUnauthorized, "invalid jwt")
			} else {
				// 通过 user_id 获取 user
				tx := database.DB.First(&user, "id = ?", userIdInt)
				if tx.Error != nil {
					c.JSON(http.StatusOK, gin.H{
						"success": true,
						"data":    user,
					})
				}
			}
		}
	} else {
		fmt.Println(err)
	}
}

func (m *MeController) GetPaged(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
