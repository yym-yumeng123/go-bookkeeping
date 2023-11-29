package controller

import (
	"bookkeeping/internal/email"
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

	if err := email.SendValidationCode(body.Email, "123456"); err != nil {
		log.Println(err, "------")
		c.String(500, "发送失败")
		return
	}

	log.Println("------------")
	log.Println(body.Email)
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// })
	c.String(200, "ok")
}
