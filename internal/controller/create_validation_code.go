package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/email"
	"bookkeeping/internal/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateValidationCode godoc
// @Summary      用邮箱发送验证码
// @Description  接受邮箱地址, 发送验证码
// @Tags         code
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router      /validation_codes [post]
func CreateValidationCode(c *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "参数错误")
		return
	}

	code := model.ValidationCode{Email: body.Email, Code: "123456"}
	tx := database.DB.Create(&code)

	if err := email.SendValidationCode(body.Email, "123456"); err != nil {
		log.Println(err, "------")
		c.String(500, "发送失败")
		return
	}

	if tx.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}

}
