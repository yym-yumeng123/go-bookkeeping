package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"fmt"
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
			"msg":     "参数错误",
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
			"msg":     "无效的验证码",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"jwt":     "xxx",
	})
}
