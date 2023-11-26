package controllertest_test

import (
	"bookkeeping/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	r := router.New()
	response := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	r.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "pong", response.Body.String())
}
