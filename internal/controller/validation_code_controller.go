package controller

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/email"
	"bookkeeping/internal/model"
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Body struct {
	Email string `json:"email" binding:"required"`
}

type ValidationCodeController struct{}

func (v *ValidationCodeController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	{
		v1.POST("/validation_codes", v.Create)
	}
}

func (v *ValidationCodeController) Create(c *gin.Context) {
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "params err")
		return
	}

	str, err := generateDigits()

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

func (v *ValidationCodeController) Destroy(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (v *ValidationCodeController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (v *ValidationCodeController) Get(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (v *ValidationCodeController) GetPaged(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func generateDigits() (string, error) {
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
