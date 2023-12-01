package controllertest

import (
	"bookkeeping/internal/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMe(t *testing.T) {
	r := router.New()
	response := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/me", nil)

	r.ServeHTTP(response, req)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "pong", response.Body.String())
}
