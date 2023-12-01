package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"bookkeeping/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type ResponseData struct {
	ID    int
	Email string
}

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
	response := ResponseData{}
	// 客户端发送请求携带 http header Authorization = Bearer <jwt>
	auth := c.GetHeader("Authorization")
	jwtString := auth[7:]

	if auth == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token字符串为空"})
	}

	// 解析 jwtString 获取 user_id
	token, err := utils.Parse(jwtString)

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 通过 user_id 获取 user
		tx := database.DB.Model(&model.User{}).First(&user, "id = ?", claims["user_id"]).Scan(&response)
		if tx.Error == nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"data":    response,
			})
		}
	} else {
		fmt.Println(err)
	}

}

func (m *MeController) GetPaged(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
