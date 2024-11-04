package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"shortener-auth/internal/routing"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	router := gin.Default()

	routing.Register(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
