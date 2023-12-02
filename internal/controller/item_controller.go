package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ItemRequest struct {
	Amount int32 `json:"amount" binding:"required"`
	Kind   int   `json:"kind" binding:"required"`
	//HappenedAt time.Time `json:"happened_at" binding:"required"`
}

type ItemController struct{}

func (i *ItemController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	{
		v1.GET("/items", i.Get)
		v1.POST("/items", i.Create)
		v1.DELETE("/items", i.Destroy)
	}
}

func (i *ItemController) Create(c *gin.Context) {
	rBody := ItemRequest{}
	if err := c.ShouldBindJSON(&rBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"msg":     "params error",
		})
		return
	}

	me, _ := c.Get("me")
	user, _ := me.(model.User)
	tx := database.DB.Create(&model.Item{
		Amount:     rBody.Amount,
		Kind:       rBody.Kind,
		HappenedAt: time.Now(),
		UserId:     user.ID,
	})

	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, tx.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
	})

}

func (i *ItemController) Destroy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (i *ItemController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (i *ItemController) Get(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (i *ItemController) GetPaged(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
