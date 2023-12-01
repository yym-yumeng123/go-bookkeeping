package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/email"
	"bookkeeping/internal/model"
	"crypto/rand"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Email string `json:"email" binding:"required"`
}

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
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "params err")
		return
	}

	str, err := generateDigist()

	if err != nil {
		c.String(200, "code fail")
		return
	}

	if err := email.SendValidationCode(body.Email, str); err != nil {
		log.Println(err, "------")
		c.String(500, "send fail")
		return
	}

	code := model.ValidationCode{Email: body.Email, Code: str}
	tx := database.DB.Create(&code)

	if tx.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"code":    str,
		})
	}

}

func generateDigist() (string, error) {
	len := 4
	b := make([]byte, len)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	digits := make([]byte, len)
	for i := range b {
		digits[i] = b[i]%10 + 48
	}

	return string(digits), nil
}
