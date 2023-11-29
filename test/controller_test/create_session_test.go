package controllertest_test

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"bookkeeping/internal/router"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	r := router.New()
	email := "12@qq.com"
	code := "123456"
	database.DB.Create(&model.ValidationCode{Email: email, Code: code})
	w := httptest.NewRecorder()
	requestParams := gin.H{
		"email": email,
		"code":  code,
	}
	bytes, _ := json.Marshal(requestParams)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/validation_codes",
		strings.NewReader(string(bytes)),
	)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	log.Println("==============")
	log.Println(strings.NewReader(string(bytes)))
	assert.Equal(t, 200, w.Code)
}
