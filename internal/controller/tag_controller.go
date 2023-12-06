package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagController struct{}

func (t *TagController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	{
		v1.POST("/tags", t.Create)
	}
}

func (t *TagController) Create(c *gin.Context) {
	var body model.CreateTagRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusUnprocessableEntity, "params error")
	}

	me, _ := c.Get("me")
	user, _ := me.(model.User)
	tx := database.DB.Create(&model.Tag{
		Name:   body.Name,
		UserId: user.ID,
	})

	if tx.Error != nil {
		c.String(http.StatusInternalServerError, "net error")
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "创建成功",
	})
}

func (t *TagController) Destroy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *TagController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *TagController) Get(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *TagController) GetPaged(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
