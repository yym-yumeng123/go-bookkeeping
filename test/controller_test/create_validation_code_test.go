package controllertest_test

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"bookkeeping/internal/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidationCode(t *testing.T) {
	r := router.New()
	w := httptest.NewRecorder()
	var count1 int64
	var count2 int64
	database.DB.Find(&model.ValidationCode{}).Count(&count1)
	// 设置 header 的 content-type 为 json
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/validation_codes",
		strings.NewReader(`{"email": "18026493873@163.com"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	database.DB.Find(&model.ValidationCode{}).Count(&count2)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, count2-1, count1)
}
