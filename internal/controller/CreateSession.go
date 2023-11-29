package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
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
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	// select * from validation_codes where email = '' and code = '' order by created_at asc;
	codes := model.ValidationCode{}
	tx := database.DB.Where("email = ? and code = ?", requestBody.Email, requestBody.Code).First(&codes)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"jwt":     "xxx",
	})
}
