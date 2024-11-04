package test

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"shortener-auth/database"
	"shortener-auth/internal/common"
	"shortener-auth/internal/common/http_actions"
	"shortener-auth/internal/common/repository"
	"shortener-auth/internal/routing"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	router := gin.Default()

	conn, err := database.GetConnection()
	if err != nil {
		return
	}

	ctx := common.NewApplicationContext("1", "", "secret")
	repo := repository.NewUserRepository(conn)
	registerAction := http_actions.NewRegisterAction(repo, ctx)
	routing.Register(router, registerAction)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
