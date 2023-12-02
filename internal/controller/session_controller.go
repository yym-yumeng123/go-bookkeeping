package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"bookkeeping/internal/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type RequestBody struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

type SessionController struct{}

func (s SessionController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	{
		v1.POST("/session", s.Create)
	}
}

func (s SessionController) Create(c *gin.Context) {
	var requestBody RequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "params error",
		})
		return
	}

	fmt.Println(requestBody)

	// select * from validation_codes where email = '' and code = '' order by created_at asc;
	tx := database.DB.
		Where(&model.ValidationCode{Email: requestBody.Email, Code: requestBody.Code}).
		First(&model.ValidationCode{})

	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "invalid code",
		})
		return
	}

	// 找到对应的用户, 有则使用, 无则创建
	user := model.User{}
	tx = database.DB.Where("email = ?", requestBody.Email).First(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		database.DB.Create(&model.User{Email: requestBody.Email}).First(&user)
	}

	jwt, err := utils.GenerateJWT(user.ID)
	if err != nil {
		log.Println("GenerateJWT fail", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "try",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"jwt":     jwt,
	})
}

func (s SessionController) Destroy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s SessionController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s SessionController) Get(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s SessionController) GetPaged(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
