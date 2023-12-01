package controller

import (
	"github.com/gin-gonic/gin"
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
	user, _ := c.Get("me")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})

}

func (m *MeController) GetPaged(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
