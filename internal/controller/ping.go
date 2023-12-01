package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      测试 API 是否正常工作
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router      /ping [get]
func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
