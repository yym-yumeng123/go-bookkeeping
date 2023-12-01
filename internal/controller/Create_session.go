package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"bookkeeping/internal/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func CreateSession(c *gin.Context) {
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
	tx = database.DB.Model(&user).Where("email = ?", requestBody.Email)
	if tx.Error != nil {
		database.DB.Create(&model.User{Email: requestBody.Email})
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
		"user_id": user.ID,
	})
}
