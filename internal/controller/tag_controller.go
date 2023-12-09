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
		v1.DELETE("/tags/:id", t.Destroy)
		v1.GET("/tags/:id", t.Get)
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
	idString := c.Param("id")

	tx := database.DB.Where("id = ?", idString).Delete(&model.Tag{})

	if tx.Error != nil {
		c.String(http.StatusUnprocessableEntity, "参数有误")
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "删除成功",
	})
}

func (t *TagController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (t *TagController) Get(c *gin.Context) {
	tag := model.Tag{}
	me, _ := c.Get("me")
	user, _ := me.(model.User)
	idString := c.Param("id")

	database.DB.Where("user_id = ?", user.ID).Where("id =?", idString).First(&tag)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tag,
	})

}

func (t *TagController) GetPaged(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
