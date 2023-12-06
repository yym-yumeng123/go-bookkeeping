package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CreateItemRequest struct {
	Amount int32 `json:"amount" binding:"required"`
	Kind   int   `json:"kind" binding:"required"`
	//HappenedAt time.Time `json:"happened_at" binding:"required"`
}

type GetPagedItemsRequest struct {
	Page int `json:"page"`
}

type ItemController struct{}

func (i *ItemController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	{
		v1.GET("/items", i.Get)
		v1.GET("/items/paged", i.GetPaged)
		v1.POST("/items", i.Create)
		v1.DELETE("/items", i.Destroy)
	}
}

func (i *ItemController) Create(c *gin.Context) {
	rBody := CreateItemRequest{}
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
	var items []model.Item
	me, _ := c.Get("me")
	user, _ := me.(model.User)
	tx := database.DB.Where(&model.Item{UserId: user.ID}).Find(&items)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		fmt.Println("data null")
	} else if tx.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"msg":     "获取失败",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    items,
	})
}

func (i *ItemController) GetPaged(c *gin.Context) {
	var params GetPagedItemsRequest
	var items []model.Item
	var count int64
	me, _ := c.Get("me")
	user, _ := me.(model.User)
	if err := c.ShouldBindJSON(&params); err != nil {
		c.String(http.StatusBadRequest, "params error")
	}

	if params.Page == 0 {
		params.Page = 1
	}

	tx := database.DB.Where(&model.Item{UserId: user.ID}).Count(&count).Limit(10).Offset((params.Page - 1) * 10).Find(&items)
	if tx.Error != nil {
		c.String(http.StatusInternalServerError, "internet error")
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    items,
		"pager": gin.H{
			"page":     params.Page,
			"per_page": 10,
			"count":    count,
		},
	})
}
