package controllertest_test

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/model"
	"bookkeeping/internal/router"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	r := router.New()
	email := "12@qq.com"
	code := "1234"
	if tx := database.DB.Create(&model.ValidationCode{Email: email, Code: code}); tx.Error != nil {
		log.Fatalln(tx.Error)
	}
	w := httptest.NewRecorder() // 记录器

	mcPostBody := map[string]any{
		"email": email,
		"code":  code,
	}
	body, _ := json.Marshal(mcPostBody)

	req, _ := http.NewRequest(
		"POST",
		"/api/v1/session",
		bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req) // server 一下记录器
	assert.Equal(t, 200, w.Code)
}
