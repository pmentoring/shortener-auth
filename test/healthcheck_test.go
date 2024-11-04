package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"shortener-auth/database"
	"shortener-auth/internal/common"
	"shortener-auth/internal/routing"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	router := gin.Default()
	db, err := database.GetConnection()
	if err != nil {
		return
	}
	routing.Register(router, db, common.NewApplicationContext("1", "", "trololo"))

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
