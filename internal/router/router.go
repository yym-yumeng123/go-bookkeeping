package router

import (
	"bookkeeping/internal/controller/ping"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", ping.Ping)

	return r
}
